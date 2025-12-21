package dtos

type CreateDamageLogDTO struct {
	Quantity  int `json:"Quantity" validate:"required,gt=0"`
	Reason    int `json:"Reason" validate:"required,gte=1"`
	ProductId int `json:"ProductId" validate:"required,gte=0"`
}
