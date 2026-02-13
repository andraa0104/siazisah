# 📦 SI-AZISAH - Installation Guide

Panduan instalasi lengkap untuk Sistem Informasi Zakat.

---

## 📋 Prerequisites

Sebelum memulai, pastikan sudah terinstall:

### Required:
- ✅ **Go 1.21+** - [Download](https://golang.org/dl/)
- ✅ **MySQL 5.7+** atau **MariaDB 10.3+**
- ✅ **Web Browser** (Chrome, Firefox, Edge, Safari)

### Optional:
- ✅ **Laragon** (untuk Windows) - [Download](https://laragon.org/)
- ✅ **XAMPP** (alternatif untuk Windows/Mac/Linux)
- ✅ **Git** (untuk version control)

---

## 🚀 Installation Steps

### Step 1: Download/Clone Project

#### Option A: Download ZIP
1. Download project sebagai ZIP
2. Extract ke folder `c:\laragon\www\zakat1`

#### Option B: Git Clone
```bash
cd c:\laragon\www
git clone <repository-url> zakat1
cd zakat1
```

---

### Step 2: Setup Database

#### Option A: Menggunakan Script (Recommended)
```bash
# Double click file ini:
setup-database.bat

# Ikuti instruksi di layar
```

#### Option B: Manual Setup
```bash
# 1. Buka MySQL
mysql -u root -p

# 2. Buat database
CREATE DATABASE siazisah;
exit

# 3. Import schema
mysql -u root -p siazisah < database/schema.sql

# 4. (Optional) Import sample data
mysql -u root -p siazisah < database/sample_data.sql
```

---

### Step 3: Generate Password Superadmin

```bash
# Masuk ke folder backend
cd backend

# Generate password hash
go run cmd/generate_password.go admin123

# Output akan seperti ini:
# $2a$14$xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

**Copy hash yang dihasilkan**, lalu update di database:

```bash
# Buka MySQL
mysql -u root -p siazisah

# Update password
UPDATE users SET password = '$2a$14$xxxxx...' WHERE username = 'superadmin';

# Keluar
exit
```

---

### Step 4: Configure Backend

Edit file `backend/.env`:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=          # Isi jika ada password
DB_NAME=siazisah

# Server Configuration
SERVER_PORT=8080
SERVER_HOST=localhost

# JWT Configuration
JWT_SECRET=your-secret-key-change-this-in-production
JWT_EXPIRATION=24h

# Application
APP_ENV=development
APP_NAME=SI-AZISAH
```

**Penting:** Ganti `JWT_SECRET` dengan string random yang aman!

---

### Step 5: Install Backend Dependencies

```bash
cd backend

# Download dependencies
go mod download

# Verify installation
go mod verify
```

---

### Step 6: Configure Frontend

Edit file `frontend/js/config.js`:

```javascript
// Sesuaikan dengan URL backend Anda
const API_BASE_URL = 'http://localhost:8080/api';
```

Jika menggunakan port lain, sesuaikan URL-nya.

---

### Step 7: Start Backend Server

#### Option A: Menggunakan Script (Recommended)
```bash
# Double click file ini:
start-backend.bat
```

#### Option B: Manual Start
```bash
cd backend
go run cmd/main.go
```

**Output yang benar:**
```
Database connected successfully
Server running on port 8080
```

Jika ada error, cek:
- Database sudah running?
- Konfigurasi `.env` sudah benar?
- Port 8080 tidak digunakan aplikasi lain?

---

### Step 8: Access Frontend

#### Option A: Menggunakan Laragon
```
http://localhost/zakat1/frontend/
```

#### Option B: Menggunakan XAMPP
```
http://localhost/zakat1/frontend/
```

#### Option C: Direct File
Buka file `frontend/index.html` di browser.

**Note:** Jika buka direct file, pastikan backend sudah running dan CORS sudah enabled.

---

## ✅ Verification

### 1. Test Backend API

Buka browser dan akses:
```
http://localhost:8080/api/public/dashboard
```

**Response yang benar:**
```json
{
  "success": true,
  "data": {
    "total_masjid": 0,
    "total_muzakki": 0,
    "total_mustahiq": 0,
    "total_zakat_fitrah": 0,
    "total_zakat_mal": 0,
    "total_infaq": 0
  }
}
```

### 2. Test Login

1. Buka `http://localhost/zakat1/frontend/`
2. Login dengan:
   - Username: `superadmin`
   - Password: `admin123`
3. Jika berhasil, akan redirect ke dashboard superadmin

### 3. Test Public Dashboard

1. Buka `http://localhost/zakat1/frontend/public.html`
2. Harus bisa melihat statistik tanpa login

---

## 🔧 Troubleshooting

### Problem 1: Backend tidak bisa connect ke database

**Error:**
```
Error connecting to database: Access denied for user 'root'@'localhost'
```

**Solution:**
1. Cek username & password di `.env`
2. Pastikan MySQL sudah running
3. Test koneksi manual:
   ```bash
   mysql -u root -p
   ```

---

### Problem 2: Port 8080 sudah digunakan

**Error:**
```
listen tcp :8080: bind: address already in use
```

**Solution:**
1. Ganti port di `.env`:
   ```env
   SERVER_PORT=8081
   ```
2. Update juga di `frontend/js/config.js`:
   ```javascript
   const API_BASE_URL = 'http://localhost:8081/api';
   ```

---

### Problem 3: CORS Error di Browser

**Error:**
```
Access to fetch at 'http://localhost:8080/api/...' from origin 'http://localhost' has been blocked by CORS policy
```

**Solution:**
Backend sudah include CORS middleware. Jika masih error:
1. Pastikan backend running
2. Clear browser cache
3. Restart browser

---

### Problem 4: Login gagal

**Error:**
```
Invalid username or password
```

**Solution:**
1. Pastikan password sudah di-generate dan diupdate di database
2. Cek di database:
   ```sql
   SELECT username, password FROM users WHERE username = 'superadmin';
   ```
3. Password hash harus dimulai dengan `$2a$14$`

---

### Problem 5: Go command not found

**Error:**
```
'go' is not recognized as an internal or external command
```

**Solution:**
1. Install Go dari https://golang.org/dl/
2. Restart terminal/command prompt
3. Verify: `go version`

---

## 🎯 Post-Installation

### 1. Create Sample Data (Optional)

Jika belum import sample data:
```bash
mysql -u root -p siazisah < database/sample_data.sql
```

**Note:** Jangan lupa generate password untuk petugas sample:
```bash
cd backend
go run cmd/generate_password.go petugas123
# Update password di database untuk user petugas
```

---

### 2. Setup Superadmin

1. Login sebagai superadmin
2. Lengkapi profile jika perlu
3. Ubah password default

---

### 3. Create First Masjid

1. Login sebagai superadmin
2. Masuk ke menu "Masjid/Langgar"
3. Klik "Tambah Masjid"
4. Isi data lengkap:
   - Nama masjid
   - Alamat lengkap
   - Desa (default: Purwajaya)
   - Kecamatan, Kabupaten, Provinsi
   - Telepon
5. Simpan

---

### 4. Create First Petugas

1. Masih di dashboard superadmin
2. Masuk ke menu "Petugas"
3. Klik "Tambah Petugas"
4. Isi data:
   - Username (unique)
   - Email
   - Password
   - Nama lengkap
   - Pilih masjid
5. Simpan

---

### 5. Test Petugas Login

1. Logout dari superadmin
2. Login dengan akun petugas yang baru dibuat
3. Harus masuk ke dashboard petugas
4. Test fitur-fitur:
   - Lihat data masjid
   - Tambah mustahiq
   - Tambah transaksi

---

## 📱 Mobile Access

### Local Network Access

Untuk akses dari HP/tablet di jaringan yang sama:

1. **Cari IP komputer:**
   ```bash
   ipconfig
   # Cari IPv4 Address, contoh: 192.168.1.100
   ```

2. **Update backend `.env`:**
   ```env
   SERVER_HOST=0.0.0.0
   ```

3. **Restart backend**

4. **Akses dari HP:**
   ```
   http://192.168.1.100:8080/api/public/dashboard
   http://192.168.1.100/zakat1/frontend/
   ```

5. **Update frontend config di HP:**
   Edit `frontend/js/config.js`:
   ```javascript
   const API_BASE_URL = 'http://192.168.1.100:8080/api';
   ```

---

## 🔒 Security Checklist

### Before Production:

- [ ] Ganti `JWT_SECRET` dengan string random yang kuat
- [ ] Ganti password default superadmin
- [ ] Disable sample data di production
- [ ] Enable HTTPS
- [ ] Setup firewall
- [ ] Regular database backup
- [ ] Update Go & dependencies
- [ ] Enable rate limiting
- [ ] Setup logging
- [ ] Monitor server resources

---

## 📊 Performance Tips

### Database Optimization:
```sql
-- Add indexes jika data sudah banyak
CREATE INDEX idx_transaksi_tanggal ON transaksi_zakat(tanggal_bayar);
CREATE INDEX idx_muzakki_nama ON muzakki(nama);
CREATE INDEX idx_mustahiq_nama ON mustahiq(nama);
```

### Backend Optimization:
- Increase connection pool di `config.go`
- Enable caching untuk public API
- Use prepared statements

### Frontend Optimization:
- Minify CSS/JS
- Enable browser caching
- Compress images
- Use CDN untuk libraries

---

## 🔄 Update & Maintenance

### Update Backend:
```bash
cd backend
go get -u ./...
go mod tidy
```

### Backup Database:
```bash
# Backup
mysqldump -u root -p siazisah > backup_$(date +%Y%m%d).sql

# Restore
mysql -u root -p siazisah < backup_20240101.sql
```

### Check Logs:
```bash
# Backend logs
tail -f backend/logs/app.log

# MySQL logs
tail -f /var/log/mysql/error.log
```

---

## 📞 Support

Jika mengalami masalah:

1. **Check Documentation:**
   - README.md
   - QUICKSTART.md
   - API_DOCUMENTATION.md

2. **Check Logs:**
   - Backend terminal output
   - Browser console (F12)
   - MySQL error logs

3. **Common Issues:**
   - Database connection
   - Port conflicts
   - CORS errors
   - Authentication errors

---

## ✅ Installation Complete!

Jika semua step sudah dilakukan dengan benar, Anda sekarang memiliki:

- ✅ Backend API running di port 8080
- ✅ Database siazisah dengan schema lengkap
- ✅ Frontend accessible via browser
- ✅ Superadmin account ready
- ✅ Sample data (optional)

**Next Steps:**
1. Baca USER_GUIDE.md (jika ada)
2. Explore fitur-fitur
3. Customize sesuai kebutuhan
4. Deploy ke production (optional)

---

**Selamat! SI-AZISAH sudah siap digunakan!** 🎉

**Happy Coding!** 💻☕
