package permission

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID          uuid.UUID
	Code        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateRequest struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Response struct {
	ID          uuid.UUID `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
}
