# Inventory Management System

## Deskripsi Proyek

### Tujuan Aplikasi

Aplikasi ini adalah sistem manajemen backend untuk mengelola inventaris produk. Sistem ini memungkinkan pengguna untuk melakukan operasi CRUD (Create, Read, Update, Delete) pada data produk, inventaris, dan gambar produk. Tujuan utamanya adalah untuk menyediakan API yang dapat digunakan untuk mengelola produk dan inventaris secara efisien.

### Fitur Aplikasi

- Menambahkan, mengubah, dan menghapus produk
- Mengelola inventaris dengan informasi kuantitas dan lokasi
- Mengunggah dan mengunduh gambar produk
- Mengakses data produk dan inventaris melalui API RESTful

### Pengguna yang Dimaksud

Aplikasi ini ditujukan untuk administrator atau staf yang bertanggung jawab atas manajemen produk dan inventaris di sistem backend.

---

## Persyaratan Sistem

- **Bahasa Pemrograman**: Golang
- **Framework**: Gin
- **Database**: MySQL
- **Alat Pengujian API**: REST Client Extension pada VS Code

---

## Instruksi Pengaturan

### 1. Menyiapkan Database

1.  **Buat Database**:
    Gunakan perintah SQL berikut untuk membuat database baru:

    ```sql
    CREATE DATABASE inventory_db;

    ```

2.  **Buat Tabel**:
    Setelah membuat database, jalankan skrip SQL berikut untuk membuat tabel yang diperlukan:

	* Tabel Produk:

    ```sql

    CREATE TABLE products (
    	id INT AUTO_INCREMENT PRIMARY KEY,
    	name VARCHAR(255),
    	description TEXT,
    	price DECIMAL(10, 2)
    );

    ```

	* Tabel Inventory:

    ```sql

    CREATE TABLE inventory (
    	product_id INT PRIMARY KEY,
    	quantity INT,
    	location VARCHAR(255),
    	FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
	);

    ```
	* Tabel Order:

    ```sql

    CREATE TABLE orders (
    	id INT AUTO_INCREMENT PRIMARY KEY,
    	product_id INT,
    	quantity INT,
    	date DATE,
    	FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
	);

    ```
2.  **Sesuaikan Konfigurasi Database**:
    Pastikan Anda mengonfigurasi koneksi database di dalam aplikasi sesuai dengan kredensial MySQL Anda.

---

## Instruksi Pengaturan

### 1. Menjalankan Server

Setelah mengikuti langkah-langkah di bagian Instruksi Pengaturan, Anda bisa mulai menjalankan server backend menggunakan Golang.
	* __Jalankan Server__  : Untuk memulai server, gunakan perintah berikut di terminal:

	```bash
	go run main.go
	```
	Server akan berjalan di http://localhost:8080.

### 2. Mengakses Endpoint API

Aplikasi ini memiliki beberapa endpoint RESTful API yang dapat diakses untuk melakukan operasi CRUD (Create, Read, Update, Delete) pada data produk dan inventaris. Berikut adalah daftar endpoint dan contoh penggunaan:

a. Mendapatkan Daftar Produk

* GET /products
* Deskripsi: Mengambil daftar semua produk yang ada di database.
* Contoh Penggunaan:

```http
GET http://localhost:8080/products
```

b. Menambahkan Produk Baru

* POST /products
* Deskripsi: Menambahkan produk baru ke dalam sistem.
* Body Request (JSON):
```json
{
  "name": "Product Name",
  "description": "Product Description",
  "price": 99.99
}
```
* Contoh Penggunaan:
```http
POST http://localhost:8080/products
Content-Type: application/json

{
  "name": "Product Name",
  "description": "Product Description",
  "price": 99.99
}
```

c. Mendapatkan Detail Produk Berdasarkan ID

* GET /products/{id}
* Deskripsi: Mengambil detail produk berdasarkan ID produk.
* Contoh Penggunaan:
```http
GET http://localhost:8080/products/1
```

d. Mengupdate Produk
* PUT /products/{id}
* Deskripsi: Mengupdate informasi produk berdasarkan ID produk.
* Body Request (JSON):
```json
{
  "name": "Product Name",
  "description": "Product Description",
  "price": 99.99
}
```
* Contoh Penggunaan:
```http
PUT http://localhost:8080/products/1
Content-Type: application/json

{
  "name": "Updated Product Name",
  "description": "Updated Product Description",
  "price": 199.99
}
```

e. Menghapus Produk										
* DELETE /products/{id}
* Deskripsi: Menghapus produk berdasarkan ID produk.
* Contoh Penggunaan:
```http
DELETE http://localhost:8080/products/1
```
f. Mengunggah Gambar Produk
* POST /products/{id}/upload
* Deskripsi: Mengunggah gambar untuk produk berdasarkan ID produk.
* Body Request: Gunakan multipart/form-data untuk mengirim file gambar.
* Contoh Penggunaan:
```http
POST http://localhost:8080/products/1/upload
Content-Type: multipart/form-data

--boundary
Content-Disposition: form-data; name="file"; filename="product_image.jpg"
Content-Type: image/jpeg

< actual binary content of the product_image.jpg file here >
--boundary--
```

g. Mengunduh Gambar Produk
* GET /products/{id}/image
* Deskripsi: Mengunduh gambar produk berdasarkan ID produk.
* Contoh Penggunaan:
```http
GET http://localhost:8080/products/1/image
```

h. Mendapatkan Daftar Inventaris
* GET /inventory
* Deskripsi: Mengambil daftar semua item inventaris di database.
* Contoh Penggunaan:
```http
GET http://localhost:8080/inventory
```

i. Menambahkan Item ke Inventaris
* POST /inventory
* Deskripsi: Menambahkan item baru ke inventaris.
* Body Request (JSON):
```json
{
  "product_id": 1,
  "quantity": 50,
  "location": "Warehouse A"
}
```
* Contoh Penggunaan:
```http
POST http://localhost:8080/inventory
Content-Type: application/json

{
  "product_id": 1,
  "quantity": 50,
  "location": "Warehouse A"
}
```

j. Mengupdate Inventaris
* PUT /inventory
* Deskripsi: Mengupdate jumlah atau lokasi item di inventaris.
* Body Request (JSON):
```json
{
  "product_id": 1,
  "quantity": 100,
  "location": "Warehouse B"
}
```
* Contoh Penggunaan:
```http
PUT http://localhost:8080/inventory
Content-Type: application/json

{
  "product_id": 1,
  "quantity": 100,
  "location": "Warehouse B"
}
```
k. Menghapus Item dari Inventaris
* DELETE /inventory/{id}
* Deskripsi: Menghapus item dari inventaris berdasarkan ID produk.
* Contoh Penggunaan:
```http
DELETE http://localhost:8080/inventory/1
```