package handler

import (
	"go-sandbox/example/failure/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morikuni/failure"
)

func UserHandleFunc(ctx *gin.Context) {
	app := service.User{}
	_, err := app.Search("id1")
	if err != nil {
		errorLog(err)
		// w.WriteHeader(httpStatus(err))
		// status: 500 でmessageを返す
		// ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		// 	"Error": "Request failed in AbortWithStatusJSON",
		// })
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Request failed in JSON",
		})
		return
	}
	ctx.JSON(http.StatusOK, "pong")
}

func errorLog(err error) {
	code, ok := failure.CodeOf(err)
	if !ok {
		log.Printf("unexpected error: %v\n", err)
		return
	}
	log.Printf("[%v] %v\n", code, err)
}

func FileHandleFunc(c *gin.Context) {
	file, _ := c.FormFile("file")
	dst := "./uploads/" + file.Filename
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to save",
		})
		return
	}
	id := c.DefaultQuery("id", "nil")

	log.Println(file.Filename)
	log.Println(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "uploaded",
	})
}
