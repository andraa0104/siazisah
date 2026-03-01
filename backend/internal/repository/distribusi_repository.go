package repository

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/yourusername/siazisah/internal/models"
)

type DistribusiZakatRepository struct {
	DB *sql.DB
}

func NewDistribusiZakatRepository(db *sql.DB) *DistribusiZakatRepository {
	return &DistribusiZakatRepository{DB: db}
}

var (
	metaBerasRegex    = regexp.MustCompile(`(?i)\[BERAS_KG=([\d.]+)\]`)
	metaModeRegex     = regexp.MustCompile(`(?i)\[MODE=([^\]]+)\]`)
	keteranganKgRegex = regexp.MustCompile(`(?i)(?:Beras|Fidyah Beras):\s*([\d.]+)\s*kg`)
	metaStripRegex    = regexp.MustCompile(`(?i)\[(?:BERAS_KG|MODE)=[^\]]+\]\s*`)
)

func composeDistribusiKeterangan(note string, berasKg float64, mode string) string {
	cleanMode := strings.TrimSpace(mode)
	if cleanMode == "" {
		cleanMode = "manual"
	}
	meta := fmt.Sprintf("[BERAS_KG=%.2f][MODE=%s]", berasKg, cleanMode)
	cleanNote := strings.TrimSpace(note)
	if cleanNote == "" {
		return meta
	}
	return meta + " " + cleanNote
}

func extractBerasKgAndMode(keterangan string) (float64, string) {
	var berasKg float64
	var mode string

	if matches := metaBerasRegex.FindStringSubmatch(keterangan); len(matches) > 1 {
		if val, err := strconv.ParseFloat(matches[1], 64); err == nil {
			berasKg = val
		}
	} else if matches := keteranganKgRegex.FindStringSubmatch(keterangan); len(matches) > 1 {
		if val, err := strconv.ParseFloat(matches[1], 64); err == nil {
			berasKg = val
		}
	}

	if matches := metaModeRegex.FindStringSubmatch(keterangan); len(matches) > 1 {
		mode = strings.TrimSpace(matches[1])
	}

	return berasKg, mode
}

func cleanKeterangan(keterangan string) string {
	clean := metaStripRegex.ReplaceAllString(keterangan, "")
	return strings.TrimSpace(clean)
}

func (r *DistribusiZakatRepository) Create(distribusi *models.DistribusiZakat) error {
	query := `INSERT INTO distribusi_zakat (masjid_id, mustahiq_id, nominal, tanggal_distribusi, keterangan)
			  VALUES (?, ?, ?, ?, ?)`
	payloadKeterangan := composeDistribusiKeterangan(distribusi.Keterangan, distribusi.BerasKg, distribusi.Mode)
	result, err := r.DB.Exec(query,
		distribusi.MasjidID,
		distribusi.MustahiqID,
		distribusi.Nominal,
		distribusi.TanggalDistribusi,
		payloadKeterangan,
	)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	distribusi.ID = int(id)
	return nil
}

func (r *DistribusiZakatRepository) Update(distribusi *models.DistribusiZakat) error {
	query := `UPDATE distribusi_zakat
			  SET mustahiq_id = ?, nominal = ?, tanggal_distribusi = ?, keterangan = ?
			  WHERE id = ? AND masjid_id = ?`
	payloadKeterangan := composeDistribusiKeterangan(distribusi.Keterangan, distribusi.BerasKg, distribusi.Mode)
	_, err := r.DB.Exec(query,
		distribusi.MustahiqID,
		distribusi.Nominal,
		distribusi.TanggalDistribusi,
		payloadKeterangan,
		distribusi.ID,
		distribusi.MasjidID,
	)
	return err
}

func (r *DistribusiZakatRepository) Delete(id int, masjidID int) error {
	query := `DELETE FROM distribusi_zakat WHERE id = ? AND masjid_id = ?`
	_, err := r.DB.Exec(query, id, masjidID)
	return err
}

func (r *DistribusiZakatRepository) GetByMasjidID(masjidID int) ([]models.DistribusiZakat, error) {
	query := `SELECT d.id, d.masjid_id, d.mustahiq_id, d.nominal,
			  DATE_FORMAT(d.tanggal_distribusi, '%Y-%m-%d') AS tanggal_distribusi,
			  DATE_FORMAT(d.created_at, '%H:%i:%s') AS waktu_input,
			  COALESCE(d.keterangan, ''), d.created_at, d.updated_at,
			  COALESCE(m.nama, 'Unknown'), COALESCE(m.jenis_penerima, '')
			  FROM distribusi_zakat d
			  LEFT JOIN mustahiq m ON d.mustahiq_id = m.id
			  WHERE d.masjid_id = ?
			  ORDER BY d.tanggal_distribusi DESC, d.created_at DESC`

	rows, err := r.DB.Query(query, masjidID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.DistribusiZakat
	for rows.Next() {
		var item models.DistribusiZakat
		var rawKeterangan string
		err := rows.Scan(
			&item.ID,
			&item.MasjidID,
			&item.MustahiqID,
			&item.Nominal,
			&item.TanggalDistribusi,
			&item.WaktuInput,
			&rawKeterangan,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.MustahiqNama,
			&item.JenisPenerima,
		)
		if err != nil {
			continue
		}

		item.BerasKg, item.Mode = extractBerasKgAndMode(rawKeterangan)
		item.Keterangan = cleanKeterangan(rawKeterangan)
		result = append(result, item)
	}

	return result, nil
}

func (r *DistribusiZakatRepository) GetByMasjidIDPaginated(masjidID, limit, offset int) ([]models.DistribusiZakat, error) {
	query := `SELECT d.id, d.masjid_id, d.mustahiq_id, d.nominal,
			  DATE_FORMAT(d.tanggal_distribusi, '%Y-%m-%d') AS tanggal_distribusi,
			  DATE_FORMAT(d.created_at, '%H:%i:%s') AS waktu_input,
			  COALESCE(d.keterangan, ''), d.created_at, d.updated_at,
			  COALESCE(m.nama, 'Unknown'), COALESCE(m.jenis_penerima, '')
			  FROM distribusi_zakat d
			  LEFT JOIN mustahiq m ON d.mustahiq_id = m.id
			  WHERE d.masjid_id = ?
			  ORDER BY d.tanggal_distribusi DESC, d.created_at DESC
			  LIMIT ? OFFSET ?`

	rows, err := r.DB.Query(query, masjidID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.DistribusiZakat
	for rows.Next() {
		var item models.DistribusiZakat
		var rawKeterangan string
		err := rows.Scan(
			&item.ID,
			&item.MasjidID,
			&item.MustahiqID,
			&item.Nominal,
			&item.TanggalDistribusi,
			&item.WaktuInput,
			&rawKeterangan,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.MustahiqNama,
			&item.JenisPenerima,
		)
		if err != nil {
			continue
		}

		item.BerasKg, item.Mode = extractBerasKgAndMode(rawKeterangan)
		item.Keterangan = cleanKeterangan(rawKeterangan)
		result = append(result, item)
	}

	return result, nil
}

func (r *DistribusiZakatRepository) CountByMasjidID(masjidID int) (int, error) {
	var total int
	query := `SELECT COUNT(*) FROM distribusi_zakat WHERE masjid_id = ?`
	if err := r.DB.QueryRow(query, masjidID).Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}

func (r *DistribusiZakatRepository) sumBerasFromTransaksi(masjidID int, jenisZakat string) (float64, error) {
	query := `SELECT kg_beras_dibayar, COALESCE(keterangan, '')
			  FROM transaksi_zakat
			  WHERE masjid_id = ? AND jenis_zakat = ? AND bentuk_zakat = 'beras'`
	rows, err := r.DB.Query(query, masjidID, jenisZakat)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	total := 0.0
	for rows.Next() {
		var kg sql.NullFloat64
		var keterangan string
		if err := rows.Scan(&kg, &keterangan); err != nil {
			continue
		}
		if kg.Valid && kg.Float64 > 0 {
			total += kg.Float64
		} else if matches := keteranganKgRegex.FindStringSubmatch(keterangan); len(matches) > 1 {
			if parsed, err := strconv.ParseFloat(matches[1], 64); err == nil {
				total += parsed
			}
		}
	}

	return total, nil
}

func (r *DistribusiZakatRepository) GetInsight(masjidID int) (*models.DistribusiInsight, error) {
	insight := &models.DistribusiInsight{}

	if err := r.DB.QueryRow(`SELECT COUNT(*) FROM mustahiq WHERE masjid_id = ? AND is_active = 1`, masjidID).Scan(&insight.TotalMustahiqAktif); err != nil {
		return nil, err
	}

	r.DB.QueryRow(`SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'fitrah' AND bentuk_zakat = 'uang'`, masjidID).Scan(&insight.TotalFitrahUang)
	r.DB.QueryRow(`SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'fidyah' AND bentuk_zakat = 'uang'`, masjidID).Scan(&insight.TotalFidyahUang)
	r.DB.QueryRow(`SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'mal'`, masjidID).Scan(&insight.TotalZakatMalUang)

	fitrahKg, err := r.sumBerasFromTransaksi(masjidID, "fitrah")
	if err != nil {
		return nil, err
	}
	fidyahKg, err := r.sumBerasFromTransaksi(masjidID, "fidyah")
	if err != nil {
		return nil, err
	}
	insight.TotalFitrahBerasKg = fitrahKg
	insight.TotalFidyahBerasKg = fidyahKg

	insight.TotalPoolBerasKg = insight.TotalFitrahBerasKg + insight.TotalFidyahBerasKg
	insight.TotalPoolUang = insight.TotalFitrahUang + insight.TotalFidyahUang + insight.TotalZakatMalUang

	rows, err := r.DB.Query(`SELECT nominal, COALESCE(keterangan, '') FROM distribusi_zakat WHERE masjid_id = ?`, masjidID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var nominal float64
		var keterangan string
		if err := rows.Scan(&nominal, &keterangan); err != nil {
			continue
		}
		insight.TotalDistribusiUang += nominal
		kg, _ := extractBerasKgAndMode(keterangan)
		insight.TotalDistribusiBerasKg += kg
	}

	insight.SisaUang = insight.TotalPoolUang - insight.TotalDistribusiUang
	insight.SisaBerasKg = insight.TotalPoolBerasKg - insight.TotalDistribusiBerasKg

	if insight.TotalMustahiqAktif > 0 {
		div := float64(insight.TotalMustahiqAktif)
		insight.RekomendasiUangPerOrang = insight.TotalPoolUang / div
		insight.RekomendasiBerasPerOrang = insight.TotalPoolBerasKg / div
	}

	return insight, nil
}

func defaultTanggalDistribusi(tanggal string) string {
	if strings.TrimSpace(tanggal) != "" {
		return tanggal
	}
	return time.Now().Format("2006-01-02")
}
