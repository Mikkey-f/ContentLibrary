package services

import (
	"demoProject/internal/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContentDeleteReq struct {
	Id int `json:"id" binding:"required"`
}

type ContentDeleteRsp struct {
	Message string `json:"message"`
}

func (cmsApp *CmsApp) ContentDelete(ctx *gin.Context) {
	var req ContentDeleteReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contentDao := dao.NewContentDao(cmsApp.db)
	exist, err := contentDao.IsExist(req.Id)
	if err != nil || !exist {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "内容不存在"})
		return
	}
	err = contentDao.Delete(req.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "删除错误"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
		"data": &ContentCreateRsp{
			Message: "success",
		},
	})
	return
}
