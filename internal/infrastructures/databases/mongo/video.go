package mongo

import (
	"context"

	"canty/internal/core/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoVideoRepository struct {
	Collection *mongo.Collection
}

func NewMongoVideoRepository(client *mongo.Client, dbName, collectionName string) *MongoVideoRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &MongoVideoRepository{
		Collection: collection,
	}
}

func (repo *MongoVideoRepository) Create(video *entities.Video) error {
	_, err := repo.Collection.InsertOne(context.Background(), video)
	return err
}

func (repo *MongoVideoRepository) Read(id string) (*entities.Video, error) {
	filter := bson.M{"id": id}
	video := &entities.Video{}
	err := repo.Collection.FindOne(context.Background(), filter).Decode(video)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (repo *MongoVideoRepository) Update(video *entities.Video) error {
	filter := bson.M{"id": video.ID}
	update := bson.M{"$set": video}
	_, err := repo.Collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (repo *MongoVideoRepository) Delete(id string) error {
	filter := bson.M{"id": id}
	_, err := repo.Collection.DeleteOne(context.Background(), filter)
	return err
}
