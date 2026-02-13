-- Sample Data untuk Testing SI-AZISAH
USE siazisah;

-- Insert sample masjid
INSERT INTO masjid (nama, alamat, desa, kecamatan, kabupaten, provinsi, telepon) VALUES
('Masjid Al-Ikhlas', 'Jl. Raya Purwajaya No. 123', 'Purwajaya', 'Kecamatan ABC', 'Kabupaten XYZ', 'Jawa Barat', '081234567890'),
('Langgar An-Nur', 'Jl. Masjid No. 45', 'Purwajaya', 'Kecamatan ABC', 'Kabupaten XYZ', 'Jawa Barat', '081234567891'),
('Masjid Ar-Rahman', 'Jl. Pesantren No. 78', 'Purwajaya', 'Kecamatan ABC', 'Kabupaten XYZ', 'Jawa Barat', '081234567892');

-- Insert sample petugas (password: petugas123)
-- Hash ini adalah contoh, generate ulang dengan cmd/generate_password.go
INSERT INTO users (username, email, password, full_name, role, masjid_id) VALUES
('petugas1', 'petugas1@siazisah.com', '$2a$14$YourHashedPasswordHere', 'Ahmad Petugas', 'petugas', 1),
('petugas2', 'petugas2@siazisah.com', '$2a$14$YourHashedPasswordHere', 'Budi Petugas', 'petugas', 2),
('petugas3', 'petugas3@siazisah.com', '$2a$14$YourHashedPasswordHere', 'Citra Petugas', 'petugas', 3);

-- Insert sample pengurus masjid
INSERT INTO pengurus_masjid (masjid_id, nama, jabatan, telepon) VALUES
(1, 'H. Abdullah', 'Ketua Takmir', '081234567893'),
(1, 'H. Usman', 'Sekretaris', '081234567894'),
(1, 'H. Ali', 'Bendahara', '081234567895');

-- Insert sample pengurus zakat
INSERT INTO pengurus_zakat (masjid_id, nama, jabatan, telepon) VALUES
(1, 'Ustadz Mahmud', 'Ketua Amil', '081234567896'),
(1, 'Ustadz Yusuf', 'Sekretaris Amil', '081234567897'),
(1, 'Ustadz Ibrahim', 'Bendahara Amil', '081234567898');

-- Insert sample kadar zakat (tahun 2024)
INSERT INTO kadar_zakat (masjid_id, tahun, kelas, nominal, keterangan) VALUES
(1, 2024, '1', 35000, 'Kelas 1 - Standar'),
(1, 2024, '2', 45000, 'Kelas 2 - Menengah'),
(1, 2024, '3', 55000, 'Kelas 3 - Tinggi'),
(2, 2024, '1', 35000, 'Kelas 1 - Standar'),
(2, 2024, '2', 45000, 'Kelas 2 - Menengah'),
(2, 2024, '3', 55000, 'Kelas 3 - Tinggi');

-- Insert sample jenis zakat tersedia
INSERT INTO jenis_zakat_tersedia (masjid_id, jenis_zakat) VALUES
(1, 'fitrah'),
(1, 'mal'),
(1, 'infaq'),
(2, 'fitrah'),
(2, 'infaq');

-- Insert sample mustahiq (penerima zakat)
INSERT INTO mustahiq (masjid_id, nama, jenis_penerima, alamat, lokasi, rt, telepon) VALUES
(1, 'Pak Slamet', 'fakir', 'Jl. Kampung No. 12', 'dalam_desa', '01', '081234567899'),
(1, 'Bu Siti', 'miskin', 'Jl. Kampung No. 34', 'dalam_desa', '02', '081234567800'),
(1, 'Pak Joko', 'fakir', 'Jl. Kampung No. 56', 'dalam_desa', '03', '081234567801'),
(1, 'Bu Aminah', 'miskin', 'Jl. Luar Desa No. 78', 'luar_desa', '5', '081234567802'),
(1, 'Pak Hasan', 'gharim', 'Jl. Kampung No. 90', 'dalam_desa', '05', '081234567803');

-- Insert sample muzakki (pemberi zakat)
INSERT INTO muzakki (masjid_id, nama, alamat, telepon) VALUES
(1, 'H. Muhammad', 'Jl. Makmur No. 11', '081234567804'),
(1, 'Hj. Fatimah', 'Jl. Sejahtera No. 22', '081234567805'),
(1, 'Bapak Ridwan', 'Jl. Bahagia No. 33', '081234567806'),
(1, 'Ibu Khadijah', 'Jl. Tentram No. 44', '081234567807'),
(1, 'Pak Umar', 'Jl. Damai No. 55', '081234567808');

-- Insert sample transaksi zakat
INSERT INTO transaksi_zakat (masjid_id, muzakki_id, jenis_zakat, bentuk_zakat, kelas_zakat, jumlah_orang, nominal_per_orang, total_wajib, total_dibayar, infaq_tambahan, tahun, tanggal_bayar) VALUES
(1, 1, 'fitrah', 'uang', '2', 4, 45000, 180000, 200000, 20000, 2024, '2024-04-10'),
(1, 2, 'fitrah', 'uang', '1', 3, 35000, 105000, 105000, 0, 2024, '2024-04-11'),
(1, 3, 'fitrah', 'uang', '3', 5, 55000, 275000, 300000, 25000, 2024, '2024-04-12'),
(1, 4, 'mal', NULL, NULL, 1, 0, 0, 500000, 0, 2024, '2024-04-13'),
(1, 5, 'infaq', NULL, NULL, 1, 0, 0, 100000, 0, 2024, '2024-04-14');

-- Note: Untuk password users, generate dengan:
-- cd backend
-- go run cmd/generate_password.go petugas123
-- Lalu update hash di database
