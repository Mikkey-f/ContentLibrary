package api

import (
	"demoProject/internal/services"
	"github.com/gin-gonic/gin"
)

const (
	rootPath   = "/api/"
	noAuthPath = "/out/api"
)

func CmsRouters(router *gin.Engine) {
	cmsApp := services.NewCmsApp()
	session := &SessionAuth{}
	// Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
	// For example, all the routes that use a common middleware for authorization could be grouped.
	// 使用路由群统一管理一部分业务路由
	root := router.Group(rootPath).Use(session.Auth)
	{
		// GET is a shortcut for router.Handle("GET", path, handlers).
		// The last handler should be the real handler, the other ones should be middleware that can and should be shared among different routes.
		root.GET("/cms/hello", cmsApp.Hello)
	}

	noAuth := router.Group(noAuthPath)
	{
		// /out/api/cms/register
		noAuth.POST("/cms/register", cmsApp.Register)
		noAuth.POST("/cms/login", cmsApp.Login)
	}
}
