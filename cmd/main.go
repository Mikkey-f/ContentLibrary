package main

import (
	"demoProject/internal/api"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.CmsRouters(router)
	err := router.Run() // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		fmt.Println("r run err:", err)
		return
	}
}
