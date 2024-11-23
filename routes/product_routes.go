package routes

import (
	"inventory-management/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	products := router.Group("/products")
	{
		products.GET("/", controllers.GetProducts)          // Get all products
		products.POST("/", controllers.CreateProduct)       // Create a new product
		products.PUT("/:id", controllers.UpdateProduct)     // Update product by ID
		products.DELETE("/:id", controllers.DeleteProduct)  // Delete product by ID

		// Untuk upload dan download gambar produk
		products.POST("/:id/image", controllers.UploadProductImage)   // Upload product image
		products.GET("/:id/image", controllers.DownloadProductImage)  // Download product image by ID
	}
}
