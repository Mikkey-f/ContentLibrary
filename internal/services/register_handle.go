package services

import (
	"demoProject/internal/dao"
	"demoProject/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type RegisterReq struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

type RegisterRsp struct {
	Message string `json:"message" binding:"required"`
}

func (cmsApp *CmsApp) Register(ctx *gin.Context) {
	var req RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//密码加密
	hashedPassword, err := encryptPassword(req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(hashedPassword)
	//账号校验
	accountDao := dao.NewAccountDao(cmsApp.db)
	exist, err := accountDao.IsExist(req.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if exist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "账号已经存在"})
		return
	}

	//账号信息持久化
	if _, err := accountDao.Create(model.Account{
		UserId:    req.UserId,
		Password:  hashedPassword,
		Nickname:  req.Nickname,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
		"data": &RegisterRsp{
			Message: fmt.Sprintf("注册成功"),
		},
	})
}

func encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("bcrypt generate from password error = %v", err)
		return "", err
	}
	return string(hashedPassword), nil
}
