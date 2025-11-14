package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Welcome to Electronic Shop API",
        "version": "1.0",
        "endpoints": gin.H{
            "categories":  "/categories",
            "suppliers":   "/suppliers", 
            "products":    "/products",
            "positions":   "/positions",
            "stores":      "/stores",
            "employees":   "/employees",
            "clients":     "/clients",
            "sales":       "/sales",
            "sale_items":  "/sale-items",
            "reviews":     "/reviews",
        },
    })
}