package main

import (
	"github.com/gin-gonic/gin"

	routes "service/routes"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})

	//localhost:8080/getOneBook?id=
	router.GET("/getOneBook", routes.GetBookByID)

	//localhost:8080/getAllBooks
	router.GET("/getAllBooks", routes.GetBooks)

	router.POST("/createBook", routes.CreateBook)
	// router.POST("/create_books",createBooks)

	// router.DELETE("/delete", DeleteBook)

	router.Run()
}
