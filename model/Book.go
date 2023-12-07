package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"title"`
	Author     string             `bson:"author"`
	Price      float64            `bson:"price"`
	Is_Deleted bool               `bson:"is_deleted"`
	Created_By string             `bson:"created_by"`
	Updated_By string             `bson:"updated_by"`
}
