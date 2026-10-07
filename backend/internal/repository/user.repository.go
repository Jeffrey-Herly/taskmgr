package repository

import (
	"errors"

	"taskmgr/internal/models"

	"gorm.io/gorm"
)

// ErrDatabaseUnavailable dikembalikan bila repository dipanggil saat db nil
// (Postgres mati — server jalan mode degraded).
var ErrDatabaseUnavailable = errors.New("database tidak tersedia")

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create menyimpan user baru. ID diisi Postgres via default gen_random_uuid()
// dan dikembalikan GORM lewat RETURNING.
func (r *UserRepository) Create(user *models.User) error {
	if r.db == nil {
		return ErrDatabaseUnavailable
	}
	return r.db.Create(user).Error
}

// FindByUsernameOrEmail mencari user berdasarkan username ATAU email.
// Mengembalikan gorm.ErrRecordNotFound bila tidak ada.
func (r *UserRepository) FindByUsernameOrEmail(username, email string) (*models.User, error) {
	if r.db == nil {
		return nil, ErrDatabaseUnavailable
	}
	var user models.User
	if err := r.db.Where("username = ? OR email = ?", username, email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
