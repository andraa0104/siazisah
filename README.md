# SI-AZISAH - Sistem Informasi Zakat

Sistem Informasi Zakat yang lengkap dengan fitur untuk Superadmin, Petugas Zakat, dan User Umum.

## 🚀 Fitur Utama

### 👨‍💼 Superadmin
- ✅ Kelola data masjid/langgar
- ✅ Kelola akun petugas zakat
- ✅ Dashboard statistik keseluruhan
- ✅ CRUD lengkap untuk masjid dan users

### 👤 Petugas Zakat
- ✅ Kelola data masjid/langgar sendiri
- ✅ Input data pengurus masjid dan pengurus zakat
- ✅ Atur kadar zakat fitrah (Kelas 1-3)
- ✅ Pilih jenis zakat yang disediakan (Fitrah, Mal, Infaq)
- ✅ CRUD data penerima zakat (Mustahiq)
  - Pilihan lokasi: Dalam/Luar Desa Purwajaya
  - RT 01-21 untuk dalam desa (dropdown)
  - Input manual RT untuk luar desa
- ✅ Input data muzakki dengan perhitungan otomatis
  - Zakat Fitrah: Uang atau Beras
  - Pilih kelas zakat (1-3)
  - Jumlah orang yang dizakati
  - Perhitungan otomatis total wajib
  - Deteksi infaq tambahan otomatis
- ✅ Cetak bukti pembayaran (Print Receipt)
- ✅ Dashboard statistik masjid

### 🌐 User Umum (Tanpa Login)
- ✅ Lihat dashboard publik
- ✅ Lihat total zakat keseluruhan
- ✅ Lihat data per masjid/langgar
- ✅ Lihat jumlah muzakki dan mustahiq
- ✅ Transparansi pengelolaan zakat

## 🛠️ Teknologi

### Backend
- **Golang 1.21+**
- **Gin Web Framework**
- **MySQL Database**
- **JWT Authentication**
- **Bcrypt Password Hashing**

### Frontend
- **HTML5, CSS3, JavaScript**
- **Tailwind CSS** (Modern & Responsive)
- **Font Awesome Icons**
- **Fetch API** untuk komunikasi dengan backend

## 📁 Struktur Project

```
zakat1/
├── backend/                    # Backend Golang
│   ├── cmd/
│   │   └── main.go            # Entry point
│   ├── config/
│   │   └── config.go          # Database config
│   ├── internal/
│   │   ├── handlers/          # HTTP handlers/controllers
│   │   │   ├── auth_handler.go
│   │   │   ├── user_handler.go
│   │   │   ├── masjid_handler.go
│   │   │   ├── muzakki_handler.go
│   │   │   ├── mustahiq_handler.go
│   │   │   ├── transaksi_handler.go
│   │   │   └── public_handler.go
│   │   ├── middleware/        # Auth middleware
│   │   │   └── auth.go
│   │   ├── models/            # Data models
│   │   │   └── models.go
│   │   ├── repository/        # Database operations
│   │   │   ├── user_repository.go
│   │   │   ├── masjid_repository.go
│   │   │   ├── muzakki_repository.go
│   │   │   ├── mustahiq_repository.go
│   │   │   └── transaksi_repository.go
│   │   └── utils/             # Utility functions
│   │       └── auth.go
│   ├── .env                   # Environment variables
│   ├── go.mod                 # Go dependencies
│   └── README.md
│
├── frontend/                   # Frontend HTML/CSS/JS
│   ├── pages/
│   │   ├── superadmin/
│   │   │   ├── dashboard.html
│   │   │   ├── masjid.html
│   │   │   └── users.html
│   │   └── petugas/
│   │       ├── dashboard.html
│   │       ├── muzakki.html
│   │       ├── mustahiq.html
│   │       └── transaksi.html
│   ├── js/
│   │   ├── config.js          # API configuration
│   │   ├── auth.js            # Authentication
│   │   └── public.js          # Public dashboard
│   ├── index.html             # Login page
│   └── public.html            # Public dashboard
│
└── database/
    └── schema.sql             # Database schema

```

## 🚀 Cara Menjalankan

### 1. Setup Database

```bash
# Buat database dan import schema
mysql -u root -p
CREATE DATABASE siazisah;
exit

# Import schema
mysql -u root -p siazisah < database/schema.sql
```

### 2. Generate Password Superadmin

Buat file `backend/cmd/generate_password.go`:

```go
package main

import (
    "fmt"
    "os"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run generate_password.go <password>")
        return
    }
    
    password := os.Args[1]
    hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
    fmt.Println(string(hash))
}
```

Jalankan:
```bash
cd backend
go run cmd/generate_password.go admin123
```

Copy hash yang dihasilkan dan update di database:
```sql
UPDATE users SET password = '<hash_yang_dihasilkan>' WHERE username = 'superadmin';
```

### 3. Jalankan Backend

```bash
cd backend
go mod download
go run cmd/main.go
```

Backend akan berjalan di `http://localhost:8080`

### 4. Jalankan Frontend

Buka `frontend/index.html` di browser atau gunakan live server.

Jika menggunakan Laragon, akses:
```
http://localhost/zakat1/frontend/
```

## 🔐 Default Login

**Superadmin:**
- Username: `superadmin`
- Password: `admin123`

## 📊 Database Schema

### Tables:
1. **users** - Superadmin dan Petugas
2. **masjid** - Data masjid/langgar
3. **pengurus_masjid** - Pengurus masjid
4. **pengurus_zakat** - Pengurus zakat
5. **kadar_zakat** - Kadar zakat fitrah (Kelas 1-3)
6. **jenis_zakat_tersedia** - Jenis zakat yang disediakan
7. **mustahiq** - Penerima zakat
8. **muzakki** - Pemberi zakat
9. **transaksi_zakat** - Transaksi pembayaran zakat
10. **distribusi_zakat** - Distribusi zakat ke mustahiq

## 🎨 Design Features

- ✨ Modern & Elegant UI
- 📱 Fully Responsive
- 🎯 User-friendly UX
- 🎨 Gradient Colors
- 💫 Smooth Transitions
- 📊 Interactive Dashboard
- 🖨️ Print-ready Receipts

## 🔄 Workflow

### Superadmin:
1. Login ke sistem
2. Tambah data masjid/langgar
3. Buat akun petugas untuk setiap masjid
4. Monitor statistik keseluruhan

### Petugas:
1. Login dengan akun yang dibuat superadmin
2. Lengkapi data masjid/langgar
3. Input pengurus masjid dan pengurus zakat
4. Atur kadar zakat fitrah (Kelas 1-3)
5. Pilih jenis zakat yang disediakan
6. Input data mustahiq (penerima zakat)
7. Input data muzakki dan transaksi zakat
8. Cetak bukti pembayaran

### User Umum:
1. Buka halaman publik (tanpa login)
2. Lihat statistik keseluruhan
3. Lihat data per masjid
4. Transparansi pengelolaan zakat

## 📝 Fitur Khusus

### Perhitungan Zakat Otomatis
- Input jumlah orang yang dizakati
- Pilih kelas zakat (1, 2, atau 3)
- Sistem menghitung total wajib otomatis
- Jika bayar lebih, otomatis masuk infaq

### Lokasi Mustahiq
- **Dalam Desa Purwajaya**: Dropdown RT 01-21
- **Luar Desa Purwajaya**: Input manual RT

### Print Receipt
- Bukti pembayaran profesional
- Siap cetak
- Include semua detail transaksi

## 🔧 Konfigurasi

Edit `backend/.env`:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=siazisah
SERVER_PORT=8080
JWT_SECRET=your-secret-key-change-this
```

Edit `frontend/js/config.js`:
```javascript
const API_BASE_URL = 'http://localhost:8080/api';
```

## 📞 Support

Untuk pertanyaan atau bantuan, silakan hubungi tim development.

## 📄 License

© 2024 SI-AZISAH. All rights reserved.
