package controllers

import (
	"backend/config"
	"backend/models"
	"context"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/api/oauth2/v1"
)

// GoogleLogin - Redirect ke Google OAuth
func GoogleLogin(c *gin.Context) {
	url := config.GoogleOAuthConfig.AuthCodeURL("randomstate")
	c.Redirect(http.StatusFound, url)
}

// GoogleCallback - Handle callback dari Google
func GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := config.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan token"})
		return
	}

	client := config.GoogleOAuthConfig.Client(context.Background(), token)
	service, err := oauth2.New(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat service OAuth2"})
		return
	}

	userInfo, err := service.Userinfo.Get().Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan user info"})
		return
	}

	// Cek apakah user sudah ada di database
	var user models.Player
	result := config.DB.Where("email = ?", userInfo.Email).First(&user)
	if result.RowsAffected == 0 {
		// Jika belum ada, buat akun baru
		user = models.Player{
			Name:      userInfo.Name,
			Email:     userInfo.Email,
			AvatarURL: userInfo.Picture,
		}
		config.DB.Create(&user)
	}

	// Buat token JWT
	tokenString, err := GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
		return
	}

	// Kirim token ke client
	c.JSON(http.StatusOK, gin.H{"token": tokenString, "user": user})
}

// GenerateJWT - Buat token JWT untuk user
func GenerateJWT(user models.Player) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func RegisterOrUpdateUser(c *gin.Context) {
	var user models.Player
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah user sudah ada berdasarkan email
	var existingUser models.Player
	result := config.DB.Where("email = ?", user.Email).First(&existingUser)

	if result.Error != nil { // Jika tidak ditemukan, buat user baru
		user.ID = uuid.New()
		if err := config.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, user)
	} else { // Jika sudah ada, update data user
		existingUser.Name = user.Name
		existingUser.AvatarURL = user.AvatarURL
		if err := config.DB.Save(&existingUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, existingUser)
	}
}