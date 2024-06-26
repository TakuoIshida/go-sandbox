package main

import (
	"fmt"
	"go-sandbox/config"
	"go-sandbox/controller"
	"go-sandbox/infrastructure/database"
	repositoryimpl "go-sandbox/infrastructure/repository"
	todoserviceimpl "go-sandbox/service"
	todousecaseimpl "go-sandbox/usecase"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	if config.Conf.GIN_MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("run in production")
	}
	router := gin.New()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// TODO: request ctxにuserIdを入れたい
	// router.Use(common.BasicAuthRequired) // Protect these resources with basic auth.

	conn := database.NewDBClientConnector()

	todoRepository := repositoryimpl.NewTodoRepositoryImpl(conn.DB)
	todoService := todoserviceimpl.NewTodoServiceImpl(todoRepository)
	todoUsecase := todousecaseimpl.NewTodoUsecaseImpl(todoService)
	todoController := controller.NewTodoController(todoUsecase)

	todoGroup := router.Group("/todo")
	{
		todoGroup.GET("/", todoController.FindList)
		todoGroup.GET("/:id", todoController.FindById)
		todoGroup.POST("/", todoController.Create)
		// todoGroup.PUT("/", todoController.Update)
		todoGroup.DELETE("/", todoController.Delete)
	}

	fileGroup := router.Group("/file")
	{
		fileGroup.POST("/upload", controller.Upload)
		fileGroup.POST("/download/:filename", controller.Download)
	}

	router.GET("/", func(ctx *gin.Context) {
		fmt.Println(os.Getenv("BUCKET"))
		ctx.JSON(http.StatusOK, "Hello, World!")
	})
	fmt.Println("Listen on http://localhost:8080")
	router.Run()
}
