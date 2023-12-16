package controller

import (
	"fmt"
	cloudstorage "go-sandbox/infrastructure"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	fmt.Println("Upload file")
	// マルチパートフォーム
	form, _ := ctx.MultipartForm()
	//NOTE: form-data {key: upload[], value: File[]}
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
	filename := ctx.Query("filename")
	bucket := os.Getenv("BUCKET")
	client := cloudstorage.New(ctx)

	rc, err := client.Bucket(bucket).Object(filename).NewReader(ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to read file: "+err.Error())
		return
	}
	defer rc.Close()
	fmt.Println("Download file")
	// ファイルの内容をHTTPレスポンスとして書き込む
	ctx.Writer.WriteHeader(http.StatusOK)
	if _, err := io.Copy(ctx.Writer, rc); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to copy content: "+err.Error())
	}
}
