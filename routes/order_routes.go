package routes

import (
	"inventory-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	orders := router.Group("/orders")
	{
		orders.POST("/", controllers.CreateOrder)
		orders.GET("/:id", controllers.GetOrder)
		orders.GET("/", controllers.GetAllOrders)
	}
}
