package main

import (
	"fmt"
	"go-sandbox/config"
	todocontroller "go-sandbox/controller"
	database "go-sandbox/infrastructure"
	repositoryimpl "go-sandbox/infrastructure/repository"
	todoserviceimpl "go-sandbox/service"
	todousecaseimpl "go-sandbox/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	router := gin.New() // TODO: router　configについて調査
	// TODO: request ctxにuserIdを入れたい
	// router.Use(common.BasicAuthRequired) // Protect these resources with basic auth.

	conn := database.NewDBClientConnector()

	todoRepository := repositoryimpl.NewTodoRepositoryImpl(conn.DB)
	todoService := todoserviceimpl.NewTodoServiceImpl(todoRepository)
	todoUsecase := todousecaseimpl.NewTodoUsecaseImpl(todoService)
	todoController := todocontroller.NewTodoController(todoUsecase)

	todoGroup := router.Group("/todo")
	{
		todoGroup.GET("/", todoController.FindList)
		todoGroup.GET("/:id", todoController.FindById)
		todoGroup.POST("/", todoController.Create)
		// todoGroup.PUT("/", todoController.Update)
		todoGroup.DELETE("/", todoController.Delete)
	}

	fmt.Println("Listen on http://localhost:8080")
	router.Run()
}
