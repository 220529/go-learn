package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World ~~",
		})
	})
	fmt.Println("服务已成功启动 ================================> ")
	// 启动服务，并检查启动状态
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("启动服务失败：%v\n", err)
	}
}
