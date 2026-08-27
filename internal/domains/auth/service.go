package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/wunlight/hermes/internal/domains/refresh_token"
	"github.com/wunlight/hermes/internal/domains/user"
	"github.com/wunlight/hermes/internal/infrastructure/security/jwt"
	"github.com/wunlight/hermes/internal/infrastructure/security/password"
)

type Service struct {
	userRepo             user.Repository
	refreshTokenRepo     refresh_token.Repository
	passwordHasher       password.Hasher
	tokenManager         jwt.TokenManager
	refreshTokenDuration time.Duration
}

func NewService(
	userRepo user.Repository,
	refreshTokenRepo refresh_token.Repository,
	passwordHasher password.Hasher,
	tokenManager jwt.TokenManager,
	refreshTokenDuration time.Duration,
) *Service {
	return &Service{
		userRepo:             userRepo,
		refreshTokenRepo:     refreshTokenRepo,
		passwordHasher:       passwordHasher,
		tokenManager:         tokenManager,
		refreshTokenDuration: refreshTokenDuration,
	}
}

func (s *Service) Register(ctx context.Context, req RegisterRequest) (*user.User, error) {
	email := strings.TrimSpace(strings.ToLower(req.Email))

	if err := validateRegisterRequest(email, req.Password, req.Name); err != nil {
		return nil, err
	}

	existingUser, err := s.userRepo.GetByEmail(ctx, email)
	if err == nil && existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}

	if err != nil && !errors.Is(err, user.ErrNotFound) {
		return nil, fmt.Errorf("check existing user: %w", err)
	}

	passwordHash, err := s.passwordHasher.Hash(req.Password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	newUser := &user.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		Name:         strings.TrimSpace(req.Name),
	}

	createdUser, err := s.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return createdUser, nil
}

func validateRegisterRequest(email string, password string, name string) error {
	if email == "" {
		return ErrEmailRequired
	}

	if !isValidEmail(email) {
		return ErrInvalidEmail
	}

	if password == "" {
		return ErrPasswordRequired
	}

	if len(password) < 8 {
		return ErrPasswordTooShort
	}

	if len(password) > 128 {
		return ErrPasswordTooLong
	}

	name = strings.TrimSpace(name)

	if name == "" {
		return ErrNameRequired
	}

	if len(name) > 255 {
		return ErrNameTooLong
	}

	return nil
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (*AuthResponse, error) {
	email := strings.TrimSpace(strings.ToLower(req.Email))

	currentUser, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return nil, ErrInvalidCredentials
		}

		return nil, fmt.Errorf("get user by email: %w", err)
	}

	if err := s.passwordHasher.Compare(
		req.Password,
		currentUser.PasswordHash,
	); err != nil {
		return nil, ErrInvalidCredentials
	}

	accessToken, err := s.tokenManager.CreateAccessToken(currentUser.ID)
	if err != nil {
		return nil, fmt.Errorf("create access token: %w", err)
	}

	rawRefreshToken, refreshTokenHash, err := generateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("generate refresh token: %w", err)
	}

	refreshToken := &refresh_token.RefreshToken{
		UserID:    currentUser.ID,
		TokenHash: refreshTokenHash,
		ExpiresAt: time.Now().Add(s.refreshTokenDuration),
	}

	if _, err := s.refreshTokenRepo.Create(ctx, refreshToken); err != nil {
		return nil, fmt.Errorf("create refresh token: %w", err)
	}

	return &AuthResponse{
		UserID:       currentUser.ID,
		AccessToken:  accessToken,
		RefreshToken: rawRefreshToken,
	}, nil
}

func (s *Service) Refresh(ctx context.Context, rawRefreshToken string) (*AuthResponse, error) {
	if strings.TrimSpace(rawRefreshToken) == "" {
		return nil, ErrInvalidRefreshToken
	}

	tokenHash := hashToken(rawRefreshToken)

	currentToken, err := s.refreshTokenRepo.GetByHash(ctx, tokenHash)
	if err != nil {
		if errors.Is(err, refresh_token.ErrNotFound) {
			return nil, ErrInvalidRefreshToken
		}

		return nil, fmt.Errorf("get refresh token: %w", err)
	}

	if currentToken.RevokedAt != nil {
		return nil, ErrRefreshTokenRevoked
	}

	if !currentToken.ExpiresAt.After(time.Now()) {
		return nil, ErrRefreshTokenExpired
	}

	if err := s.refreshTokenRepo.Revoke(ctx, currentToken.ID); err != nil {
		return nil, fmt.Errorf("revoke old refresh token: %w", err)
	}

	accessToken, err := s.tokenManager.CreateAccessToken(currentToken.UserID)
	if err != nil {
		return nil, fmt.Errorf("create access token: %w", err)
	}

	rawRefreshToken, newRefreshTokenHash, err := generateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("generate refresh token: %w", err)
	}

	newRefreshToken := &refresh_token.RefreshToken{
		UserID:    currentToken.UserID,
		TokenHash: newRefreshTokenHash,
		ExpiresAt: time.Now().Add(s.refreshTokenDuration),
	}

	if _, err := s.refreshTokenRepo.Create(ctx, newRefreshToken); err != nil {
		return nil, fmt.Errorf("create refresh token: %w", err)
	}

	return &AuthResponse{
		UserID:       currentToken.UserID,
		AccessToken:  accessToken,
		RefreshToken: rawRefreshToken,
	}, nil
}

func generateRefreshToken() (string, string, error) {
	rawToken, err := generateRandomToken()
	if err != nil {
		return "", "", err
	}

	return rawToken, hashToken(rawToken), nil
}

func generateRandomToken() (string, error) {
	const tokenSize = 32

	raw := make([]byte, tokenSize)

	if _, err := rand.Read(raw); err != nil {
		return "", fmt.Errorf("generate random token: %w", err)
	}

	return base64.RawURLEncoding.EncodeToString(raw), nil
}

func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))

	return hex.EncodeToString(hash[:])
}

func (s *Service) Logout(ctx context.Context, rawRefreshToken string) error {
	return s.Revoke(ctx, rawRefreshToken)
}

func (s *Service) Revoke(ctx context.Context, rawRefreshToken string) error {
	if strings.TrimSpace(rawRefreshToken) == "" {
		return nil
	}

	tokenHash := hashToken(rawRefreshToken)

	token, err := s.refreshTokenRepo.GetByHash(
		ctx,
		tokenHash,
	)
	if err != nil {
		if errors.Is(err, refresh_token.ErrNotFound) {
			return nil
		}

		return fmt.Errorf("get refresh token: %w", err)
	}

	if token.RevokedAt != nil {
		return nil
	}

	if err := s.refreshTokenRepo.Revoke(
		ctx,
		token.ID,
	); err != nil {
		return fmt.Errorf("revoke refresh token: %w", err)
	}

	return nil
}
