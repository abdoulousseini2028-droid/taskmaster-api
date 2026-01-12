package models

import "time"

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      string     `json:"status"`
	Priority    string     `json:"priority"`
	ProjectID   *int       `json:"project_id"`
	AssignedTo  *int       `json:"assigned_to"`
	DueDate     *time.Time `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type CreateTaskRequest struct {
	Title       string     `json:"title" binding:"required,min=3"`
	Description *string    `json:"description"`
	Status      string     `json:"status" binding:"omitempty,oneof=todo in_progress done archived"`
	Priority    string     `json:"priority" binding:"omitempty,oneof=low medium high urgent"`
	ProjectID   *int       `json:"project_id"`
	AssignedTo  *int       `json:"assigned_to"`
	DueDate     *time.Time `json:"due_date"`
}

type UpdateTaskRequest struct {
	Title       *string    `json:"title" binding:"omitempty,min=3"`
	Description *string    `json:"description"`
	Status      *string    `json:"status" binding:"omitempty,oneof=todo in_progress done archived"`
	Priority    *string    `json:"priority" binding:"omitempty,oneof=low medium high urgent"`
	ProjectID   *int       `json:"project_id"`
	AssignedTo  *int       `json:"assigned_to"`
	DueDate     *time.Time `json:"due_date"`
}
