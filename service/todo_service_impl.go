package todoserviceimpl

import (
	todo_model "go-sandbox/domain/model"
	todorepository "go-sandbox/domain/repository"
	todoservice "go-sandbox/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoServiceImpl struct {
	TodoRepository todorepository.ITodoRepository
}

func NewTodoServiceImpl(tr todorepository.ITodoRepository) todoservice.ITodoService {
	return &TodoServiceImpl{
		TodoRepository: tr,
	}
}

// Create implements TodoService
func (t *TodoServiceImpl) Create(ctx *gin.Context, todo todo_model.Todo) {
	t.TodoRepository.Save(ctx, todo)
}

// Delete implements TodoService
func (t *TodoServiceImpl) Delete(ctx *gin.Context, id uuid.UUID) {
	t.TodoRepository.Delete(ctx, id)
}

// FindAll implements TodoService
func (t *TodoServiceImpl) FindAll(ctx *gin.Context) []todo_model.Todo {
	return t.TodoRepository.FindAll(ctx)
}

// FindById implements TodoService
func (t *TodoServiceImpl) FindById(ctx *gin.Context, id uuid.UUID) todo_model.Todo {
	return t.TodoRepository.FindById(ctx, id)
}

// // Update implements TodoService
// func (t *TodoServiceImpl) Update(ctx *gin.Context, Todo request.UpdateTodoRequest) {
// 	tagData, err := t.TodoRepository.FindById(Todo.Id)
// 	helper.ErrorPanic(err)
// 	tagData.Name = Todo.Name
// 	t.TodoRepository.Update(tagData)
// }
