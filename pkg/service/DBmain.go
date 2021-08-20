package service

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"restapi/pkg/config"
)

var (
	ConnectionString string
	DatabaseName     string
)

var clientGlobal *mongo.Client

func Connect() error {
	ConnectionString = config.ResultConfig.DatabaseConfig.ConnectionString
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
