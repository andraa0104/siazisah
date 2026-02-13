package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/siazisah/internal/models"
	"github.com/yourusername/siazisah/internal/repository"
)

type TransaksiHandler struct {
	transaksiRepo *repository.TransaksiZakatRepository
	muzakkiRepo   *repository.MuzakkiRepository
}

func NewTransaksiHandler(transaksiRepo *repository.TransaksiZakatRepository, muzakkiRepo *repository.MuzakkiRepository) *TransaksiHandler {
	return &TransaksiHandler{
		transaksiRepo: transaksiRepo,
		muzakkiRepo:   muzakkiRepo,
	}
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

	// Hitung total wajib dan infaq tambahan
	if transaksi.JenisZakat == "fitrah" && transaksi.BentukZakat != nil && *transaksi.BentukZakat == "uang" {
		transaksi.TotalWajib = transaksi.NominalPerOrang * float64(transaksi.JumlahOrang)
		if transaksi.TotalDibayar > transaksi.TotalWajib {
			transaksi.InfaqTambahan = transaksi.TotalDibayar - transaksi.TotalWajib
		}
	}

	if err := h.transaksiRepo.Create(&transaksi); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to create transaksi"})
		return
	}

	c.JSON(http.StatusCreated, models.Response{Success: true, Message: "Transaksi created successfully", Data: transaksi})
}

func (h *TransaksiHandler) GetAll(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	transaksis, err := h.transaksiRepo.GetByMasjidID(*masjidID.(*int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get transaksis"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: transaksis})
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
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.transaksiRepo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to delete transaksi"})
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

func (h *TransaksiHandler) GetDashboard(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	
	// Implementasi dashboard stats akan ditambahkan
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data: gin.H{
			"masjid_id": *masjidID.(*int),
			"message": "Dashboard data",
		},
	})
}
