package models

import "time"

type ProjectStatus string

const (
	ProjectStatusActive   ProjectStatus = "active"
	ProjectStatusArchived ProjectStatus = "archived"
)

type Project struct {
	ID          string        `db:"id"          json:"id"`
	OwnerID     string        `db:"owner_id"    json:"owner_id"`
	Name        string        `db:"name"        json:"name"`
	Description string        `db:"description" json:"description"`
	Status      ProjectStatus `db:"status"      json:"status"`
	CreatedAt   time.Time     `db:"created_at"  json:"created_at"`

	// Joined fields (opsional, diisi saat query WITH detail)
	Owner   *UserResponse           `db:"-" json:"owner,omitempty"`
	Members []ProjectMemberResponse `db:"-" json:"members,omitempty"`
}

// DTO

type CreateProjectRequest struct {
	Name        string        `json:"name"        binding:"required,min=3,max=100"`
	Description string        `json:"description" binding:"omitempty"`
	Status      ProjectStatus `json:"status"      binding:"omitempty,oneof=active archived"`
}

type UpdateProjectRequest struct {
	Name        string        `json:"name"        binding:"omitempty,min=3,max=100"`
	Description string        `json:"description" binding:"omitempty"`
	Status      ProjectStatus `json:"status"      binding:"omitempty,oneof=active archived"`
}

type ProjectResponse struct {
	ID          string                  `json:"id"`
	OwnerID     string                  `json:"owner_id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Status      ProjectStatus           `json:"status"`
	CreatedAt   time.Time               `json:"created_at"`
	Owner       *UserResponse           `json:"owner,omitempty"`
	Members     []ProjectMemberResponse `json:"members,omitempty"`
}

func (p *Project) ToResponse() ProjectResponse {
	return ProjectResponse{
		ID:          p.ID,
		OwnerID:     p.OwnerID,
		Name:        p.Name,
		Description: p.Description,
		Status:      p.Status,
		CreatedAt:   p.CreatedAt,
		Owner:       p.Owner,
		Members:     p.Members,
	}
}
