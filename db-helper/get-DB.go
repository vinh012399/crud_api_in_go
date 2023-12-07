package db_helper

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDatabase(client *mongo.Client, databaseName string) *mongo.Database {
	database := client.Database(databaseName)
	return database
}
