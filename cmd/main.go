package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-learn/internal/api"
)

func main() {
	r := gin.Default()
	api.CmsRouters(r)
	fmt.Println("cmd.服务已成功启动 ================================>>>>> ")
	// 启动服务，并检查启动状态
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("启动服务失败：%v\n", err)
	}
}
