package models

import "time"

type TaskComment struct {
	ID        string    `db:"id"         json:"id"`
	TaskID    string    `db:"task_id"    json:"task_id"`
	UserID    string    `db:"user_id"    json:"user_id"`
	Content   string    `db:"content"    json:"content"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`

	// Joined
	Author *UserResponse `db:"-" json:"author,omitempty"`
}

// DTO

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required,min=1"`
}

type UpdateCommentRequest struct {
	Content string `json:"content" binding:"required,min=1"`
}

type CommentResponse struct {
	ID        string        `json:"id"`
	TaskID    string        `json:"task_id"`
	UserID    string        `json:"user_id"`
	Content   string        `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
	Author    *UserResponse `json:"author,omitempty"`
}

func (c *TaskComment) ToResponse() CommentResponse {
	return CommentResponse{
		ID:        c.ID,
		TaskID:    c.TaskID,
		UserID:    c.UserID,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
		Author:    c.Author,
	}
}
