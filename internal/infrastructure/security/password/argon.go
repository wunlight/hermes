package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	version     = 19
	memory      = 64 * 1024
	iterations  = 3
	parallelism = 4
	keyLength   = 32
	saltLength  = 16
)

type Argon struct{}

func NewArgon() *Argon {
	return &Argon{}
}

func (a *Argon) Hash(password string) (string, error) {
	salt := make([]byte, saltLength)

	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("generate password salt: %w", err)
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		iterations,
		memory,
		parallelism,
		keyLength,
	)

	encoded := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		version,
		memory,
		iterations,
		parallelism,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)

	return encoded, nil
}

func (a *Argon) Compare(password, encodedHash string) error {
	params, salt, expectedHash, err := decode(encodedHash)
	if err != nil {
		return fmt.Errorf("decode password hash: %w", err)
	}

	actualHash := argon2.IDKey(
		[]byte(password),
		salt,
		params.iterations,
		params.memory,
		params.parallelism,
		params.keyLength,
	)

	if subtle.ConstantTimeCompare(actualHash, expectedHash) != 1 {
		return errors.New("invalid password")
	}

	return nil
}

type parameters struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	keyLength   uint32
}

func decode(encodedHash string) (parameters, []byte, []byte, error) {
	var params parameters

	parts := strings.Split(encodedHash, "$")

	if len(parts) != 6 {
		return params, nil, nil, errors.New("invalid argon2id hash format")
	}

	if parts[1] != "argon2id" {
		return params, nil, nil, errors.New("unsupported password hash algorithm")
	}

	if parts[2] != "v=19" {
		return params, nil, nil, errors.New("unsupported argon2 version")
	}

	_, err := fmt.Sscanf(
		parts[3],
		"m=%d,t=%d,p=%d",
		&params.memory,
		&params.iterations,
		&params.parallelism,
	)
	if err != nil {
		return params, nil, nil, errors.New("invalid argon2 parameters")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return params, nil, nil, errors.New("invalid salt encoding")
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return params, nil, nil, errors.New("invalid hash encoding")
	}

	params.keyLength = uint32(len(expectedHash))

	if params.memory == 0 ||
		params.iterations == 0 ||
		params.parallelism == 0 ||
		params.keyLength == 0 {
		return params, nil, nil, errors.New("invalid argon2 parameters")
	}

	return params, salt, expectedHash, nil
}
