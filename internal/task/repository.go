package task

import (
	"context"
	"time"

	"github.com/DevDaniloFerrari/homeworke-done-api/internal"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	Conn *pgxpool.Pool
}

func (r *Repository) Insert(task internal.TaskModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := r.Conn.Exec(ctx,
		"INSERT INTO tasks (description, isdone) VALUES ($1, $2) RETURNING id, description, isdone;",
		task.Description,
		task.IsDone)

	return err
}

func (r *Repository) FindAll() []internal.TaskModel {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var tasks []internal.TaskModel
	rows, err := r.Conn.Query(
		ctx,
		"SELECT id, description, isdone FROM tasks;")

	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var task internal.TaskModel
		err := rows.Scan(&task.ID, &task.Description, &task.IsDone)
		if err != nil {
			return nil
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil
	}

	return tasks
}
