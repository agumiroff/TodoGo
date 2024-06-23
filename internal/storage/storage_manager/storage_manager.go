package storage_manager

import (
	"TodoList/internal/storage/sqlite"
	_ "TodoList/internal/storage/sqlite"
	"database/sql"
)

type StorageManager struct {
	storage sqlite.StorageInterface
}

type StorageManagerInterface interface {
	AddTodo(head string, todo string) error
	DeleteTodo(id int) error
	EditTodo(id int, head string, newTodo string) error
	GetTodo(id int) (*sql.Row, error)
	GetAllTodos() (*sql.Rows, error)
}

func NewStorageManager(storage sqlite.StorageInterface) *StorageManager {
	return &StorageManager{
		storage: storage,
	}
}

func (m *StorageManager) AddTodo(head string, todo string) error {
	return m.storage.CreateTodo(head, todo)
}

func (m *StorageManager) DeleteTodo(id int) error {
	return m.storage.DeleteTodo(id)
}

func (m *StorageManager) EditTodo(id int, head string, newTodo string) error {
	return m.storage.EditTodo(id, head, newTodo)
}

func (m *StorageManager) GetAllTodos() (*sql.Rows, error) {
	return m.storage.GetAllTodos()
}

func (m *StorageManager) GetTodo(id int) (*sql.Row, error) {
	return m.storage.GetTodo(id)
}
