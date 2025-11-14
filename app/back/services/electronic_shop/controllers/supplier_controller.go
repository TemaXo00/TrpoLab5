package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"electronic_shop/models"
)

func GetSuppliers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var suppliers []models.Supplier
		if err := db.Find(&suppliers).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, suppliers)
	}
}

func GetSupplier(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var supplier models.Supplier
		if err := db.First(&supplier, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
			return
		}
		c.JSON(http.StatusOK, supplier)
	}
}

func CreateSupplier(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var supplier models.Supplier
		if err := c.ShouldBindJSON(&supplier); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&supplier).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, supplier)
	}
}

func UpdateSupplier(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var supplier models.Supplier
		if err := db.First(&supplier, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
			return
		}
		if err := c.ShouldBindJSON(&supplier); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&supplier).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, supplier)
	}
}

func DeleteSupplier(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Supplier{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Supplier deleted"})
	}
}