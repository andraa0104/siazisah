-- Migration: Add missing columns to masjid table
USE siazisah;

-- Check current columns
SHOW COLUMNS FROM masjid;

-- Add missing columns if they don't exist
ALTER TABLE masjid ADD COLUMN kode_pos VARCHAR(10) NULL AFTER provinsi;

-- Verify columns were added
SHOW COLUMNS FROM masjid;

-- Check masjid data
SELECT COUNT(*) as total_masjid FROM masjid;
SELECT id, nama, kode_pos FROM masjid LIMIT 10;
