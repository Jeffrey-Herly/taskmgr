package service

import (
	"errors"
	"time"

	"taskmgr/internal/models"
	"taskmgr/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ErrInvalidCredentials dikembalikan bila username tidak dikenal atau password salah.
var ErrInvalidCredentials = errors.New("username atau password salah")

// ErrUserExists dikembalikan bila username/email sudah terdaftar.
var ErrUserExists = errors.New("username atau email sudah terdaftar")

// ErrDatabaseUnavailable diteruskan dari repository saat Postgres mati.
var ErrDatabaseUnavailable = repository.ErrDatabaseUnavailable

// Kredensial statis sementara untuk LOGIN (belum pindah DB).
// TODO: login via repository + bcrypt.CompareHashAndPassword + JWT.
const (
	demoUsername = "admin"
	demoPassword = "12345678"
)

type LoginResult struct {
	Token     string
	ExpiresAt time.Time
	User      models.UserResponse
}

type AuthService struct {
	users *repository.UserRepository
}

func NewAuthService(users *repository.UserRepository) *AuthService {
	return &AuthService{users: users}
}

// Login memverifikasi kredensial dan mengembalikan token + user.
// Mengembalikan ErrInvalidCredentials bila tidak cocok.
func (s *AuthService) Login(username, password string) (LoginResult, error) {
	if username != demoUsername || password != demoPassword {
		return LoginResult{}, ErrInvalidCredentials
	}

	user := models.User{
		ID:        "00000000-0000-4000-8000-000000000001",
		Username:  "admin",
		Email:     "admin@taskmgr.local",
		Role:      models.UserRoleAdmin,
		CreatedAt: time.Now().UTC(),
	}

	return LoginResult{
		Token:     "mock-token-ganti-dengan-jwt-nanti",
		ExpiresAt: time.Now().UTC().Add(24 * time.Hour),
		User:      user.ToResponse(),
	}, nil
}

// Register mendaftarkan user baru: cek duplikat -> hash bcrypt -> simpan.
// Role selalu "user" (dibatasi server, bukan dari input). ID diisi Postgres.
func (s *AuthService) Register(username, email, password string) (models.UserResponse, error) {
	existing, err := s.users.FindByUsernameOrEmail(username, email)
	if err == nil && existing != nil {
		return models.UserResponse{}, ErrUserExists
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.UserResponse{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.UserResponse{}, err
	}

	user := models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hash),
		Role:         models.UserRoleUser,
	}

	if err := s.users.Create(&user); err != nil {
		return models.UserResponse{}, err
	}

	return user.ToResponse(), nil
}
