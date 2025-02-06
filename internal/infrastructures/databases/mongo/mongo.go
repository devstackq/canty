package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	URI    string
	DBName string
}

type Database struct {
	Config Config
}

func (mg *Database) Connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mg.Config.URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
