package models

import "time"

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

type User struct {
	ID           string    `db:"id"            json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username     string    `db:"username"      json:"username" gorm:"size:50;uniqueIndex;not null"`
	Email        string    `db:"email"         json:"email" gorm:"size:255;uniqueIndex;not null"`
	PasswordHash string    `db:"password_hash" json:"-" gorm:"column:password_hash;not null"`
	Role         UserRole  `db:"role"          json:"role" gorm:"size:20;not null;default:user"`
	CreatedAt    time.Time `db:"created_at"    json:"created_at" gorm:"autoCreateTime"`
}

// TableName menetapkan nama tabel eksplisit untuk AutoMigrate.
func (User) TableName() string { return "users" }

// DTO

type CreateUserRequest struct {
	Username string   `json:"username" binding:"required,min=3,max=50"`
	Email    string   `json:"email"    binding:"required,email"`
	Password string   `json:"password" binding:"required,min=8"`
	Role     UserRole `json:"role"     binding:"omitempty,oneof=admin user"`
}

type UpdateUserRequest struct {
	Username string `json:"username" binding:"omitempty,min=3,max=50"`
	Email    string `json:"email"    binding:"omitempty,email"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}
