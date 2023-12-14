package repository

import (
	todo_model "go-sandbox/domain/model"

	"github.com/gin-gonic/gin"
)

type ITodoRepository interface {
	Save(ctx *gin.Context, todo todo_model.Todo)
	// Update(ctx *gin.Context, todo todo_model.Todo)
	Delete(ctx *gin.Context, id int64)
	FindById(ctx *gin.Context, id int64) todo_model.Todo
	FindAll(ctx *gin.Context) []todo_model.Todo
}
