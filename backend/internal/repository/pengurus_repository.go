package repository

import (
	"database/sql"

	"github.com/yourusername/siazisah/internal/models"
)

type PengurusRepository struct {
	DB *sql.DB
}

func NewPengurusRepository(db *sql.DB) *PengurusRepository {
	return &PengurusRepository{DB: db}
}

// Pengurus Masjid methods
func (r *PengurusRepository) SavePengurusMasjid(pengurus *models.PengurusMasjid) error {
	// Check if exists
	var existingID int
	checkQuery := `SELECT id FROM pengurus_masjid WHERE masjid_id=? AND jabatan=?`
	err := r.DB.QueryRow(checkQuery, pengurus.MasjidID, pengurus.Jabatan).Scan(&existingID)
	
	if err == sql.ErrNoRows {
		// Insert new
		insertQuery := `INSERT INTO pengurus_masjid (masjid_id, nama, jabatan, telepon, alamat) 
		                VALUES (?, ?, ?, ?, ?)`
		result, err := r.DB.Exec(insertQuery, pengurus.MasjidID, pengurus.Nama, 
			pengurus.Jabatan, pengurus.Telepon, pengurus.Alamat)
		if err != nil {
			return err
		}
		id, _ := result.LastInsertId()
		pengurus.ID = int(id)
		return nil
	} else if err != nil {
		return err
	}
	
	// Update existing
	updateQuery := `UPDATE pengurus_masjid SET nama=?, telepon=?, alamat=? 
	                WHERE id=?`
	_, err = r.DB.Exec(updateQuery, pengurus.Nama, pengurus.Telepon, 
		pengurus.Alamat, existingID)
	pengurus.ID = existingID
	return err
}

func (r *PengurusRepository) GetPengurusMasjidByMasjidID(masjidID int) ([]models.PengurusMasjid, error) {
	query := `SELECT id, masjid_id, nama, jabatan, telepon, alamat, created_at, updated_at 
	          FROM pengurus_masjid WHERE masjid_id=? ORDER BY 
	          CASE jabatan 
	            WHEN 'Ketua' THEN 1 
	            WHEN 'Sekretaris' THEN 2 
	            ELSE 3 
	          END`
	
	rows, err := r.DB.Query(query, masjidID)
	if err != nil {
		return []models.PengurusMasjid{}, nil
	}
	defer rows.Close()

	var pengurusList []models.PengurusMasjid
	for rows.Next() {
		var p models.PengurusMasjid
		err := rows.Scan(&p.ID, &p.MasjidID, &p.Nama, &p.Jabatan, &p.Telepon, 
			&p.Alamat, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			continue
		}
		pengurusList = append(pengurusList, p)
	}
	return pengurusList, nil
}

// Pengurus Zakat (UPZ) methods
func (r *PengurusRepository) SavePengurusZakat(pengurus *models.PengurusZakat) error {
	// Check if exists
	var existingID int
	checkQuery := `SELECT id FROM pengurus_zakat WHERE masjid_id=? AND jabatan=?`
	err := r.DB.QueryRow(checkQuery, pengurus.MasjidID, pengurus.Jabatan).Scan(&existingID)
	
	if err == sql.ErrNoRows {
		// Insert new
		insertQuery := `INSERT INTO pengurus_zakat (masjid_id, nama, jabatan, telepon, alamat) 
		                VALUES (?, ?, ?, ?, ?)`
		result, err := r.DB.Exec(insertQuery, pengurus.MasjidID, pengurus.Nama, 
			pengurus.Jabatan, pengurus.Telepon, pengurus.Alamat)
		if err != nil {
			return err
		}
		id, _ := result.LastInsertId()
		pengurus.ID = int(id)
		return nil
	} else if err != nil {
		return err
	}
	
	// Update existing
	updateQuery := `UPDATE pengurus_zakat SET nama=?, telepon=?, alamat=? 
	                WHERE id=?`
	_, err = r.DB.Exec(updateQuery, pengurus.Nama, pengurus.Telepon, 
		pengurus.Alamat, existingID)
	pengurus.ID = existingID
	return err
}

func (r *PengurusRepository) GetPengurusZakatByMasjidID(masjidID int) ([]models.PengurusZakat, error) {
	query := `SELECT id, masjid_id, nama, jabatan, telepon, alamat, created_at, updated_at 
	          FROM pengurus_zakat WHERE masjid_id=? ORDER BY 
	          CASE jabatan 
	            WHEN 'Ketua UPZ' THEN 1 
	            WHEN 'Sekretaris UPZ' THEN 2
	            WHEN 'Anggota UPZ' THEN 3 
	            ELSE 4 
	          END, nama ASC`
	
	rows, err := r.DB.Query(query, masjidID)
	if err != nil {
		return []models.PengurusZakat{}, nil
	}
	defer rows.Close()

	var pengurusList []models.PengurusZakat
	for rows.Next() {
		var p models.PengurusZakat
		err := rows.Scan(&p.ID, &p.MasjidID, &p.Nama, &p.Jabatan, &p.Telepon, 
			&p.Alamat, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			continue
		}
		pengurusList = append(pengurusList, p)
	}
	return pengurusList, nil
}

func (r *PengurusRepository) DeletePengurusZakat(id int) error {
	query := `DELETE FROM pengurus_zakat WHERE id=?`
	_, err := r.DB.Exec(query, id)
	return err
}
