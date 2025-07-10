package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := router.Run() // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		fmt.Println("r run err:", err)
		return
	}
}
