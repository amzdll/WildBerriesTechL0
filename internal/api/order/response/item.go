package response

import "wb/internal/model"

type item struct {
	ChrtID      int     `json:"chrt_id"`
	TrackNumber string  `json:"track_number"`
	Price       float64 `json:"price"`
	RID         string  `json:"rid"`
	Name        string  `json:"name"`
	Sale        int     `json:"sale"`
	Size        string  `json:"size"`
	TotalPrice  float64 `json:"total_price"`
	NmID        int     `json:"nm_id"`
	Brand       string  `json:"brand"`
	Status      int     `json:"status"`
}

func (i *item) fromModel(m model.Item) {
	i.ChrtID = m.ChrtID
	i.TrackNumber = m.TrackNumber
	i.Price = m.Price
	i.RID = m.RID
	i.Name = m.Name
	i.Sale = m.Sale
	i.Size = m.Size
	i.TotalPrice = m.TotalPrice
	i.NmID = m.NmID
	i.Brand = m.Brand
	i.Status = m.Status
}
