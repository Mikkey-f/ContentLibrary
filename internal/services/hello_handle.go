package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloReq struct {
	Name string `json:"name" binding:"required"`
}

type HelloRsp struct {
	Message string `json:"message" binding:"required"`
}

func (cmsApp *CmsApp) Hello(ctx *gin.Context) {
	var req HelloReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
		"data": &HelloRsp{
			Message: fmt.Sprintf("hello %s", req.Name),
		},
	})
}
