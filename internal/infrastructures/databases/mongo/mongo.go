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
	conn   *mongo.Client
}

func (mg *Database) Connect() (interface{}, error) {
	clientOptions := options.Client().ApplyURI(mg.Config.URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	mg.conn = client
	return client, nil
}

func (db *Database) Close() error {
	if db.conn != nil {
		return db.conn.Disconnect(context.Background())
	}
	return nil
}
