package product

import (
	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID
	SKU         string
	Name        string
	CategoryID  uuid.UUID
	BrandID     uuid.UUID
	UnitID      uuid.UUID
	MinStock    int32
	Weight      float32
	Length      float32
	Width       float32
	Description *string
	Status      string
}

type CreateReq struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	CategoryID  string  `json:"category_id"`
	BrandID     string  `json:"brand_id"`
	UnitID      string  `json:"unit_id"`
	MinStock    int32   `json:"min_stock"`
	Weight      float32 `json:"weight"`
	Length      float32 `json:"length"`
	Width       float32 `json:"width"`
	Description *string `json:"description"`
}

type UpdateReq struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	CategoryID  string  `json:"category_id"`
	BrandID     string  `json:"brand_id"`
	UnitID      string  `json:"unit_id"`
	MinStock    int32   `json:"min_stock"`
	Weight      float32 `json:"weight"`
	Length      float32 `json:"length"`
	Width       float32 `json:"width"`
	Description *string `json:"description"`
}

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	SKU         string    `json:"sku"`
	Name        string    `json:"name"`
	CategoryID  uuid.UUID `json:"category_id"`
	BrandID     uuid.UUID `json:"brand_id"`
	UnitID      uuid.UUID `json:"unit_id"`
	MinStock    int32     `json:"min_stock"`
	Weight      float32   `json:"weight"`
	Length      float32   `json:"length"`
	Width       float32   `json:"width"`
	Description *string   `json:"description"`
}
