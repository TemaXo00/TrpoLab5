package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"electronic_shop/models"
)

func GetSales(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var sales []models.Sale
		if err := db.Preload("Store").Preload("Customer").Preload("Employee").Preload("SaleItems").Preload("SaleItems.Product").Find(&sales).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, sales)
	}
}

func GetSale(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var sale models.Sale
		if err := db.Preload("Store").Preload("Customer").Preload("Employee").Preload("SaleItems").Preload("SaleItems.Product").First(&sale, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Sale not found"})
			return
		}
		c.JSON(http.StatusOK, sale)
	}
}

func CreateSale(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var sale models.Sale
		if err := c.ShouldBindJSON(&sale); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&sale).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		db.Preload("Store").Preload("Customer").Preload("Employee").Preload("SaleItems").Preload("SaleItems.Product").First(&sale, sale.ID)
		c.JSON(http.StatusCreated, sale)
	}
}

func UpdateSale(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var sale models.Sale
		if err := db.First(&sale, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Sale not found"})
			return
		}
		if err := c.ShouldBindJSON(&sale); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&sale).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		db.Preload("Store").Preload("Customer").Preload("Employee").Preload("SaleItems").Preload("SaleItems.Product").First(&sale, sale.ID)
		c.JSON(http.StatusOK, sale)
	}
}

func PatchSale(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var sale models.Sale
        if err := db.First(&sale, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Sale not found"})
            return
        }

        var updateData map[string]interface{}
        if err := c.ShouldBindJSON(&updateData); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := db.Model(&sale).Updates(updateData).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        
        db.Preload("Store").Preload("Customer").Preload("Employee").Preload("SaleItems").Preload("SaleItems.Product").First(&sale, sale.ID)
        c.JSON(http.StatusOK, sale)
    }
}

func DeleteSale(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Sale{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Sale deleted"})
	}
}