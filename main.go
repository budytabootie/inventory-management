package main

import (
	"inventory-management/config"
	"inventory-management/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	// Untuk inisialisasi router
	router := gin.Default()

	// Register product ke router
	routes.ProductRoutes(router)
	// Register inventory ke router
    routes.InventoryRoutes(router)
	// Register order ke router
    routes.OrderRoutes(router)

	// Menjalankan server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}