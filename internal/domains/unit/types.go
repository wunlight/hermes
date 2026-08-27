package unit

import (
	"github.com/google/uuid"
)

type Unit struct {
	ID   uuid.UUID
	Code string
	Name string
}

type CreateReq struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type UpdateReq struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type UnitResponse struct {
	ID   uuid.UUID `json:"id"`
	Code string    `json:"code"`
	Name string    `json:"name"`
}
