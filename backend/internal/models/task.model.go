package models

import "time"

type TaskPriority string
type TaskStatus string

const (
	TaskPriorityLow    TaskPriority = "low"
	TaskPriorityMedium TaskPriority = "medium"
	TaskPriorityHigh   TaskPriority = "high"

	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusDone       TaskStatus = "done"
)

type Task struct {
	ID          string       `db:"id"          json:"id"`
	ProjectID   string       `db:"project_id"  json:"project_id"`
	AssignedTo  *string      `db:"assigned_to" json:"assigned_to"`
	Title       string       `db:"title"       json:"title"`
	Description string       `db:"description" json:"description"`
	Priority    TaskPriority `db:"priority"    json:"priority"`
	Status      TaskStatus   `db:"status"      json:"status"`
	DueDate     *time.Time   `db:"due_date"    json:"due_date"`
	CreatedAt   time.Time    `db:"created_at"  json:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"  json:"updated_at"`

	// Joined
	Assignee *UserResponse     `db:"-" json:"assignee,omitempty"`
	Labels   []LabelResponse   `db:"-" json:"labels,omitempty"`
	Comments []CommentResponse `db:"-" json:"comments,omitempty"`
}

// DTO

type CreateTaskRequest struct {
	AssignedTo  *string      `json:"assigned_to" binding:"omitempty,uuid"`
	Title       string       `json:"title"       binding:"required,min=3,max=200"`
	Description string       `json:"description" binding:"omitempty"`
	Priority    TaskPriority `json:"priority"    binding:"required,oneof=low medium high"`
	Status      TaskStatus   `json:"status"      binding:"omitempty,oneof=todo in_progress done"`
	DueDate     *time.Time   `json:"due_date"    binding:"omitempty"`
	LabelIDs    []string     `json:"label_ids"   binding:"omitempty,dive,uuid"`
}

type UpdateTaskRequest struct {
	AssignedTo  *string      `json:"assigned_to" binding:"omitempty,uuid"`
	Title       string       `json:"title"       binding:"omitempty,min=3,max=200"`
	Description string       `json:"description" binding:"omitempty"`
	Priority    TaskPriority `json:"priority"    binding:"omitempty,oneof=low medium high"`
	Status      TaskStatus   `json:"status"      binding:"omitempty,oneof=todo in_progress done"`
	DueDate     *time.Time   `json:"due_date"    binding:"omitempty"`
}

type UpdateTaskLabelsRequest struct {
	LabelIDs []string `json:"label_ids" binding:"required,dive,uuid"`
}

type TaskResponse struct {
	ID          string            `json:"id"`
	ProjectID   string            `json:"project_id"`
	AssignedTo  *string           `json:"assigned_to"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Priority    TaskPriority      `json:"priority"`
	Status      TaskStatus        `json:"status"`
	DueDate     *time.Time        `json:"due_date"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	Assignee    *UserResponse     `json:"assignee,omitempty"`
	Labels      []LabelResponse   `json:"labels,omitempty"`
	Comments    []CommentResponse `json:"comments,omitempty"`
}

func (t *Task) ToResponse() TaskResponse {
	return TaskResponse{
		ID:          t.ID,
		ProjectID:   t.ProjectID,
		AssignedTo:  t.AssignedTo,
		Title:       t.Title,
		Description: t.Description,
		Priority:    t.Priority,
		Status:      t.Status,
		DueDate:     t.DueDate,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		Assignee:    t.Assignee,
		Labels:      t.Labels,
		Comments:    t.Comments,
	}
}
