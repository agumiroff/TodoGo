// wire.go
//go:build wireinject
// +build wireinject

package wire

import (
	"TodoList/cmd/internal/config"
	"TodoList/cmd/internal/entities"
	"TodoList/cmd/internal/logger"
	"TodoList/cmd/internal/routes"
	"TodoList/cmd/internal/services"
	"TodoList/cmd/internal/storage/sqlite"
	"TodoList/cmd/internal/storage/storage_manager"
	"TodoList/web"
	"github.com/google/wire"
)

func BuildApp() (*entities.App, error) {
	wire.Build(
		config.MustLoad,
		logger.SetupLogger,
		routes.NewRouter,
		web.NewServer,
		sqlite.NewStorage,
		storage_manager.NewStorageManager,
		services.NewTodoService,
		wire.Bind(new(sqlite.StorageInterface), new(*sqlite.Storage)),
		wire.Bind(new(storage_manager.StorageManagerInterface), new(*storage_manager.StorageManager)),
		wire.Bind(new(services.TodoServiceInterface), new(*services.TodoService)),
		wire.Struct(new(entities.App), "*"),
	)
	return nil, nil
}
