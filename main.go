package main

import (
	"fmt"
	"go-sandbox/config"
	todo_controller "go-sandbox/controller"
	"go-sandbox/domain/service"
	database "go-sandbox/infrastructure"
	"go-sandbox/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	router := gin.New() // TODO
	// router.Use(common.BasicAuthRequired) // Protect these resources with basic auth.

	conn := database.NewDBClientConnector()

	todoRepository := repository.NewTodoRepositoryImpl(conn.DB)
	todoService := service.NewTagsServiceImpl(todoRepository)
	todoController := todo_controller.TodoController(todoService)

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
