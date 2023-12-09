package main

import (
	"fmt"
	todo "go-sandbox/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New() // TODO
	// router.Use(common.BasicAuthRequired) // Protect these resources with basic auth.

	todoGroup := router.Group("/todo")
	{
		todoGroup.GET("/", todo.FindList)
		todoGroup.GET("/:id", todo.Find)
		todoGroup.POST("/", todo.Create)
		todoGroup.PUT("/:id", todo.Update)
		todoGroup.DELETE("/:id", todo.Delete)

	}

	fmt.Println("Listen on http://localhost:8080")
	router.Run()
}
