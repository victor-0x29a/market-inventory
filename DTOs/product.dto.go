package dtos

type CreateProductDTO struct {
	Title             string  `json:"title" validate:"required,min=1,max=64"`
	Description       *string `json:"description" validate:"omitempty,max=256"`
	Price             int64   `json:"price" validate:"required,gt=0"`
	InventoryQuantity int64   `json:"quantity" validate:"required,gt=0"`
}

type FindOneProductDTO struct {
	ID int `validate:"required,gt=0"`
}
