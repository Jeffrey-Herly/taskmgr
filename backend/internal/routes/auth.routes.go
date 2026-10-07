package routes

import (
	"taskmgr/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes mendaftarkan endpoint autentikasi JSON di bawah /api.
func RegisterAuthRoutes(router *gin.Engine, appHandler *handler.AppHandler) {
	api := router.Group("/api")
	api.POST("/login", appHandler.Auth.Login)
}
