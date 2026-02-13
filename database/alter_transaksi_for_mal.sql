-- Alter table transaksi_zakat untuk support zakat mal
ALTER TABLE transaksi_zakat 
ADD COLUMN jenis_harta VARCHAR(50) NULL AFTER bentuk_zakat,
ADD COLUMN nominal_harta DECIMAL(15,2) NULL AFTER jenis_harta,
ADD COLUMN persentase_zakat DECIMAL(5,2) NULL AFTER nominal_harta;

-- Update keterangan kolom
-- jenis_harta: emas, perak, uang, perdagangan, pertanian, peternakan
-- nominal_harta: total harta yang dizakati (untuk zakat mal)
-- persentase_zakat: 2.5, 5, atau 10 (untuk zakat mal)
