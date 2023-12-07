package routes

import (
	"context"
	"net/http"
	db_helper "service/db_helper"
	model "service/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteBook(c *gin.Context) {

	var ctx = context.TODO()
	client := db_helper.ConnectDB()
	database := db_helper.GetDatabase(client, "Book_Store")
	collection := db_helper.GetCollection(database, "tbl_Book")

	var book model.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.Id = primitive.NewObjectID()

	result, err := collection.InsertOne(ctx, book)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inserted document with _id: " + result.InsertedID.(primitive.ObjectID).Hex()})

}
