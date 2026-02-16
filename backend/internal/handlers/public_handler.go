package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/siazisah/internal/models"
)

type PublicHandler struct {
	DB *sql.DB
}

func NewPublicHandler(db *sql.DB) *PublicHandler {
	return &PublicHandler{DB: db}
}

var berasKgRegex = regexp.MustCompile(`(?i)(?:Beras|Fidyah Beras):\s*([\d.]+)\s*kg`)

func parseKgFromKeterangan(keterangan string) float64 {
	matches := berasKgRegex.FindStringSubmatch(keterangan)
	if len(matches) < 2 {
		return 0
	}
	kg, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0
	}
	return kg
}

func (h *PublicHandler) GetPublicDashboard(c *gin.Context) {
	var stats models.DashboardStats

	// Total masjid
	h.DB.QueryRow("SELECT COUNT(*) FROM masjid WHERE is_active = 1").Scan(&stats.TotalMasjid)

	// Total orang yang dizakati
	h.DB.QueryRow("SELECT COALESCE(SUM(jumlah_orang), 0) FROM transaksi_zakat WHERE jenis_zakat = 'fitrah'").Scan(&stats.TotalMuzakki)

	// Total mustahiq
	h.DB.QueryRow("SELECT COUNT(*) FROM mustahiq WHERE is_active = 1").Scan(&stats.TotalMustahiq)

	// Total zakat fitrah (semua bentuk)
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE jenis_zakat = 'fitrah'").Scan(&stats.TotalZakatFitrah)

	// Total zakat fitrah uang
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE jenis_zakat = 'fitrah' AND bentuk_zakat = 'uang'").Scan(&stats.TotalZakatFitrahUang)

	// Total zakat fitrah beras (dalam rupiah)
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE jenis_zakat = 'fitrah' AND bentuk_zakat = 'beras'").Scan(&stats.TotalZakatFitrahBerasRp)

	// Total zakat fitrah beras (dalam kg)
	h.DB.QueryRow("SELECT COALESCE(SUM(kg_beras_dibayar), 0) FROM transaksi_zakat WHERE jenis_zakat = 'fitrah' AND bentuk_zakat = 'beras'").Scan(&stats.TotalZakatFitrahBerasKg)
	rows, err := h.DB.Query("SELECT keterangan FROM transaksi_zakat WHERE jenis_zakat = 'fitrah' AND bentuk_zakat = 'beras' AND (kg_beras_dibayar IS NULL OR kg_beras_dibayar = 0) AND keterangan IS NOT NULL")
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var keterangan string
			if err := rows.Scan(&keterangan); err == nil {
				stats.TotalZakatFitrahBerasKg += parseKgFromKeterangan(keterangan)
			}
		}
	}

	// Total zakat mal
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE jenis_zakat = 'mal'").Scan(&stats.TotalZakatMal)

	// Total fidyah
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE jenis_zakat = 'fidyah'").Scan(&stats.TotalFidyah)

	// Total fidyah uang
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE jenis_zakat = 'fidyah' AND bentuk_zakat = 'uang'").Scan(&stats.TotalFidyahUang)

	// Total fidyah beras (dalam rupiah)
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE jenis_zakat = 'fidyah' AND bentuk_zakat = 'beras'").Scan(&stats.TotalFidyahBerasRp)

	// Total fidyah beras (dalam kg)
	h.DB.QueryRow("SELECT COALESCE(SUM(kg_beras_dibayar), 0) FROM transaksi_zakat WHERE jenis_zakat = 'fidyah' AND bentuk_zakat = 'beras'").Scan(&stats.TotalFidyahBerasKg)
	fidyahRows, err := h.DB.Query("SELECT keterangan FROM transaksi_zakat WHERE jenis_zakat = 'fidyah' AND bentuk_zakat = 'beras' AND (kg_beras_dibayar IS NULL OR kg_beras_dibayar = 0) AND keterangan IS NOT NULL")
	if err == nil {
		defer fidyahRows.Close()
		for fidyahRows.Next() {
			var keterangan string
			if err := fidyahRows.Scan(&keterangan); err == nil {
				stats.TotalFidyahBerasKg += parseKgFromKeterangan(keterangan)
			}
		}
	}

	// Total infaq
	h.DB.QueryRow("SELECT COALESCE(SUM(infaq_tambahan), 0) + COALESCE(SUM(CASE WHEN jenis_zakat = 'infaq' THEN total_dibayar ELSE 0 END), 0) FROM transaksi_zakat").Scan(&stats.TotalInfaq)

	// Last update (tanggal transaksi terakhir)
	var lastUpdate sql.NullString
	h.DB.QueryRow("SELECT MAX(tanggal_bayar) FROM transaksi_zakat").Scan(&lastUpdate)
	if lastUpdate.Valid {
		stats.LastUpdate = lastUpdate.String
	} else {
		stats.LastUpdate = ""
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: stats})
}

func (h *PublicHandler) GetAllMasjid(c *gin.Context) {
	rows, err := h.DB.Query(`
		SELECT id, nama, alamat, telepon, is_active, created_at, updated_at
		FROM masjid 
		WHERE is_active = 1 
		ORDER BY nama ASC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get masjids"})
		return
	}
	defer rows.Close()

	var masjids []models.Masjid
	for rows.Next() {
		var masjid models.Masjid
		rows.Scan(&masjid.ID, &masjid.Nama, &masjid.Alamat, &masjid.Telepon, &masjid.IsActive, &masjid.CreatedAt, &masjid.UpdatedAt)
		masjids = append(masjids, masjid)
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: masjids})
}

func (h *PublicHandler) GetMasjidStats(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var stats struct {
		Masjid                   models.Masjid            `json:"masjid"`
		TotalMuzakki             int                      `json:"total_muzakki"`
		TotalMustahiq            int                      `json:"total_mustahiq"`
		TotalZakatFitrah         float64                  `json:"total_zakat_fitrah"`
		TotalZakatFitrahUang     float64                  `json:"total_zakat_fitrah_uang"`
		TotalZakatFitrahBerasKg  float64                  `json:"total_zakat_fitrah_beras_kg"`
		TotalZakatFitrahBerasRp  float64                  `json:"total_zakat_fitrah_beras_rupiah"`
		TotalZakatMal            float64                  `json:"total_zakat_mal"`
		TotalFidyah              float64                  `json:"total_fidyah"`
		TotalFidyahUang          float64                  `json:"total_fidyah_uang"`
		TotalFidyahBerasKg       float64                  `json:"total_fidyah_beras_kg"`
		TotalFidyahBerasRp       float64                  `json:"total_fidyah_beras_rupiah"`
		TotalInfaq               float64                  `json:"total_infaq"`
		LastUpdate               string                   `json:"last_update"`
		Transaksi                []map[string]interface{} `json:"transaksi"`
		Mustahiq                 []map[string]interface{} `json:"mustahiq"`
	}

	// Get masjid info
	err := h.DB.QueryRow(`
		SELECT id, nama, alamat, telepon, is_active, created_at, updated_at
		FROM masjid WHERE id = ?
	`, id).Scan(&stats.Masjid.ID, &stats.Masjid.Nama, &stats.Masjid.Alamat,
		&stats.Masjid.Telepon, &stats.Masjid.IsActive, &stats.Masjid.CreatedAt, &stats.Masjid.UpdatedAt)

	if err != nil {
		log.Printf("Error get masjid: %v", err)
		c.JSON(http.StatusNotFound, models.Response{Success: false, Message: "Masjid not found"})
		return
	}

	// Get stats
	h.DB.QueryRow("SELECT COALESCE(SUM(jumlah_orang), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'fitrah'", id).Scan(&stats.TotalMuzakki)
	h.DB.QueryRow("SELECT COUNT(*) FROM mustahiq WHERE masjid_id = ?", id).Scan(&stats.TotalMustahiq)
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'fitrah'", id).Scan(&stats.TotalZakatFitrah)
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'fitrah' AND bentuk_zakat = 'uang'", id).Scan(&stats.TotalZakatFitrahUang)
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'fitrah' AND bentuk_zakat = 'beras'", id).Scan(&stats.TotalZakatFitrahBerasRp)
	
	// Total kg beras fitrah
	h.DB.QueryRow("SELECT COALESCE(SUM(kg_beras_dibayar), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'fitrah' AND bentuk_zakat = 'beras'", id).Scan(&stats.TotalZakatFitrahBerasKg)
	berasRows, err := h.DB.Query("SELECT keterangan FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'fitrah' AND bentuk_zakat = 'beras' AND (kg_beras_dibayar IS NULL OR kg_beras_dibayar = 0) AND keterangan IS NOT NULL", id)
	if err == nil {
		defer berasRows.Close()
		for berasRows.Next() {
			var keterangan string
			if err := berasRows.Scan(&keterangan); err == nil {
				stats.TotalZakatFitrahBerasKg += parseKgFromKeterangan(keterangan)
			}
		}
	}
	
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'mal'", id).Scan(&stats.TotalZakatMal)
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'fidyah'", id).Scan(&stats.TotalFidyah)
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'fidyah' AND bentuk_zakat = 'uang'", id).Scan(&stats.TotalFidyahUang)
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'fidyah' AND bentuk_zakat = 'beras'", id).Scan(&stats.TotalFidyahBerasRp)
	
	// Total kg beras fidyah
	h.DB.QueryRow("SELECT COALESCE(SUM(kg_beras_dibayar), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'fidyah' AND bentuk_zakat = 'beras'", id).Scan(&stats.TotalFidyahBerasKg)
	fidyahRows, err := h.DB.Query("SELECT keterangan FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'fidyah' AND bentuk_zakat = 'beras' AND (kg_beras_dibayar IS NULL OR kg_beras_dibayar = 0) AND keterangan IS NOT NULL", id)
	if err == nil {
		defer fidyahRows.Close()
		for fidyahRows.Next() {
			var keterangan string
			if err := fidyahRows.Scan(&keterangan); err == nil {
				stats.TotalFidyahBerasKg += parseKgFromKeterangan(keterangan)
			}
		}
	}
	
	h.DB.QueryRow("SELECT COALESCE(SUM(infaq_tambahan), 0) + COALESCE(SUM(CASE WHEN jenis_zakat = 'infaq' THEN total_dibayar ELSE 0 END), 0) FROM transaksi_zakat WHERE masjid_id = ?", id).Scan(&stats.TotalInfaq)

	// Get transaksi list
	stats.Transaksi = []map[string]interface{}{}
	tRows, err := h.DB.Query(`
		SELECT t.id, m.nama as muzakki_nama, t.jenis_zakat, COALESCE(t.bentuk_zakat, ''), COALESCE(t.jumlah_orang, 0), COALESCE(t.jumlah_hari_fidyah, 0), t.total_dibayar, COALESCE(t.infaq_tambahan, 0), t.tanggal_bayar, COALESCE(t.keterangan, '')
		FROM transaksi_zakat t
		JOIN muzakki m ON t.muzakki_id = m.id
		WHERE t.masjid_id = ?
		ORDER BY t.tanggal_bayar DESC
	`, id)
	if err != nil {
		log.Printf("Error query transaksi: %v", err)
	} else if tRows != nil {
		defer tRows.Close()
		for tRows.Next() {
			var tid int
			var nama, jenis, bentuk, tanggal, keterangan string
			var jumlah int
			var jumlahHari int
			var total, infaq float64
			err := tRows.Scan(&tid, &nama, &jenis, &bentuk, &jumlah, &jumlahHari, &total, &infaq, &tanggal, &keterangan)
			if err != nil {
				log.Printf("Error scan transaksi: %v", err)
				continue
			}
			stats.Transaksi = append(stats.Transaksi, map[string]interface{}{
				"id":             tid,
				"muzakki_nama":   nama,
				"jenis_zakat":    jenis,
				"bentuk_zakat":   bentuk,
				"jumlah_orang":   jumlah,
				"jumlah_hari_fidyah": jumlahHari,
				"total_dibayar":       total,
				"infaq_tambahan": infaq,
				"tanggal_bayar":  tanggal,
				"keterangan":     keterangan,
			})
		}
	}

	// Get mustahiq list
	stats.Mustahiq = []map[string]interface{}{}
	mRows, err := h.DB.Query(`
		SELECT id, nama, alamat, jenis_penerima, rt, COALESCE(keterangan, '')
		FROM mustahiq
		WHERE masjid_id = ?
		ORDER BY nama ASC
	`, id)
	if err != nil {
		log.Printf("Error query mustahiq: %v", err)
	} else if mRows != nil {
		defer mRows.Close()
		count := 0
		for mRows.Next() {
			var mid int
			var nama, alamat, kategori, rt, keterangan string
			err := mRows.Scan(&mid, &nama, &alamat, &kategori, &rt, &keterangan)
			if err != nil {
				log.Printf("Error scan mustahiq: %v", err)
				continue
			}
			stats.Mustahiq = append(stats.Mustahiq, map[string]interface{}{
				"id":         mid,
				"nama":       nama,
				"alamat":     alamat,
				"kategori":   kategori,
				"rt":         rt,
				"keterangan": keterangan,
			})
			count++
		}
		log.Printf("Masjid ID %d: Found %d mustahiq", id, count)
	}

	// Last update (tanggal transaksi terakhir untuk masjid ini)
	var lastUpdate sql.NullString
	h.DB.QueryRow("SELECT MAX(tanggal_bayar) FROM transaksi_zakat WHERE masjid_id = ?", id).Scan(&lastUpdate)
	if lastUpdate.Valid {
		stats.LastUpdate = lastUpdate.String
	} else {
		stats.LastUpdate = ""
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: stats})
}

// GetZakatFitrahStats - Get stats for Zakat Fitrah
func (h *PublicHandler) GetZakatFitrahStats(c *gin.Context) {
	stats := map[string]interface{}{}

	// Count unique muzakki (wajib zakat) per transaction
	var totalMuzakki int
	h.DB.QueryRow(`
		SELECT COUNT(DISTINCT muzakki_id) 
		FROM transaksi_zakat 
		WHERE jenis_zakat = 'fitrah'
	`).Scan(&totalMuzakki)
	stats["total_muzakki"] = totalMuzakki

	// Total beras (kg) - parsing from keterangan field
	var totalBerasKg float64 = 0
	// For now, we'll count the number of beras transactions * assume 2.5kg each
	berasCount := 0
	h.DB.QueryRow(`
		SELECT COUNT(*) 
		FROM transaksi_zakat 
		WHERE jenis_zakat = 'fitrah' AND bentuk_zakat = 'beras'
	`).Scan(&berasCount)
	totalBerasKg = float64(berasCount) * 2.5 // Default kadar zakat fitrah = 2.5kg
	stats["total_beras"] = totalBerasKg

	// Total uang (only where bentuk_zakat = 'uang')
	var totalUang float64 = 0
	err := h.DB.QueryRow(`
		SELECT COALESCE(SUM(total_dibayar), 0) 
		FROM transaksi_zakat 
		WHERE jenis_zakat = 'fitrah' AND bentuk_zakat = 'uang'
	`).Scan(&totalUang)
	if err != nil {
		log.Printf("Error querying fitrah uang: %v", err)
	}
	log.Printf("Fitrah - Total Uang from DB: %.2f", totalUang)
	stats["total_uang"] = totalUang

	// Count total active mustahiq (penerima)
	var totalMustahiq int = 0
	h.DB.QueryRow(`
		SELECT COUNT(*) 
		FROM mustahiq 
		WHERE is_active = 1
	`).Scan(&totalMustahiq)
	stats["total_mustahiq"] = totalMustahiq

	c.JSON(http.StatusOK, models.Response{Success: true, Data: stats})
}

// GetZakatMalStats - Get stats for Zakat Mal
func (h *PublicHandler) GetZakatMalStats(c *gin.Context) {
	stats := map[string]interface{}{}

	// Count unique muzakki
	var totalMuzakki int
	h.DB.QueryRow(`
		SELECT COUNT(DISTINCT muzakki_id) 
		FROM transaksi_zakat 
		WHERE jenis_zakat = 'mal'
	`).Scan(&totalMuzakki)
	stats["total_muzakki"] = totalMuzakki

	// Total beras (kg) - same approach as fitrah
	var totalBerasKg float64 = 0
	berasCount := 0
	h.DB.QueryRow(`
		SELECT COUNT(*) 
		FROM transaksi_zakat 
		WHERE jenis_zakat = 'mal' AND bentuk_zakat = 'beras'
	`).Scan(&berasCount)
	totalBerasKg = float64(berasCount) * 2.5
	stats["total_beras"] = totalBerasKg

	// Total uang (for mal, typically all transactions are monetary)
	var totalUang float64 = 0
	h.DB.QueryRow(`
		SELECT COALESCE(SUM(total_dibayar), 0) 
		FROM transaksi_zakat 
		WHERE jenis_zakat = 'mal'
	`).Scan(&totalUang)
	stats["total_uang"] = totalUang

	// Count total active mustahiq
	var totalMustahiq int = 0
	h.DB.QueryRow(`
		SELECT COUNT(*) 
		FROM mustahiq 
		WHERE is_active = 1
	`).Scan(&totalMustahiq)
	stats["total_mustahiq"] = totalMustahiq

	c.JSON(http.StatusOK, models.Response{Success: true, Data: stats})
}

// GetFidyahStats - Get stats for Fidyah
func (h *PublicHandler) GetFidyahStats(c *gin.Context) {
	stats := map[string]interface{}{}

	// Count unique muzakki
	var totalMuzakki int
	h.DB.QueryRow(`
		SELECT COUNT(DISTINCT muzakki_id) 
		FROM transaksi_zakat 
		WHERE jenis_zakat = 'fidyah'
	`).Scan(&totalMuzakki)
	stats["total_muzakki"] = totalMuzakki

	// Total beras (kg)
	var totalBerasKg float64 = 0
	berasCount := 0
	h.DB.QueryRow(`
		SELECT COUNT(*) 
		FROM transaksi_zakat 
		WHERE jenis_zakat = 'fidyah' AND bentuk_zakat = 'beras'
	`).Scan(&berasCount)
	totalBerasKg = float64(berasCount) * 2.5
	stats["total_beras"] = totalBerasKg

	// Total uang (for fidyah, can be either uang or beras)
	var totalUang float64 = 0
	h.DB.QueryRow(`
		SELECT COALESCE(SUM(total_dibayar), 0) 
		FROM transaksi_zakat 
		WHERE jenis_zakat = 'fidyah'
	`).Scan(&totalUang)
	stats["total_uang"] = totalUang

	// Count total active mustahiq
	var totalMustahiq int = 0
	h.DB.QueryRow(`
		SELECT COUNT(*) 
		FROM mustahiq 
		WHERE is_active = 1
	`).Scan(&totalMustahiq)
	stats["total_mustahiq"] = totalMustahiq

	c.JSON(http.StatusOK, models.Response{Success: true, Data: stats})
}
