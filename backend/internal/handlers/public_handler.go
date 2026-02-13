package handlers

import (
	"database/sql"
	"log"
	"net/http"
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

func (h *PublicHandler) GetPublicDashboard(c *gin.Context) {
	var stats models.DashboardStats

	// Total masjid
	h.DB.QueryRow("SELECT COUNT(*) FROM masjid WHERE is_active = 1").Scan(&stats.TotalMasjid)

	// Total orang yang dizakati
	h.DB.QueryRow("SELECT COALESCE(SUM(jumlah_orang), 0) FROM transaksi_zakat WHERE jenis_zakat = 'fitrah'").Scan(&stats.TotalMuzakki)

	// Total mustahiq
	h.DB.QueryRow("SELECT COUNT(*) FROM mustahiq WHERE is_active = 1").Scan(&stats.TotalMustahiq)

	// Total zakat fitrah
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE jenis_zakat = 'fitrah'").Scan(&stats.TotalZakatFitrah)

	// Total zakat mal
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE jenis_zakat = 'mal'").Scan(&stats.TotalZakatMal)

	// Total infaq
	h.DB.QueryRow("SELECT COALESCE(SUM(infaq_tambahan), 0) + COALESCE(SUM(CASE WHEN jenis_zakat = 'infaq' THEN total_dibayar ELSE 0 END), 0) FROM transaksi_zakat").Scan(&stats.TotalInfaq)

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
		Masjid           models.Masjid            `json:"masjid"`
		TotalMuzakki     int                      `json:"total_muzakki"`
		TotalMustahiq    int                      `json:"total_mustahiq"`
		TotalZakatFitrah float64                  `json:"total_zakat_fitrah"`
		TotalZakatMal    float64                  `json:"total_zakat_mal"`
		TotalInfaq       float64                  `json:"total_infaq"`
		Transaksi        []map[string]interface{} `json:"transaksi"`
		Mustahiq         []map[string]interface{} `json:"mustahiq"`
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
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'mal'", id).Scan(&stats.TotalZakatMal)
	h.DB.QueryRow("SELECT COALESCE(SUM(infaq_tambahan), 0) + COALESCE(SUM(CASE WHEN jenis_zakat = 'infaq' THEN total_dibayar ELSE 0 END), 0) FROM transaksi_zakat WHERE masjid_id = ?", id).Scan(&stats.TotalInfaq)

	// Get transaksi list
	stats.Transaksi = []map[string]interface{}{}
	tRows, err := h.DB.Query(`
		SELECT t.id, m.nama as muzakki_nama, t.jenis_zakat, COALESCE(t.jumlah_orang, 0), t.total_dibayar, COALESCE(t.infaq_tambahan, 0), t.tanggal_bayar
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
			var nama, jenis, tanggal string
			var jumlah int
			var total, infaq float64
			err := tRows.Scan(&tid, &nama, &jenis, &jumlah, &total, &infaq, &tanggal)
			if err != nil {
				log.Printf("Error scan transaksi: %v", err)
				continue
			}
			stats.Transaksi = append(stats.Transaksi, map[string]interface{}{
				"id":             tid,
				"muzakki_nama":   nama,
				"jenis_zakat":    jenis,
				"jumlah_orang":   jumlah,
				"total_dibayar":  total,
				"infaq_tambahan": infaq,
				"tanggal_bayar":  tanggal,
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

	c.JSON(http.StatusOK, models.Response{Success: true, Data: stats})
}
