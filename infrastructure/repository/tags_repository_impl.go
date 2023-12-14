package repository

import (
	"fmt"
	todo_model "go-sandbox/domain/model"
	"go-sandbox/helper"
	"go-sandbox/query"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	Db *gorm.DB
}

func NewTodoRepositoryImpl(Db *gorm.DB) ITodoRepository {
	return &TodoRepositoryImpl{Db: Db}
}

// Delete implements TodoRepository
func (t *TodoRepositoryImpl) Delete(ctx *gin.Context, todoId int64) {
	var todo todo_model.Todo
	result := t.Db.Where("id = ?", todoId).Delete(&todo)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TodoRepository
func (t *TodoRepositoryImpl) FindAll(ctx *gin.Context) []todo_model.Todo {
	var todos []todo_model.Todo
	result := t.Db.Find(&todos)
	helper.ErrorPanic(result.Error)
	return todos
}

// FindById implements TodoRepository
func (t *TodoRepositoryImpl) FindById(ctx *gin.Context, id int64) todo_model.Todo {
	// genから生成したmodelでも取得できるがmappingが大変。
	// DDDでentity = tableの場合 => GORMのdomain/modelのentityのまま利用した方が良さそう
	// DDDでentity = tableの場合 => GORMのdomain/modelのentityのまま利用した方が良さそう
	result, err := query.Todo.Where(query.Todo.ID.Eq(id)).First()
	if err != nil {
		panic(err)
	}
	// var tag todo_model.Todo
	// result := t.Db.Find(&tag, id)
	fmt.Println(result)

	return todo_model.Todo{
		ID:         result.ID,
		Title:      result.Title,
		Content:    result.Content,
		CreatedAt:  result.CreatedAt,
		UpdatedAt:  result.UpdatedAt,
		DeleteFlag: result.DeleteFlag,
	}
}

// Save implements TodoRepository
func (t *TodoRepositoryImpl) Save(ctx *gin.Context, todo todo_model.Todo) {
	result := t.Db.Create(&todo)
	helper.ErrorPanic(result.Error)
}
