package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/abdoulousseini2028-droid/taskmaster-api/internal/models"
)

type TaskRepository struct {
	db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(ctx context.Context, req *models.CreateTaskRequest) (*models.Task, error) {
	query := `
		INSERT INTO tasks (title, description, status, priority, project_id, assigned_to, due_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, title, description, status, priority, project_id, assigned_to, due_date, created_at, updated_at
	`
	
	status := "todo"
	if req.Status != "" {
		status = req.Status
	}
	
	priority := "medium"
	if req.Priority != "" {
		priority = req.Priority
	}
	
	var task models.Task
	err := r.db.QueryRow(ctx, query,
		req.Title,
		req.Description,
		status,
		priority,
		req.ProjectID,
		req.AssignedTo,
		req.DueDate,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.Priority,
		&task.ProjectID,
		&task.AssignedTo,
		&task.DueDate,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}
	
	return &task, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id int) (*models.Task, error) {
	query := `
		SELECT id, title, description, status, priority, project_id, assigned_to, due_date, created_at, updated_at
		FROM tasks
		WHERE id = $1
	`
	
	var task models.Task
	err := r.db.QueryRow(ctx, query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.Priority,
		&task.ProjectID,
		&task.AssignedTo,
		&task.DueDate,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	
	if err != nil {
		return nil, fmt.Errorf("task not found: %w", err)
	}
	
	return &task, nil
}

func (r *TaskRepository) GetAll(ctx context.Context, status string, limit, offset int) ([]models.Task, error) {
	query := `
		SELECT id, title, description, status, priority, project_id, assigned_to, due_date, created_at, updated_at
		FROM tasks
		WHERE ($1 = '' OR status = $1)
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	
	rows, err := r.db.Query(ctx, query, status, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}
	defer rows.Close()
	
	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.Priority,
			&task.ProjectID,
			&task.AssignedTo,
			&task.DueDate,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}
	
	return tasks, nil
}

func (r *TaskRepository) Update(ctx context.Context, id int, req *models.UpdateTaskRequest) (*models.Task, error) {
	query := `
		UPDATE tasks
		SET title = COALESCE($1, title),
		    description = COALESCE($2, description),
		    status = COALESCE($3, status),
		    priority = COALESCE($4, priority),
		    project_id = COALESCE($5, project_id),
		    assigned_to = COALESCE($6, assigned_to),
		    due_date = COALESCE($7, due_date),
		    updated_at = NOW()
		WHERE id = $8
		RETURNING id, title, description, status, priority, project_id, assigned_to, due_date, created_at, updated_at
	`
	
	var task models.Task
	err := r.db.QueryRow(ctx, query,
		req.Title,
		req.Description,
		req.Status,
		req.Priority,
		req.ProjectID,
		req.AssignedTo,
		req.DueDate,
		id,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.Priority,
		&task.ProjectID,
		&task.AssignedTo,
		&task.DueDate,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	
	if err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}
	
	return &task, nil
}

func (r *TaskRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	
	if result.RowsAffected() == 0 {
		return fmt.Errorf("task not found")
	}
	
	return nil
}
