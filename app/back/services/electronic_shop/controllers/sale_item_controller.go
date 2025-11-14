package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"electronic_shop/models"
)

func GetSaleItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var saleItems []models.SaleItem
		if err := db.Preload("Sale").Preload("Product").Find(&saleItems).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, saleItems)
	}
}

func GetSaleItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var saleItem models.SaleItem
		if err := db.Preload("Sale").Preload("Product").First(&saleItem, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Sale item not found"})
			return
		}
		c.JSON(http.StatusOK, saleItem)
	}
}

func CreateSaleItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var saleItem models.SaleItem
		if err := c.ShouldBindJSON(&saleItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&saleItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		db.Preload("Sale").Preload("Product").First(&saleItem, saleItem.ID)
		c.JSON(http.StatusCreated, saleItem)
	}
}

func UpdateSaleItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var saleItem models.SaleItem
		if err := db.First(&saleItem, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Sale item not found"})
			return
		}
		if err := c.ShouldBindJSON(&saleItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&saleItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		db.Preload("Sale").Preload("Product").First(&saleItem, saleItem.ID)
		c.JSON(http.StatusOK, saleItem)
	}
}

func DeleteSaleItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.SaleItem{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Sale item deleted"})
	}
}