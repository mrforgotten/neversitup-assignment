package handler

import (
	"net/http"
	"simple-template/server/model"
	"simple-template/server/repo"
	"simple-template/server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{
		service: s,
	}
}

func (handler *TodoHandler) FindAll(ctx *gin.Context) {
	data := handler.service.FindAll()
	ctx.JSON(http.StatusOK, data)
}

func (handler *TodoHandler) FindById(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := handler.service.FindUnique(&model.TodoUnique{
		Id: &id,
	})

	if err != nil {
		switch err.Error() {
		case repo.ERROR_MISSING_UNIQUE:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": repo.ERROR_MISSING_UNIQUE,
			})
		case repo.ERROR_NOT_FOUND:
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": repo.ERROR_NOT_FOUND,
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, data)

}

func (handler *TodoHandler) Create(ctx *gin.Context) {
	var createData *model.TodoData
	err := ctx.ShouldBindJSON(createData)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = handler.service.Create(createData)
	if err != nil {
		switch msg := err.Error(); msg {
		case repo.ERROR_DUPLICATE_UNIQUE:
			ctx.JSON(http.StatusBadRequest, repo.ERROR_DUPLICATE_UNIQUE)

		default:
			ctx.JSON(http.StatusInternalServerError, msg)
		}
		return
	}
}

func (handler *TodoHandler) Delete(ctx *gin.Context) {
	var unique model.TodoUnique
	err := ctx.ShouldBindJSON(&unique)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = handler.service.Delete(unique)
	if err != nil {

		switch msg := err.Error(); msg {
		case repo.ERROR_MISSING_UNIQUE:
			ctx.JSON(http.StatusBadRequest, repo.ERROR_MISSING_UNIQUE)

		case repo.ERROR_NOT_FOUND:
			ctx.JSON(http.StatusNotFound, repo.ERROR_NOT_FOUND)
		default:
			ctx.JSON(http.StatusInternalServerError, msg)
		}
		return

	}
	ctx.JSON(http.StatusOK, gin.H{})
}
