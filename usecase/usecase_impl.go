package todousecaseimpl

import (
	todomodel "go-sandbox/domain/model"
	todoservice "go-sandbox/domain/service"
	todousecase "go-sandbox/domain/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoUsecaseImpl struct {
	TodoService todoservice.ITodoService
}

func NewTodoUsecaseImpl(ts todoservice.ITodoService) todousecase.ITodoUsecase {
	return &TodoUsecaseImpl{
		TodoService: ts,
	}
}

// Create implements todoservice.ITodoService.
func (tu *TodoUsecaseImpl) Create(ctx *gin.Context, req todousecase.CreateTodoRequest) {
	new := todomodel.New(req.Title, req.Content, req.UserId)
	tu.TodoService.Create(ctx, new)
}

// Delete implements TodoService
func (tu *TodoUsecaseImpl) Delete(ctx *gin.Context, id uuid.UUID) {
	tu.TodoService.Delete(ctx, id)
}

// FindAll implements TodoService
func (tu *TodoUsecaseImpl) FindAll(ctx *gin.Context) []todomodel.Todo {
	return tu.TodoService.FindAll(ctx)
}

// FindById implements TodoService
func (tu *TodoUsecaseImpl) FindById(ctx *gin.Context, id uuid.UUID) todomodel.Todo {
	return tu.TodoService.FindById(ctx, id)
}

// // Update implements TodoService
// func (tu *TodoUsecaseImpl) Update(ctx *gin.Context, Todo request.UpdateTodoRequest) {
// 	tagData, err := t.TodoRepository.FindById(Todo.Id)
// 	helper.ErrorPanic(err)
// 	tagData.Name = Todo.Name
// 	t.TodoRepository.Update(tagData)
// }
