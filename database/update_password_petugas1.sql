-- Update password untuk petugas1
-- Password: andra123
UPDATE users 
SET password = '$2a$14$x0.8KZy18CI/HNV1Vv5eUuxUlPRvmuvv4H/xXZ8KM9tZ.s94sydYi' 
WHERE username = 'petugas1';

-- Cek hasil update
SELECT id, username, role, is_active FROM users WHERE username = 'petugas1';
