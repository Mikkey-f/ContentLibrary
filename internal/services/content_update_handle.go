package services

import (
	"demoProject/internal/dao"
	"demoProject/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type ContentUpdateReq struct {
	Id             int           `json:"id"`
	Title          string        `json:"title"`
	VideoURL       string        `json:"video_url"`
	Author         string        `json:"author"`
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
type ContentUpdateRsp struct {
	Message string `json:"message"`
}

func (cmsApp *CmsApp) ContentUpdate(ctx *gin.Context) {
	var req ContentUpdateReq
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
	if err = contentDao.Update(req.Id, model.ContentDetail{
		Title:          req.Title,
		VideoURL:       req.VideoURL,
		Author:         req.Author,
		Description:    req.Description,
		Thumbnail:      req.Thumbnail,
		Category:       req.Category,
		Duration:       strconv.FormatInt(int64(req.Duration), 10),
		Resolution:     req.Resolution,
		FileSize:       req.FileSize,
		Format:         req.Format,
		Quality:        req.Quality,
		ApprovalStatus: req.ApprovalStatus,
		CreatedAt:      req.CreatedAt,
		UpdatedAt:      req.UpdatedAt,
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
		"data": &ContentUpdateRsp{
			Message: "success",
		},
	})
	return
}
