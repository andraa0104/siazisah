-- Add missing columns for Fidyah and Zakat Fitrah Beras
ALTER TABLE transaksi_zakat 
ADD COLUMN jumlah_hari_fidyah INT NULL AFTER jumlah_orang,
ADD COLUMN standar_beras_per_jiwa DECIMAL(5,2) NULL AFTER jumlah_hari_fidyah,
ADD COLUMN kg_beras_dibayar DECIMAL(10,2) NULL AFTER standar_beras_per_jiwa,
ADD COLUMN harga_beras_per_kg DECIMAL(15,2) NULL AFTER kg_beras_dibayar;

-- Description:
-- jumlah_hari_fidyah: untuk menghitung jumlah hari fidyah (denda puasa)
-- standar_beras_per_jiwa: standar beras per jiwa (default 2.5 kg)
-- kg_beras_dibayar: jumlah kg beras yang dibayarkan (untuk fitrah beras)
-- harga_beras_per_kg: harga beras per kg (untuk konversi ke rupiah)
