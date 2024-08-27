package storage

import (
	"sync"
	"todo_go_react/internal/models"
)

type MemoryStorage struct {
	nextID uint
	mu     sync.Mutex // mutex for the following
	todos  []models.TodoTask
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		todos: make([]models.TodoTask, 0),
	}
}

func (s *MemoryStorage) GetAll() []models.TodoTask {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.todos
}

func (s *MemoryStorage) Create(todo models.TodoTask) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nextID++
	todo.ID = s.nextID
	s.todos = append(s.todos, todo)
	return nil
}

func (s *MemoryStorage) Update(todoTask models.TodoTask) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, todo := range s.todos {
		if todo.ID == todoTask.ID {
			todo.Text = todoTask.Text
			return true
		}
	}
	return false
}

func (s *MemoryStorage) Delete(id uint) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	found := false
	for i, todo := range s.todos {
		if todo.ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			found = true
			break
		}
	}
	return found
}
