package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/siazisah/backend/internal/models"
	"github.com/yourusername/siazisah/backend/internal/repository"
)

type TransaksiHandler struct {
	transaksiRepo *repository.TransaksiZakatRepository
	muzakkiRepo   *repository.MuzakkiRepository
	masjidRepo    *repository.MasjidRepository
	pengurusRepo  *repository.PengurusRepository
}

func NewTransaksiHandler(
	transaksiRepo *repository.TransaksiZakatRepository,
	muzakkiRepo *repository.MuzakkiRepository,
	masjidRepo *repository.MasjidRepository,
	pengurusRepo *repository.PengurusRepository,
) *TransaksiHandler {
	return &TransaksiHandler{
		transaksiRepo: transaksiRepo,
		muzakkiRepo:   muzakkiRepo,
		masjidRepo:    masjidRepo,
		pengurusRepo:  pengurusRepo,
	}
}

var printKgRegex = regexp.MustCompile(`(?i)(?:Beras|Fidyah Beras):\s*([\d.]+)\s*kg`)

func parsePrintKgFromKeterangan(keterangan string) float64 {
	if strings.TrimSpace(keterangan) == "" {
		return 0
	}
	matches := printKgRegex.FindStringSubmatch(keterangan)
	if len(matches) < 2 {
		return 0
	}
	kg, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0
	}
	return kg
}

func formatIDNumber(value int64) string {
	negative := value < 0
	if negative {
		value = -value
	}

	raw := strconv.FormatInt(value, 10)
	parts := []string{}
	for len(raw) > 3 {
		parts = append([]string{raw[len(raw)-3:]}, parts...)
		raw = raw[:len(raw)-3]
	}
	if raw != "" {
		parts = append([]string{raw}, parts...)
	}

	result := strings.Join(parts, ".")
	if negative {
		return "-" + result
	}
	return result
}

func formatSignedDateID(t time.Time) string {
	months := []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	month := months[int(t.Month())-1]
	return fmt.Sprintf("%d %s %d", t.Day(), month, t.Year())
}

func normalizeMasjidTitleName(name string) string {
	clean := strings.TrimSpace(name)
	if clean == "" {
		return clean
	}

	originalUpper := strings.ToUpper(clean)
	prefix := ""
	if strings.HasPrefix(originalUpper, "MASJID ") {
		prefix = "MASJID"
	}
	if strings.HasPrefix(originalUpper, "LANGGAR ") {
		prefix = "LANGGAR"
	}

	for {
		upper := strings.ToUpper(clean)
		if strings.HasPrefix(upper, "MASJID ") {
			clean = strings.TrimSpace(clean[len("MASJID "):])
			continue
		}
		if strings.HasPrefix(upper, "LANGGAR ") {
			clean = strings.TrimSpace(clean[len("LANGGAR "):])
			continue
		}
		break
	}

	if prefix != "" {
		if clean == "" {
			return prefix
		}
		return prefix + " " + clean
	}
	return clean
}

func pickKetuaPengurusNama(pengurusMasjid []models.PengurusMasjid) string {
	for _, p := range pengurusMasjid {
		if strings.EqualFold(strings.TrimSpace(p.Jabatan), "ketua") {
			return p.Nama
		}
	}
	for _, p := range pengurusMasjid {
		if strings.Contains(strings.ToLower(strings.TrimSpace(p.Jabatan)), "ketua") {
			return p.Nama
		}
	}
	return "-"
}

func pickKetuaUPZNama(pengurusZakat []models.PengurusZakat) string {
	for _, p := range pengurusZakat {
		jabatan := strings.ToLower(strings.TrimSpace(p.Jabatan))
		if jabatan == "ketua upz" {
			return p.Nama
		}
	}
	for _, p := range pengurusZakat {
		if strings.Contains(strings.ToLower(strings.TrimSpace(p.Jabatan)), "ketua") {
			return p.Nama
		}
	}
	return "-"
}

func (h *TransaksiHandler) Create(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	var transaksi models.TransaksiZakat
	if err := c.ShouldBindJSON(&transaksi); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	transaksi.MasjidID = *masjidID.(*int)
	transaksi.Tahun = time.Now().Year()

	// Hitung total wajib dan infaq tambahan (server-side source of truth)
	originalTotalWajib := transaksi.TotalWajib
	transaksi.TotalWajib = 0
	transaksi.InfaqTambahan = 0

	switch transaksi.JenisZakat {
	case "fitrah":
		if transaksi.BentukZakat != nil && *transaksi.BentukZakat == "uang" {
			transaksi.TotalWajib = transaksi.NominalPerOrang * float64(transaksi.JumlahOrang)
		}
	case "mal":
		if transaksi.NominalHarta != nil && transaksi.PersentaseZakat != nil {
			transaksi.TotalWajib = (*transaksi.NominalHarta) * (*transaksi.PersentaseZakat) / 100
		}
	case "fidyah":
		if transaksi.BentukZakat != nil && *transaksi.BentukZakat == "uang" {
			jumlahHari := 0
			if transaksi.JumlahHariFidyah != nil {
				jumlahHari = *transaksi.JumlahHariFidyah
			}
			if jumlahHari <= 0 {
				jumlahHari = transaksi.JumlahOrang
			}
			transaksi.TotalWajib = transaksi.NominalPerOrang * float64(jumlahHari)
		}
	}

	if transaksi.TotalWajib <= 0 && originalTotalWajib > 0 {
		transaksi.TotalWajib = originalTotalWajib
	}

	if transaksi.TotalWajib > 0 && transaksi.TotalDibayar > transaksi.TotalWajib {
		transaksi.InfaqTambahan = transaksi.TotalDibayar - transaksi.TotalWajib
	}

	// Atomic create: ensure muzakki + transaksi saved together in one DB transaction.
	// This prevents partial writes when client/network fails between separate requests.
	tx, err := h.transaksiRepo.DB.Begin()
	if err != nil {
		log.Printf("Create transaksi: begin tx failed: %v", err)
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to start transaction"})
		return
	}
	defer func() { _ = tx.Rollback() }()

	if transaksi.MuzakkiID <= 0 {
		nama := strings.TrimSpace(transaksi.MuzakkiNama)
		if nama == "" {
			c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Muzakki wajib diisi"})
			return
		}

		var existingID int
		err := tx.QueryRow(
			`SELECT id FROM muzakki WHERE masjid_id = ? AND LOWER(TRIM(nama)) = LOWER(TRIM(?)) LIMIT 1`,
			transaksi.MasjidID,
			nama,
		).Scan(&existingID)
		if err == nil && existingID > 0 {
			transaksi.MuzakkiID = existingID
		} else if err != nil && err != sql.ErrNoRows {
			log.Printf("Create transaksi: lookup muzakki failed (masjid_id=%d nama=%q): %v", transaksi.MasjidID, nama, err)
			c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to lookup muzakki"})
			return
		} else {
			res, err := tx.Exec(
				`INSERT INTO muzakki (masjid_id, nama, alamat, telepon) VALUES (?, ?, ?, ?)`,
				transaksi.MasjidID,
				nama,
				strings.TrimSpace(transaksi.MuzakkiAlamat),
				strings.TrimSpace(transaksi.MuzakkiTelepon),
			)
			if err != nil {
				log.Printf("Create transaksi: insert muzakki failed (masjid_id=%d nama=%q): %v", transaksi.MasjidID, nama, err)
				c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to create muzakki"})
				return
			}
			id, _ := res.LastInsertId()
			transaksi.MuzakkiID = int(id)
		}
	}

	result, err := tx.Exec(
		`INSERT INTO transaksi_zakat (masjid_id, muzakki_id, jenis_zakat, bentuk_zakat, jenis_harta,
			  nominal_harta, persentase_zakat, kelas_zakat, jumlah_orang, jumlah_hari_fidyah,
			  standar_beras_per_jiwa, kg_beras_dibayar, harga_beras_per_kg, nominal_per_orang, total_wajib,
			  total_dibayar, infaq_tambahan, keterangan, tahun, tanggal_bayar)
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		transaksi.MasjidID,
		transaksi.MuzakkiID,
		transaksi.JenisZakat,
		transaksi.BentukZakat,
		transaksi.JenisHarta,
		transaksi.NominalHarta,
		transaksi.PersentaseZakat,
		transaksi.KelasZakat,
		transaksi.JumlahOrang,
		transaksi.JumlahHariFidyah,
		transaksi.StandarBerasPerJiwa,
		transaksi.KgBerasDibayar,
		transaksi.HargaBerasPerKg,
		transaksi.NominalPerOrang,
		transaksi.TotalWajib,
		transaksi.TotalDibayar,
		transaksi.InfaqTambahan,
		transaksi.Keterangan,
		transaksi.Tahun,
		transaksi.TanggalBayar,
	)
	if err != nil {
		log.Printf("Create transaksi: insert transaksi failed (masjid_id=%d muzakki_id=%d jenis=%q): %v", transaksi.MasjidID, transaksi.MuzakkiID, transaksi.JenisZakat, err)
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to create transaksi"})
		return
	}
	newID, _ := result.LastInsertId()
	transaksi.ID = int(newID)

	if err := tx.Commit(); err != nil {
		log.Printf("Create transaksi: commit failed (masjid_id=%d muzakki_id=%d): %v", transaksi.MasjidID, transaksi.MuzakkiID, err)
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusCreated, models.Response{Success: true, Message: "Transaksi created successfully", Data: transaksi})
}

func (h *TransaksiHandler) GetAll(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	jenisZakat := strings.TrimSpace(c.Query("jenis_zakat"))
	bentukZakat := strings.TrimSpace(c.Query("bentuk_zakat"))
	muzakkiQuery := strings.TrimSpace(c.Query("q"))

	if jenisZakat != "fitrah" && jenisZakat != "fidyah" {
		bentukZakat = ""
	}

	page, pageSize, isAll := parsePagination(c)

	if isAll {
		transaksis, err := h.transaksiRepo.GetByMasjidIDFiltered(*masjidID.(*int), jenisZakat, bentukZakat, muzakkiQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get transaksis"})
			return
		}
		pagination := buildPagination(len(transaksis), page, pageSize, true)
		c.JSON(http.StatusOK, models.Response{Success: true, Data: gin.H{"items": transaksis, "pagination": pagination}})
		return
	}

	total, err := h.transaksiRepo.CountByMasjidIDFiltered(*masjidID.(*int), jenisZakat, bentukZakat, muzakkiQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get transaksis"})
		return
	}

	offset := (page - 1) * pageSize
	transaksis, err := h.transaksiRepo.GetByMasjidIDPaginatedFiltered(*masjidID.(*int), jenisZakat, bentukZakat, muzakkiQuery, pageSize, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get transaksis"})
		return
	}

	fmt.Printf("GetAll Transaksi - MasjidID: %d, Count: %d\n", *masjidID.(*int), len(transaksis))
	pagination := buildPagination(total, page, pageSize, false)
	c.JSON(http.StatusOK, models.Response{Success: true, Data: gin.H{"items": transaksis, "pagination": pagination}})
}

func (h *TransaksiHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	transaksi, err := h.transaksiRepo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{Success: false, Message: "Transaksi not found"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: transaksi})
}

func (h *TransaksiHandler) Delete(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	id, _ := strconv.Atoi(c.Param("id"))
	if id <= 0 {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid transaksi id"})
		return
	}

	tx, err := h.transaksiRepo.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to start transaction"})
		return
	}
	defer func() { _ = tx.Rollback() }()

	var muzakkiID int
	err = tx.QueryRow(
		`SELECT muzakki_id FROM transaksi_zakat WHERE id = ? AND masjid_id = ? LIMIT 1`,
		id,
		*masjidID.(*int),
	).Scan(&muzakkiID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, models.Response{Success: false, Message: "Transaksi not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to lookup transaksi"})
		return
	}

	res, err := tx.Exec(`DELETE FROM transaksi_zakat WHERE id = ? AND masjid_id = ?`, id, *masjidID.(*int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to delete transaksi"})
		return
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		c.JSON(http.StatusNotFound, models.Response{Success: false, Message: "Transaksi not found"})
		return
	}

	// Delete the linked muzakki as well (domain rule: 1 muzakki record is tied to 1 transaksi).
	// Note: FK cascade means deleting muzakki will also delete any remaining transaksi that still reference it.
	if _, err := tx.Exec(`DELETE FROM muzakki WHERE id = ? AND masjid_id = ?`, muzakkiID, *masjidID.(*int)); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to delete muzakki"})
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "Transaksi deleted successfully"})
}

func (h *TransaksiHandler) PrintReceipt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	namaPetugas := "Petugas"
	if username, exists := c.Get("username"); exists {
		if uname, ok := username.(string); ok && uname != "" {
			namaPetugas = uname
		}
	}

	transaksi, err := h.transaksiRepo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{Success: false, Message: "Transaksi not found"})
		return
	}

	// Generate HTML receipt
	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Bukti Pembayaran Zakat</title>
    <style>
        body { font-family: Arial, sans-serif; padding: 20px; }
        .header { text-align: center; margin-bottom: 30px; }
        .content { margin: 20px 0; }
        .row { display: flex; margin: 10px 0; }
        .label { width: 200px; font-weight: bold; }
        .value { flex: 1; }
        .footer { margin-top: 50px; text-align: right; }
        @media print { button { display: none; } }
    </style>
</head>
<body>
    <div class="header">
        <h2>BUKTI PEMBAYARAN ZAKAT</h2>
        <p>Tahun %d</p>
    </div>
    <div class="content">
        <div class="row"><div class="label">Nama Muzakki:</div><div class="value">%s</div></div>
        <div class="row"><div class="label">Alamat:</div><div class="value">%s</div></div>
        <div class="row"><div class="label">Jenis Zakat:</div><div class="value">%s</div></div>
        <div class="row"><div class="label">Bentuk Zakat:</div><div class="value">%s</div></div>
        <div class="row"><div class="label">Jumlah Orang:</div><div class="value">%d</div></div>
        <div class="row"><div class="label">Total Wajib:</div><div class="value">Rp %.2f</div></div>
        <div class="row"><div class="label">Total Dibayar:</div><div class="value">Rp %.2f</div></div>
        <div class="row"><div class="label">Infaq Tambahan:</div><div class="value">Rp %.2f</div></div>
        <div class="row"><div class="label">Tanggal:</div><div class="value">%s</div></div>
    </div>
    <div class="footer">
        <p>Petugas Zakat</p>
        <br><br>
        <p>( %s )</p>
    </div>
    <button onclick="window.print()">Cetak</button>
</body>
</html>
	`, transaksi.Tahun, transaksi.MuzakkiNama, transaksi.MuzakkiAlamat, transaksi.JenisZakat,
		*transaksi.BentukZakat, transaksi.JumlahOrang, transaksi.TotalWajib, transaksi.TotalDibayar,
		transaksi.InfaqTambahan, transaksi.TanggalBayar, namaPetugas)

	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, html)
}

func (h *TransaksiHandler) PrintSummary(c *gin.Context) {
	masjidIDValue, _ := c.Get("masjid_id")
	masjidID := *masjidIDValue.(*int)

	signDateRaw := strings.TrimSpace(c.Query("sign_date"))
	signDate := time.Now()
	if signDateRaw != "" {
		parsed, err := time.Parse("2006-01-02", signDateRaw)
		if err == nil {
			signDate = parsed
		}
	}

	tahun := signDate.Year()

	masjid, err := h.masjidRepo.FindByID(masjidID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{Success: false, Message: "Masjid not found"})
		return
	}

	pengurusMasjid, _ := h.pengurusRepo.GetPengurusMasjidByMasjidID(masjidID)
	pengurusZakat, _ := h.pengurusRepo.GetPengurusZakatByMasjidID(masjidID)

	ketuaPengurus := pickKetuaPengurusNama(pengurusMasjid)
	ketuaUPZ := pickKetuaUPZNama(pengurusZakat)

	transaksis, err := h.transaksiRepo.GetByMasjidIDAndYear(masjidID, tahun)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get transaksi summary"})
		return
	}

	type recapRow struct {
		No             int
		JenisLabel     string
		Muzakki        int
		JumlahDizakati int
		Uang           float64
		Beras          float64
	}

	rowTemplate := map[string]*recapRow{
		"fitrah": {No: 1, JenisLabel: "Zakat Fitrah"},
		"mal":    {No: 2, JenisLabel: "Zakat Mal"},
		"fidyah": {No: 3, JenisLabel: "Fidyah"},
		"infaq":  {No: 4, JenisLabel: "Infaq"},
	}
	muzakkiSet := map[string]map[int]struct{}{
		"fitrah": {},
		"mal":    {},
		"fidyah": {},
		"infaq":  {},
	}

	for _, transaksi := range transaksis {
		jenis := strings.ToLower(strings.TrimSpace(transaksi.JenisZakat))
		row, exists := rowTemplate[jenis]
		if !exists {
			continue
		}

		if transaksi.MuzakkiID > 0 {
			muzakkiSet[jenis][transaksi.MuzakkiID] = struct{}{}
		}

		switch jenis {
		case "fitrah":
			if transaksi.JumlahOrang > 0 {
				row.JumlahDizakati += transaksi.JumlahOrang
			}
		case "fidyah":
			if transaksi.JumlahHariFidyah != nil && *transaksi.JumlahHariFidyah > 0 {
				row.JumlahDizakati += *transaksi.JumlahHariFidyah
			} else if transaksi.JumlahOrang > 0 {
				row.JumlahDizakati += transaksi.JumlahOrang
			}
		case "mal", "infaq":
			// kolom jumlah dizakati tidak dipakai
		}

		infaqTambahan := transaksi.InfaqTambahan
		if infaqTambahan < 0 {
			infaqTambahan = 0
		}

		if jenis == "infaq" {
			row.Uang += transaksi.TotalDibayar
		} else {
			nilaiZakatMurni := transaksi.TotalDibayar - infaqTambahan
			if nilaiZakatMurni < 0 {
				nilaiZakatMurni = 0
			}
			row.Uang += nilaiZakatMurni

			if infaqTambahan > 0 {
				rowTemplate["infaq"].Uang += infaqTambahan
				if transaksi.MuzakkiID > 0 {
					muzakkiSet["infaq"][transaksi.MuzakkiID] = struct{}{}
				}
			}
		}

		kg := 0.0
		if transaksi.KgBerasDibayar != nil && *transaksi.KgBerasDibayar > 0 {
			kg = *transaksi.KgBerasDibayar
		} else {
			kg = parsePrintKgFromKeterangan(transaksi.Keterangan)
		}
		row.Beras += kg
	}

	rows := []recapRow{}
	totalMuzakki := 0
	totalDizakati := 0
	totalUang := 0.0
	totalBeras := 0.0
	for _, key := range []string{"fitrah", "mal", "fidyah", "infaq"} {
		row := rowTemplate[key]
		row.Muzakki = len(muzakkiSet[key])
		if key == "fidyah" {
			row.JumlahDizakati = row.Muzakki
		}
		rows = append(rows, *row)

		totalMuzakki += row.Muzakki
		totalDizakati += row.JumlahDizakati
		totalUang += row.Uang
		totalBeras += row.Beras
	}

	type recapTemplateData struct {
		MasjidNama        string
		MasjidNamaUpper   string
		MasjidNamaHeader  string
		AlamatMasjid      string
		UpzFontSize       string
		UpzFontWeight     int
		AreaFontSize      string
		AreaFontWeight    int
		AddressFontSize   string
		AddressFontWeight int
		Tahun             int
		TanggalTtd        string
		Rows              []recapRow
		TotalMuzakki      int
		TotalDizakati     int
		TotalUang         float64
		TotalBeras        float64
		KetuaUPZ          string
		KetuaPengurus     string
	}

	data := recapTemplateData{
		MasjidNama:        strings.TrimSpace(masjid.Nama),
		MasjidNamaUpper:   strings.ToUpper(strings.TrimSpace(masjid.Nama)),
		MasjidNamaHeader:  strings.ToUpper(normalizeMasjidTitleName(masjid.Nama)),
		AlamatMasjid:      strings.TrimSpace(masjid.Alamat),
		UpzFontSize:       "18px",
		UpzFontWeight:     700,
		AreaFontSize:      "14px",
		AreaFontWeight:    700,
		AddressFontSize:   "11px",
		AddressFontWeight: 300,
		Tahun:             tahun,
		TanggalTtd:        formatSignedDateID(signDate),
		Rows:              rows,
		TotalMuzakki:      totalMuzakki,
		TotalDizakati:     totalDizakati,
		TotalUang:         totalUang,
		TotalBeras:        totalBeras,
		KetuaUPZ:          ketuaUPZ,
		KetuaPengurus:     ketuaPengurus,
	}

	htmlTemplate := `<!DOCTYPE html>
<html lang="id">
<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Laporan Zakat, Infaq, dan Fidyah</title>
    <style>
        body { font-family: Arial, sans-serif; color: #000; margin: 0; padding: 24px 28px; }
        .wrapper { max-width: 980px; margin: 0 auto; }
		.text-center { text-align: center; }
		.header-upz { font-size: {{ .UpzFontSize }}; font-weight: {{ .UpzFontWeight }}; margin: 2px 0; text-transform: uppercase; }
		.header-area { font-size: {{ .AreaFontSize }}; font-weight: {{ .AreaFontWeight }}; margin: 2px 0; text-transform: uppercase; }
		.header-address { font-size: {{ .AddressFontSize }}; font-weight: {{ .AddressFontWeight }}; margin: 2px 0 8px; }
        .divider { border-top: 2px solid #000; margin: 8px 0 14px; }
        .report-title { font-size: 12px; font-weight: 700; margin: 0; text-transform: uppercase; }
        .report-year { font-size: 12px; font-weight: 700; margin: 4px 0 16px; text-transform: uppercase; }
        table { width: 100%; border-collapse: collapse; font-size: 11px; }
        th, td { border: 1px solid #000; padding: 4px 4px; }
        th { background: #f5f5f5; font-weight: 700; }
        .right { text-align: right; }
        .center { text-align: center; }
        .left { text-align: left; }
        .money-cell { display: flex; justify-content: space-between; gap: 8px; }
        .total-row td { font-weight: 700; }
        .signature-row { margin-top: 26px; display: flex; justify-content: space-between; gap: 24px; font-size: 12px; }
        .signature-col { width: 45%; }
		.signature-head { min-height: 18px; display: flex; flex-direction: column; justify-content: flex-start; }
		.signature-line { margin: 0; line-height: 1.2; }
        .signature-name { margin-top: 70px; font-weight: 700; text-transform: uppercase; text-decoration: underline; }
        @media print {
          body { padding: 0; }
          .wrapper { max-width: none; }
        }
    </style>
</head>
<body>
    <div class="wrapper">
		<div class="text-center">
			<p class="header-upz">UNIT PENGUMPUL ZAKAT (UPZ) {{ .MasjidNamaHeader }}</p>
			<p class="header-area">DESA PURWAJAYA KECAMATAN LOA JANAN</p>
			<p class="header-address">Alamat: {{ .AlamatMasjid }}</p>
        </div>
        <div class="divider"></div>
        <div class="text-center">
            <p class="report-title">LAPORAN ZAKAT, INFAQ, DAN FIDYAH</p>
            <p class="report-year">TAHUN {{ .Tahun }}</p>
        </div>

        <table>
            <thead>
                <tr>
                    <th rowspan="2" style="width: 42px;">No</th>
                    <th rowspan="2" style="width: 150px;">Jenis Zakat</th>
                    <th rowspan="2" style="width: 100px;">Muzakki (orang)</th>
                    <th rowspan="2" style="width: 160px;">Jumlah Orang yang Dizakati (orang)</th>
                    <th colspan="2">Jumlah Zakat/Infaq</th>
                </tr>
                <tr>
                    <th style="width: 180px;">Uang (Rp)</th>
                    <th style="width: 120px;">Beras (Kg)</th>
                </tr>
            </thead>
            <tbody>
                {{ range .Rows }}
                <tr>
                    <td class="center">{{ .No }}</td>
                    <td class="left">{{ .JenisLabel }}</td>
                    <td class="right">{{ formatInt .Muzakki }}</td>
                    <td class="right">{{ formatInt .JumlahDizakati }}</td>
                    <td class="right">
                      <div class="money-cell"><span>Rp</span><span>{{ formatMoney .Uang }}</span></div>
                    </td>
                    <td class="right">{{ formatKg .Beras }}</td>
                </tr>
                {{ end }}
                <tr class="total-row">
                    <td colspan="2" class="center">TOTAL</td>
                    <td class="right">{{ formatInt .TotalMuzakki }}</td>
                    <td class="right">{{ formatInt .TotalDizakati }}</td>
                    <td class="right">
                      <div class="money-cell"><span>Rp</span><span>{{ formatMoney .TotalUang }}</span></div>
                    </td>
                    <td class="right">{{ formatKg .TotalBeras }}</td>
                </tr>
            </tbody>
        </table>

		<div class="signature-row">
            <div class="signature-col">
				<div class="signature-head">
					<p class="signature-line">Ketua UPZ {{ .MasjidNama }}</p>
				</div>
                <p class="signature-name">{{ .KetuaUPZ }}</p>
            </div>
            <div class="signature-col">
				<div class="signature-head">
					<p class="signature-line">Purwajaya, {{ .TanggalTtd }}</p>
					<p class="signature-line">Ketua Pengurus {{ .MasjidNama }}</p>
				</div>
                <p class="signature-name">{{ .KetuaPengurus }}</p>
            </div>
        </div>
    </div>
    <script>window.onload = function(){ window.print(); };</script>
</body>
</html>`

	tmpl, err := template.New("print-summary").Funcs(template.FuncMap{
		"formatInt": func(v int) string {
			return formatIDNumber(int64(v))
		},
		"formatMoney": func(v float64) string {
			return formatIDNumber(int64(v + 0.5))
		},
		"formatKg": func(v float64) string {
			if v == float64(int64(v)) {
				return formatIDNumber(int64(v))
			}
			str := strconv.FormatFloat(v, 'f', 2, 64)
			parts := strings.Split(str, ".")
			left := parts[0]
			right := parts[1]
			right = strings.TrimRight(right, "0")
			if right == "" {
				return formatIDNumber(int64(v))
			}
			leftInt, _ := strconv.ParseInt(left, 10, 64)
			return formatIDNumber(leftInt) + "," + right
		},
	}).Parse(htmlTemplate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to build print template"})
		return
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")
	c.Status(http.StatusOK)
	_ = tmpl.Execute(c.Writer, data)
}

func (h *TransaksiHandler) GetDashboard(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")

	// Implementasi dashboard stats akan ditambahkan
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data: gin.H{
			"masjid_id": *masjidID.(*int),
			"message":   "Dashboard data",
		},
	})
}
