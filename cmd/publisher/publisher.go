package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"log"
	"math/rand"
	"time"
)

type Order struct {
	OrderUID          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          string    `json:"shardkey"`
	SMID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

type Delivery struct {
	OrderUID string `json:"order_uid"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Zip      int    `json:"zip"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Region   string `json:"region"`
	Email    string `json:"email"`
}

type Payment struct {
	Transaction  string  `json:"transaction"`
	RequestId    string  `json:"request_id"`
	Currency     string  `json:"currency"`
	Provider     string  `json:"provider"`
	Amount       int     `json:"amount"`
	PaymentDt    int64   `json:"payment_dt"`
	Bank         string  `json:"bank"`
	DeliveryCost float64 `json:"delivery_cost"`
	GoodsTotal   int     `json:"goods_total"`
	CustomFee    float64 `json:"custom_fee"`
}

type Item struct {
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

func main() {
	rand.Seed(time.Now().UnixNano())

	sc, err := stan.Connect(
		"test-cluster",
		"publisher",
		stan.NatsURL(nats.DefaultURL),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	order := generateRandomOrder()
	fmt.Println(order.Items)
	mess, _ := json.Marshal(order)

	err = sc.Publish("orders", mess)
	if err != nil {
		log.Fatal(err)
	}
}

func generateRandomOrder() Order {
	order_uid := generateRandomString(20)
	track_number := generateRandomString(12)
	return Order{
		OrderUID:    order_uid,
		TrackNumber: track_number,
		Entry:       generateRandomString(4),
		Delivery: Delivery{
			OrderUID: order_uid,
			Name:     "Test Testov",
			Phone:    "+72000000000",
			Zip:      rand.Intn(10000),
			City:     "Kiryat Mozkin",
			Address:  "Ploshad Mira 15",
			Region:   "Kraiot",
			Email:    "test@gmail.com",
		},
		Payment: Payment{
			Transaction:  order_uid,
			RequestId:    generateRandomString(20),
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       rand.Intn(1000000),
			PaymentDt:    time.Now().Unix(),
			Bank:         "alpha",
			DeliveryCost: rand.Float64() * 2000,
			GoodsTotal:   rand.Intn(1000000),
			CustomFee:    rand.Float64() * 100,
		},
		Items: []Item{
			{
				ChrtID:      rand.Intn(1000000),
				TrackNumber: track_number,
				Price:       rand.Float64() * 1000,
				RID:         generateRandomString(20),
				Name:        "Mascaras",
				Sale:        rand.Intn(50),
				Size:        "S",
				TotalPrice:  rand.Float64() * 1000,
				NmID:        rand.Intn(1000000),
				Brand:       "Vivienne Sabo",
				Status:      rand.Intn(500),
			},
		},
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		ShardKey:          generateRandomString(1),
		SMID:              rand.Intn(100),
		DateCreated:       time.Now(),
		OofShard:          generateRandomString(1),
	}
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
