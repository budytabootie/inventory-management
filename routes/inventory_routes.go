package routes

import (
	"inventory-management/controllers"

	"github.com/gin-gonic/gin"
)

func InventoryRoutes(router *gin.Engine) {
	inventory := router.Group("/inventory")
	{
		inventory.GET("/", controllers.GetInventory)
		inventory.POST("/", controllers.CreateInventory)
		inventory.PUT("/", controllers.UpdateInventory)
		inventory.DELETE("/:id", controllers.DeleteInventory)
	}
}
