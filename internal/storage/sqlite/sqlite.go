package sqlite

import (
	"TodoList/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
	"time"
)

type Storage struct {
	db     *sql.DB
	logger *slog.Logger
}

type StorageInterface interface {
	CreateTodo(head string, todo string) error
	DeleteTodo(id int) error
	EditTodo(id int, head string, newTodo string) error
	GetTodo(id int) (*sql.Row, error)
	GetAllTodos() (*sql.Rows, error)
}

func NewStorage(cfg *config.ServerConfig, logger *slog.Logger) (*Storage, error) {
	op := "storage.sqlite.NewStorage"

	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
CREATE TABLE IF NOT EXISTS todo (
    id INTEGER PRIMARY KEY,
    head TEXT NOT NULL,
    todo TEXT,
    created_at DATE NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_todo ON todo(todo);
`)

	if err != nil {
		return nil, fmt.Errorf("#{op}: #{err}")
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("#{op}: #{err}")
	}

	return &Storage{
		db:     db,
		logger: logger,
	}, nil
}

func (s *Storage) CreateTodo(head string, todo string) error {
	const op = "storage.sqlite.CreateTodo"

	date := time.Now().Format("2006-01-02 15:04:05") // Форматирование даты

	stmt, err := s.db.Prepare(`INSERT INTO todo (head, todo, created_at) VALUES (?, ?, ?)`)
	if err != nil {
		s.logger.Info("%s: %w", op, err)
	}

	_, err = stmt.Exec(head, todo, date)
	if err != nil {
		s.logger.Info("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) DeleteTodo(id int) error {
	const op = "storage.sqlite.DeleteTodo"

	stmt, err := s.db.Prepare(`DELETE FROM todo WHERE id = ?`)
	if err != nil {
		s.logger.Info("%s: %w", op, err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		s.logger.Info("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) EditTodo(id int, head string, newTodo string) error {
	const op = "storage.sqlite.EditTodo"

	stmt, err := s.db.Prepare(`UPDATE todo SET head = ?, todo = ? WHERE id = ?`)
	if err != nil {
		s.logger.Info("%s: %w", op, err)
	}

	_, err = stmt.Exec(head, newTodo, id)
	if err != nil {
		s.logger.Info("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetTodo(id int) (*sql.Row, error) {
	const op = "storage.sqlite.GetTodo"

	stmt, err := s.db.Prepare(`SELECT * FROM todo WHERE id = ?`)
	if err != nil {
		s.logger.Info("%s: %w", op, err)
	}

	row := stmt.QueryRow(id)
	return row, nil
}

func (s *Storage) GetAllTodos() (*sql.Rows, error) {
	const op = "storage.sqlite.GetAllTodos"

	rows, err := s.db.Query(`SELECT * FROM todo`)
	if err != nil {
		s.logger.Info("%s: %w", op, err)
	}

	s.logger.Info("method completed successfully", op)

	return rows, nil
}
