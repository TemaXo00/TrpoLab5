package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"electronic_shop/models"
)

func GetPositions(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var positions []models.Position
		if err := db.Find(&positions).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, positions)
	}
}

func GetPosition(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var position models.Position
		if err := db.First(&position, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Position not found"})
			return
		}
		c.JSON(http.StatusOK, position)
	}
}

func CreatePosition(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var position models.Position
		if err := c.ShouldBindJSON(&position); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&position).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, position)
	}
}

func UpdatePosition(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var position models.Position
		if err := db.First(&position, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Position not found"})
			return
		}
		if err := c.ShouldBindJSON(&position); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&position).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, position)
	}
}

func DeletePosition(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Position{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Position deleted"})
	}
}