package service

import (
	todo_model "go-sandbox/domain/model"

	"github.com/gin-gonic/gin"
)

type ITodoService interface {
	Create(ctx *gin.Context, todo todo_model.Todo)
	// Update(ctx *gin.Context, todo request.UpdatetodoRequest)
	Delete(ctx *gin.Context, id int64)
	FindById(ctx *gin.Context, id int64) todo_model.Todo
	FindAll(ctx *gin.Context) []todo_model.Todo
}
