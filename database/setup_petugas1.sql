-- Cek apakah petugas1 sudah ada
SELECT * FROM users WHERE username = 'petugas1';

-- Jika belum ada, insert user petugas1
-- Password: andra123
INSERT INTO users (username, email, password, full_name, role, masjid_id, is_active, created_at, updated_at)
VALUES (
    'petugas1',
    'petugas1@example.com',
    '$2a$14$x0.8KZy18CI/HNV1Vv5eUuxUlPRvmuvv4H/xXZ8KM9tZ.s94sydYi',
    'Petugas 1',
    'petugas',
    1,
    1,
    NOW(),
    NOW()
);

-- Jika sudah ada, update password
UPDATE users 
SET password = '$2a$14$x0.8KZy18CI/HNV1Vv5eUuxUlPRvmuvv4H/xXZ8KM9tZ.s94sydYi' 
WHERE username = 'petugas1';
