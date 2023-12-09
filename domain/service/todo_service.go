package todo_service

import (
	todo_model "go-sandbox/domain/model"
	todo_repository "go-sandbox/domain/repository"

	"github.com/gin-gonic/gin"
)

type ITodoService interface {
	FindById(ctx *gin.Context, id string) (todo_model.Todo, error)
	FindList(ctx *gin.Context) ([]todo_model.Todo, error)
	Create(ctx *gin.Context, todo todo_model.Todo) error
	Update(ctx *gin.Context, todo todo_model.Todo) error
	Delete(ctx *gin.Context, id string) error
}

// struct that meets interface
type todoService struct {
	repository todo_repository.ITodoRepository
}

// constructor
func NewTodoService(tr todo_repository.ITodoRepository) ITodoService {
	return &todoService{
		repository: tr,
	}
}

// FindById implements ITodoService.
func (ts *todoService) FindById(ctx *gin.Context, id string) (todo_model.Todo, error) {
	return ts.FindById(ctx, id)
}

// FindList implements ITodoService.
func (ts *todoService) FindList(ctx *gin.Context) ([]todo_model.Todo, error) {
	return ts.FindList(ctx)
}

// Create implements ITodoService.
func (ts *todoService) Create(ctx *gin.Context, todo todo_model.Todo) error {
	return ts.Create(ctx, todo)
}

// Update implements ITodoService.
func (ts *todoService) Update(ctx *gin.Context, todo todo_model.Todo) error {
	return ts.Update(ctx, todo)
}

// Delete implements ITodoService.
func (ts *todoService) Delete(ctx *gin.Context, id string) error {
	return ts.Delete(ctx, id)
}
