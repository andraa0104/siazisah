# API Documentation - SI-AZISAH

Base URL: `http://localhost:8080/api`

## Authentication

Semua endpoint yang memerlukan autentikasi harus menyertakan header:
```
Authorization: Bearer <token>
```

---

## Public Endpoints (No Auth Required)

### 1. Get Public Dashboard
**GET** `/public/dashboard`

Response:
```json
{
  "success": true,
  "data": {
    "total_masjid": 3,
    "total_muzakki": 15,
    "total_mustahiq": 20,
    "total_zakat_fitrah": 5000000,
    "total_zakat_mal": 2000000,
    "total_infaq": 1000000
  }
}
```

### 2. Get All Masjid
**GET** `/public/masjid`

Response:
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "nama": "Masjid Al-Ikhlas",
      "alamat": "Jl. Raya Purwajaya No. 123",
      "desa": "Purwajaya",
      "kecamatan": "Kecamatan ABC",
      "kabupaten": "Kabupaten XYZ",
      "telepon": "081234567890"
    }
  ]
}
```

### 3. Get Masjid Stats
**GET** `/public/masjid/:id/stats`

Response:
```json
{
  "success": true,
  "data": {
    "masjid": {
      "id": 1,
      "nama": "Masjid Al-Ikhlas",
      "alamat": "Jl. Raya Purwajaya No. 123"
    },
    "total_muzakki": 10,
    "total_mustahiq": 15,
    "total_zakat_fitrah": 3000000,
    "total_zakat_mal": 1000000,
    "total_infaq": 500000
  }
}
```

---

## Auth Endpoints

### 1. Login
**POST** `/auth/login`

Request Body:
```json
{
  "username": "superadmin",
  "password": "admin123"
}
```

Response:
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "superadmin",
      "email": "superadmin@siazisah.com",
      "full_name": "Super Administrator",
      "role": "superadmin",
      "masjid_id": null,
      "is_active": true
    }
  }
}
```

---

## Superadmin Endpoints

### Masjid Management

#### 1. Get All Masjid
**GET** `/superadmin/masjid`

Headers: `Authorization: Bearer <token>`

Response:
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "nama": "Masjid Al-Ikhlas",
      "alamat": "Jl. Raya Purwajaya No. 123",
      "desa": "Purwajaya",
      "kecamatan": "Kecamatan ABC",
      "kabupaten": "Kabupaten XYZ",
      "provinsi": "Jawa Barat",
      "kode_pos": "12345",
      "telepon": "081234567890",
      "is_active": true
    }
  ]
}
```

#### 2. Create Masjid
**POST** `/superadmin/masjid`

Request Body:
```json
{
  "nama": "Masjid Al-Ikhlas",
  "alamat": "Jl. Raya Purwajaya No. 123",
  "desa": "Purwajaya",
  "kecamatan": "Kecamatan ABC",
  "kabupaten": "Kabupaten XYZ",
  "provinsi": "Jawa Barat",
  "kode_pos": "12345",
  "telepon": "081234567890",
  "is_active": true
}
```

#### 3. Update Masjid
**PUT** `/superadmin/masjid/:id`

Request Body: (sama dengan Create)

#### 4. Delete Masjid
**DELETE** `/superadmin/masjid/:id`

### User Management

#### 1. Get All Users
**GET** `/superadmin/users`

#### 2. Create User (Petugas)
**POST** `/superadmin/users`

Request Body:
```json
{
  "username": "petugas1",
  "email": "petugas1@siazisah.com",
  "password": "petugas123",
  "full_name": "Ahmad Petugas",
  "role": "petugas",
  "masjid_id": 1,
  "is_active": true
}
```

#### 3. Update User
**PUT** `/superadmin/users/:id`

#### 4. Delete User
**DELETE** `/superadmin/users/:id`

---

## Petugas Endpoints

### Masjid

#### 1. Get My Masjid
**GET** `/petugas/masjid`

#### 2. Update My Masjid
**PUT** `/petugas/masjid`

### Muzakki Management

#### 1. Get All Muzakki
**GET** `/petugas/muzakki`

Response:
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "masjid_id": 1,
      "nama": "H. Muhammad",
      "alamat": "Jl. Makmur No. 11",
      "telepon": "081234567804"
    }
  ]
}
```

#### 2. Create Muzakki
**POST** `/petugas/muzakki`

Request Body:
```json
{
  "nama": "H. Muhammad",
  "alamat": "Jl. Makmur No. 11",
  "telepon": "081234567804"
}
```

#### 3. Update Muzakki
**PUT** `/petugas/muzakki/:id`

#### 4. Delete Muzakki
**DELETE** `/petugas/muzakki/:id`

### Mustahiq Management

#### 1. Get All Mustahiq
**GET** `/petugas/mustahiq`

Response:
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "masjid_id": 1,
      "nama": "Pak Slamet",
      "jenis_penerima": "fakir",
      "alamat": "Jl. Kampung No. 12",
      "lokasi": "dalam_desa",
      "rt": "01",
      "telepon": "081234567899",
      "is_active": true
    }
  ]
}
```

#### 2. Create Mustahiq
**POST** `/petugas/mustahiq`

Request Body:
```json
{
  "nama": "Pak Slamet",
  "jenis_penerima": "fakir",
  "alamat": "Jl. Kampung No. 12",
  "lokasi": "dalam_desa",
  "rt": "01",
  "telepon": "081234567899",
  "keterangan": "Penerima tetap"
}
```

Jenis Penerima: `fakir`, `miskin`, `amil`, `mualaf`, `riqab`, `gharim`, `fisabilillah`, `ibnu sabil`

Lokasi: `dalam_desa`, `luar_desa`

#### 3. Update Mustahiq
**PUT** `/petugas/mustahiq/:id`

#### 4. Delete Mustahiq
**DELETE** `/petugas/mustahiq/:id`

### Transaksi Management

#### 1. Get All Transaksi
**GET** `/petugas/transaksi`

Response:
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "masjid_id": 1,
      "muzakki_id": 1,
      "muzakki_nama": "H. Muhammad",
      "muzakki_alamat": "Jl. Makmur No. 11",
      "jenis_zakat": "fitrah",
      "bentuk_zakat": "uang",
      "kelas_zakat": "2",
      "jumlah_orang": 4,
      "nominal_per_orang": 45000,
      "total_wajib": 180000,
      "total_dibayar": 200000,
      "infaq_tambahan": 20000,
      "tahun": 2024,
      "tanggal_bayar": "2024-04-10"
    }
  ]
}
```

#### 2. Create Transaksi
**POST** `/petugas/transaksi`

Request Body (Zakat Fitrah - Uang):
```json
{
  "muzakki_id": 1,
  "jenis_zakat": "fitrah",
  "bentuk_zakat": "uang",
  "kelas_zakat": "2",
  "jumlah_orang": 4,
  "nominal_per_orang": 45000,
  "total_dibayar": 200000,
  "tanggal_bayar": "2024-04-10",
  "keterangan": "Zakat fitrah keluarga"
}
```

Request Body (Zakat Mal):
```json
{
  "muzakki_id": 1,
  "jenis_zakat": "mal",
  "total_dibayar": 500000,
  "tanggal_bayar": "2024-04-10",
  "keterangan": "Zakat mal"
}
```

Request Body (Infaq):
```json
{
  "muzakki_id": 1,
  "jenis_zakat": "infaq",
  "total_dibayar": 100000,
  "tanggal_bayar": "2024-04-10",
  "keterangan": "Infaq"
}
```

#### 3. Get Transaksi by ID
**GET** `/petugas/transaksi/:id`

#### 4. Print Receipt
**GET** `/petugas/transaksi/:id/print`

Returns HTML for printing

#### 5. Delete Transaksi
**DELETE** `/petugas/transaksi/:id`

---

## Error Responses

### 400 Bad Request
```json
{
  "success": false,
  "message": "Invalid request"
}
```

### 401 Unauthorized
```json
{
  "success": false,
  "message": "Authorization header required"
}
```

### 403 Forbidden
```json
{
  "success": false,
  "message": "Superadmin access required"
}
```

### 404 Not Found
```json
{
  "success": false,
  "message": "Resource not found"
}
```

### 500 Internal Server Error
```json
{
  "success": false,
  "message": "Internal server error"
}
```

---

## Notes

1. Semua endpoint yang memerlukan autentikasi harus menyertakan JWT token di header
2. Token expired setelah 24 jam
3. Untuk perhitungan zakat fitrah otomatis:
   - `total_wajib = nominal_per_orang * jumlah_orang`
   - `infaq_tambahan = total_dibayar - total_wajib` (jika total_dibayar > total_wajib)
4. RT untuk lokasi "dalam_desa" adalah dropdown 01-21
5. RT untuk lokasi "luar_desa" adalah input manual (number)
