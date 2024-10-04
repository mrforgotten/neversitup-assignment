package initial

import (
	"net/http"
	"simple-template/server/router"
	"simple-template/server/wire"

	"github.com/gin-gonic/gin"
)

func InitGin() *gin.Engine {
	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	todoHandler := wire.InitTodo()

	app = router.BindTodoRoute(app, *todoHandler)

	return app
}
