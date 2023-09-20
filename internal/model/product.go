package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID        primitive.ObjectID `bson:"_id"`
	Inventory int                `bson:"inventory"`
	StripeID  string             `bson:"stripe_id"`
}
