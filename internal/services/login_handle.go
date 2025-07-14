package services

import (
	"demoProject/internal/dao"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LoginReq struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRsp struct {
	SessionID string `json:"session_id"`
	UserId    string `json:"user_id"`
	Nickname  string `json:"nickname"`
}

func (cmsApp *CmsApp) Login(ctx *gin.Context) {
	var req LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var (
		userId   = req.UserId
		password = req.Password
	)

	accountDao := dao.NewAccountDao(cmsApp.db)
	account, err := accountDao.FindByUserId(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "请输入正确的账号ID"})
		return
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(account.Password),
		[]byte(password)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "密码错误"})
		return
	}
	sessionId := generateSessionId()
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
		"data": &LoginRsp{
			SessionID: sessionId,
			UserId:    account.UserId,
			Nickname:  account.Nickname,
		},
	})
	return
}

func generateSessionId() string {
	//TODO : session id 生成
	//TODO : session id 持久化
	return "session-id"
}
