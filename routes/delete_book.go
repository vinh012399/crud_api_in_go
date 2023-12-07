package routes

import (
	"context"
	"net/http"
	db_helper "service/db_helper"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteBook(c *gin.Context) {

	var ctx = context.TODO()
	client := db_helper.ConnectDB()
	database := db_helper.GetDatabase(client, "Book_Store")
	collection := db_helper.GetCollection(database, "tbl_Book")

	id := c.Query("id")

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No data to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
