package routes

import (
	"TodoList/internal/services"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewRouter(s services.TodoServiceInterface) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
		s.GetAllTodosHandler(w, r)
	})
	return router
}
