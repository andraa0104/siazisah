-- Check transaksi vs muzakki join issue
-- Cek masjid untuk Al-Ikhlas (biasanya ID 1 atau 2)

-- 1. Lihat berapa total transaksi untuk setiap masjid
SELECT masjid_id, COUNT(*) as total_transaksi
FROM transaksi_zakat
GROUP BY masjid_id;

-- 2. Lihat transaksi yang bisa JOIN dengan muzakki (valid)
SELECT COUNT(*) as valid_transaksi
FROM transaksi_zakat t
JOIN muzakki m ON t.muzakki_id = m.id
WHERE t.masjid_id = 1;

-- 3. Lihat transaksi yang TIDAK bisa JOIN dengan muzakki (orphaned)
SELECT t.id, t.masjid_id, t.muzakki_id, m.id as muzakki_exists
FROM transaksi_zakat t
LEFT JOIN muzakki m ON t.muzakki_id = m.id
WHERE t.masjid_id = 1 AND m.id IS NULL;

-- 4. Lihat semua transaksi untuk masjid 1 dengan detail muzakki
SELECT t.id, t.masjid_id, t.muzakki_id, m.id as muzakki_id_check, m.nama as muzakki_nama,
       t.jenis_zakat, t.total_dibayar, t.tanggal_bayar
FROM transaksi_zakat t
LEFT JOIN muzakki m ON t.muzakki_id = m.id AND t.masjid_id = m.masjid_id
WHERE t.masjid_id = 1
ORDER BY t.created_at DESC;

-- 5. Check muzakki table
SELECT id, masjid_id, nama FROM muzakki WHERE masjid_id = 1;
