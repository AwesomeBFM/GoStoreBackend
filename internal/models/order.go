package models

type Order struct {
	ID     int64       `bson:"id"`
	UserID int64       `bson:"user_id"`
	Items  []OrderItem `bson:"items"`
	Total  float32     `bson:"total"`
}

type OrderItem struct {
	ItemID    int64   `bson:"item_id"`
	Quantity  int     `bson:"quantity"`
	UnitPrice float32 `bson:"unit_price"`
}
