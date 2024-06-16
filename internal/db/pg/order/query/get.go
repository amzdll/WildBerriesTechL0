package query

const (
	GetOrders       = `select * from "order"`
	GetOrderById    = `select * from "order" where order_uid = $1`
	GetDeliveryById = `select * from "delivery" where order_uid = $1`
	GetPaymentById  = `select * from "payment" where transaction = $1`
	GetItemById     = `select * from "items" where track_number = $1`
)
