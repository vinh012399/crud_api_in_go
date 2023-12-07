package routes

import (
	"context"
	"net/http"
	db_helper "service/db_helper"
	"service/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateBook(c *gin.Context) {

	var ctx = context.TODO()
	client := db_helper.ConnectDB()
	database := db_helper.GetDatabase(client, "Book_Store")
	collection := db_helper.GetCollection(database, "tbl_Book")

	existBook := new(model.Book)

	// if err:=c.BindJSON(&existBook); err != nil{
	// 	c.JSON(http.StatusBadRequest, gin.H{"message":err.Error()})
	// 	return
	// }

	idParam := c.Query("id")

	// id, err := primitive.ObjectIDFromHex(string(existBook.Id))
	id, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := c.BindJSON(&existBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	filter := bson.M{"_id": id}

	edited := bson.M{"Author": existBook.Author, "Title": existBook.Title, "Price": existBook.Price, "Is_Deleted": existBook.Is_Deleted, "Updated_By": existBook.Updated_By, "Created_By": existBook.Created_By}

	result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": edited})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	res := map[string]interface{}{"data": result}

	if result.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesn't exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully", "Data": res})

}
