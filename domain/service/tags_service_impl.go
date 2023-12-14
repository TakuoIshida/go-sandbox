package service

import (
	todo_model "go-sandbox/domain/model"
	"go-sandbox/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

type TodoServiceImpl struct {
	TodoRepository repository.ITodoRepository
}

func NewTodoServiceImpl(tagRepository repository.ITodoRepository) ITodoService {
	return &TodoServiceImpl{
		TodoRepository: tagRepository,
	}
}

// Create implements TodoService
func (t *TodoServiceImpl) Create(ctx *gin.Context, todo todo_model.Todo) {
	t.TodoRepository.Save(ctx, todo)
}

// Delete implements TodoService
func (t *TodoServiceImpl) Delete(ctx *gin.Context, id int64) {
	t.TodoRepository.Delete(ctx, id)
}

// FindAll implements TodoService
func (t *TodoServiceImpl) FindAll(ctx *gin.Context) []todo_model.Todo {
	return t.TodoRepository.FindAll(ctx)
}

// FindById implements TodoService
func (t *TodoServiceImpl) FindById(ctx *gin.Context, id int64) todo_model.Todo {
	return t.TodoRepository.FindById(ctx, id)
}

// // Update implements TodoService
// func (t *TodoServiceImpl) Update(ctx *gin.Context, Todo request.UpdateTodoRequest) {
// 	tagData, err := t.TodoRepository.FindById(Todo.Id)
// 	helper.ErrorPanic(err)
// 	tagData.Name = Todo.Name
// 	t.TodoRepository.Update(tagData)
// }
