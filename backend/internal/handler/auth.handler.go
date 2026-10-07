package handler

import (
	"net/http"
	"time"

	"taskmgr/internal/models"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginData struct {
	Token     string              `json:"token"`
	ExpiresAt time.Time           `json:"expires_at"`
	User      models.UserResponse `json:"user"`
}

// Demo credential sementara (belum ada DB).
// TODO: ganti dengan lookup Postgres + verifikasi hash password + JWT.
const (
	demoEmail    = "admin@taskmgr.local"
	demoPassword = "password123"
)

// Login menerima {email, password}, mengembalikan APIResponse{success, data:{token, user}}.
// 200 + success:true bila kredensial cocok, 401 + success:false bila tidak.
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Fail("email dan password tidak valid"))
		return
	}

	if req.Email != demoEmail || req.Password != demoPassword {
		c.JSON(http.StatusUnauthorized, models.Fail("email atau password salah"))
		return
	}

	user := models.User{
		ID:        "00000000-0000-4000-8000-000000000001",
		Username:  "admin",
		Email:     demoEmail,
		Role:      models.UserRoleAdmin,
		CreatedAt: time.Now().UTC(),
	}

	c.JSON(http.StatusOK, models.OK(LoginData{
		Token:     "mock-token-ganti-dengan-jwt-nanti",
		ExpiresAt: time.Now().UTC().Add(24 * time.Hour),
		User:      user.ToResponse(),
	}))
}
