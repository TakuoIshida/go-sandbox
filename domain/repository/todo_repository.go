package todo_repository

import (
	todo_model "go-sandbox/domain/model"

	"github.com/gin-gonic/gin"
)

type ITodoRepository interface {
	FindById(ctx *gin.Context, id string) (todo_model.Todo, error)
	FindList(ctx *gin.Context) ([]todo_model.Todo, error)
	Create(ctx *gin.Context, todo todo_model.Todo) error
	Update(ctx *gin.Context, todo todo_model.Todo) error
	Delete(ctx *gin.Context, id string) error
}
