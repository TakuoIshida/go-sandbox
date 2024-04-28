package main

import (
	"fmt"
	"go-sandbox/example/failure/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(someMiddleware()) // ここで認証を行う
	router.GET("/ping", handler.UserHandleFunc)
	router.POST("/file/upload", handler.UploadFileHandleFunc)
	router.GET("/file/download", handler.ReadFileHandleFunc)
	fmt.Println("Listen on http://localhost:8080")
	router.Run()
}

func someMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before request")
		c.Next()
		log.Println("after request")
	}
}
