package entity

import "wb/internal/model"

type Item struct {
	ChrtID      int
	TrackNumber string
	Price       float64
	RID         string
	Name        string
	Sale        int
	Size        string
	TotalPrice  float64
	NmID        int
	Brand       string
	Status      int
}

func (i *Item) toModel() model.Item {
	return model.Item{
		ChrtID:      i.ChrtID,
		TrackNumber: i.TrackNumber,
		Price:       i.Price,
		RID:         i.RID,
		Name:        i.Name,
		Sale:        i.Sale,
		Size:        i.Size,
		TotalPrice:  i.TotalPrice,
		NmID:        i.NmID,
		Brand:       i.Brand,
		Status:      i.Status,
	}
}
