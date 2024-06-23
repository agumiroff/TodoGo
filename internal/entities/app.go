package entities

import (
	"TodoList/internal/config"
	"TodoList/web"
	"log/slog"
)

type App struct {
	Config    *config.ServerConfig
	Logger    *slog.Logger
	WebServer *web.Server
}

func (a *App) Start() {
	a.Logger.Info("starting server")
	a.WebServer.StartServer()
}
