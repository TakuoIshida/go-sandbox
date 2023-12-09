package todo_repository_impl

import (
	"database/sql"
	todo_model "go-sandbox/domain/model"
	todo_repository "go-sandbox/domain/repository"

	"github.com/gin-gonic/gin"
)

type todoRepository struct {
	DB *sql.DB
}

func TodoRepository(db *sql.DB) todo_repository.ITodoRepository {
	return &todoRepository{
		DB: db,
	}
}

// FindById implements todo_repository.ITodoRepository.
func (sr *todoRepository) FindById(ctx *gin.Context, id string) (todo_model.Todo, error) {
	panic("FindById unimplemented")
}

// FindList implements todo_repository.ITodoRepository.
func (sr *todoRepository) FindList(ctx *gin.Context) ([]todo_model.Todo, error) {
	panic("FindList unimplemented")
}

// Create implements todo_repository.ITodoRepository.
func (sr *todoRepository) Create(ctx *gin.Context, todo todo_model.Todo) error {
	panic("Create unimplemented")
}

// Update implements todo_repository.ITodoRepository.
func (sr *todoRepository) Update(ctx *gin.Context, todo todo_model.Todo) error {
	panic("Update unimplemented")
}

// Delete implements todo_repository.ITodoRepository.
func (sr *todoRepository) Delete(ctx *gin.Context, id string) error {
	panic("Delete unimplemented")
}
