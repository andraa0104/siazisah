package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/siazisah/backend/internal/models"
	"github.com/yourusername/siazisah/backend/internal/repository"
)

type DistribusiHandler struct {
	distribusiRepo *repository.DistribusiZakatRepository
	masjidRepo     *repository.MasjidRepository
	pengurusRepo   *repository.PengurusRepository
	mustahiqRepo   *repository.MustahiqRepository
}

func NewDistribusiHandler(
	distribusiRepo *repository.DistribusiZakatRepository,
	masjidRepo *repository.MasjidRepository,
	pengurusRepo *repository.PengurusRepository,
	mustahiqRepo *repository.MustahiqRepository,
) *DistribusiHandler {
	return &DistribusiHandler{
		distribusiRepo: distribusiRepo,
		masjidRepo:     masjidRepo,
		pengurusRepo:   pengurusRepo,
		mustahiqRepo:   mustahiqRepo,
	}
}

func normalizeJenisMustahiq(jenis string) string {
	clean := strings.ToLower(strings.TrimSpace(jenis))
	clean = strings.ReplaceAll(clean, "-", " ")
	clean = strings.Join(strings.Fields(clean), " ")
	switch clean {
	case "fakir":
		return "fakir"
	case "miskin":
		return "miskin"
	case "amil":
		return "amil"
	case "mualaf":
		return "mualaf"
	case "riqab":
		return "riqab"
	case "gharim", "gharimin":
		return "gharimin"
	case "fisabilillah", "fi sabilillah", "fisabil lah":
		return "fisabilillah"
	case "ibnu sabil", "ibnusabil":
		return "ibnu sabil"
	default:
		return clean
	}
}

func extractYearFromDate(dateStr string) int {
	trimmed := strings.TrimSpace(dateStr)
	if trimmed == "" {
		return 0
	}
	parsed, err := time.Parse("2006-01-02", trimmed)
	if err == nil {
		return parsed.Year()
	}
	if len(trimmed) >= 4 {
		year, convErr := strconv.Atoi(trimmed[:4])
		if convErr == nil {
			return year
		}
	}
	return 0
}

func (h *DistribusiHandler) GetInsight(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	insight, err := h.distribusiRepo.GetInsight(*masjidID.(*int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get distribusi insight"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: insight})
}

func (h *DistribusiHandler) GetAll(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	page, pageSize, isAll := parsePagination(c)

	if isAll {
		data, err := h.distribusiRepo.GetByMasjidID(*masjidID.(*int))
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get distribusi"})
			return
		}
		pagination := buildPagination(len(data), page, pageSize, true)
		c.JSON(http.StatusOK, models.Response{Success: true, Data: gin.H{"items": data, "pagination": pagination}})
		return
	}

	total, err := h.distribusiRepo.CountByMasjidID(*masjidID.(*int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get distribusi"})
		return
	}

	offset := (page - 1) * pageSize
	data, err := h.distribusiRepo.GetByMasjidIDPaginated(*masjidID.(*int), pageSize, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get distribusi"})
		return
	}

	pagination := buildPagination(total, page, pageSize, false)
	c.JSON(http.StatusOK, models.Response{Success: true, Data: gin.H{"items": data, "pagination": pagination}})
}

func (h *DistribusiHandler) Create(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	var payload models.DistribusiZakat
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	if payload.MustahiqID <= 0 {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Mustahiq is required"})
		return
	}

	if strings.TrimSpace(payload.TanggalDistribusi) == "" {
		payload.TanggalDistribusi = time.Now().Format("2006-01-02")
	}

	payload.MasjidID = *masjidID.(*int)
	if err := h.distribusiRepo.Create(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to create distribusi"})
		return
	}

	c.JSON(http.StatusCreated, models.Response{Success: true, Message: "Distribusi created successfully", Data: payload})
}

func (h *DistribusiHandler) Update(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	id, _ := strconv.Atoi(c.Param("id"))

	var payload models.DistribusiZakat
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	if payload.MustahiqID <= 0 {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Mustahiq is required"})
		return
	}

	if strings.TrimSpace(payload.TanggalDistribusi) == "" {
		payload.TanggalDistribusi = time.Now().Format("2006-01-02")
	}

	payload.ID = id
	payload.MasjidID = *masjidID.(*int)
	if err := h.distribusiRepo.Update(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to update distribusi"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "Distribusi updated successfully", Data: payload})
}

func (h *DistribusiHandler) Delete(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.distribusiRepo.Delete(id, *masjidID.(*int)); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to delete distribusi"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "Distribusi deleted successfully"})
}

func (h *DistribusiHandler) PrintSummary(c *gin.Context) {
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

	mustahiqList, err := h.mustahiqRepo.GetByMasjidID(masjidID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get mustahiq"})
		return
	}

	distribusiList, err := h.distribusiRepo.GetByMasjidID(masjidID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get distribusi"})
		return
	}

	type recapRow struct {
		No          int
		JenisLabel  string
		JumlahOrang int
		Uang        float64
		Beras       float64
	}

	rowOrder := []struct {
		Key   string
		Label string
	}{
		{Key: "fakir", Label: "Fakir"},
		{Key: "miskin", Label: "Miskin"},
		{Key: "amil", Label: "Amil"},
		{Key: "mualaf", Label: "Mualaf"},
		{Key: "riqab", Label: "Riqab"},
		{Key: "gharimin", Label: "Gharimin"},
		{Key: "fisabilillah", Label: "Fisabilillah"},
		{Key: "ibnu sabil", Label: "Ibnu Sabil"},
	}

	rowsMap := map[string]*recapRow{}
	for idx, item := range rowOrder {
		rowsMap[item.Key] = &recapRow{No: idx + 1, JenisLabel: item.Label}
	}

	for _, mustahiq := range mustahiqList {
		key := normalizeJenisMustahiq(mustahiq.JenisPenerima)
		if row, ok := rowsMap[key]; ok {
			row.JumlahOrang++
		}
	}

	for _, distribusi := range distribusiList {
		if extractYearFromDate(distribusi.TanggalDistribusi) != tahun {
			continue
		}
		key := normalizeJenisMustahiq(distribusi.JenisPenerima)
		if row, ok := rowsMap[key]; ok {
			row.Uang += distribusi.Nominal
			row.Beras += distribusi.BerasKg
		}
	}

	rows := make([]recapRow, 0, len(rowOrder))
	totalJumlah := 0
	totalUang := 0.0
	totalBeras := 0.0
	for _, item := range rowOrder {
		row := rowsMap[item.Key]
		rows = append(rows, *row)
		totalJumlah += row.JumlahOrang
		totalUang += row.Uang
		totalBeras += row.Beras
	}

	type recapTemplateData struct {
		MasjidNama        string
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
		TotalJumlah       int
		TotalUang         float64
		TotalBeras        float64
		KetuaUPZ          string
		KetuaPengurus     string
	}

	data := recapTemplateData{
		MasjidNama:        strings.TrimSpace(masjid.Nama),
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
		TotalJumlah:       totalJumlah,
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
    <title>Laporan Data Mustahiq dan Distribusi Zakat</title>
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
            <p class="report-title">LAPORAN DATA MUSTAHIQ DAN DISTRIBUSI ZAKAT</p>
            <p class="report-year">TAHUN {{ .Tahun }}</p>
        </div>

        <table>
            <thead>
                <tr>
                    <th rowspan="2" style="width: 42px;">No</th>
                    <th rowspan="2" style="width: 150px;">Jenis Mustahiq</th>
                    <th rowspan="2" style="width: 110px;">Jumlah (orang)</th>
                    <th colspan="2">Jumlah Zakat yang di Distribusikan</th>
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
                    <td class="right">{{ formatInt .JumlahOrang }}</td>
                    <td class="right">
                      {{ if gt .Uang 0.0 }}
                      <div class="money-cell"><span>Rp</span><span>{{ formatMoney .Uang }}</span></div>
                      {{ else }}
                      <div class="money-cell"><span>Rp</span><span>-</span></div>
                      {{ end }}
                    </td>
                    <td class="right">{{ formatKg .Beras }}</td>
                </tr>
                {{ end }}
                <tr class="total-row">
                    <td colspan="2" class="center">TOTAL</td>
                    <td class="right">{{ formatInt .TotalJumlah }}</td>
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

	tmpl, err := template.New("print-distribusi-summary").Funcs(template.FuncMap{
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
