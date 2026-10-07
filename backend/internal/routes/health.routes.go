package routes

import (
	"taskmgr/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterHealthRoutes mendaftarkan endpoint cek kesehatan di bawah /api.
func RegisterHealthRoutes(router *gin.Engine, appHandler *handler.AppHandler) {
	api := router.Group("/api")
	api.GET("/health", appHandler.Health.Check)
}
