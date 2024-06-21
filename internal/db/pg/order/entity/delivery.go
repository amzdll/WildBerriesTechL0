package entity

import "wb/internal/model"

type Delivery struct {
	OrderUID string `db:"order_uid"`
	Name     string `db:"name"`
	Phone    string `db:"phone"`
	Zip      int    `db:"zip"`
	City     string `db:"city"`
	Address  string `db:"address"`
	Region   string `db:"region"`
	Email    string `db:"email"`
}

func (e *Delivery) toModel() model.Delivery {
	return model.Delivery{
		OrderUID: e.OrderUID,
		Name:     e.Name,
		Phone:    e.Phone,
		Zip:      e.Zip,
		City:     e.City,
		Address:  e.Address,
		Region:   e.Region,
		Email:    e.Email,
	}
}
