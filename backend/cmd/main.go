package main

import (
	"taskmgr/internal/handler"
	"taskmgr/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Inisialisasi Handler & Routes (termasuk static SPA dari frontend/dist)
	appHandler := handler.NewAppHandler()
	routes.RegisterRoutes(r, appHandler)

	r.Run(":8080")
}
