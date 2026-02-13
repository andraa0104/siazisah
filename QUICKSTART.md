# Quick Start Guide - SI-AZISAH

## 🚀 Langkah Cepat Memulai

### 1. Setup Database (5 menit)

```bash
# Masuk ke MySQL
mysql -u root -p

# Buat database
CREATE DATABASE siazisah;
exit

# Import schema
mysql -u root -p siazisah < database/schema.sql
```

### 2. Generate Password Superadmin (2 menit)

```bash
cd backend
go run cmd/generate_password.go admin123
```

Copy hash yang dihasilkan, lalu update di database:

```sql
mysql -u root -p siazisah

UPDATE users SET password = '<paste_hash_disini>' WHERE username = 'superadmin';
exit
```

### 3. Jalankan Backend (1 menit)

```bash
cd backend
go mod download
go run cmd/main.go
```

✅ Backend berjalan di: `http://localhost:8080`

### 4. Buka Frontend

Jika menggunakan Laragon:
- Buka browser: `http://localhost/zakat1/frontend/`

Atau buka langsung file:
- `frontend/index.html`

### 5. Login

**Superadmin:**
- Username: `superadmin`
- Password: `admin123`

---

## 📋 Checklist Setup

- [ ] Database siazisah sudah dibuat
- [ ] Schema sudah diimport
- [ ] Password superadmin sudah di-generate dan diupdate
- [ ] Backend berjalan di port 8080
- [ ] Frontend bisa diakses
- [ ] Bisa login sebagai superadmin

---

## 🎯 Langkah Selanjutnya

### Sebagai Superadmin:

1. **Tambah Masjid/Langgar**
   - Masuk ke menu "Masjid/Langgar"
   - Klik "Tambah Masjid"
   - Isi data lengkap
   - Simpan

2. **Buat Akun Petugas**
   - Masuk ke menu "Petugas"
   - Klik "Tambah Petugas"
   - Isi data dan pilih masjid
   - Simpan

### Sebagai Petugas:

1. **Login dengan akun petugas**
2. **Lengkapi data masjid**
3. **Tambah data Muzakki** (pemberi zakat)
4. **Tambah data Mustahiq** (penerima zakat)
5. **Input transaksi zakat**
6. **Cetak bukti pembayaran**

### User Umum:

1. Buka `public.html`
2. Lihat statistik zakat
3. Lihat data per masjid

---

## 🔧 Troubleshooting

### Backend tidak bisa connect ke database

Cek file `backend/.env`:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=siazisah
```

### CORS Error

Pastikan backend sudah berjalan di port 8080.

Cek `frontend/js/config.js`:
```javascript
const API_BASE_URL = 'http://localhost:8080/api';
```

### Login gagal

1. Pastikan password sudah di-generate dan diupdate di database
2. Cek username: `superadmin`
3. Cek password: `admin123`

---

## 📞 Butuh Bantuan?

Jika ada masalah, cek:
1. Backend logs di terminal
2. Browser console (F12)
3. MySQL error logs

---

## ✨ Fitur Unggulan

- ✅ **Perhitungan Otomatis**: Zakat fitrah dihitung otomatis berdasarkan kelas dan jumlah orang
- ✅ **Infaq Otomatis**: Kelebihan bayar otomatis masuk infaq
- ✅ **RT Dinamis**: Dropdown RT 01-21 untuk dalam desa, input manual untuk luar desa
- ✅ **Print Receipt**: Cetak bukti pembayaran profesional
- ✅ **Dashboard Publik**: Transparansi pengelolaan zakat
- ✅ **Responsive Design**: Bisa diakses dari HP, tablet, atau desktop

---

Selamat menggunakan SI-AZISAH! 🎉
