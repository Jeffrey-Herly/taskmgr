package handler

import (
	"errors"
	"net/http"
	"time"

	"taskmgr/internal/models"
	"taskmgr/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	auth *service.AuthService
}

func NewAuthHandler(auth *service.AuthService) *AuthHandler {
	return &AuthHandler{auth: auth}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginData struct {
	Token     string              `json:"token"`
	ExpiresAt time.Time           `json:"expires_at"`
	User      models.UserResponse `json:"user"`
}

// Login menerima {username, password}, mengembalikan APIResponse{success, data:{token, user}}.
// 200 + success:true bila kredensial cocok, 401 + success:false bila tidak.
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Fail("username dan password tidak valid"))
		return
	}

	result, err := h.auth.Login(req.Username, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, models.Fail("username atau password salah"))
			return
		}
		c.JSON(http.StatusInternalServerError, models.Fail("terjadi kesalahan server"))
		return
	}

	c.JSON(http.StatusOK, models.OK(LoginData{
		Token:     result.Token,
		ExpiresAt: result.ExpiresAt,
		User:      result.User,
	}))
}

// Register menerima {username, email, password}, mengembalikan user baru.
// 201 bila sukses, 409 bila username/email sudah dipakai,
// 503 bila database mati, 400 bila body tidak valid.
func (h *AuthHandler) Register(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Fail("username, email, dan password tidak valid"))
		return
	}

	user, err := h.auth.Register(req.Username, req.Email, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrUserExists):
			c.JSON(http.StatusConflict, models.Fail(err.Error()))
		case errors.Is(err, service.ErrDatabaseUnavailable):
			c.JSON(http.StatusServiceUnavailable, models.Fail(err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, models.Fail("terjadi kesalahan server"))
		}
		return
	}

	c.JSON(http.StatusCreated, models.OK(user))
}
