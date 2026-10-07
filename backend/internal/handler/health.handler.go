package handler

import (
	"net/http"

	"taskmgr/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

type HealthData struct {
	Status   string `json:"status"`
	Database string `json:"database"`
}

// Check melapor status server + konektivitas Postgres.
// 200 bila DB terhubung, 503 bila tidak (db nil atau Ping gagal).
func (h *HealthHandler) Check(c *gin.Context) {
	// Handler jika db nil
	if h.db == nil {
		c.JSON(http.StatusServiceUnavailable, models.OK(HealthData{
			Status:   "degraded",
			Database: "disconnected",
		}))
		return
	}

	// jika db tidak nil, cek koneksi Postgres
	sqlDB, err := h.db.DB()
	if err != nil || sqlDB.Ping() != nil {
		c.JSON(http.StatusServiceUnavailable, models.OK(HealthData{
			Status:   "degraded",
			Database: "disconnected",
		}))
		return
	}

	c.JSON(http.StatusOK, models.OK(HealthData{
		Status:   "ok",
		Database: "connected",
	}))
}
