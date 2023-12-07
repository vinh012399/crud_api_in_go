package db_helper

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb+srv://duyvinh012399:01869768563@cluster0.wptzgjf.mongodb.net/?retryWrites=true&w=majority"

var Client *mongo.Client

func ConnectDB() error {
	bsonOpts := &options.BSONOptions{
		OmitZeroStruct: true,
	}
	clientOptions := options.Client().ApplyURI(uri).SetBSONOptions(bsonOpts)

	var err error

	Client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return err
	}

	// _, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel()

	return nil
}
