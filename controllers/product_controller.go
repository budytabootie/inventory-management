package controllers

import (
	"fmt"
	"inventory-management/config"
	"inventory-management/models"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const uploadDir = "./uploads"


// Fungsi untuk mendapatkan daftar produk
func GetProducts(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name, description, price, category FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Desc, &product.Price, &product.Category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, product)
	}
	c.JSON(http.StatusOK, products)
}

// Fungsi untuk membuat produk baru
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("INSERT INTO products (name, description, price, category) VALUES (?, ?, ?, ?)",
		product.Name, product.Desc, product.Price, product.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully!"})
}

// Fungsi untuk memperbarui produk
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("UPDATE products SET name = ?, description = ?, price = ?, category = ? WHERE id = ?",
		product.Name, product.Desc, product.Price, product.Category, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully!"})
}

// Fungsi untuk menghapus produk
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	_, err := config.DB.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully!"})
}

// Fungsi untuk mengunggah gambar produk
func UploadProductImage(c *gin.Context) {
    productID := c.Param("id")

    // Ambil file dari form-data
    file, err := c.FormFile("file") // ubah "image" menjadi "file"
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file: " + err.Error()})
        return
    }

    // Tentukan direktori upload
    const uploadDir = "./uploads"

    // Buat folder upload jika belum ada
    if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
        os.MkdirAll(uploadDir, os.ModePerm)
    }

    // Simpan file dengan nama berdasarkan ID produk
    filePath := filepath.Join(uploadDir, productID+filepath.Ext(file.Filename))
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully!", "path": filePath})
}

// Fungsi untuk mengunduh gambar produk
func DownloadProductImage(c *gin.Context) {
    productID := c.Param("id")

    // Cari file berdasarkan ID produk
    files, err := filepath.Glob(filepath.Join(uploadDir, productID+".*"))
    if err != nil || len(files) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
        return
    }

    // Log file path untuk debugging
    fmt.Println("Menampilkan file:", files[0])

    // Set header untuk memicu unduhan dan menampilkan tipe gambar
    c.Header("Content-Type", "image/jpeg") // atau secara dinamis berdasarkan tipe file
    c.Header("Content-Disposition", "attachment; filename="+filepath.Base(files[0]))

    // Mengirim file gambar
    c.File(files[0])
}

