package models

type Order struct {
	Id     int64       `bson:"id"`
	UserId int64       `bson:"user_id"`
	Items  []OrderItem `bson:"items"`
	Total  float32     `bson:"total"`
}

type OrderItem struct {
	ItemID    int64   `bson:"item_id"`
	Quantity  int     `bson:"quantity"`
	UnitPrice float32 `bson:"unit_price"`
}
