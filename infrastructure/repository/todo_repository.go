package todo_repository_impl

import (
	"fmt"
	todo_model "go-sandbox/domain/model"
	todo_repository "go-sandbox/domain/repository"
	"go-sandbox/model"
	"go-sandbox/query"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type todoRepository struct {
	DB *gorm.DB
}

func TodoRepository(db *gorm.DB) todo_repository.ITodoRepository {
	return &todoRepository{
		DB: db,
	}
}

// FindById implements todo_repository.ITodoRepository.
func (sr *todoRepository) FindById(ctx *gin.Context, id int64) (todo_model.Todo, error) {
	todo := query.Todo
	result, err := todo.WithContext(ctx).Where(todo.ID.Eq(id)).First()
	if err != nil {
		log.Fatalf("todo is not found")
	}

	return todo_model.Todo{
		ID:         result.ID,
		Title:      result.Title,
		Content:    result.Content,
		CreatedAt:  result.CreatedAt,
		UpdatedAt:  result.UpdatedAt,
		DeleteFlag: result.DeleteFlag,
	}, err
}

// FindList implements todo_repository.ITodoRepository.
func (sr *todoRepository) FindList(ctx *gin.Context) ([]todo_model.Todo, error) {
	todo := query.Todo
	result, err := todo.WithContext(ctx).Where(todo.DeleteFlag.Is(false)).Find()
	if err != nil {
		log.Fatalf("todo is not found")
	}

	var todos []todo_model.Todo
	for _, todo := range result {
		todos = append(todos, todo_model.Todo{
			ID:         todo.ID,
			Title:      todo.Title,
			Content:    todo.Content,
			DeleteFlag: todo.DeleteFlag,
			CreatedAt:  todo.CreatedAt,
			UpdatedAt:  todo.UpdatedAt,
		})
	}
	return todos, err
}

// Create implements todo_repository.ITodoRepository.
func (sr *todoRepository) Create(ctx *gin.Context, todo todo_model.Todo) error {
	data := model.Todo{
		Title:   todo.Title,
		Content: todo.Content,
		UserID:  1,
	}
	return query.Todo.WithContext(ctx).Create(&data)
}

// Update implements todo_repository.ITodoRepository.
func (sr *todoRepository) Update(ctx *gin.Context, todo todo_model.Todo) error {
	data := model.Todo{
		Title:   todo.Title,
		Content: todo.Content,
		UserID:  1,
	}
	return query.Todo.WithContext(ctx).Save(&data)
}

// Delete implements todo_repository.ITodoRepository.
func (sr *todoRepository) Delete(ctx *gin.Context, id int64) error {
	todo := query.Todo

	result, err := todo.WithContext(ctx).Where(todo.ID.Eq(id)).Delete()
	fmt.Println(result)
	return err
}
