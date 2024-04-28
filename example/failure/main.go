package main

import (
	"fmt"
	"go-sandbox/example/failure/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/ping", handler.UserHandleFunc)
	router.POST("/file/upload", handler.UploadFileHandleFunc)
	router.GET("/file/download", handler.ReadFileHandleFunc)
	fmt.Println("Listen on http://localhost:8080")
	router.Run()
}
