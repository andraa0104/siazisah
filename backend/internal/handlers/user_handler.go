package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/siazisah/internal/models"
	"github.com/yourusername/siazisah/internal/repository"
	"github.com/yourusername/siazisah/internal/utils"
)

type UserHandler struct {
	userRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

func (h *UserHandler) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to hash password"})
		return
	}
	user.Password = hashedPassword
	user.IsActive = true

	if err := h.userRepo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, models.Response{Success: true, Message: "User created successfully", Data: user})
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.userRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to get users"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: users})
}

func (h *UserHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.userRepo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{Success: false, Message: "User not found"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: user})
}

func (h *UserHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Invalid request"})
		return
	}

	user.ID = id

	// Update user data
	if err := h.userRepo.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to update user"})
		return
	}

	// Update password if provided
	if user.Password != "" {
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to hash password"})
			return
		}
		if err := h.userRepo.UpdatePassword(id, hashedPassword); err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to update password"})
			return
		}
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "User updated successfully", Data: user})
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.userRepo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Message: "User deleted successfully"})
}
