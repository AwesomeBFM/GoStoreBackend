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
var orders *mongo.Collection

func Init(uri string, database string) error {
	var err error = nil
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	products = client.Database(database).Collection("products")
	orders = client.Database(database).Collection("orders")
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

func GetProductByPriceID(priceID string) (model.Product, error) {
	var product model.Product

	filter := bson.M{"stripe_id": priceID}
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

func CreateOrder(orderDto model.CreateOrderDto) error {
	// Create a new MongoDB document for the order
	order := bson.M{
		"customer_id": orderDto.CustomerID,
		"total":       orderDto.Total,
		"items":       orderDto.Items,
	}

	// Insert the order document into the MongoDB collection
	_, err := orders.InsertOne(context.TODO(), order)
	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	err := client.Disconnect(context.TODO())
	return err
}
