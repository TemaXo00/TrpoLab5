package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"electronic_shop/models"
)

func GetReviews(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reviews []models.Review
		if err := db.Preload("Customer").Preload("Product").Preload("Employee").Preload("Store").Find(&reviews).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, reviews)
	}
}

func GetReview(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var review models.Review
		if err := db.Preload("Customer").Preload("Product").Preload("Employee").Preload("Store").First(&review, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
			return
		}
		c.JSON(http.StatusOK, review)
	}
}

func CreateReview(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var review models.Review
		if err := c.ShouldBindJSON(&review); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&review).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		db.Preload("Customer").Preload("Product").Preload("Employee").Preload("Store").First(&review, review.ID)
		c.JSON(http.StatusCreated, review)
	}
}

func UpdateReview(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var review models.Review
		if err := db.First(&review, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
			return
		}
		if err := c.ShouldBindJSON(&review); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&review).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		db.Preload("Customer").Preload("Product").Preload("Employee").Preload("Store").First(&review, review.ID)
		c.JSON(http.StatusOK, review)
	}
}

func PatchReview(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var review models.Review
        if err := db.First(&review, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
            return
        }

        var updateData map[string]interface{}
        if err := c.ShouldBindJSON(&updateData); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := db.Model(&review).Updates(updateData).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        
        db.Preload("Customer").Preload("Product").Preload("Employee").Preload("Store").First(&review, review.ID)
        c.JSON(http.StatusOK, review)
    }
}

func DeleteReview(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Review{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Review deleted"})
	}
}