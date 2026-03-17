package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/siazisah/backend/internal/models"
)

type PublicHandler struct {
	DB *sql.DB
}

func NewPublicHandler(db *sql.DB) *PublicHandler {
	return &PublicHandler{DB: db}
}

var berasKgRegex = regexp.MustCompile(`(?i)(?:Beras|Fidyah Beras):\s*([\d.]+)\s*kg`)
var distribusiBerasRegex = regexp.MustCompile(`(?i)\[BERAS_KG=([\d.]+)\]`)
var distribusiMetaStripRegex = regexp.MustCompile(`(?i)\[(?:BERAS_KG|MODE)=[^\]]+\]\s*`)

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

func parseDistribusiKg(keterangan string) float64 {
	if matches := distribusiBerasRegex.FindStringSubmatch(keterangan); len(matches) > 1 {
		kg, err := strconv.ParseFloat(matches[1], 64)
		if err == nil {
			return kg
		}
	}
	return parseKgFromKeterangan(keterangan)
}

func cleanDistribusiKeterangan(keterangan string) string {
	clean := distribusiMetaStripRegex.ReplaceAllString(keterangan, "")
	return strings.TrimSpace(clean)
}

func (h *PublicHandler) GetPublicDashboard(c *gin.Context) {
	var stats models.DashboardStats

	// Total masjid
	h.DB.QueryRow("SELECT COUNT(*) FROM masjid WHERE is_active = 1").Scan(&stats.TotalMasjid)

	// Total muzakki (pemberi zakat)
	h.DB.QueryRow("SELECT COUNT(*) FROM muzakki").Scan(&stats.TotalMuzakki)

	// Total mustahiq
	h.DB.QueryRow("SELECT COUNT(*) FROM mustahiq").Scan(&stats.TotalMustahiq)

	// Total orang yang dizakati (fitrah by jumlah_orang, mal/fidyah count per transaksi, exclude infaq)
	h.DB.QueryRow(`
		SELECT COALESCE(SUM(CASE
			WHEN jenis_zakat = 'fitrah' THEN jumlah_orang
			WHEN jenis_zakat IN ('mal', 'fidyah') THEN 1
			ELSE 0
		END), 0)
		FROM transaksi_zakat
	`).Scan(&stats.TotalOrangDizakati)

	// Total zakat fitrah (semua bentuk)
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE jenis_zakat = 'fitrah'").Scan(&stats.TotalZakatFitrah)

	// Total zakat fitrah uang
	// Exclude infaq surplus: zakat = total_dibayar - infaq_tambahan
	h.DB.QueryRow(`
		SELECT COALESCE(SUM(total_dibayar - COALESCE(infaq_tambahan, 0)), 0)
		FROM transaksi_zakat
		WHERE jenis_zakat = 'fitrah' AND bentuk_zakat = 'uang'
	`).Scan(&stats.TotalZakatFitrahUang)

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
	// Exclude surplus/infaq: zakat mal uang = min(total_dibayar, total_wajib) when total_wajib is available.
	h.DB.QueryRow(`
		SELECT COALESCE(SUM(
			CASE
				WHEN COALESCE(total_wajib, 0) > 0 THEN LEAST(total_dibayar, total_wajib)
				ELSE (total_dibayar - COALESCE(infaq_tambahan, 0))
			END
		), 0)
		FROM transaksi_zakat
		WHERE jenis_zakat = 'mal'
	`).Scan(&stats.TotalZakatMal)

	// Total fidyah
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE jenis_zakat = 'fidyah'").Scan(&stats.TotalFidyah)

	// Total fidyah uang
	// Exclude infaq surplus: fidyah uang = min(total_dibayar, total_wajib) when total_wajib is available.
	h.DB.QueryRow(`
		SELECT COALESCE(SUM(
			CASE
				WHEN COALESCE(total_wajib, 0) > 0 THEN LEAST(total_dibayar, total_wajib)
				ELSE (total_dibayar - COALESCE(infaq_tambahan, 0))
			END
		), 0)
		FROM transaksi_zakat
		WHERE jenis_zakat = 'fidyah' AND bentuk_zakat = 'uang'
	`).Scan(&stats.TotalFidyahUang)

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

	// Mustahiq breakdown by jenis_penerima
	stats.MustahiqPerJenis = []map[string]interface{}{}
	jenisRows, err := h.DB.Query(`
		SELECT jenis_penerima, COUNT(*) as total
		FROM mustahiq
		GROUP BY jenis_penerima
		ORDER BY jenis_penerima ASC
	`)
	if err == nil {
		defer jenisRows.Close()
		for jenisRows.Next() {
			var jenis string
			var total int
			if err := jenisRows.Scan(&jenis, &total); err == nil {
				stats.MustahiqPerJenis = append(stats.MustahiqPerJenis, map[string]interface{}{
					"jenis_penerima": jenis,
					"total":          total,
				})
			}
		}
	}

	// Last update (tanggal terakhir ada perubahan data)
	var lastUpdate sql.NullString
	h.DB.QueryRow(`
		SELECT GREATEST(
			COALESCE((SELECT MAX(tanggal_bayar) FROM transaksi_zakat), '0000-00-00'),
			COALESCE((SELECT MAX(updated_at) FROM mustahiq), '0000-00-00'),
			COALESCE((SELECT MAX(updated_at) FROM muzakki), '0000-00-00'),
			COALESCE((SELECT MAX(updated_at) FROM masjid), '0000-00-00')
		) as last_update
	`).Scan(&lastUpdate)
	if lastUpdate.Valid && lastUpdate.String != "0000-00-00" {
		stats.LastUpdate = lastUpdate.String
	} else {
		// Jika tidak ada data sama sekali, gunakan tanggal hari ini
		stats.LastUpdate = time.Now().Format("2006-01-02")
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
	includeLists := strings.EqualFold(c.Query("include_lists"), "1") || strings.EqualFold(c.Query("include_lists"), "true")

	var stats struct {
		Masjid                  models.Masjid            `json:"masjid"`
		TotalMuzakki            int                      `json:"total_muzakki"`
		TotalMustahiq           int                      `json:"total_mustahiq"`
		TotalOrangDizakati      int                      `json:"total_orang_dizakati"`
		TotalZakatFitrah        float64                  `json:"total_zakat_fitrah"`
		TotalZakatFitrahUang    float64                  `json:"total_zakat_fitrah_uang"`
		TotalZakatFitrahBerasKg float64                  `json:"total_zakat_fitrah_beras_kg"`
		TotalZakatFitrahBerasRp float64                  `json:"total_zakat_fitrah_beras_rupiah"`
		TotalZakatMal           float64                  `json:"total_zakat_mal"`
		TotalFidyah             float64                  `json:"total_fidyah"`
		TotalFidyahUang         float64                  `json:"total_fidyah_uang"`
		TotalFidyahBerasKg      float64                  `json:"total_fidyah_beras_kg"`
		TotalFidyahBerasRp      float64                  `json:"total_fidyah_beras_rupiah"`
		TotalInfaq              float64                  `json:"total_infaq"`
		LastUpdate              string                   `json:"last_update"`
		Transaksi               []map[string]interface{} `json:"transaksi"`
		Mustahiq                []map[string]interface{} `json:"mustahiq"`
		Distribusi              []map[string]interface{} `json:"distribusi"`
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
	h.DB.QueryRow("SELECT COUNT(*) FROM muzakki WHERE masjid_id = ?", id).Scan(&stats.TotalMuzakki)
	h.DB.QueryRow("SELECT COUNT(*) FROM mustahiq WHERE masjid_id = ?", id).Scan(&stats.TotalMustahiq)
	h.DB.QueryRow(`
		SELECT COALESCE(SUM(CASE
			WHEN jenis_zakat = 'fitrah' THEN jumlah_orang
			WHEN jenis_zakat IN ('mal', 'fidyah') THEN 1
			ELSE 0
		END), 0)
		FROM transaksi_zakat
		WHERE masjid_id = ?
	`, id).Scan(&stats.TotalOrangDizakati)
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'fitrah'", id).Scan(&stats.TotalZakatFitrah)
	// Exclude infaq surplus: zakat = total_dibayar - infaq_tambahan
	h.DB.QueryRow(`
		SELECT COALESCE(SUM(total_dibayar - COALESCE(infaq_tambahan, 0)), 0)
		FROM transaksi_zakat
		WHERE masjid_id = ? AND jenis_zakat = 'fitrah' AND bentuk_zakat = 'uang'
	`, id).Scan(&stats.TotalZakatFitrahUang)
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

	// Exclude surplus/infaq: zakat mal uang = min(total_dibayar, total_wajib) when total_wajib is available.
	h.DB.QueryRow(`
		SELECT COALESCE(SUM(
			CASE
				WHEN COALESCE(total_wajib, 0) > 0 THEN LEAST(total_dibayar, total_wajib)
				ELSE (total_dibayar - COALESCE(infaq_tambahan, 0))
			END
		), 0)
		FROM transaksi_zakat
		WHERE masjid_id = ? AND jenis_zakat = 'mal'
	`, id).Scan(&stats.TotalZakatMal)
	h.DB.QueryRow("SELECT COALESCE(SUM(total_dibayar), 0) FROM transaksi_zakat WHERE masjid_id = ? AND jenis_zakat = 'fidyah'", id).Scan(&stats.TotalFidyah)
	// Exclude infaq surplus: fidyah uang = min(total_dibayar, total_wajib) when total_wajib is available.
	h.DB.QueryRow(`
		SELECT COALESCE(SUM(
			CASE
				WHEN COALESCE(total_wajib, 0) > 0 THEN LEAST(total_dibayar, total_wajib)
				ELSE (total_dibayar - COALESCE(infaq_tambahan, 0))
			END
		), 0)
		FROM transaksi_zakat
		WHERE masjid_id = ? AND jenis_zakat = 'fidyah' AND bentuk_zakat = 'uang'
	`, id).Scan(&stats.TotalFidyahUang)
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

	stats.Transaksi = []map[string]interface{}{}
	stats.Mustahiq = []map[string]interface{}{}
	stats.Distribusi = []map[string]interface{}{}
	if includeLists {
		// Get transaksi list
		tRows, err := h.DB.Query(`
			SELECT t.id, COALESCE(m.nama, 'Unknown') as muzakki_nama, t.jenis_zakat, COALESCE(t.bentuk_zakat, ''),
				COALESCE(t.jumlah_orang, 0), COALESCE(t.jumlah_hari_fidyah, 0), COALESCE(t.kelas_zakat, ''), COALESCE(t.total_wajib, 0),
				t.total_dibayar, COALESCE(t.infaq_tambahan, 0), t.tanggal_bayar, COALESCE(t.keterangan, '')
			FROM transaksi_zakat t
			LEFT JOIN muzakki m ON t.muzakki_id = m.id
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
				var kelasZakat string
				var jumlah int
				var jumlahHari int
				var totalWajib float64
				var total, infaq float64
				err := tRows.Scan(&tid, &nama, &jenis, &bentuk, &jumlah, &jumlahHari, &kelasZakat, &totalWajib, &total, &infaq, &tanggal, &keterangan)
				if err != nil {
					log.Printf("Error scan transaksi: %v", err)
					continue
				}
				stats.Transaksi = append(stats.Transaksi, map[string]interface{}{
					"id":                 tid,
					"muzakki_nama":       nama,
					"jenis_zakat":        jenis,
					"bentuk_zakat":       bentuk,
					"kelas_zakat":        kelasZakat,
					"jumlah_orang":       jumlah,
					"jumlah_hari_fidyah": jumlahHari,
					"total_wajib":        totalWajib,
					"total_dibayar":      total,
					"infaq_tambahan":     infaq,
					"tanggal_bayar":      tanggal,
					"keterangan":         keterangan,
				})
			}
		}

		// Get mustahiq list
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

		// Get distribusi list
		dRows, err := h.DB.Query(`
			SELECT d.id, COALESCE(m.nama, 'Unknown'), d.nominal, d.tanggal_distribusi, COALESCE(d.keterangan, '')
			FROM distribusi_zakat d
			LEFT JOIN mustahiq m ON d.mustahiq_id = m.id
			WHERE d.masjid_id = ?
			ORDER BY d.tanggal_distribusi DESC, d.created_at DESC
		`, id)
		if err != nil {
			log.Printf("Error query distribusi: %v", err)
		} else if dRows != nil {
			defer dRows.Close()
			for dRows.Next() {
				var did int
				var nama, tanggal, keterangan string
				var nominal float64
				if err := dRows.Scan(&did, &nama, &nominal, &tanggal, &keterangan); err != nil {
					log.Printf("Error scan distribusi: %v", err)
					continue
				}
				stats.Distribusi = append(stats.Distribusi, map[string]interface{}{
					"id":                 did,
					"mustahiq_nama":      nama,
					"nominal":            nominal,
					"tanggal_distribusi": tanggal,
					"beras_kg":           parseDistribusiKg(keterangan),
					"keterangan":         cleanDistribusiKeterangan(keterangan),
				})
			}
		}
	}

	// Last update (tanggal transaksi terakhir untuk masjid ini)
	var lastUpdate sql.NullString
	// Use updated_at timestamps to include time component.
	h.DB.QueryRow(`
		SELECT GREATEST(
			COALESCE((SELECT MAX(updated_at) FROM transaksi_zakat WHERE masjid_id = ?), '0000-00-00 00:00:00'),
			COALESCE((SELECT MAX(updated_at) FROM distribusi_zakat WHERE masjid_id = ?), '0000-00-00 00:00:00'),
			COALESCE((SELECT MAX(updated_at) FROM mustahiq WHERE masjid_id = ?), '0000-00-00 00:00:00'),
			COALESCE((SELECT MAX(updated_at) FROM muzakki WHERE masjid_id = ?), '0000-00-00 00:00:00'),
			COALESCE((SELECT updated_at FROM masjid WHERE id = ?), '0000-00-00 00:00:00')
		) as last_update
	`, id, id, id, id, id).Scan(&lastUpdate)
	if lastUpdate.Valid && lastUpdate.String != "0000-00-00 00:00:00" {
		stats.LastUpdate = lastUpdate.String
	} else {
		stats.LastUpdate = ""
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: stats})
}

func (h *PublicHandler) GetMasjidTransaksi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	page, pageSize, isAll := parsePagination(c)
	jenisFilter := strings.TrimSpace(strings.ToLower(c.Query("jenis_zakat")))
	bentukFilter := strings.TrimSpace(strings.ToLower(c.Query("bentuk_zakat")))
	qFilter := strings.TrimSpace(strings.ToLower(c.Query("q")))

	if jenisFilter != "fitrah" && jenisFilter != "fidyah" {
		bentukFilter = ""
	}

	if isAll {
		page = 1
	}

	whereClauses := []string{"t.masjid_id = ?"}
	whereArgs := []interface{}{id}

	if jenisFilter != "" {
		whereClauses = append(whereClauses, "t.jenis_zakat = ?")
		whereArgs = append(whereArgs, jenisFilter)
	}

	if bentukFilter != "" {
		whereClauses = append(whereClauses, "t.bentuk_zakat = ?")
		whereArgs = append(whereArgs, bentukFilter)
	}

	if qFilter != "" {
		whereClauses = append(whereClauses, "LOWER(COALESCE(m.nama, '')) LIKE ?")
		whereArgs = append(whereArgs, "%"+qFilter+"%")
	}

	whereSQL := strings.Join(whereClauses, " AND ")

	var total int
	countQuery := "SELECT COUNT(*) FROM transaksi_zakat t LEFT JOIN muzakki m ON t.muzakki_id = m.id WHERE " + whereSQL
	if err := h.DB.QueryRow(countQuery, whereArgs...).Scan(&total); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get transaksi"})
		return
	}

	query := `
		SELECT t.id, COALESCE(m.nama, 'Unknown') as muzakki_nama, t.jenis_zakat, COALESCE(t.bentuk_zakat, ''),
			COALESCE(t.jumlah_orang, 0), COALESCE(t.jumlah_hari_fidyah, 0), COALESCE(t.kg_beras_dibayar, 0),
			COALESCE(t.kelas_zakat, ''), COALESCE(t.total_wajib, 0), t.total_dibayar,
			COALESCE(t.infaq_tambahan, 0), t.tanggal_bayar, COALESCE(t.keterangan, '')
		FROM transaksi_zakat t
		LEFT JOIN muzakki m ON t.muzakki_id = m.id
		WHERE ` + whereSQL + `
		ORDER BY t.tanggal_bayar DESC
	`

	args := append([]interface{}{}, whereArgs...)
	if !isAll {
		query += " LIMIT ? OFFSET ?"
		offset := (page - 1) * pageSize
		args = append(args, pageSize, offset)
	}

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get transaksi"})
		return
	}
	defer rows.Close()

	items := []map[string]interface{}{}
	for rows.Next() {
		var tid int
		var nama, jenis, bentuk, tanggal, keterangan string
		var jumlah int
		var jumlahHari int
		var kgBeras float64
		var kelasZakat string
		var totalWajib float64
		var totalDibayar, infaq float64
		if err := rows.Scan(&tid, &nama, &jenis, &bentuk, &jumlah, &jumlahHari, &kgBeras, &kelasZakat, &totalWajib, &totalDibayar, &infaq, &tanggal, &keterangan); err != nil {
			continue
		}
		items = append(items, map[string]interface{}{
			"id":                 tid,
			"muzakki_nama":       nama,
			"jenis_zakat":        jenis,
			"bentuk_zakat":       bentuk,
			"kelas_zakat":        kelasZakat,
			"jumlah_orang":       jumlah,
			"jumlah_hari_fidyah": jumlahHari,
			"kg_beras_dibayar":   kgBeras,
			"total_wajib":        totalWajib,
			"total_dibayar":      totalDibayar,
			"infaq_tambahan":     infaq,
			"tanggal_bayar":      tanggal,
			"keterangan":         keterangan,
		})
	}

	pagination := buildPagination(total, page, pageSize, isAll)
	c.JSON(http.StatusOK, models.Response{Success: true, Data: gin.H{"items": items, "pagination": pagination}})
}

func (h *PublicHandler) GetMasjidMustahiq(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	page, pageSize, isAll := parsePagination(c)
	jenisFilter := strings.TrimSpace(strings.ToLower(c.Query("jenis_penerima")))
	qFilter := strings.TrimSpace(strings.ToLower(c.Query("q")))

	whereClauses := []string{"masjid_id = ?"}
	whereArgs := []interface{}{id}

	if jenisFilter != "" {
		whereClauses = append(whereClauses, "LOWER(jenis_penerima) = ?")
		whereArgs = append(whereArgs, jenisFilter)
	}

	if qFilter != "" {
		whereClauses = append(whereClauses, "LOWER(nama) LIKE ?")
		whereArgs = append(whereArgs, "%"+qFilter+"%")
	}

	whereSQL := strings.Join(whereClauses, " AND ")

	var total int
	countQuery := "SELECT COUNT(*) FROM mustahiq WHERE " + whereSQL
	if err := h.DB.QueryRow(countQuery, whereArgs...).Scan(&total); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get mustahiq"})
		return
	}

	query := `
		SELECT id, nama, alamat, jenis_penerima, rt, COALESCE(keterangan, '')
		FROM mustahiq
		WHERE ` + whereSQL + `
		ORDER BY nama ASC
	`

	args := append([]interface{}{}, whereArgs...)
	if !isAll {
		query += " LIMIT ? OFFSET ?"
		offset := (page - 1) * pageSize
		args = append(args, pageSize, offset)
	}

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get mustahiq"})
		return
	}
	defer rows.Close()

	items := []map[string]interface{}{}
	for rows.Next() {
		var mid int
		var nama, alamat, kategori, rt, keterangan string
		if err := rows.Scan(&mid, &nama, &alamat, &kategori, &rt, &keterangan); err != nil {
			continue
		}
		items = append(items, map[string]interface{}{
			"id":             mid,
			"nama":           nama,
			"alamat":         alamat,
			"kategori":       kategori,
			"jenis_penerima": kategori,
			"rt":             rt,
			"keterangan":     keterangan,
		})
	}

	pagination := buildPagination(total, page, pageSize, isAll)
	c.JSON(http.StatusOK, models.Response{Success: true, Data: gin.H{"items": items, "pagination": pagination}})
}

func (h *PublicHandler) GetMasjidDistribusi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	page, pageSize, isAll := parsePagination(c)

	var total int
	if err := h.DB.QueryRow("SELECT COUNT(*) FROM distribusi_zakat WHERE masjid_id = ?", id).Scan(&total); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get distribusi"})
		return
	}

	query := `
		SELECT d.id, COALESCE(m.nama, 'Unknown'), d.nominal, d.tanggal_distribusi, COALESCE(d.keterangan, '')
		FROM distribusi_zakat d
		LEFT JOIN mustahiq m ON d.mustahiq_id = m.id
		WHERE d.masjid_id = ?
		ORDER BY d.tanggal_distribusi DESC, d.created_at DESC
	`

	args := []interface{}{id}
	if !isAll {
		query += " LIMIT ? OFFSET ?"
		offset := (page - 1) * pageSize
		args = append(args, pageSize, offset)
	}

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get distribusi"})
		return
	}
	defer rows.Close()

	items := []map[string]interface{}{}
	for rows.Next() {
		var did int
		var nama, tanggal, keterangan string
		var nominal float64
		if err := rows.Scan(&did, &nama, &nominal, &tanggal, &keterangan); err != nil {
			continue
		}
		items = append(items, map[string]interface{}{
			"id":                 did,
			"mustahiq_nama":      nama,
			"nominal":            nominal,
			"tanggal_distribusi": tanggal,
			"beras_kg":           parseDistribusiKg(keterangan),
			"keterangan":         cleanDistribusiKeterangan(keterangan),
		})
	}

	pagination := buildPagination(total, page, pageSize, isAll)
	c.JSON(http.StatusOK, models.Response{Success: true, Data: gin.H{"items": items, "pagination": pagination}})
}

// GetZakatFitrahStats - Get stats for Zakat Fitrah
func (h *PublicHandler) GetZakatFitrahStats(c *gin.Context) {
	stats := map[string]interface{}{}

	// Count muzakki (pemberi zakat)
	var totalMuzakki int
	h.DB.QueryRow("SELECT COUNT(*) FROM muzakki").Scan(&totalMuzakki)
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

	// Count muzakki (pemberi zakat)
	var totalMuzakki int
	h.DB.QueryRow("SELECT COUNT(*) FROM muzakki").Scan(&totalMuzakki)
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

	// Count muzakki (pemberi zakat)
	var totalMuzakki int
	h.DB.QueryRow("SELECT COUNT(*) FROM muzakki").Scan(&totalMuzakki)
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

// GetMustahiqStats - Get global mustahiq stats
func (h *PublicHandler) GetMustahiqStats(c *gin.Context) {
	stats := map[string]interface{}{}

	var totalMustahiq int
	h.DB.QueryRow(`
		SELECT COUNT(*)
		FROM mustahiq
	`).Scan(&totalMustahiq)

	rows, err := h.DB.Query(`
		SELECT LOWER(COALESCE(jenis_penerima, 'lainnya')) AS jenis_penerima, COUNT(*) AS total
		FROM mustahiq
		GROUP BY LOWER(COALESCE(jenis_penerima, 'lainnya'))
		ORDER BY jenis_penerima ASC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get mustahiq stats"})
		return
	}
	defer rows.Close()

	perJenis := []map[string]interface{}{}
	for rows.Next() {
		var jenis string
		var total int
		if err := rows.Scan(&jenis, &total); err != nil {
			continue
		}
		perJenis = append(perJenis, map[string]interface{}{
			"jenis_penerima": jenis,
			"total":          total,
		})
	}

	stats["total_mustahiq"] = totalMustahiq
	stats["per_jenis"] = perJenis

	c.JSON(http.StatusOK, models.Response{Success: true, Data: stats})
}
