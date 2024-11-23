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

	-Tabel Produk:

    ```sql

    CREATE TABLE products (
    	id INT AUTO_INCREMENT PRIMARY KEY,
    	name VARCHAR(255),
    	description TEXT,
    	price DECIMAL(10, 2)
    );

    ```

	-Tabel Inventaris:

    ```sql

    CREATE TABLE inventory (
    	product_id INT PRIMARY KEY,
    	quantity INT,
    	location VARCHAR(255),
    	FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
	);

    ```
2.  **Sesuaikan Konfigurasi Database**:
    Pastikan Anda mengonfigurasi koneksi database di dalam aplikasi sesuai dengan kredensial MySQL Anda.
