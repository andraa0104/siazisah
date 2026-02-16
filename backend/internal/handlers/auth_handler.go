package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/siazisah/internal/models"
	"github.com/yourusername/siazisah/internal/repository"
	"github.com/yourusername/siazisah/internal/utils"
)

type AuthHandler struct {
	userRepo *repository.UserRepository
}

func NewAuthHandler(userRepo *repository.UserRepository) *AuthHandler {
	return &AuthHandler{userRepo: userRepo}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	user, err := h.userRepo.FindByUsername(req.Username)
	if err != nil {
		// Log the actual error for debugging
		println("FindByUsername error:", err.Error())
		c.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Invalid username or password",
		})
		return
	}

	if !user.IsActive {
		c.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Account is inactive",
		})
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		// Log for debugging
		println("Password check failed for user:", user.Username)
		println("Input password:", req.Password)
		println("Stored hash exists:", len(user.Password) > 0)
		c.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Invalid username or password",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Role, user.MasjidID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Login successful",
		Data: models.LoginResponse{
			Token: token,
			User:  *user,
		},
	})
}
