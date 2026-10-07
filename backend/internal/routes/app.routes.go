package routes

import (
	"taskmgr/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, appHandler *handler.AppHandler) {
	RegisterAuthRoutes(router, appHandler)
	RegisterHealthRoutes(router, appHandler)
	RegisterViewRoutes(router, appHandler)
}
