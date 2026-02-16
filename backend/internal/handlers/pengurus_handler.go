package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/siazisah/internal/models"
	"github.com/yourusername/siazisah/internal/repository"
)

type PengurusHandler struct {
	pengurusRepo *repository.PengurusRepository
}

func NewPengurusHandler(pengurusRepo *repository.PengurusRepository) *PengurusHandler {
	return &PengurusHandler{pengurusRepo: pengurusRepo}
}

// Pengurus Masjid handlers
func (h *PengurusHandler) SavePengurusMasjid(c *gin.Context) {
	masjidID, exists := c.Get("masjid_id")
	if !exists || masjidID == nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Masjid ID not found"})
		return
	}

	id, ok := masjidID.(*int)
	if !ok || id == nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid masjid ID"})
		return
	}

	var pengurus models.PengurusMasjid
	if err := c.ShouldBindJSON(&pengurus); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	pengurus.MasjidID = *id

	if err := h.pengurusRepo.SavePengurusMasjid(&pengurus); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to save pengurus masjid"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "Pengurus masjid saved successfully", Data: pengurus})
}

func (h *PengurusHandler) GetPengurusMasjid(c *gin.Context) {
	masjidID, exists := c.Get("masjid_id")
	if !exists || masjidID == nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Masjid ID not found"})
		return
	}

	id, ok := masjidID.(*int)
	if !ok || id == nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid masjid ID"})
		return
	}

	pengurusList, err := h.pengurusRepo.GetPengurusMasjidByMasjidID(*id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get pengurus masjid"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: pengurusList})
}

// Pengurus Zakat (UPZ) handlers
func (h *PengurusHandler) SavePengurusZakat(c *gin.Context) {
	masjidID, exists := c.Get("masjid_id")
	if !exists || masjidID == nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Masjid ID not found"})
		return
	}

	id, ok := masjidID.(*int)
	if !ok || id == nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid masjid ID"})
		return
	}

	var pengurus models.PengurusZakat
	if err := c.ShouldBindJSON(&pengurus); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	pengurus.MasjidID = *id

	if err := h.pengurusRepo.SavePengurusZakat(&pengurus); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to save pengurus zakat"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "Pengurus zakat saved successfully", Data: pengurus})
}

func (h *PengurusHandler) GetPengurusZakat(c *gin.Context) {
	masjidID, exists := c.Get("masjid_id")
	if !exists || masjidID == nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Masjid ID not found"})
		return
	}

	id, ok := masjidID.(*int)
	if !ok || id == nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid masjid ID"})
		return
	}

	pengurusList, err := h.pengurusRepo.GetPengurusZakatByMasjidID(*id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get pengurus zakat"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: pengurusList})
}

func (h *PengurusHandler) DeletePengurusZakat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid ID"})
		return
	}

	if err := h.pengurusRepo.DeletePengurusZakat(id); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to delete pengurus zakat"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "Pengurus zakat deleted successfully"})
}

// Public handlers untuk menampilkan di dashboard publik
func (h *PengurusHandler) GetPengurusMasjidPublic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid ID"})
		return
	}

	pengurusList, err := h.pengurusRepo.GetPengurusMasjidByMasjidID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get pengurus masjid"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: pengurusList})
}

func (h *PengurusHandler) GetPengurusZakatPublic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid ID"})
		return
	}

	pengurusList, err := h.pengurusRepo.GetPengurusZakatByMasjidID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get pengurus zakat"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: pengurusList})
}
