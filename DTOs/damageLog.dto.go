package dtos

type CreateDamageLogDTO struct {
	Quantity  int `json:"Quantity" validate:"required,gte=1"`
	Reason    int `json:"Reason" validate:"required,gte=1,is_valid_damage_reason"`
	ProductId int `json:"ProductId" validate:"required,gte=1"`
}
