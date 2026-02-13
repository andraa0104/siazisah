package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/siazisah/internal/models"
	"github.com/yourusername/siazisah/internal/repository"
)

type MuzakkiHandler struct {
	muzakkiRepo *repository.MuzakkiRepository
}

func NewMuzakkiHandler(muzakkiRepo *repository.MuzakkiRepository) *MuzakkiHandler {
	return &MuzakkiHandler{muzakkiRepo: muzakkiRepo}
}

func (h *MuzakkiHandler) Create(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	var muzakki models.Muzakki
	if err := c.ShouldBindJSON(&muzakki); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	muzakki.MasjidID = *masjidID.(*int)
	if err := h.muzakkiRepo.Create(&muzakki); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to create muzakki"})
		return
	}

	c.JSON(http.StatusCreated, models.Response{Success: true, Message: "Muzakki created successfully", Data: muzakki})
}

func (h *MuzakkiHandler) GetAll(c *gin.Context) {
	masjidID, _ := c.Get("masjid_id")
	muzakkis, err := h.muzakkiRepo.GetByMasjidID(*masjidID.(*int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get muzakkis"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: muzakkis})
}

func (h *MuzakkiHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	muzakki, err := h.muzakkiRepo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{Success: false, Message: "Muzakki not found"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: muzakki})
}

func (h *MuzakkiHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var muzakki models.Muzakki
	if err := c.ShouldBindJSON(&muzakki); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	muzakki.ID = id
	if err := h.muzakkiRepo.Update(&muzakki); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to update muzakki"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "Muzakki updated successfully", Data: muzakki})
}

func (h *MuzakkiHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.muzakkiRepo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to delete muzakki"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "Muzakki deleted successfully"})
}
