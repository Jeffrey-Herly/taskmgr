package infrastructures

import (
	"time"

	"taskmgr/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectPostgres membuka koneksi Postgres via GORM dan memverifikasinya dengan Ping.
// Kembalikan *gorm.DB untuk dipakai repository/service.
// (Jangan log cfg.DSN() di sini — berisi password.)
func ConnectPostgres(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
