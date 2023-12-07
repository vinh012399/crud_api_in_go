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
	// client := db_helper.ConnectDB()
	database := db_helper.GetDatabase(db_helper.Client, "Book_Store")
	collection := db_helper.GetCollection(database, "tbl_Book")

	bookRequest := new(model.BookRequest)

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

	if err := c.BindJSON(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	existedBook := model.Book{
		Id: id,
	}

	if bookRequest.Author != nil {
		existedBook.Author = *bookRequest.Author
	}

	if bookRequest.Title != nil {
		existedBook.Title = *bookRequest.Title
	}

	if bookRequest.Price != nil {
		existedBook.Price = *bookRequest.Price
	}

	if bookRequest.Created_By != nil {
		existedBook.Created_By = *bookRequest.Created_By
	}

	if bookRequest.Updated_By != nil {
		existedBook.Updated_By = *bookRequest.Updated_By
	}

	filter := bson.M{"_id": id}

	result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": existedBook})

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
