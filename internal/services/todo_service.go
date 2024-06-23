package services

import (
	"TodoList/internal/entities"
	"TodoList/internal/storage/storage_manager"
	"encoding/json"
	"log/slog"
	"net/http"
)

type TodoServiceInterface interface {
	//AddTodoHandler(w http.ResponseWriter, r *http.Request)
	//DeleteTodoHandler(w http.ResponseWriter, r *http.Request)
	//EditTodoHandler(w http.ResponseWriter, r *http.Request)
	//GetTodoHandler(w http.ResponseWriter, r *http.Request)
	GetAllTodosHandler(w http.ResponseWriter, r *http.Request)
}

type TodoService struct {
	manager storage_manager.StorageManagerInterface
	logger  *slog.Logger
}

func NewTodoService(
	manager storage_manager.StorageManagerInterface,
	logger *slog.Logger,
) *TodoService {
	return &TodoService{
		manager: manager,
		logger:  logger,
	}
}

func (s *TodoService) GetAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	const op = "services/todo_service/GetAllTodosHandler"

	var todos []entities.Todo

	rows, err := s.manager.GetAllTodos()
	if err != nil {
		http.Error(w, "Failed to retrieve todos", http.StatusInternalServerError)
		s.logger.Info("%s: %v", op, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var todo entities.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Todo, &todo.Date); err != nil {
			http.Error(w, "Failed to scan todo", http.StatusInternalServerError)
			s.logger.Info("%s: %v", op, err)
			return
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating over todos", http.StatusInternalServerError)
		s.logger.Info("%s: %v", op, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		http.Error(w, "Failed to encode todos", http.StatusInternalServerError)
		s.logger.Info("%s: %v", op, err)
		return
	}
}
