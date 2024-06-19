package message

import "wb/internal/model"

type Delivery struct {
	OrderUID string `json:"order_uid"  validate:"required"`
	Name     string `json:"name"  validate:"required"`
	Phone    string `json:"phone"  validate:"required"`
	Zip      int    `json:"zip"  validate:"required"`
	City     string `json:"city"  validate:"required"`
	Address  string `json:"address"  validate:"required"`
	Region   string `json:"region"  validate:"required"`
	Email    string `json:"email"  validate:"required"`
}

func (d *Delivery) toModel() model.Delivery {
	return model.Delivery{
		OrderUID: d.OrderUID,
		Name:     d.Name,
		Phone:    d.Phone,
		Zip:      d.Zip,
		City:     d.City,
		Address:  d.Address,
		Region:   d.Region,
		Email:    d.Email,
	}
}
