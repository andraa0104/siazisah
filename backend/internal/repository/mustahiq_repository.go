package repository

import (
	"database/sql"
	"github.com/yourusername/siazisah/internal/models"
)

type MustahiqRepository struct {
	DB *sql.DB
}

func NewMustahiqRepository(db *sql.DB) *MustahiqRepository {
	return &MustahiqRepository{DB: db}
}

func (r *MustahiqRepository) Create(mustahiq *models.Mustahiq) error {
	query := `INSERT INTO mustahiq (masjid_id, nama, jenis_penerima, alamat, lokasi, rt, telepon, keterangan, is_active) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := r.DB.Exec(query, mustahiq.MasjidID, mustahiq.Nama, mustahiq.JenisPenerima, 
		mustahiq.Alamat, mustahiq.Lokasi, mustahiq.RT, mustahiq.Telepon, mustahiq.Keterangan, mustahiq.IsActive)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	mustahiq.ID = int(id)
	return nil
}

func (r *MustahiqRepository) FindByID(id int) (*models.Mustahiq, error) {
	mustahiq := &models.Mustahiq{}
	query := `SELECT id, masjid_id, nama, jenis_penerima, alamat, lokasi, rt, telepon, keterangan, is_active, created_at, updated_at 
			  FROM mustahiq WHERE id = ?`
	err := r.DB.QueryRow(query, id).Scan(
		&mustahiq.ID, &mustahiq.MasjidID, &mustahiq.Nama, &mustahiq.JenisPenerima,
		&mustahiq.Alamat, &mustahiq.Lokasi, &mustahiq.RT, &mustahiq.Telepon,
		&mustahiq.Keterangan, &mustahiq.IsActive, &mustahiq.CreatedAt, &mustahiq.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return mustahiq, nil
}

func (r *MustahiqRepository) GetByMasjidID(masjidID int) ([]models.Mustahiq, error) {
	query := `SELECT id, masjid_id, nama, jenis_penerima, alamat, lokasi, rt, telepon, keterangan, is_active, created_at, updated_at 
			  FROM mustahiq WHERE masjid_id = ? ORDER BY nama ASC`
	rows, err := r.DB.Query(query, masjidID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mustahiqs []models.Mustahiq
	for rows.Next() {
		var mustahiq models.Mustahiq
		err := rows.Scan(&mustahiq.ID, &mustahiq.MasjidID, &mustahiq.Nama, &mustahiq.JenisPenerima,
			&mustahiq.Alamat, &mustahiq.Lokasi, &mustahiq.RT, &mustahiq.Telepon,
			&mustahiq.Keterangan, &mustahiq.IsActive, &mustahiq.CreatedAt, &mustahiq.UpdatedAt)
		if err != nil {
			continue
		}
		mustahiqs = append(mustahiqs, mustahiq)
	}
	return mustahiqs, nil
}

func (r *MustahiqRepository) Update(mustahiq *models.Mustahiq) error {
	query := `UPDATE mustahiq SET nama=?, jenis_penerima=?, alamat=?, lokasi=?, rt=?, telepon=?, keterangan=?, is_active=? WHERE id=?`
	_, err := r.DB.Exec(query, mustahiq.Nama, mustahiq.JenisPenerima, mustahiq.Alamat, 
		mustahiq.Lokasi, mustahiq.RT, mustahiq.Telepon, mustahiq.Keterangan, mustahiq.IsActive, mustahiq.ID)
	return err
}

func (r *MustahiqRepository) Delete(id int) error {
	query := `DELETE FROM mustahiq WHERE id=?`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *MustahiqRepository) FindDuplicate(nama, lokasi, rt string) (*struct {
	MasjidNama string
}, error) {
	var result struct {
		MasjidNama string
	}
	query := `SELECT m.nama FROM mustahiq mu JOIN masjid m ON mu.masjid_id = m.id WHERE mu.nama = ? AND mu.lokasi = ? AND mu.rt = ? LIMIT 1`
	err := r.DB.QueryRow(query, nama, lokasi, rt).Scan(&result.MasjidNama)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
}
