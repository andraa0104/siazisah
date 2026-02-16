-- Add fidyah to jenis_zakat enum in transaksi_zakat and distribusi_zakat tables

-- Update transaksi_zakat table
ALTER TABLE transaksi_zakat 
MODIFY COLUMN jenis_zakat ENUM('fitrah', 'mal', 'fidyah', 'infaq') NOT NULL;

-- Update distribusi_zakat table
ALTER TABLE distribusi_zakat 
MODIFY COLUMN jenis_zakat ENUM('fitrah', 'mal', 'fidyah', 'infaq') NOT NULL;
