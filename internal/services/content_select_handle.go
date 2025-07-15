package services

import (
	"demoProject/internal/dao"
	"demoProject/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContentSelectReq struct {
	Id       int `json:"id"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type ContentSelectRsp struct {
	Message     string                `json:"message"`
	Total       int64                 `json:"total"`
	ContentList []model.ContentDetail `json:"contentList"`
}

func (cmsApp *CmsApp) ContentSelect(ctx *gin.Context) {
	var req ContentSelectReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	contentDao := dao.NewContentDao(cmsApp.db)
	contentList, total, err := contentDao.Select(&dao.ContentSelectReq{
		Id:       req.Id,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "内部错误"})
		return
	}
	contents := make([]model.ContentDetail, 0, len(contentList))
	for _, content := range contentList {
		contents = append(contents, model.ContentDetail{
			ID:             content.ID,
			Title:          content.Title,
			VideoURL:       content.VideoURL,
			Author:         content.Author,
			Description:    content.Description,
			Thumbnail:      content.Thumbnail,
			Category:       content.Category,
			Duration:       content.Duration,
			Resolution:     content.Resolution,
			FileSize:       content.FileSize,
			Format:         content.Format,
			Quality:        content.Quality,
			ApprovalStatus: content.ApprovalStatus,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
		"data": &ContentSelectRsp{
			Message:     fmt.Sprintf("ok"),
			Total:       total,
			ContentList: contents,
		},
	})
	return
}
