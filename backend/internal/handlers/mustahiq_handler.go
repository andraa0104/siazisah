package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/siazisah/internal/models"
	"github.com/yourusername/siazisah/internal/repository"
)

type MustahiqHandler struct {
	mustahiqRepo *repository.MustahiqRepository
}

func NewMustahiqHandler(mustahiqRepo *repository.MustahiqRepository) *MustahiqHandler {
	return &MustahiqHandler{mustahiqRepo: mustahiqRepo}
}

func (h *MustahiqHandler) Create(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	var mustahiq models.Mustahiq
	if err := c.ShouldBindJSON(&mustahiq); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	mustahiq.MasjidID = *masjidID.(*int)
	mustahiq.IsActive = true
	if err := h.mustahiqRepo.Create(&mustahiq); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to create mustahiq"})
		return
	}

	c.JSON(http.StatusCreated, models.Response{Success: true, Message: "Mustahiq created successfully", Data: mustahiq})
}

func (h *MustahiqHandler) GetAll(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	mustahiqs, err := h.mustahiqRepo.GetByMasjidID(*masjidID.(*int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get mustahiqs"})
		return
	}

	fmt.Printf("GetAll Mustahiq - MasjidID: %d, Count: %d\n", *masjidID.(*int), len(mustahiqs))
	c.JSON(http.StatusOK, models.Response{Success: true, Data: mustahiqs})
}

func (h *MustahiqHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	mustahiq, err := h.mustahiqRepo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{Success: false, Message: "Mustahiq not found"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: mustahiq})
}

func (h *MustahiqHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var mustahiq models.Mustahiq
	if err := c.ShouldBindJSON(&mustahiq); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	mustahiq.ID = id
	if err := h.mustahiqRepo.Update(&mustahiq); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to update mustahiq"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "Mustahiq updated successfully", Data: mustahiq})
}

func (h *MustahiqHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.mustahiqRepo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to delete mustahiq"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "Mustahiq deleted successfully"})
}

func (h *MustahiqHandler) CheckDuplicate(c *gin.Context) {
	var req struct {
		Nama   string `json:"nama"`
		Lokasi string `json:"lokasi"`
		RT     string `json:"rt"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	existing, err := h.mustahiqRepo.FindDuplicate(req.Nama, req.Lokasi, req.RT)
	if err != nil || existing == nil {
		c.JSON(http.StatusOK, models.Response{Success: true, Data: gin.H{"exists": false}})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: gin.H{"exists": true, "masjid_nama": existing.MasjidNama}})
}
