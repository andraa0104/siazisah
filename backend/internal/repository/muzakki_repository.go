package repository

import (
	"database/sql"
	"github.com/yourusername/siazisah/internal/models"
)

type MuzakkiRepository struct {
	DB *sql.DB
}

func NewMuzakkiRepository(db *sql.DB) *MuzakkiRepository {
	return &MuzakkiRepository{DB: db}
}

func (r *MuzakkiRepository) Create(muzakki *models.Muzakki) error {
	query := `INSERT INTO muzakki (masjid_id, nama, alamat, telepon) VALUES (?, ?, ?, ?)`
	result, err := r.DB.Exec(query, muzakki.MasjidID, muzakki.Nama, muzakki.Alamat, muzakki.Telepon)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	muzakki.ID = int(id)
	return nil
}

func (r *MuzakkiRepository) FindByID(id int) (*models.Muzakki, error) {
	muzakki := &models.Muzakki{}
	query := `SELECT id, masjid_id, nama, alamat, telepon, created_at, updated_at FROM muzakki WHERE id = ?`
	err := r.DB.QueryRow(query, id).Scan(&muzakki.ID, &muzakki.MasjidID, &muzakki.Nama, &muzakki.Alamat, &muzakki.Telepon, &muzakki.CreatedAt, &muzakki.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return muzakki, nil
}

func (r *MuzakkiRepository) GetByMasjidID(masjidID int) ([]models.Muzakki, error) {
	query := `SELECT id, masjid_id, nama, alamat, telepon, created_at, updated_at FROM muzakki WHERE masjid_id = ? ORDER BY nama ASC`
	rows, err := r.DB.Query(query, masjidID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var muzakkis []models.Muzakki
	for rows.Next() {
		var muzakki models.Muzakki
		err := rows.Scan(&muzakki.ID, &muzakki.MasjidID, &muzakki.Nama, &muzakki.Alamat, &muzakki.Telepon, &muzakki.CreatedAt, &muzakki.UpdatedAt)
		if err != nil {
			continue
		}
		muzakkis = append(muzakkis, muzakki)
	}
	return muzakkis, nil
}

func (r *MuzakkiRepository) Update(muzakki *models.Muzakki) error {
	query := `UPDATE muzakki SET nama=?, alamat=?, telepon=? WHERE id=?`
	_, err := r.DB.Exec(query, muzakki.Nama, muzakki.Alamat, muzakki.Telepon, muzakki.ID)
	return err
}

func (r *MuzakkiRepository) Delete(id int) error {
	query := `DELETE FROM muzakki WHERE id=?`
	_, err := r.DB.Exec(query, id)
	return err
}
