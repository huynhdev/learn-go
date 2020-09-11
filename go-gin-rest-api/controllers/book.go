package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhdev/go-gin-rest-api/models"
)

func CreateBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)

	models.DB.Create(&book)

	c.JSON(http.StatusOK, book)
}

func ShowBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func ListBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

func UpdateBooks(c *gin.Context) {
	var book models.Book

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
	}
	c.BindJSON(&book)

	models.DB.Save(&book)

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	var book models.Book

	id := c.Param("id")

	models.DB.Where("id = ?", id).Delete(&book)
	c.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
}
