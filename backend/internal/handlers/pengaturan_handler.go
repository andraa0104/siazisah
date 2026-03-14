package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/siazisah/backend/internal/models"
	"github.com/yourusername/siazisah/backend/internal/repository"
)

type PengaturanHandler struct {
	pengaturanRepo *repository.PengaturanRepository
}

func NewPengaturanHandler(pengaturanRepo *repository.PengaturanRepository) *PengaturanHandler {
	return &PengaturanHandler{pengaturanRepo: pengaturanRepo}
}

func (h *PengaturanHandler) GetMyPengaturan(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	data, err := h.pengaturanRepo.GetByMasjidID(*masjidID.(*int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get pengaturan zakat"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: data})
}

func (h *PengaturanHandler) SaveMyPengaturan(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	var payload models.PengaturanZakat
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	if payload.Kelas1 <= 0 || payload.Kelas2 <= 0 || payload.Kelas3 <= 0 || payload.FidyahPerHari <= 0 || payload.FitrahBerasPerJiwa <= 0 || payload.FidyahBerasPerHari <= 0 {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Nilai kadar harus lebih dari 0"})
		return
	}

	if payload.MalRates == nil {
		payload.MalRates = map[string]float64{}
	}
	for key, value := range payload.MalRates {
		if value <= 0 {
			delete(payload.MalRates, key)
		}
	}

	payload.MasjidID = *masjidID.(*int)
	if err := h.pengaturanRepo.Upsert(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to save pengaturan zakat"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "Pengaturan zakat berhasil disimpan", Data: payload})
}
