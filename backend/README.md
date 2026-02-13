# SI-AZISAH Backend

Backend API untuk Sistem Informasi Zakat menggunakan Golang, Gin Framework, dan MySQL.

## Fitur

- **Superadmin**: Kelola masjid/langgar dan akun petugas
- **Petugas**: Kelola data masjid, muzakki, mustahiq, dan transaksi zakat
- **Public API**: Dashboard publik untuk user umum (tanpa login)

## Teknologi

- Golang 1.21+
- Gin Web Framework
- MySQL Database
- JWT Authentication
- Bcrypt Password Hashing

## Setup

### 1. Install Dependencies

```bash
cd backend
go mod download
```

### 2. Setup Database

Jalankan script SQL di `../database/schema.sql` ke MySQL:

```bash
mysql -u root -p < ../database/schema.sql
```

### 3. Konfigurasi Environment

Edit file `.env` sesuai dengan konfigurasi database Anda:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=siazisah
SERVER_PORT=8080
JWT_SECRET=your-secret-key-change-this
```

### 4. Generate Password untuk Superadmin

Jalankan script untuk generate password hash:

```bash
go run cmd/generate_password.go admin123
```

Copy hash yang dihasilkan dan update di database.

### 5. Jalankan Server

```bash
go run cmd/main.go
```

Server akan berjalan di `http://localhost:8080`

## API Endpoints

### Public API (No Auth)

- `GET /api/public/dashboard` - Dashboard publik
- `GET /api/public/masjid` - List semua masjid
- `GET /api/public/masjid/:id/stats` - Statistik per masjid

### Authentication

- `POST /api/auth/login` - Login

### Superadmin Routes

- `GET /api/superadmin/masjid` - List masjid
- `POST /api/superadmin/masjid` - Create masjid
- `PUT /api/superadmin/masjid/:id` - Update masjid
- `DELETE /api/superadmin/masjid/:id` - Delete masjid
- `GET /api/superadmin/users` - List users
- `POST /api/superadmin/users` - Create user (petugas)
- `PUT /api/superadmin/users/:id` - Update user
- `DELETE /api/superadmin/users/:id` - Delete user

### Petugas Routes

- `GET /api/petugas/masjid` - Get my masjid
- `PUT /api/petugas/masjid` - Update my masjid
- `GET /api/petugas/muzakki` - List muzakki
- `POST /api/petugas/muzakki` - Create muzakki
- `PUT /api/petugas/muzakki/:id` - Update muzakki
- `DELETE /api/petugas/muzakki/:id` - Delete muzakki
- `GET /api/petugas/mustahiq` - List mustahiq
- `POST /api/petugas/mustahiq` - Create mustahiq
- `PUT /api/petugas/mustahiq/:id` - Update mustahiq
- `DELETE /api/petugas/mustahiq/:id` - Delete mustahiq
- `GET /api/petugas/transaksi` - List transaksi
- `POST /api/petugas/transaksi` - Create transaksi
- `GET /api/petugas/transaksi/:id/print` - Print receipt
- `DELETE /api/petugas/transaksi/:id` - Delete transaksi

## Default Login

**Superadmin:**
- Username: `superadmin`
- Password: `admin123`

## Struktur Project

```
backend/
├── cmd/
│   └── main.go              # Entry point
├── config/
│   └── config.go            # Database & app config
├── internal/
│   ├── handlers/            # HTTP handlers
│   ├── middleware/          # Auth middleware
│   ├── models/              # Data models
│   ├── repository/          # Database operations
│   └── utils/               # Utility functions
├── .env                     # Environment variables
└── go.mod                   # Go dependencies
```
