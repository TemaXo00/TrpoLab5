package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"electronic_shop/models"
)

func GetEmployees(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var employees []models.Employee
		if err := db.Preload("Position").Preload("Store").Find(&employees).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, employees)
	}
}

func GetEmployee(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var employee models.Employee
		if err := db.Preload("Position").Preload("Store").First(&employee, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		c.JSON(http.StatusOK, employee)
	}
}

func CreateEmployee(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var employee models.Employee
		if err := c.ShouldBindJSON(&employee); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&employee).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		db.Preload("Position").Preload("Store").First(&employee, employee.ID)
		c.JSON(http.StatusCreated, employee)
	}
}

func UpdateEmployee(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var employee models.Employee
		if err := db.First(&employee, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		if err := c.ShouldBindJSON(&employee); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Save(&employee).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		db.Preload("Position").Preload("Store").First(&employee, employee.ID)
		c.JSON(http.StatusOK, employee)
	}
}

func PatchEmployee(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var employee models.Employee
        if err := db.First(&employee, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
            return
        }

        var updateData map[string]interface{}
        if err := c.ShouldBindJSON(&updateData); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := db.Model(&employee).Updates(updateData).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        
        db.Preload("Position").Preload("Store").First(&employee, employee.ID)
        c.JSON(http.StatusOK, employee)
    }
}

func DeleteEmployee(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Employee{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
	}
}