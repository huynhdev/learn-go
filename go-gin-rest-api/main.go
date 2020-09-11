package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhdev/go-gin-rest-api/controllers"
	"github.com/huynhdev/go-gin-rest-api/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Hello World",
		})
	})

	r.GET("/books", controllers.ListBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.ShowBook)
	r.PUT("/books/:id", controllers.UpdateBooks)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
