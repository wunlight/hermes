package product

import "errors"

var (
	ErrNotFound          = errors.New("product not found")
	ErrSKUAlreadyExists  = errors.New("product sku already exists")
	ErrInvalidID         = errors.New("invalid product id")
	ErrSKURequired       = errors.New("product sku required")
	ErrNameRequired      = errors.New("product name required")
	ErrInvalidReq        = errors.New("invalid product request")
	ErrInvalidCategoryID = errors.New("invalid category id")
	ErrInvalidBrandID    = errors.New("invalid brand id")
	ErrInvalidUnitID     = errors.New("invalid unit id")
	ErrInvalidMinStock   = errors.New("invalid min stock")
	ErrInvalidWeight     = errors.New("invalid weight")
	ErrInvalidLength     = errors.New("invalid length")
	ErrInvalidWidth      = errors.New("invalid width")
)
