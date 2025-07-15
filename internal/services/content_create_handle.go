package services

import (
	"demoProject/internal/dao"
	"demoProject/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type ContentCreateReq struct {
	Title          string        `json:"title" binding:"required"`
	VideoURL       string        `json:"video_url" binding:"required"`
	Author         string        `json:"author" binding:"required"`
	Description    string        `json:"description"`
	Thumbnail      string        `json:"thumbnail"`
	Category       string        `json:"category"`
	Duration       time.Duration `json:"duration"`
	Resolution     string        `json:"resolution"`
	FileSize       int64         `json:"fileSize"`
	Format         string        `json:"format"`
	Quality        int           `json:"quality"`
	ApprovalStatus int           `json:"approval_status"`
	UpdatedAt      time.Time     `json:"updated_at"`
	CreatedAt      time.Time     `json:"created_at"`
}
type ContentCreateRsp struct {
	Message string `json:"message"`
}

func (cmsApp *CmsApp) ContentCreate(ctx *gin.Context) {
	var req ContentCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contentDao := dao.NewContentDao(cmsApp.db)
	err := contentDao.Create(model.ContentDetail{
		Title:          req.Title,
		Description:    req.Description,
		Author:         req.Author,
		VideoURL:       req.VideoURL,
		Thumbnail:      req.Thumbnail,
		Category:       req.Category,
		Duration:       strconv.FormatInt(int64(req.Duration), 10),
		Resolution:     req.Resolution,
		FileSize:       req.FileSize,
		Format:         req.Format,
		Quality:        req.Quality,
		ApprovalStatus: req.ApprovalStatus,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
