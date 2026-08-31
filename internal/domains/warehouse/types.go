package warehouse

import (
	"github.com/google/uuid"
)

type Warehouse struct {
	ID          uuid.UUID
	Code        string
	Name        string
	Description *string
}

type CreateReq struct {
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type UpdateReq struct {
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type WarehouseResponse struct {
	ID          uuid.UUID `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}
