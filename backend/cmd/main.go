package main

import (
	"log"

	"taskmgr/internal/config"
	"taskmgr/internal/handler"
	"taskmgr/internal/infrastructures"
	"taskmgr/internal/models"
	"taskmgr/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("[1/4] Load config (.env)")
	cfg := config.Load()
	log.Printf("      server=%s:%s  db=%s:%s/%s", cfg.Host, cfg.Port, cfg.DBHost, cfg.DBPort, cfg.DBDatabase)

	log.Println("[2/4] Connect Postgres")
	// Postgres opsional saat developing: gagal konek -> server tetap jalan,
	// /api/health melapor degraded. db diteruskan ke handler yang butuh.
	db, err := infrastructures.ConnectPostgres(cfg)
	if err != nil {
		log.Printf("      [WARN] Postgres tidak terjangkau, jalan tanpa DB: %v", err)
		db = nil
	} else {
		log.Println("      [OK] Postgres connected")
		log.Println("      migrate tabel users ...")
		if err := db.AutoMigrate(&models.User{}); err != nil {
			log.Fatalf("gagal migrate users: %v", err)
		}
		log.Println("      [OK] migrate users")
	}

	log.Println("[3/4] Init handlers & routes")
	r := gin.Default()
	appHandler := handler.NewAppHandler(db)
	routes.RegisterRoutes(r, appHandler)
	log.Println("      [OK] API: POST /api/login, GET /api/health")
	log.Println("      [OK] SPA: /, /login, /dashboard, /products/:id + fallback")

	log.Printf("[4/4] Listening on http://%s:%s", cfg.Host, cfg.Port)
	if err := r.Run(cfg.Host + ":" + cfg.Port); err != nil {
		log.Fatalf("gagal jalan server: %v", err)
	}
}
