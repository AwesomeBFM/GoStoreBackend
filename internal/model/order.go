package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID         primitive.ObjectID  `bson:"_id"`
	Timestamp  primitive.Timestamp `bson:"timestamp"`
	CustomerID primitive.ObjectID  `bson:"customer_id"`
	Total      float64             `bson:"total"`
	Items      []OrderItem         `bson:"items"`
}

type OrderItem struct {
	ItemID   primitive.ObjectID `bson:"item_id"`
	Quantity int32              `bson:"quantity"`
}
