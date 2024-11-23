package controllers

import (
	"inventory-management/config"
	"inventory-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrder: Endpoint to create a new order
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("INSERT INTO orders (product_id, quantity, date) VALUES (?, ?, ?)",
		order.ProductID, order.Quantity, order.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully!"})
}

// GetOrder: Endpoint to retrieve order details by ID
func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	err := config.DB.QueryRow("SELECT id, product_id, quantity, date FROM orders WHERE id = ?", id).
		Scan(&order.ID, &order.ProductID, &order.Quantity, &order.Date)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// GetAllOrders: Endpoint to retrieve all orders
func GetAllOrders(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, product_id, quantity, date FROM orders")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.ProductID, &order.Quantity, &order.Date); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		orders = append(orders, order)
	}

	c.JSON(http.StatusOK, orders)
}
