package models

type Item struct {
	ID        int64   `bson:"id"`
	Inventory int     `bson:"inventory"`
	Price     float32 `bson:"total"`
}
