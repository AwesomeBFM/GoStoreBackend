package database

import (
	"context"
	"github.com/awesomebfm/go-store-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var products *mongo.Collection

func Init(uri string, database string) error {
	var err error = nil
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	products = client.Database(database).Collection("products")
	return nil
}

func GetProductByID(id primitive.ObjectID) (model.Product, error) {
	var product model.Product

	filter := bson.M{"_id": id}
	err := products.FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func GetOrderByID(id primitive.ObjectID) (model.Order, error) {
	var order model.Order

	filter := bson.M{"_id": id}
	err := products.FindOne(context.TODO(), filter).Decode(&order)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

func Close() error {
	err := client.Disconnect(context.TODO())
	return err
}
