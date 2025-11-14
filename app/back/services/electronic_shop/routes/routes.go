package routes

import (
	"electronic_shop/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupCategoryRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/categories", controllers.GetCategories(db))
	router.GET("/categories/:id", controllers.GetCategory(db))
	router.POST("/categories", controllers.CreateCategory(db))
	router.PUT("/categories/:id", controllers.UpdateCategory(db))
	router.DELETE("/categories/:id", controllers.DeleteCategory(db))
}

func setupSupplierRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/suppliers", controllers.GetSuppliers(db))
	router.GET("/suppliers/:id", controllers.GetSupplier(db))
	router.POST("/suppliers", controllers.CreateSupplier(db))
	router.PUT("/suppliers/:id", controllers.UpdateSupplier(db))
	router.DELETE("/suppliers/:id", controllers.DeleteSupplier(db))
}

func setupProductRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/products", controllers.GetProducts(db))
	router.GET("/products/:id", controllers.GetProduct(db))
	router.POST("/products", controllers.CreateProduct(db))
	router.PUT("/products/:id", controllers.UpdateProduct(db))
	router.PATCH("/products/:id", controllers.PatchProduct(db))
	router.DELETE("/products/:id", controllers.DeleteProduct(db))
}

func setupPositionRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/positions", controllers.GetPositions(db))
	router.GET("/positions/:id", controllers.GetPosition(db))
	router.POST("/positions", controllers.CreatePosition(db))
	router.PUT("/positions/:id", controllers.UpdatePosition(db))
	router.DELETE("/positions/:id", controllers.DeletePosition(db))
}

func setupStoreRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/stores", controllers.GetStores(db))
	router.GET("/stores/:id", controllers.GetStore(db))
	router.POST("/stores", controllers.CreateStore(db))
	router.PUT("/stores/:id", controllers.UpdateStore(db))
	router.PATCH("/stores/:id", controllers.PatchStore(db))
	router.DELETE("/stores/:id", controllers.DeleteStore(db))
}

func setupEmployeeRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/employees", controllers.GetEmployees(db))
	router.GET("/employees/:id", controllers.GetEmployee(db))
	router.POST("/employees", controllers.CreateEmployee(db))
	router.PUT("/employees/:id", controllers.UpdateEmployee(db))
	router.PATCH("/employees/:id", controllers.PatchEmployee(db))
	router.DELETE("/employees/:id", controllers.DeleteEmployee(db))
}

func setupClientRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/clients", controllers.GetClients(db))
	router.GET("/clients/:id", controllers.GetClient(db))
	router.POST("/clients", controllers.CreateClient(db))
	router.PUT("/clients/:id", controllers.UpdateClient(db))
	router.PATCH("/clients/:id", controllers.PatchClient(db))
	router.DELETE("/clients/:id", controllers.DeleteClient(db))
}

func setupSaleRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/sales", controllers.GetSales(db))
	router.GET("/sales/:id", controllers.GetSale(db))
	router.POST("/sales", controllers.CreateSale(db))
	router.PUT("/sales/:id", controllers.UpdateSale(db))
	router.PATCH("/sales/:id", controllers.PatchSale(db))
	router.DELETE("/sales/:id", controllers.DeleteSale(db))
}

func setupSaleItemRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/sale-items", controllers.GetSaleItems(db))
	router.GET("/sale-items/:id", controllers.GetSaleItem(db))
	router.POST("/sale-items", controllers.CreateSaleItem(db))
	router.PUT("/sale-items/:id", controllers.UpdateSaleItem(db))
	router.DELETE("/sale-items/:id", controllers.DeleteSaleItem(db))
}

func setupReviewRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/reviews", controllers.GetReviews(db))
	router.GET("/reviews/:id", controllers.GetReview(db))
	router.POST("/reviews", controllers.CreateReview(db))
	router.PUT("/reviews/:id", controllers.UpdateReview(db))
	router.PATCH("/reviews/:id", controllers.PatchReview(db))
	router.DELETE("/reviews/:id", controllers.DeleteReview(db))
}

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/", controllers.Welcome)
	setupCategoryRoutes(router, db)
	setupSupplierRoutes(router, db)
	setupProductRoutes(router, db)
	setupPositionRoutes(router, db)
	setupStoreRoutes(router, db)
	setupEmployeeRoutes(router, db)
	setupClientRoutes(router, db)
	setupSaleRoutes(router, db)
	setupSaleItemRoutes(router, db)
	setupReviewRoutes(router, db)
}