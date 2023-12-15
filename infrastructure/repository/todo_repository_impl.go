package repositoryimpl

import (
	"fmt"
	todomodel "go-sandbox/domain/model"
	todorepository "go-sandbox/domain/repository"
	"go-sandbox/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	Db *gorm.DB
}

func NewTodoRepositoryImpl(Db *gorm.DB) todorepository.ITodoRepository {
	return &TodoRepositoryImpl{Db: Db}
}

// Delete implements TodoRepository
func (t *TodoRepositoryImpl) Delete(ctx *gin.Context, id uuid.UUID) {
	var todo todomodel.Todo
	result := t.Db.Where("id = ?", id).Delete(&todo)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TodoRepository
func (t *TodoRepositoryImpl) FindAll(ctx *gin.Context) []todomodel.Todo {
	var todos []todomodel.Todo
	result := t.Db.Find(&todos)
	helper.ErrorPanic(result.Error)
	return todos
}

// FindById implements TodoRepository
func (t *TodoRepositoryImpl) FindById(ctx *gin.Context, id uuid.UUID) todomodel.Todo {
	// genから生成したmodelでも取得できるがmappingが大変。
	// DDDでentity = tableの場合 => GORMのdomain/modelのentityのまま利用した方が良さそう
	// DDDでentity = tableの場合 => GORMのdomain/modelのentityのまま利用した方が良さそう
	// result, err := query.Todo.Where(query.Todo.ID.Eq(id)).First()
	// if err != nil {
	// 	panic(err)
	// }
	var todo todomodel.Todo
	result := t.Db.Find(&todo, id)
	fmt.Println(result)

	return todo
}

// Save implements TodoRepository
func (t *TodoRepositoryImpl) Save(ctx *gin.Context, todo todomodel.Todo) {
	result := t.Db.Create(&todo)
	helper.ErrorPanic(result.Error)
}
