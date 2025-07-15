package services

import (
	"context"
	"demoProject/internal/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
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
	sessionId, err := cmsApp.generateSessionId(context.Background(), userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "系统错误，稍后重试"})
		return
	}
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

func (cmsApp *CmsApp) generateSessionId(ctx context.Context, userId string) (string, error) {
	sessionId := uuid.New().String()
	// key : session_id:{user_id} val : session_id
	sessionKey := fmt.Sprintf("session_id:%s", userId)
	err := cmsApp.rdb.Set(ctx, sessionKey, sessionId, time.Hour*8).Err()
	if err != nil {
		fmt.Printf("rdb set error = %v \n", err)
		return "", err
	}

	authKey := fmt.Sprintf("session_auth:%s", sessionId)
	err = cmsApp.rdb.Set(ctx, authKey, time.Now().Unix(), time.Hour*8).Err()
	if err != nil {
		fmt.Printf("rdb set error = %v \n", err)
		return "", err
	}
	fmt.Println(sessionKey)
	fmt.Println(authKey)
	return sessionId, nil
}
