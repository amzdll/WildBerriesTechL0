package response

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

func (d *delivery) fromModel(m model.Delivery) {
	d.OrderUID = m.OrderUID
	d.Name = m.Name
	d.Phone = m.Phone
	d.Zip = m.Zip
	d.City = m.City
	d.Address = m.Address
	d.Region = m.Region
	d.Email = m.Email
}
