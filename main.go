package main

import (
	"fmt"
	todo "go-sandbox/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	todo.CreateTodo()
	engine := gin.Default()
	// htmlのディレクトリを指定
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			// htmlに渡す変数を定義
			"message": "hello gin",
		})
	})
	fmt.Println("Listen on http://localhost:8080")
	engine.Run()
}
