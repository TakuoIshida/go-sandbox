package todousecase

import (
	todomodel "go-sandbox/domain/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ITodoUsecase interface {
	Create(ctx *gin.Context, req CreateTodoRequest)
	// Update(ctx *gin.Context, todo request.UpdatetodoRequest)
	Delete(ctx *gin.Context, id uuid.UUID)
	FindById(ctx *gin.Context, id uuid.UUID) todomodel.Todo
	FindAll(ctx *gin.Context) []todomodel.Todo
}

// 大規模ならusecaseのメソッドごとに、usecase.create, usecase.updateのようにpackageから分離するべき
type CreateTodoRequest struct {
	Title   string
	Content string
	UserId  uuid.UUID
}

type UpdateTodoRequest struct {
	Title   string
	Content string
	UserId  uuid.UUID
}
