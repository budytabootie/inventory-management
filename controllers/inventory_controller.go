package controllers

import (
	"inventory-management/config"
	"inventory-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInventory(c *gin.Context) {
	rows, err := config.DB.Query("SELECT product_id, quantity, location FROM inventory")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var inventory []models.Inventory
	for rows.Next() {
		var item models.Inventory
		if err := rows.Scan(&item.ProductID, &item.Quantity, &item.Location); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		inventory = append(inventory, item)
	}
	c.JSON(http.StatusOK, inventory)
}

func CreateInventory(c *gin.Context) {
	var inventory models.Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("INSERT INTO inventory (product_id, quantity, location) VALUES (?, ?, ?)",
		inventory.ProductID, inventory.Quantity, inventory.Location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory created successfully!"})
}

func UpdateInventory(c *gin.Context) {
	var inventory models.Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("UPDATE inventory SET quantity = ?, location = ? WHERE product_id = ?",
		inventory.Quantity, inventory.Location, inventory.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory updated successfully!"})
}

func DeleteInventory(c *gin.Context) {
	productID := c.Param("id")

	_, err := config.DB.Exec("DELETE FROM inventory WHERE product_id = ?", productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory deleted successfully!"})
}
