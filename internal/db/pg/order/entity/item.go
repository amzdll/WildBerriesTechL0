package entity

import "wb/internal/model"

type Item struct {
	ChrtID      int     `db:"chrt_id"`
	TrackNumber string  `db:"track_number"`
	Price       float64 `db:"price"`
	RID         string  `db:"rid"`
	Name        string  `db:"name"`
	Sale        int     `db:"sale"`
	Size        string  `db:"size"`
	TotalPrice  float64 `db:"total_price"`
	NmID        int     `db:"nm_id"`
	Brand       string  `db:"brand"`
	Status      int     `db:"status"`
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
