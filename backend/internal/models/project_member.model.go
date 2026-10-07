package models

import "time"

type MemberRole string

const (
	MemberRoleManager MemberRole = "manager"
	MemberRoleMember  MemberRole = "member"
)

type ProjectMember struct {
	ID        string     `db:"id"         json:"id"`
	ProjectID string     `db:"project_id" json:"project_id"`
	UserID    string     `db:"user_id"    json:"user_id"`
	Role      MemberRole `db:"role"       json:"role"`
	JoinedAt  time.Time  `db:"joined_at"  json:"joined_at"`

	// Joined
	User *UserResponse `db:"-" json:"user,omitempty"`
}

// DTO

type AddProjectMemberRequest struct {
	UserID string     `json:"user_id" binding:"required,uuid"`
	Role   MemberRole `json:"role"    binding:"required,oneof=manager member"`
}

type UpdateProjectMemberRequest struct {
	Role MemberRole `json:"role" binding:"required,oneof=manager member"`
}

type ProjectMemberResponse struct {
	ID        string        `json:"id"`
	ProjectID string        `json:"project_id"`
	UserID    string        `json:"user_id"`
	Role      MemberRole    `json:"role"`
	JoinedAt  time.Time     `json:"joined_at"`
	User      *UserResponse `json:"user,omitempty"`
}

func (pm *ProjectMember) ToResponse() ProjectMemberResponse {
	return ProjectMemberResponse{
		ID:        pm.ID,
		ProjectID: pm.ProjectID,
		UserID:    pm.UserID,
		Role:      pm.Role,
		JoinedAt:  pm.JoinedAt,
		User:      pm.User,
	}
}
