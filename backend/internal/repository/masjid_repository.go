package repository

import (
	"database/sql"

	"github.com/yourusername/siazisah/internal/models"
)

type MasjidRepository struct {
	DB *sql.DB
}

func NewMasjidRepository(db *sql.DB) *MasjidRepository {
	return &MasjidRepository{DB: db}
}

func (r *MasjidRepository) Create(masjid *models.Masjid) error {
	query := `INSERT INTO masjid (nama, alamat, telepon, is_active) 
			  VALUES (?, ?, ?, ?)`
	result, err := r.DB.Exec(query, masjid.Nama, masjid.Alamat,
		masjid.Telepon, masjid.IsActive)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	masjid.ID = int(id)
	return nil
}

func (r *MasjidRepository) FindByID(id int) (*models.Masjid, error) {
	masjid := &models.Masjid{}
	query := `SELECT id, nama, alamat, telepon, is_active, created_at, updated_at 
			  FROM masjid WHERE id = ?`
	err := r.DB.QueryRow(query, id).Scan(
		&masjid.ID, &masjid.Nama, &masjid.Alamat, &masjid.Telepon,
		&masjid.IsActive, &masjid.CreatedAt, &masjid.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return masjid, nil
}

func (r *MasjidRepository) GetAll() ([]models.Masjid, error) {
	query := `SELECT id, nama, alamat, telepon, is_active, created_at, updated_at 
			  FROM masjid ORDER BY nama ASC`
	rows, err := r.DB.Query(query)
	if err != nil {
		return []models.Masjid{}, nil
	}
	defer rows.Close()

	var masjids []models.Masjid
	for rows.Next() {
		var masjid models.Masjid
		err := rows.Scan(&masjid.ID, &masjid.Nama, &masjid.Alamat, &masjid.Telepon,
			&masjid.IsActive, &masjid.CreatedAt, &masjid.UpdatedAt)
		if err != nil {
			continue
		}
		masjids = append(masjids, masjid)
	}
	return masjids, nil
}

func (r *MasjidRepository) Update(masjid *models.Masjid) error {
	query := `UPDATE masjid SET nama=?, alamat=?, telepon=?, is_active=? WHERE id=?`
	_, err := r.DB.Exec(query, masjid.Nama, masjid.Alamat, masjid.Telepon, masjid.IsActive, masjid.ID)
	return err
}

func (r *MasjidRepository) Delete(id int) error {
	query := `DELETE FROM masjid WHERE id=?`
	_, err := r.DB.Exec(query, id)
	return err
}
