package db_helper

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb+srv://duyvinh012399:01869768563@cluster0.wptzgjf.mongodb.net/?retryWrites=true&w=majority"

func ConnectDB() *mongo.Client {

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println(err.Error())
	}

	// _, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel()

	return client
}
