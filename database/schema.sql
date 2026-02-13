-- Database Schema untuk Sistem Informasi Zakat
-- Database: siazisah

CREATE DATABASE IF NOT EXISTS siazisah;
USE siazisah;

-- Table: users (untuk superadmin dan petugas)
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(150) NOT NULL,
    role ENUM('superadmin', 'petugas') NOT NULL,
    masjid_id INT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_role (role),
    INDEX idx_masjid (masjid_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Table: masjid (data masjid/langgar)
CREATE TABLE masjid (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(200) NOT NULL,
    alamat TEXT NOT NULL,
    desa VARCHAR(100) DEFAULT 'Purwajaya',
    kecamatan VARCHAR(100),
    kabupaten VARCHAR(100),
    provinsi VARCHAR(100),
    kode_pos VARCHAR(10),
    telepon VARCHAR(20),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_nama (nama)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Table: pengurus_masjid
CREATE TABLE pengurus_masjid (
    id INT AUTO_INCREMENT PRIMARY KEY,
    masjid_id INT NOT NULL,
    nama VARCHAR(150) NOT NULL,
    jabatan VARCHAR(100) NOT NULL,
    telepon VARCHAR(20),
    alamat TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (masjid_id) REFERENCES masjid(id) ON DELETE CASCADE,
    INDEX idx_masjid (masjid_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Table: pengurus_zakat
CREATE TABLE pengurus_zakat (
    id INT AUTO_INCREMENT PRIMARY KEY,
    masjid_id INT NOT NULL,
    nama VARCHAR(150) NOT NULL,
    jabatan VARCHAR(100) NOT NULL,
    telepon VARCHAR(20),
    alamat TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (masjid_id) REFERENCES masjid(id) ON DELETE CASCADE,
    INDEX idx_masjid (masjid_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Table: kadar_zakat (kadar zakat fitrah berupa uang)
CREATE TABLE kadar_zakat (
    id INT AUTO_INCREMENT PRIMARY KEY,
    masjid_id INT NOT NULL,
    tahun INT NOT NULL,
    kelas ENUM('1', '2', '3') NOT NULL,
    nominal DECIMAL(15,2) NOT NULL,
    keterangan TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (masjid_id) REFERENCES masjid(id) ON DELETE CASCADE,
    INDEX idx_masjid_tahun (masjid_id, tahun),
    UNIQUE KEY unique_masjid_tahun_kelas (masjid_id, tahun, kelas)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Table: jenis_zakat_tersedia (zakat yang disediakan per masjid)
CREATE TABLE jenis_zakat_tersedia (
    id INT AUTO_INCREMENT PRIMARY KEY,
    masjid_id INT NOT NULL,
    jenis_zakat ENUM('fitrah', 'mal', 'infaq') NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (masjid_id) REFERENCES masjid(id) ON DELETE CASCADE,
    UNIQUE KEY unique_masjid_jenis (masjid_id, jenis_zakat),
    INDEX idx_masjid (masjid_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Table: mustahiq (penerima zakat)
CREATE TABLE mustahiq (
    id INT AUTO_INCREMENT PRIMARY KEY,
    masjid_id INT NOT NULL,
    nama VARCHAR(150) NOT NULL,
    jenis_penerima ENUM('fakir', 'miskin', 'amil', 'mualaf', 'riqab', 'gharim', 'fisabilillah', 'ibnu sabil') NOT NULL,
    alamat TEXT NOT NULL,
    lokasi ENUM('dalam_desa', 'luar_desa') DEFAULT 'dalam_desa',
    rt VARCHAR(10),
    telepon VARCHAR(20),
    keterangan TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (masjid_id) REFERENCES masjid(id) ON DELETE CASCADE,
    INDEX idx_masjid (masjid_id),
    INDEX idx_jenis (jenis_penerima)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Table: muzakki (pemberi zakat)
CREATE TABLE muzakki (
    id INT AUTO_INCREMENT PRIMARY KEY,
    masjid_id INT NOT NULL,
    nama VARCHAR(150) NOT NULL,
    alamat TEXT NOT NULL,
    telepon VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (masjid_id) REFERENCES masjid(id) ON DELETE CASCADE,
    INDEX idx_masjid (masjid_id),
    INDEX idx_nama (nama)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Table: transaksi_zakat
CREATE TABLE transaksi_zakat (
    id INT AUTO_INCREMENT PRIMARY KEY,
    masjid_id INT NOT NULL,
    muzakki_id INT NOT NULL,
    jenis_zakat ENUM('fitrah', 'mal', 'infaq') NOT NULL,
    bentuk_zakat ENUM('uang', 'beras') NULL,
    kelas_zakat ENUM('1', '2', '3') NULL,
    jumlah_orang INT DEFAULT 1,
    nominal_per_orang DECIMAL(15,2) DEFAULT 0,
    total_wajib DECIMAL(15,2) DEFAULT 0,
    total_dibayar DECIMAL(15,2) NOT NULL,
    infaq_tambahan DECIMAL(15,2) DEFAULT 0,
    keterangan TEXT,
    tahun INT NOT NULL,
    tanggal_bayar DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (masjid_id) REFERENCES masjid(id) ON DELETE CASCADE,
    FOREIGN KEY (muzakki_id) REFERENCES muzakki(id) ON DELETE CASCADE,
    INDEX idx_masjid (masjid_id),
    INDEX idx_muzakki (muzakki_id),
    INDEX idx_jenis (jenis_zakat),
    INDEX idx_tahun (tahun),
    INDEX idx_tanggal (tanggal_bayar)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Table: distribusi_zakat (penyaluran zakat ke mustahiq)
CREATE TABLE distribusi_zakat (
    id INT AUTO_INCREMENT PRIMARY KEY,
    masjid_id INT NOT NULL,
    mustahiq_id INT NOT NULL,
    jenis_zakat ENUM('fitrah', 'mal', 'infaq') NOT NULL,
    nominal DECIMAL(15,2) NOT NULL,
    tanggal_distribusi DATE NOT NULL,
    keterangan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (masjid_id) REFERENCES masjid(id) ON DELETE CASCADE,
    FOREIGN KEY (mustahiq_id) REFERENCES mustahiq(id) ON DELETE CASCADE,
    INDEX idx_masjid (masjid_id),
    INDEX idx_mustahiq (mustahiq_id),
    INDEX idx_tanggal (tanggal_distribusi)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert default superadmin
-- Password: admin123 (hashed dengan bcrypt)
INSERT INTO users (username, email, password, full_name, role) VALUES 
('superadmin', 'superadmin@siazisah.com', '$2a$10$YourHashedPasswordHere', 'Super Administrator', 'superadmin');
