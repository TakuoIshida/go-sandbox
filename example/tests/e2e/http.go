package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello, World!")
	})
	fmt.Println("Listen on http://localhost:8080")
	router.Run()
}
