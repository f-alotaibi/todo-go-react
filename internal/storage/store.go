// An interface used for selecting various storage methods (memory, postgres, sqlite etc)
package storage

import "todo_go_react/internal/models"

type Storage interface {
	GetAll() []models.TodoTask
	Create(todo models.TodoTask) error
	Update(todo models.TodoTask) bool
	Delete(id uint) bool
}
