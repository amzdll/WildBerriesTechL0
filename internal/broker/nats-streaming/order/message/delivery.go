package message

import "wb/internal/model"

type delivery struct {
	OrderUID string `json:"order_uid"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Zip      int    `json:"zip"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Region   string `json:"region"`
	Email    string `json:"email"`
}

func (d *delivery) toModel() model.Delivery {
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
