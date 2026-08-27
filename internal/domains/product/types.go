package product

import (
	"github.com/google/uuid"
)

type Product struct {
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

type ProductResponse struct {
	ID   uuid.UUID `json:"id"`
	Code string    `json:"code"`
	Name string    `json:"name"`
}
