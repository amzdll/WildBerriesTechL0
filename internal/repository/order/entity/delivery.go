package entity

import "wb/internal/model"

type Delivery struct {
	OrderUID string
	Name     string
	Phone    string
	Zip      int
	City     string
	Address  string
	Region   string
	Email    string
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
