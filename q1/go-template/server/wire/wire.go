//go:build wireinject
// +build wireinject

package wire

import (
	"simple-template/server/handler"
	"simple-template/server/repo"
	"simple-template/server/service"

	"github.com/google/wire"
)

var TodoSet = wire.NewSet(repo.NewTodoRepo, service.NewTodoService, handler.NewTodoHandler)

func InitTodo() *handler.TodoHandler {

	wire.Build(TodoSet)

	return &handler.TodoHandler{}
}
