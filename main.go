package main

import (
	"fmt"
	"go-sandbox/config"
	todo_controller "go-sandbox/controller"
	todo_service "go-sandbox/domain/service"
	database "go-sandbox/infrastructure"
	todo_repository_impl "go-sandbox/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	router := gin.New() // TODO
	// router.Use(common.BasicAuthRequired) // Protect these resources with basic auth.

	conn := database.NewDBClientConnector()
	todoRepository := todo_repository_impl.TodoRepository(conn.DB)
	todoService := todo_service.NewTodoService(todoRepository)
	todoController := todo_controller.TodoController(todoService)

	todoGroup := router.Group("/todo")
	{
		todoGroup.GET("/", todoController.FindList)
		todoGroup.GET("/:id", todoController.FindById)
		todoGroup.POST("/", todoController.Create)
		todoGroup.PUT("/", todoController.Update)
		todoGroup.DELETE("/", todoController.Delete)
	}

	fmt.Println("Listen on http://localhost:8080")
	router.Run()
}
