# Backend Inventory Management System

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

1. **Buat Database**:
   Gunakan perintah SQL berikut untuk membuat database baru:

   ```sql
   CREATE DATABASE inventory_db;
