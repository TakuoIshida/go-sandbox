// Package: e2e testの動作確認をおこなうためのサンプルコード
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
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})
	fmt.Println("Listen on http://localhost:8080")
	router.Run()
}
