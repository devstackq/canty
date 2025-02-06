package mongo

import (
	"context"

	"canty/internal/core/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoAdvertisementRepository struct {
	Collection *mongo.Collection
}

func NewMongoAdvertisementRepository(client *mongo.Client, dbName, collectionName string) *MongoAdvertisementRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &MongoAdvertisementRepository{
		Collection: collection,
	}
}

func (repo *MongoAdvertisementRepository) Create(ad *entities.Advertisement) error {
	_, err := repo.Collection.InsertOne(context.Background(), ad)
	return err
}

func (repo *MongoAdvertisementRepository) Read(id string) (*entities.Advertisement, error) {
	filter := bson.M{"_id": id}
	ad := &entities.Advertisement{}
	err := repo.Collection.FindOne(context.Background(), filter).Decode(ad)
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func (repo *MongoAdvertisementRepository) Update(ad *entities.Advertisement) error {
	filter := bson.M{"_id": ad.ID}
	update := bson.M{"$set": ad}
	_, err := repo.Collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (repo *MongoAdvertisementRepository) Delete(id string) error {
	filter := bson.M{"_id": id}
	_, err := repo.Collection.DeleteOne(context.Background(), filter)
	return err
}
