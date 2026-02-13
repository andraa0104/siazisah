-- Cek data user petugas1
SELECT id, username, role, masjid_id, is_active FROM users WHERE username = 'petugas1';

-- Cek data masjid yang tersedia
SELECT id, nama FROM masjid LIMIT 5;

-- Update petugas1 dengan masjid_id = 1 (sesuaikan dengan ID masjid yang ada)
UPDATE users SET masjid_id = 1 WHERE username = 'petugas1';

-- Verifikasi hasil update
SELECT id, username, role, masjid_id, is_active FROM users WHERE username = 'petugas1';
