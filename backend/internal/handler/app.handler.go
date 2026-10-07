package handler

import (
	"taskmgr/internal/repository"
	"taskmgr/internal/service"

	"gorm.io/gorm"
)

type AppHandler struct {
	Views  *ViewsHandler
	Auth   *AuthHandler
	Health *HealthHandler
}

// NewAppHandler merakit layer: repository -> service -> handler.
// db boleh nil (Postgres mati): repository mengembalikan
// ErrDatabaseUnavailable, server tetap jalan, /api/health degraded.
func NewAppHandler(db *gorm.DB) *AppHandler {
	userRepo := repository.NewUserRepository(db)
	authSvc := service.NewAuthService(userRepo)

	return &AppHandler{
		Views:  NewViewsHandler(),
		Auth:   NewAuthHandler(authSvc),
		Health: NewHealthHandler(db),
	}
}
