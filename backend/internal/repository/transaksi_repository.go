package repository

import (
	"database/sql"
	"github.com/yourusername/siazisah/internal/models"
)

type TransaksiZakatRepository struct {
	DB *sql.DB
}

func NewTransaksiZakatRepository(db *sql.DB) *TransaksiZakatRepository {
	return &TransaksiZakatRepository{DB: db}
}

func (r *TransaksiZakatRepository) Create(transaksi *models.TransaksiZakat) error {
	query := `INSERT INTO transaksi_zakat (masjid_id, muzakki_id, jenis_zakat, bentuk_zakat, jenis_harta, 
			  nominal_harta, persentase_zakat, kelas_zakat, jumlah_orang, nominal_per_orang, total_wajib, 
			  total_dibayar, infaq_tambahan, keterangan, tahun, tanggal_bayar) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := r.DB.Exec(query, transaksi.MasjidID, transaksi.MuzakkiID, transaksi.JenisZakat,
		transaksi.BentukZakat, transaksi.JenisHarta, transaksi.NominalHarta, transaksi.PersentaseZakat,
		transaksi.KelasZakat, transaksi.JumlahOrang, transaksi.NominalPerOrang,
		transaksi.TotalWajib, transaksi.TotalDibayar, transaksi.InfaqTambahan, transaksi.Keterangan,
		transaksi.Tahun, transaksi.TanggalBayar)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	transaksi.ID = int(id)
	return nil
}

func (r *TransaksiZakatRepository) FindByID(id int) (*models.TransaksiZakat, error) {
	transaksi := &models.TransaksiZakat{}
	query := `SELECT t.id, t.masjid_id, t.muzakki_id, t.jenis_zakat, t.bentuk_zakat, t.jenis_harta,
			  t.nominal_harta, t.persentase_zakat, t.kelas_zakat, t.jumlah_orang, t.nominal_per_orang, 
			  t.total_wajib, t.total_dibayar, t.infaq_tambahan, t.keterangan, t.tahun, t.tanggal_bayar, 
			  t.created_at, t.updated_at, m.nama, m.alamat
			  FROM transaksi_zakat t
			  JOIN muzakki m ON t.muzakki_id = m.id
			  WHERE t.id = ?`
	err := r.DB.QueryRow(query, id).Scan(
		&transaksi.ID, &transaksi.MasjidID, &transaksi.MuzakkiID, &transaksi.JenisZakat,
		&transaksi.BentukZakat, &transaksi.JenisHarta, &transaksi.NominalHarta, &transaksi.PersentaseZakat,
		&transaksi.KelasZakat, &transaksi.JumlahOrang, &transaksi.NominalPerOrang,
		&transaksi.TotalWajib, &transaksi.TotalDibayar, &transaksi.InfaqTambahan, &transaksi.Keterangan,
		&transaksi.Tahun, &transaksi.TanggalBayar, &transaksi.CreatedAt, &transaksi.UpdatedAt,
		&transaksi.MuzakkiNama, &transaksi.MuzakkiAlamat,
	)
	if err != nil {
		return nil, err
	}
	return transaksi, nil
}

func (r *TransaksiZakatRepository) GetByMasjidID(masjidID int) ([]models.TransaksiZakat, error) {
	query := `SELECT t.id, t.masjid_id, t.muzakki_id, t.jenis_zakat, t.bentuk_zakat, t.jenis_harta,
			  t.nominal_harta, t.persentase_zakat, t.kelas_zakat, t.jumlah_orang, t.nominal_per_orang, 
			  t.total_wajib, t.total_dibayar, t.infaq_tambahan, t.keterangan, t.tahun, t.tanggal_bayar, 
			  t.created_at, t.updated_at, m.nama, m.alamat
			  FROM transaksi_zakat t
			  JOIN muzakki m ON t.muzakki_id = m.id
			  WHERE t.masjid_id = ? 
			  ORDER BY t.tanggal_bayar DESC, t.created_at DESC`
	rows, err := r.DB.Query(query, masjidID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transaksis []models.TransaksiZakat
	for rows.Next() {
		var transaksi models.TransaksiZakat
		err := rows.Scan(&transaksi.ID, &transaksi.MasjidID, &transaksi.MuzakkiID, &transaksi.JenisZakat,
			&transaksi.BentukZakat, &transaksi.JenisHarta, &transaksi.NominalHarta, &transaksi.PersentaseZakat,
			&transaksi.KelasZakat, &transaksi.JumlahOrang, &transaksi.NominalPerOrang,
			&transaksi.TotalWajib, &transaksi.TotalDibayar, &transaksi.InfaqTambahan, &transaksi.Keterangan,
			&transaksi.Tahun, &transaksi.TanggalBayar, &transaksi.CreatedAt, &transaksi.UpdatedAt,
			&transaksi.MuzakkiNama, &transaksi.MuzakkiAlamat)
		if err != nil {
			continue
		}
		transaksis = append(transaksis, transaksi)
	}
	return transaksis, nil
}

func (r *TransaksiZakatRepository) GetByMasjidIDAndYear(masjidID, tahun int) ([]models.TransaksiZakat, error) {
	query := `SELECT t.id, t.masjid_id, t.muzakki_id, t.jenis_zakat, t.bentuk_zakat, t.jenis_harta,
			  t.nominal_harta, t.persentase_zakat, t.kelas_zakat, t.jumlah_orang, t.nominal_per_orang, 
			  t.total_wajib, t.total_dibayar, t.infaq_tambahan, t.keterangan, t.tahun, t.tanggal_bayar, 
			  t.created_at, t.updated_at, m.nama, m.alamat
			  FROM transaksi_zakat t
			  JOIN muzakki m ON t.muzakki_id = m.id
			  WHERE t.masjid_id = ? AND t.tahun = ?
			  ORDER BY t.tanggal_bayar DESC`
	rows, err := r.DB.Query(query, masjidID, tahun)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transaksis []models.TransaksiZakat
	for rows.Next() {
		var transaksi models.TransaksiZakat
		err := rows.Scan(&transaksi.ID, &transaksi.MasjidID, &transaksi.MuzakkiID, &transaksi.JenisZakat,
			&transaksi.BentukZakat, &transaksi.JenisHarta, &transaksi.NominalHarta, &transaksi.PersentaseZakat,
			&transaksi.KelasZakat, &transaksi.JumlahOrang, &transaksi.NominalPerOrang,
			&transaksi.TotalWajib, &transaksi.TotalDibayar, &transaksi.InfaqTambahan, &transaksi.Keterangan,
			&transaksi.Tahun, &transaksi.TanggalBayar, &transaksi.CreatedAt, &transaksi.UpdatedAt,
			&transaksi.MuzakkiNama, &transaksi.MuzakkiAlamat)
		if err != nil {
			continue
		}
		transaksis = append(transaksis, transaksi)
	}
	return transaksis, nil
}

func (r *TransaksiZakatRepository) Delete(id int) error {
	query := `DELETE FROM transaksi_zakat WHERE id=?`
	_, err := r.DB.Exec(query, id)
	return err
}
