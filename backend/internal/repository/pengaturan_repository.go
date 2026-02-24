package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/yourusername/siazisah/internal/models"
)

type PengaturanRepository struct {
	DB *sql.DB
}

func NewPengaturanRepository(db *sql.DB) *PengaturanRepository {
	repo := &PengaturanRepository{DB: db}
	repo.ensureTable()
	return repo
}

func (r *PengaturanRepository) ensureTable() {
	query := `CREATE TABLE IF NOT EXISTS pengaturan_zakat (
		id INT AUTO_INCREMENT PRIMARY KEY,
		masjid_id INT NOT NULL UNIQUE,
		kelas1 DECIMAL(15,2) NOT NULL DEFAULT 35000,
		kelas2 DECIMAL(15,2) NOT NULL DEFAULT 45000,
		kelas3 DECIMAL(15,2) NOT NULL DEFAULT 55000,
		fitrah_beras_per_jiwa DECIMAL(6,2) NOT NULL DEFAULT 2.5,
		fidyah_per_hari DECIMAL(15,2) NOT NULL DEFAULT 30000,
		fidyah_beras_per_hari DECIMAL(6,2) NOT NULL DEFAULT 0.6,
		mal_rates_json TEXT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (masjid_id) REFERENCES masjid(id) ON DELETE CASCADE,
		INDEX idx_masjid (masjid_id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`
	_, _ = r.DB.Exec(query)
	_, _ = r.DB.Exec(`ALTER TABLE pengaturan_zakat ADD COLUMN mal_rates_json TEXT NULL`)
}

func defaultMalRates() map[string]float64 {
	return map[string]float64{}
}

func defaultPengaturan(masjidID int) *models.PengaturanZakat {
	return &models.PengaturanZakat{
		MasjidID:           masjidID,
		Kelas1:             35000,
		Kelas2:             45000,
		Kelas3:             55000,
		FitrahBerasPerJiwa: 2.5,
		FidyahPerHari:      30000,
		FidyahBerasPerHari: 0.6,
		MalRates:           defaultMalRates(),
	}
}

func (r *PengaturanRepository) GetByMasjidID(masjidID int) (*models.PengaturanZakat, error) {
	item := defaultPengaturan(masjidID)
	query := `SELECT masjid_id, kelas1, kelas2, kelas3, fitrah_beras_per_jiwa, fidyah_per_hari, fidyah_beras_per_hari, COALESCE(mal_rates_json, '')
			  FROM pengaturan_zakat
			  WHERE masjid_id = ?
			  LIMIT 1`
	var malRatesJSON string
	err := r.DB.QueryRow(query, masjidID).Scan(
		&item.MasjidID,
		&item.Kelas1,
		&item.Kelas2,
		&item.Kelas3,
		&item.FitrahBerasPerJiwa,
		&item.FidyahPerHari,
		&item.FidyahBerasPerHari,
		&malRatesJSON,
	)
	if err == sql.ErrNoRows {
		return item, nil
	}
	if err != nil {
		return nil, err
	}

	mergedRates := defaultMalRates()
	if malRatesJSON != "" {
		parsed := map[string]float64{}
		if err := json.Unmarshal([]byte(malRatesJSON), &parsed); err == nil {
			for key, value := range parsed {
				if value > 0 {
					mergedRates[key] = value
				}
			}
		}
	}
	item.MalRates = mergedRates

	return item, nil
}

func (r *PengaturanRepository) Upsert(data *models.PengaturanZakat) error {
	malRatesJSON, _ := json.Marshal(data.MalRates)
	query := `INSERT INTO pengaturan_zakat
			  (masjid_id, kelas1, kelas2, kelas3, fitrah_beras_per_jiwa, fidyah_per_hari, fidyah_beras_per_hari, mal_rates_json)
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?)
			  ON DUPLICATE KEY UPDATE
			  kelas1 = VALUES(kelas1),
			  kelas2 = VALUES(kelas2),
			  kelas3 = VALUES(kelas3),
			  fitrah_beras_per_jiwa = VALUES(fitrah_beras_per_jiwa),
			  fidyah_per_hari = VALUES(fidyah_per_hari),
			  fidyah_beras_per_hari = VALUES(fidyah_beras_per_hari),
			  mal_rates_json = VALUES(mal_rates_json)`
	_, err := r.DB.Exec(query,
		data.MasjidID,
		data.Kelas1,
		data.Kelas2,
		data.Kelas3,
		data.FitrahBerasPerJiwa,
		data.FidyahPerHari,
		data.FidyahBerasPerHari,
		string(malRatesJSON),
	)
	return err
}
