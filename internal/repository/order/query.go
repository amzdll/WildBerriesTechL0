package order

// Create queries
const (
	createOrderQuery = ` 
		insert into "order" 
		(
			order_uid,
			track_number,
			entry,
			locale,
			internal_signature,
			customer_id,
			delivery_service,
			shardkey,
			sm_id,
			date_created,
			oof_shard
		) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	createDeliveryQuery = `
		insert into "delivery" 
    (
			order_uid,
			name,
			phone,
			zip,
			city,
			address,
			region,
			email
		)
		values($1, $2, $3, $4, $5, $6, $7, $8)`

	createPaymentQuery = `
	insert into "payment"
	(
	 	transaction,
		request_id,
		currency,
		provider,
		amount,
		payment_dt,
		bank,
		delivery_cost,
		goods_total,
		custom_fee
	)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	createItemQuery = `
	insert into "items"
	(	
		chrt_id,
		track_number,
		price,
		rid,
		name,
		sale,
		size,
		total_price,
		nm_id,
		brand,
		status
	)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
)

// Get queries
const (
	getOrderQuery    = `select * from "order" where order_uid = $1`
	getDeliveryQuery = `select * from "delivery" where order_uid = $1`
	getPaymentQuery  = `select * from "payment" where transaction = $1`
	getItemQuery     = `select * from "items" where track_number = $1`
)
