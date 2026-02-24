package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/siazisah/internal/models"
	"github.com/yourusername/siazisah/internal/repository"
)

type DistribusiHandler struct {
	distribusiRepo *repository.DistribusiZakatRepository
}

func NewDistribusiHandler(distribusiRepo *repository.DistribusiZakatRepository) *DistribusiHandler {
	return &DistribusiHandler{distribusiRepo: distribusiRepo}
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
	data, err := h.distribusiRepo.GetByMasjidID(*masjidID.(*int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get distribusi"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: data})
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
