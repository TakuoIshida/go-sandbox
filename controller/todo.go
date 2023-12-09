package todo

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Printf("find, %s", id)
}

func FindList(ctx *gin.Context) {
	name := ctx.Query("name")
	fmt.Printf("find, %s", name)
}

func Create(ctx *gin.Context) {
	fmt.Println("create")
}

func Update(ctx *gin.Context) {
	fmt.Println("update")
}

func Delete(ctx *gin.Context) {
	fmt.Println("delete")
}
