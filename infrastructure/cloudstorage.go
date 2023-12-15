package cloudstorage

import (
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func New(ctx *gin.Context) *storage.Client {
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect cloud storage")
	}
	defer client.Close()

	return client
}

func UploadFile(ctx *gin.Context, client *storage.Client, bucket string, path string) {
	writer := client.Bucket(bucket).Object(path).NewWriter(ctx)
	if err := writer.Close(); err != nil {
		panic(err)
	}

	log.Println("UploadFile done")
}
