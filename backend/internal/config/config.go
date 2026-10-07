package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string

	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBDatabase string
	DBSSLMode  string
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// Load membaca .env (diabaikan bila tidak ada, misal di VPS pakai env asli)
// lalu mengembalikan Config dengan default sesuai .env.example.
func Load() Config {
	_ = godotenv.Load()

	return Config{
		Host: getenv("HOST", "localhost"),
		Port: getenv("PORT", "8080"),

		DBHost:     getenv("DB_HOST", "localhost"),
		DBPort:     getenv("DB_PORT", "5432"),
		DBUsername: getenv("DB_USERNAME", "postgres"),
		DBPassword: getenv("DB_PASSWORD", ""),
		DBDatabase: getenv("DB_DATABASE", "task_manager_db"),
		DBSSLMode:  getenv("DB_SSLMODE", "disable"),
	}
}

// DSN merangkai connection string Postgres untuk driver GORM.
func (c Config) DSN() string {
	return "host=" + c.DBHost +
		" port=" + c.DBPort +
		" user=" + c.DBUsername +
		" password=" + c.DBPassword +
		" dbname=" + c.DBDatabase +
		" sslmode=" + c.DBSSLMode
}
