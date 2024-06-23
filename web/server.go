package web

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

type Server struct {
	engine *chi.Mux
	logger *slog.Logger
}

func NewServer(engine *chi.Mux, logger *slog.Logger) *Server {
	return &Server{
		engine: engine,
		logger: logger,
	}
}

func (s *Server) StartServer() {
	err := http.ListenAndServe(":8080", s.engine)
	if err != nil {
		s.logger.Info("starting server failed", "error", err)
	}
}
