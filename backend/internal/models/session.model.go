package models

import "time"

type Session struct {
	ID        string    `db:"id"         json:"id"`
	UserID    string    `db:"user_id"    json:"user_id"`
	Token     string    `db:"token"      json:"token"`
	IPAddress string    `db:"ip_address" json:"ip_address"`
	ExpiresAt time.Time `db:"expires_at" json:"expires_at"`
}

// DTO

type CreateSessionRequest struct {
	UserID    string `json:"user_id"    binding:"required,uuid"`
	Token     string `json:"token"      binding:"required"`
	IPAddress string `json:"ip_address" binding:"omitempty,ip"`
	ExpiresAt string `json:"expires_at" binding:"required"`
}

type SessionResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Token     string    `json:"token"`
	IPAddress string    `json:"ip_address"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (s *Session) ToResponse() SessionResponse {
	return SessionResponse{
		ID:        s.ID,
		UserID:    s.UserID,
		Token:     s.Token,
		IPAddress: s.IPAddress,
		ExpiresAt: s.ExpiresAt,
	}
}
