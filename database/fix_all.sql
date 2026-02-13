-- 1. Alter table transaksi_zakat untuk support zakat mal
ALTER TABLE transaksi_zakat 
ADD COLUMN jenis_harta VARCHAR(50) NULL AFTER bentuk_zakat,
ADD COLUMN nominal_harta DECIMAL(15,2) NULL AFTER jenis_harta,
ADD COLUMN persentase_zakat DECIMAL(5,2) NULL AFTER nominal_harta;

-- 2. Pastikan ada masjid untuk testing
INSERT INTO masjid (nama, alamat, desa, kecamatan, kabupaten, provinsi, telepon, is_active) 
VALUES ('Masjid Al-Ikhlas', 'Jl. Raya Purwajaya', 'Purwajaya', 'Kecamatan Test', 'Kabupaten Test', 'Jawa Barat', '08123456789', 1)
ON DUPLICATE KEY UPDATE nama=nama;

-- 3. Update petugas1 dengan masjid_id
UPDATE users SET masjid_id = 1 WHERE username = 'petugas1';

-- 4. Verifikasi
SELECT id, username, role, masjid_id FROM users WHERE username = 'petugas1';
SELECT id, nama FROM masjid LIMIT 1;
