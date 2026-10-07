package routes

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"taskmgr/internal/handler"

	"github.com/gin-gonic/gin"
)

// frontendDist mengembalikan folder frontend/dist.
// Bisa dioverride lewat env FRONTEND_DIST, default ../frontend/dist
// (relatif dari folder backend saat go run ./cmd).
func frontendDist() string {
	if d := os.Getenv("FRONTEND_DIST"); d != "" {
		return d
	}
	return filepath.Join("..", "frontend", "dist")
}

// RegisterViewRoutes menyajikan SPA Vue dari frontend/dist:
// - /assets/* dan /favicon.ico sebagai file statis
// - /, /login, /dashboard, /products/:id dan route frontend lain
//   mengembalikan index.html agar Vue Router yang menangani.
func RegisterViewRoutes(router *gin.Engine, _ *handler.AppHandler) {
	dist := frontendDist()
	index := filepath.Join(dist, "index.html")

	if abs, err := filepath.Abs(dist); err == nil {
		log.Printf("      [OK] SPA dir: %s", abs)
	}
	if _, err := os.Stat(index); err != nil {
		log.Printf("      [WARN] %s tidak ditemukan, halaman akan 404. Jalankan `npm.cmd run build` di frontend/", index)
	}

	router.Static("/assets", filepath.Join(dist, "assets"))
	router.StaticFile("/favicon.ico", filepath.Join(dist, "favicon.ico"))

	serveSPA := func(c *gin.Context) {
		c.File(index)
	}

	router.GET("/", serveSPA)
	router.GET("/login", serveSPA)
	router.GET("/register", serveSPA)
	router.GET("/dashboard", serveSPA)
	router.GET("/products/:id", serveSPA)

	// Fallback SPA: semua path non-API yang tidak cocok dikembalikan ke index.html.
	// Path /api/* tetap 404 JSON agar tidak tertelan SPA.
	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "not found"})
			return
		}
		c.File(index)
	})
}
