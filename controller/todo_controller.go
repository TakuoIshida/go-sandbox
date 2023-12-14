package todo_controller

import (
	"fmt"
	todo_model "go-sandbox/domain/model"
	todo_service "go-sandbox/domain/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type todoController struct {
	todoService todo_service.ITodoService
}

func TodoController(ts todo_service.ITodoService) *todoController {
	return &todoController{
		todoService: ts,
	}
}

func (ts *todoController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseInt(fmt.Sprintf("%s", ctx.Param("id")), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todo := ts.todoService.FindById(ctx, id)

	ctx.JSON(http.StatusOK, todo)
}

func (ts *todoController) FindList(ctx *gin.Context) {
	todoList := ts.todoService.FindAll(ctx)
	ctx.JSON(http.StatusOK, todoList)
}

type CreateTodoDto struct {
	Title   string
	Content string
}

func (ts *todoController) Create(ctx *gin.Context) {
	var body CreateTodoDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	new := todo_model.Todo{
		Title:      body.Title,
		Content:    body.Content,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeleteFlag: false,
	}
	ts.todoService.Create(ctx, new)
	ctx.JSON(http.StatusCreated, new)
}

type UpdateTodoDto struct {
	ID        int64
	Title     string
	Content   string
	CreatedAt time.Time
}

// func (ts *todoController) Update(ctx *gin.Context) {
// 	var body UpdateTodoDto
// 	if err := ctx.ShouldBindJSON(&body); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	edit := todo_model.Todo{
// 		ID:         body.ID,
// 		Title:      body.Title,
// 		Content:    body.Content,
// 		CreatedAt:  body.CreatedAt,
// 		UpdatedAt:  time.Now(),
// 		DeleteFlag: false,
// 	}
// 	ts.todoService.Create(ctx, edit)
// 	ctx.JSON(http.StatusCreated, edit)
// }

type DeleteTodoDto struct {
	ID int64
}

func (ts *todoController) Delete(ctx *gin.Context) {
	var body DeleteTodoDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ts.todoService.Delete(ctx, body.ID)
	ctx.JSON(http.StatusCreated, body.ID)
}
