package database

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var orders *mongo.Collection

func Init(uri string, database string) error {
	var err error = nil
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	orders = client.Database(database).Collection("orders")

	return err
}

func GetOrderByID(id int64) []byte {
	var result bson.M

	err := orders.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&result)
	if err != nil {
		return nil
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return nil
	}

	return jsonData
}

func Close() error {
	err := client.Disconnect(context.TODO())
	return err
}
