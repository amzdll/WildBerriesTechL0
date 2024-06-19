package message

import "wb/internal/model"

type Item struct {
	ChrtID      int     `json:"chrt_id"  validate:"required"`
	TrackNumber string  `json:"track_number"  validate:"required"`
	Price       float64 `json:"price"  validate:"required"`
	RID         string  `json:"rid"  validate:"required"`
	Name        string  `json:"name"  validate:"required"`
	Sale        int     `json:"sale"  validate:"required"`
	Size        string  `json:"size"  validate:"required"`
	TotalPrice  float64 `json:"total_price"  validate:"required"`
	NmID        int     `json:"nm_id"  validate:"required"`
	Brand       string  `json:"brand"  validate:"required"`
	Status      int     `json:"status"  validate:"required"`
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
