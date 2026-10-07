package models

type Label struct {
	ID        string `db:"id"         json:"id"`
	ProjectID string `db:"project_id" json:"project_id"`
	Name      string `db:"name"       json:"name"`
	Color     string `db:"color"      json:"color"`
}

// DTO

type CreateLabelRequest struct {
	Name  string `json:"name"  binding:"required,min=1,max=50"`
	Color string `json:"color" binding:"required,hexcolor"`
}

type UpdateLabelRequest struct {
	Name  string `json:"name"  binding:"omitempty,min=1,max=50"`
	Color string `json:"color" binding:"omitempty,hexcolor"`
}

type LabelResponse struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
}

func (l *Label) ToResponse() LabelResponse {
	return LabelResponse{
		ID:        l.ID,
		ProjectID: l.ProjectID,
		Name:      l.Name,
		Color:     l.Color,
	}
}
