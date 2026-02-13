package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

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
	masjids, err := h.masjidRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get masjids"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: masjids})
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
