package service

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ConnectionString = "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"

var clientGlobal *mongo.Client

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(ConnectionString))
	if err != nil {
		return err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return err
	}

	clientGlobal = client
	return nil
}

func CloseConnect() error {

	return nil
}
