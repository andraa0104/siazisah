-- Add missing columns to masjid table if they don't exist
ALTER TABLE masjid
ADD COLUMN IF NOT EXISTS kode_pos VARCHAR(10) NULL AFTER provinsi;

-- Verify columns exist
SHOW COLUMNS FROM masjid WHERE Field IN ('kode_pos', 'desa', 'kecamatan', 'kabupaten', 'provinsi', 'telepon');

-- Check if any data exists
SELECT COUNT(*) as total_masjid FROM masjid;
SELECT id, nama, kode_pos FROM masjid LIMIT 5;
