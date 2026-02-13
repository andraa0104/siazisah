package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yourusername/siazisah/config"
	"github.com/yourusername/siazisah/internal/middleware"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(userID int, username, role string, masjidID *int) (string, error) {
	claims := middleware.Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		MasjidID: masjidID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}
