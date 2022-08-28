package controllers

import (
	"net/http"
	"testAPI/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FindProducts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var product []models.Product
	db.Find(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}
