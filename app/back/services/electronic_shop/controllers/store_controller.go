package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"electronic_shop/models"
)

func GetStores(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var stores []models.Store
		if err := db.Find(&stores).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, stores)
	}
}

func GetStore(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var store models.Store
		if err := db.First(&store, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
			return
		}
		c.JSON(http.StatusOK, store)
	}
}

func CreateStore(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var store models.Store
		if err := c.ShouldBindJSON(&store); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&store).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, store)
	}
}

func UpdateStore(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var store models.Store
		if err := db.First(&store, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
			return
		}
		if err := c.ShouldBindJSON(&store); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&store).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, store)
	}
}

func PatchStore(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var store models.Store
        if err := db.First(&store, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
            return
        }

        var updateData map[string]interface{}
        if err := c.ShouldBindJSON(&updateData); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := db.Model(&store).Updates(updateData).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        
        c.JSON(http.StatusOK, store)
    }
}

func DeleteStore(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Store{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Store deleted"})
	}
}