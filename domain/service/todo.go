package todoservice

import (
	todomodel "go-sandbox/domain/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ITodoService interface {
	Create(ctx *gin.Context, t todomodel.Todo)
	// Update(ctx *gin.Context, todo request.UpdatetodoRequest)
	Delete(ctx *gin.Context, id uuid.UUID)
	FindById(ctx *gin.Context, id uuid.UUID) todomodel.Todo
	FindAll(ctx *gin.Context) []todomodel.Todo
}
