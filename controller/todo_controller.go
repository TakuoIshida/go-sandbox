package todo_controller

import (
	todo_model "go-sandbox/domain/model"
	todo_service "go-sandbox/domain/service"
	"net/http"
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
	id := ctx.Query("id")

	todo, err := ts.todoService.FindById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (ts *todoController) FindList(ctx *gin.Context) {
	todoList, err := ts.todoService.FindList(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
		ID:         1, //TODO
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
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
}

func (ts *todoController) Update(ctx *gin.Context) {
	var body UpdateTodoDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	edit := todo_model.Todo{
		ID:         body.ID,
		Title:      body.Title,
		Content:    body.Content,
		CreatedAt:  body.CreatedAt,
		UpdatedAt:  time.Now(),
		DeleteFlag: false,
	}
	ts.todoService.Create(ctx, edit)
	ctx.JSON(http.StatusCreated, edit)
}

type DeleteTodoDto struct {
	ID int
}

func (ts *todoController) Delete(ctx *gin.Context) {
	var body DeleteTodoDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	delete := todo_model.Todo{
		ID: body.ID,
	}
	ts.todoService.Create(ctx, delete)
	ctx.JSON(http.StatusCreated, delete)
}
