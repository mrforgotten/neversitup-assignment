package router

import (
	"simple-template/server/handler"

	"github.com/gin-gonic/gin"
)

func BindTodoRoute(g *gin.Engine, h handler.TodoHandler) *gin.Engine {
	group := g.Group("todo")
	{
		group.GET("/GetAll", h.FindAll)
		group.GET("/:id", h.FindById)
		group.POST("/Create", h.Create)
		group.POST("/", h.Delete)
	}
	return g
}
