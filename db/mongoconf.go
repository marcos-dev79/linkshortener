package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connects to mongo database and returns a client
func Mongo_connect() mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongols:27017"))
	if err != nil {
		panic(err)
	}

	return *client
}
