package controller

import (
	"fmt"
	cloudstorage "go-sandbox/infrastructure"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	fmt.Println("Upload file")
	// マルチパートフォーム
	form, _ := ctx.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)
		// upload file
		bucket := os.Getenv("BUCKET")
		client := cloudstorage.New(ctx)
		cloudstorage.UploadFile(ctx, client, bucket, file.Filename)
		log.Println("file.Filename uploaded")
	}
	ctx.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func Download(ctx *gin.Context) {
	fmt.Println("Download file")
	ctx.Status(http.StatusOK)
}
