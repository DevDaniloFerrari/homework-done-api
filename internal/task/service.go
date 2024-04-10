package task

import (
	"github.com/DevDaniloFerrari/homeworke-done-api/internal"
)

type Service struct {
	Repository Repository
}

func (p Service) Create(task internal.TaskModel) error {
	return p.Repository.Insert(task)
}

func (p Service) Update(task internal.TaskModel) error {
	return p.Repository.Update(task)
}

func (p Service) Delete(taskID int) error {
	return p.Repository.Delete(taskID)
}

func (p Service) FindAll() []internal.TaskModel {
	return p.Repository.FindAll()
}
