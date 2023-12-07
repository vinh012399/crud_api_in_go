package routes

import (
	"context"
	"net/http"
	db_helper "service/db_helper"
	model "service/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBooks(c *gin.Context) {

	var ctx = context.TODO()
	client := db_helper.ConnectDB()
	database := db_helper.GetDatabase(client, "Book_Store")
	collection := db_helper.GetCollection(database, "tbl_Book")

	//find books
	cursor, err := collection.Find(ctx, bson.D{{}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//map results
	var books []model.Book
	if err = cursor.All(ctx, &books); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//return books
	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {

	var ctx = context.TODO()
	client := db_helper.ConnectDB()
	database := db_helper.GetDatabase(client, "Book_Store")
	collection := db_helper.GetCollection(database, "tbl_Book")

	idParam := c.Query("id")

	newId, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book bson.M

	// Find book by ID
	err = collection.FindOne(ctx, bson.D{{Key: "_id", Value: newId}}).Decode(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return book
	c.JSON(http.StatusOK, book)
}
