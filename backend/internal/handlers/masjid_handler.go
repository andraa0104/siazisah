package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/siazisah/internal/models"
	"github.com/yourusername/siazisah/internal/repository"
)

type MasjidHandler struct {
	masjidRepo *repository.MasjidRepository
}

func NewMasjidHandler(masjidRepo *repository.MasjidRepository) *MasjidHandler {
	return &MasjidHandler{masjidRepo: masjidRepo}
}

func (h *MasjidHandler) Create(c *gin.Context) {
	var masjid models.Masjid
	if err := c.ShouldBindJSON(&masjid); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	masjid.IsActive = true
	if err := h.masjidRepo.Create(&masjid); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to create masjid"})
		return
	}

	c.JSON(http.StatusCreated, models.Response{Success: true, Message: "Masjid created successfully", Data: masjid})
}

func (h *MasjidHandler) GetAll(c *gin.Context) {
	page, pageSize, isAll := parsePagination(c)

	if isAll {
		masjids, err := h.masjidRepo.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get masjids"})
			return
		}
		pagination := buildPagination(len(masjids), page, pageSize, true)
		c.JSON(http.StatusOK, models.Response{Success: true, Data: gin.H{"items": masjids, "pagination": pagination}})
		return
	}

	total, err := h.masjidRepo.CountAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get masjids"})
		return
	}

	offset := (page - 1) * pageSize
	masjids, err := h.masjidRepo.GetAllPaginated(pageSize, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get masjids"})
		return
	}

	pagination := buildPagination(total, page, pageSize, false)
	c.JSON(http.StatusOK, models.Response{Success: true, Data: gin.H{"items": masjids, "pagination": pagination}})
}

func (h *MasjidHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	masjid, err := h.masjidRepo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{Success: false, Message: "Masjid not found"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: masjid})
}

func (h *MasjidHandler) GetMyMasjid(c *gin.Context) {
	masjidID, exists := c.Get("masjid_id")
	log.Printf("[DEBUG] GetMyMasjid - exists: %v, masjidID: %v", exists, masjidID)

	if !exists {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Masjid ID not found in token"})
		return
	}

	if masjidID == nil {
		c.JSON(http.StatusOK, models.Response{Success: true, Data: nil, Message: "User not assigned to masjid yet"})
		return
	}

	id, ok := masjidID.(*int)
	log.Printf("[DEBUG] GetMyMasjid - type assertion ok: %v, id: %v", ok, id)

	if !ok || id == nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid masjid ID format"})
		return
	}

	log.Printf("[DEBUG] GetMyMasjid - Finding masjid with ID: %d", *id)
	masjid, err := h.masjidRepo.FindByID(*id)
	if err != nil {
		log.Printf("[ERROR] GetMyMasjid - Error finding masjid: %v", err)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, models.Response{Success: false, Message: "Masjid not found"})
		} else {
			c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Error retrieving masjid"})
		}
		return
	}

	log.Printf("[DEBUG] GetMyMasjid - Found masjid: %s", masjid.Nama)
	c.JSON(http.StatusOK, models.Response{Success: true, Data: masjid})
}

func (h *MasjidHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var masjid models.Masjid
	if err := c.ShouldBindJSON(&masjid); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	masjid.ID = id
	if err := h.masjidRepo.Update(&masjid); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to update masjid"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "Masjid updated successfully", Data: masjid})
}

func (h *MasjidHandler) UpdateMyMasjid(c *gin.Context) {
	masjidID, exists := c.Get("masjid_id")
	if !exists || masjidID == nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Masjid ID not found"})
		return
	}

	var masjid models.Masjid
	if err := c.ShouldBindJSON(&masjid); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	id := masjidID.(*int)
	masjid.ID = *id
	if err := h.masjidRepo.Update(&masjid); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to update masjid"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "Masjid updated successfully", Data: masjid})
}

func (h *MasjidHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.masjidRepo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to delete masjid"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "Masjid deleted successfully"})
}

func (h *MasjidHandler) PrintZakatSummary(c *gin.Context) {
	signDateRaw := strings.TrimSpace(c.Query("sign_date"))
	signDate := time.Now()
	if signDateRaw != "" {
		parsed, err := time.Parse("2006-01-02", signDateRaw)
		if err == nil {
			signDate = parsed
		}
	}

	tahun := signDate.Year()

	masjids, err := h.masjidRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get masjid"})
		return
	}

	type reportRow struct {
		No              int
		MasjidID        int
		NamaMasjid      string
		JumlahMuzakki   int
		JumlahDizakati  int
		FitrahUang      float64
		FitrahBeras     float64
		MalUang         float64
		FidyahUang      float64
		FidyahBeras     float64
		TotalZakatUang  float64
		TotalZakatBeras float64
		TotalInfaq      float64
		muzakkiSet      map[int]struct{}
	}

	rows := make([]*reportRow, 0, len(masjids))
	rowsByMasjidID := map[int]*reportRow{}
	for idx, masjid := range masjids {
		row := &reportRow{
			No:         idx + 1,
			MasjidID:   masjid.ID,
			NamaMasjid: strings.TrimSpace(masjid.Nama),
			muzakkiSet: map[int]struct{}{},
		}
		rows = append(rows, row)
		rowsByMasjidID[masjid.ID] = row
	}

	query := `SELECT
			masjid_id,
			COALESCE(jenis_zakat, ''),
			COALESCE(bentuk_zakat, ''),
			COALESCE(total_dibayar, 0),
			COALESCE(infaq_tambahan, 0),
			COALESCE(kg_beras_dibayar, 0),
			COALESCE(keterangan, ''),
			COALESCE(jumlah_orang, 0),
			COALESCE(jumlah_hari_fidyah, 0),
			COALESCE(muzakki_id, 0)
		FROM transaksi_zakat
		WHERE tahun = ?`

	txRows, err := h.masjidRepo.DB.Query(query, tahun)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get transaksi summary"})
		return
	}
	defer txRows.Close()

	for txRows.Next() {
		var masjidID int
		var jenis string
		var bentuk string
		var totalDibayar float64
		var infaqTambahan float64
		var kgBeras float64
		var keterangan string
		var jumlahOrang int
		var jumlahHariFidyah int
		var muzakkiID int

		if err := txRows.Scan(
			&masjidID,
			&jenis,
			&bentuk,
			&totalDibayar,
			&infaqTambahan,
			&kgBeras,
			&keterangan,
			&jumlahOrang,
			&jumlahHariFidyah,
			&muzakkiID,
		); err != nil {
			continue
		}

		row, exists := rowsByMasjidID[masjidID]
		if !exists {
			continue
		}

		if muzakkiID > 0 {
			row.muzakkiSet[muzakkiID] = struct{}{}
		}

		jenis = strings.ToLower(strings.TrimSpace(jenis))
		bentuk = strings.ToLower(strings.TrimSpace(bentuk))

		if infaqTambahan < 0 {
			infaqTambahan = 0
		}

		if jenis == "infaq" {
			row.TotalInfaq += totalDibayar
		} else {
			nilaiZakatMurni := totalDibayar - infaqTambahan
			if nilaiZakatMurni < 0 {
				nilaiZakatMurni = 0
			}

			switch jenis {
			case "fitrah":
				if bentuk == "uang" {
					row.FitrahUang += nilaiZakatMurni
				}
				if bentuk == "beras" {
					if kgBeras <= 0 {
						kgBeras = parsePrintKgFromKeterangan(keterangan)
					}
					row.FitrahBeras += kgBeras
				}
				if jumlahOrang > 0 {
					row.JumlahDizakati += jumlahOrang
				}
			case "mal":
				row.MalUang += nilaiZakatMurni
			case "fidyah":
				if bentuk == "uang" {
					row.FidyahUang += nilaiZakatMurni
				}
				if bentuk == "beras" {
					if kgBeras <= 0 {
						kgBeras = parsePrintKgFromKeterangan(keterangan)
					}
					row.FidyahBeras += kgBeras
				}
				if jumlahHariFidyah > 0 {
					row.JumlahDizakati += jumlahHariFidyah
				} else if jumlahOrang > 0 {
					row.JumlahDizakati += jumlahOrang
				}
			}

			if infaqTambahan > 0 {
				row.TotalInfaq += infaqTambahan
			}
		}
	}

	type finalRow struct {
		No              int
		NamaMasjid      string
		JumlahMuzakki   int
		JumlahDizakati  int
		FitrahUang      float64
		FitrahBeras     float64
		MalUang         float64
		FidyahUang      float64
		FidyahBeras     float64
		TotalZakatUang  float64
		TotalZakatBeras float64
		TotalInfaq      float64
	}

	finalRows := make([]finalRow, 0, len(rows))
	grand := finalRow{}
	for _, row := range rows {
		row.JumlahMuzakki = len(row.muzakkiSet)
		row.TotalZakatUang = row.FitrahUang + row.MalUang + row.FidyahUang
		row.TotalZakatBeras = row.FitrahBeras + row.FidyahBeras

		finalRows = append(finalRows, finalRow{
			No:              row.No,
			NamaMasjid:      row.NamaMasjid,
			JumlahMuzakki:   row.JumlahMuzakki,
			JumlahDizakati:  row.JumlahDizakati,
			FitrahUang:      row.FitrahUang,
			FitrahBeras:     row.FitrahBeras,
			MalUang:         row.MalUang,
			FidyahUang:      row.FidyahUang,
			FidyahBeras:     row.FidyahBeras,
			TotalZakatUang:  row.TotalZakatUang,
			TotalZakatBeras: row.TotalZakatBeras,
			TotalInfaq:      row.TotalInfaq,
		})

		grand.JumlahMuzakki += row.JumlahMuzakki
		grand.JumlahDizakati += row.JumlahDizakati
		grand.FitrahUang += row.FitrahUang
		grand.FitrahBeras += row.FitrahBeras
		grand.MalUang += row.MalUang
		grand.FidyahUang += row.FidyahUang
		grand.FidyahBeras += row.FidyahBeras
		grand.TotalZakatUang += row.TotalZakatUang
		grand.TotalZakatBeras += row.TotalZakatBeras
		grand.TotalInfaq += row.TotalInfaq
	}

	type templateData struct {
		Tahun                 int
		TanggalTtd            string
		Rows                  []finalRow
		Grand                 finalRow
		ColNoWidth            string
		ColMasjidWidth        string
		ColMuzakkiWidth       string
		ColDizakatiWidth      string
		ColFitrahUangWidth    string
		ColFitrahBerasWidth   string
		ColMalUangWidth       string
		ColFidyahUangWidth    string
		ColFidyahBerasWidth   string
		ColTotalUangWidth     string
		ColTotalBerasWidth    string
		ColTotalInfaqWidth    string
		MainTitleFontSize     string
		MainTitleFontWeight   int
		AreaTitleFontSize     string
		AreaTitleFontWeight   int
		AddressFontSize       string
		AddressFontWeight     int
		ReportTitleFontSize   string
		ReportTitleFontWeight int
		YearTitleFontSize     string
		YearTitleFontWeight   int
	}

	data := templateData{
		Tahun:                 tahun,
		TanggalTtd:            formatSignedDateID(signDate),
		Rows:                  finalRows,
		Grand:                 grand,
		ColNoWidth:            "30px",
		ColMasjidWidth:        "130px",
		ColMuzakkiWidth:       "70px",
		ColDizakatiWidth:      "70px",
		ColFitrahUangWidth:    "110px",
		ColFitrahBerasWidth:   "50px",
		ColMalUangWidth:       "90px",
		ColFidyahUangWidth:    "90px",
		ColFidyahBerasWidth:   "50px",
		ColTotalUangWidth:     "110px",
		ColTotalBerasWidth:    "50px",
		ColTotalInfaqWidth:    "80px",
		MainTitleFontSize:     "30px",
		MainTitleFontWeight:   800,
		AreaTitleFontSize:     "25px",
		AreaTitleFontWeight:   600,
		AddressFontSize:       "16px",
		AddressFontWeight:     300,
		ReportTitleFontSize:   "20px",
		ReportTitleFontWeight: 700,
		YearTitleFontSize:     "20px",
		YearTitleFontWeight:   700,
	}

	htmlTemplate := `<!DOCTYPE html>
<html lang="id">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Laporan Data Zakat</title>
  <style>
    body { font-family: Arial, sans-serif; color: #000; margin: 0; padding: 24px 28px; }
    .wrapper { max-width: 1400px; margin: 0 auto; }
    .text-center { text-align: center; }
			.title-main { font-size: {{ .MainTitleFontSize }}; font-weight: {{ .MainTitleFontWeight }}; margin: 2px 0; text-transform: uppercase; }
			.title-area { font-size: {{ .AreaTitleFontSize }}; font-weight: {{ .AreaTitleFontWeight }}; margin: 2px 0; text-transform: uppercase; }
			.title-address { font-size: {{ .AddressFontSize }}; font-weight: {{ .AddressFontWeight }}; margin: 2px 0 8px; }
    .divider { border-top: 3px solid #000; margin: 8px 0 14px; }
			.report-title { font-size: {{ .ReportTitleFontSize }}; font-weight: {{ .ReportTitleFontWeight }}; margin: 0; text-transform: uppercase; }
			.report-year { font-size: {{ .YearTitleFontSize }}; font-weight: {{ .YearTitleFontWeight }}; margin: 4px 0 16px; text-transform: uppercase; }
    table { width: 100%; border-collapse: collapse; font-size: 13px; }
    th, td { border: 1px solid #000; padding: 4px 4px; }
    th { background: #f5f5f5; font-weight: 700; }
    .right { text-align: right; }
    .center { text-align: center; }
    .left { text-align: left; }
    .money-cell { display: flex; justify-content: space-between; gap: 8px; }
    .total-row td { font-weight: 700; }
	.signature-row { margin-top: 38px; display: flex; justify-content: space-between; gap: 24px; font-size: 14px; }
	.signature-col { width: 45%; }
	.signature-head { min-height: 42px; display: flex; flex-direction: column; justify-content: flex-start; }
	.signature-line { margin: 0; line-height: 1.3; }
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
      <p class="title-main">PENGURUS PERINGATAN HARI BESAR ISLAM</p>
      <p class="title-area">DESA PURWAJAYA KECAMATAN LOA JANAN</p>
      <p class="title-address">Alamat: Jalan Pembangunan RT. 03 Desa Purwajaya</p>
    </div>
    <div class="divider"></div>
    <div class="text-center">
      <p class="report-title">LAPORAN ZAKAT, INFAQ, DAN FIDYAH</p>
      <p class="report-year">TAHUN {{ .Tahun }}</p>
    </div>

		<table>
			<colgroup>
				<col style="width: {{ .ColNoWidth }};" />
				<col style="width: {{ .ColMasjidWidth }};" />
				<col style="width: {{ .ColMuzakkiWidth }};" />
				<col style="width: {{ .ColDizakatiWidth }};" />
				<col style="width: {{ .ColFitrahUangWidth }};" />
				<col style="width: {{ .ColFitrahBerasWidth }};" />
				<col style="width: {{ .ColMalUangWidth }};" />
				<col style="width: {{ .ColFidyahUangWidth }};" />
				<col style="width: {{ .ColFidyahBerasWidth }};" />
				<col style="width: {{ .ColTotalUangWidth }};" />
				<col style="width: {{ .ColTotalBerasWidth }};" />
				<col style="width: {{ .ColTotalInfaqWidth }};" />
			</colgroup>
			<thead>
				<tr>
					<th rowspan="3">No</th>
					<th rowspan="3">Nama Masjid/Langgar</th>
					<th rowspan="3">Jumlah Muzakki (orang)</th>
					<th rowspan="3">Jumlah Orang yang Dizakati (orang)</th>
					<th colspan="5">Jenis Zakat</th>
					<th colspan="2">Total Zakat dan Fidyah</th>
					<th rowspan="3">Total Infaq (Rp)</th>
				</tr>
				<tr>
					<th colspan="2">Fitrah</th>
					<th rowspan="2">Mal<br/>Uang (Rp)</th>
					<th colspan="2">Fidyah</th>
					<th rowspan="2">Uang (Rp)</th>
					<th rowspan="2">Beras (kg)</th>
				</tr>
				<tr>
					<th>Uang (Rp)</th>
					<th>Beras (kg)</th>
					<th>Uang (Rp)</th>
					<th>Beras (kg)</th>
				</tr>
			</thead>
      <tbody>
        {{ range .Rows }}
        <tr>
          <td class="center">{{ .No }}</td>
          <td class="left">{{ .NamaMasjid }}</td>
          <td class="right">{{ formatInt .JumlahMuzakki }}</td>
          <td class="right">{{ formatInt .JumlahDizakati }}</td>
          <td class="right">{{ moneyOrDash .FitrahUang }}</td>
          <td class="right">{{ kgOrDash .FitrahBeras }}</td>
          <td class="right">{{ moneyOrDash .MalUang }}</td>
          <td class="right">{{ moneyOrDash .FidyahUang }}</td>
          <td class="right">{{ kgOrDash .FidyahBeras }}</td>
          <td class="right">{{ moneyOrDash .TotalZakatUang }}</td>
          <td class="right">{{ kgOrDash .TotalZakatBeras }}</td>
          <td class="right">{{ moneyOrDash .TotalInfaq }}</td>
        </tr>
        {{ end }}
        <tr class="total-row">
          <td colspan="2" class="center">TOTAL</td>
          <td class="right">{{ formatInt .Grand.JumlahMuzakki }}</td>
          <td class="right">{{ formatInt .Grand.JumlahDizakati }}</td>
          <td class="right">{{ money .Grand.FitrahUang }}</td>
          <td class="right">{{ kg .Grand.FitrahBeras }}</td>
          <td class="right">{{ money .Grand.MalUang }}</td>
          <td class="right">{{ money .Grand.FidyahUang }}</td>
          <td class="right">{{ kg .Grand.FidyahBeras }}</td>
          <td class="right">{{ money .Grand.TotalZakatUang }}</td>
          <td class="right">{{ kg .Grand.TotalZakatBeras }}</td>
          <td class="right">{{ money .Grand.TotalInfaq }}</td>
        </tr>
      </tbody>
    </table>

    <div class="signature-row">
      <div class="signature-col">
				<div class="signature-head">
					<p class="signature-line">Mengetahui,</p>
					<p class="signature-line">Kepala Desa Purwajaya</p>
				</div>
        <p class="signature-name">ADI SUCIPTO</p>
      </div>
      <div class="signature-col">
				<div class="signature-head">
					<p class="signature-line">Purwajaya, {{ .TanggalTtd }}</p>
					<p class="signature-line">Ketua Pengurus PHBI</p>
				</div>
        <p class="signature-name">SLAMET JUDIONO S,Pd</p>
      </div>
    </div>
  </div>
  <script>window.onload = function(){ window.print(); };</script>
</body>
</html>`

	tmpl, err := template.New("print-superadmin-zakat").Funcs(template.FuncMap{
		"formatInt": func(v int) string {
			return formatIDNumber(int64(v))
		},
		"money": func(v float64) string {
			return "Rp " + formatIDNumber(int64(v+0.5))
		},
		"moneyOrDash": func(v float64) string {
			if v <= 0 {
				return "-"
			}
			return "Rp " + formatIDNumber(int64(v+0.5))
		},
		"kg": func(v float64) string {
			if v == float64(int64(v)) {
				return formatIDNumber(int64(v))
			}
			str := strconv.FormatFloat(v, 'f', 2, 64)
			parts := strings.Split(str, ".")
			leftInt, _ := strconv.ParseInt(parts[0], 10, 64)
			right := strings.TrimRight(parts[1], "0")
			if right == "" {
				return formatIDNumber(leftInt)
			}
			return fmt.Sprintf("%s,%s", formatIDNumber(leftInt), right)
		},
		"kgOrDash": func(v float64) string {
			if v <= 0 {
				return "-"
			}
			if v == float64(int64(v)) {
				return formatIDNumber(int64(v))
			}
			str := strconv.FormatFloat(v, 'f', 2, 64)
			parts := strings.Split(str, ".")
			leftInt, _ := strconv.ParseInt(parts[0], 10, 64)
			right := strings.TrimRight(parts[1], "0")
			if right == "" {
				return formatIDNumber(leftInt)
			}
			return fmt.Sprintf("%s,%s", formatIDNumber(leftInt), right)
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
