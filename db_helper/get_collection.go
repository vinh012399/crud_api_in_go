package db_helper

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(database *mongo.Database, collectionName string) *mongo.Collection {
	collection := database.Collection(collectionName)
	return collection
}
