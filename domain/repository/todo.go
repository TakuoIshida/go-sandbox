package todorepository

import (
	todo "go-sandbox/domain/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ITodoRepository interface {
	Save(ctx *gin.Context, t todo.Todo)
	// Update(ctx *gin.Context, todo todo_model.Todo)
	Delete(ctx *gin.Context, id uuid.UUID)
	FindById(ctx *gin.Context, id uuid.UUID) todo.Todo
	FindAll(ctx *gin.Context) []todo.Todo
}
