package todocontroller

import (
	todousecase "go-sandbox/domain/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoControllerImpl struct {
	TodoUsecase todousecase.ITodoUsecase
}

func NewTodoController(tu todousecase.ITodoUsecase) *TodoControllerImpl {
	return &TodoControllerImpl{
		TodoUsecase: tu,
	}
}

func (tc *TodoControllerImpl) FindById(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todo := tc.TodoUsecase.FindById(ctx, id)

	ctx.JSON(http.StatusOK, todo)
}

func (tc *TodoControllerImpl) FindList(ctx *gin.Context) {
	todoList := tc.TodoUsecase.FindAll(ctx)
	ctx.JSON(http.StatusOK, todoList)
}

type CreateTodoDto struct {
	Title   string
	Content string
	UserId  uuid.UUID
}

func (tc *TodoControllerImpl) Create(ctx *gin.Context) {
	var body CreateTodoDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tc.TodoUsecase.Create(ctx, todousecase.CreateTodoRequest{
		Title:   body.Title,
		Content: body.Content,
		UserId:  body.UserId,
	})
	ctx.Status(http.StatusCreated)
}

type UpdateTodoDto struct {
	ID        int64
	Title     string
	Content   string
	CreatedAt time.Time
}

// func (tc *TodoControllerImpl) Update(ctx *gin.Context) {
// 	var body UpdateTodoDto
// 	if err := ctx.ShouldBindJSON(&body); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	edit := todomodel.Todo{
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
	ID uuid.UUID
}

func (tc *TodoControllerImpl) Delete(ctx *gin.Context) {
	var body DeleteTodoDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tc.TodoUsecase.Delete(ctx, body.ID)
	ctx.JSON(http.StatusCreated, body.ID)
}
