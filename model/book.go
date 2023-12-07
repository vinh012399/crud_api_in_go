package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id"`
	Title      string             `bson:"title" json:"title"`
	Author     string             `bson:"author" json:"author"`
	Price      float64            `bson:"price" json:"price"`
	Is_Deleted bool               `bson:"is_deleted" json:"is_deleted"`
	Created_By string             `bson:"created_by" json:"created_by"`
	Updated_By string             `bson:"updated_by" json:"updated_by"`
}

func NewBook() Book {

	instance := Book{}
	instance.Id = primitive.NewObjectID()

	return instance
}
