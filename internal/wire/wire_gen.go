// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"TodoList/internal/config"
	"TodoList/internal/entities"
	"TodoList/internal/logger"
	"TodoList/internal/routes"
	"TodoList/internal/services"
	"TodoList/internal/storage/sqlite"
	"TodoList/internal/storage/storage_manager"
	"TodoList/web"
)

// Injectors from wire.go:

func BuildApp() (*entities.App, error) {
	serverConfig := config.MustLoad()
	slogLogger := logger.SetupLogger(serverConfig)
	storage, err := sqlite.NewStorage(serverConfig, slogLogger)
	if err != nil {
		return nil, err
	}
	storageManager := storage_manager.NewStorageManager(storage)
	todoService := services.NewTodoService(storageManager, slogLogger)
	mux := routes.NewRouter(todoService)
	server := web.NewServer(mux, slogLogger)
	app := &entities.App{
		Config:    serverConfig,
		Logger:    slogLogger,
		WebServer: server,
	}
	return app, nil
}
