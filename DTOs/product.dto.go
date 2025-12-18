package dtos

type CreateProductDTO struct {
	Title             string  `json:"Title" validate:"required,min=1,max=64"`
	Description       *string `json:"Description" validate:"omitempty,max=256"`
	Price             int64   `json:"Price" validate:"required,gt=0"`
	InventoryQuantity int64   `json:"Quantity" validate:"required,gt=0"`
}

type FindOneProductDTO struct {
	ID int `validate:"required,gt=0"`
}
