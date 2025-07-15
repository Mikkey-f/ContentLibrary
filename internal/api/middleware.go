package api

import (
	"context"
	"demoProject/internal/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
)

const SessionKey = "session_id"

type SessionAuth struct {
	rdb *redis.Client
}

func NewSessionAuth() *SessionAuth {
	s := &SessionAuth{}
	connRdb(s)
	return s
}

func (s *SessionAuth) Auth(ctx *gin.Context) {
	sessionID := ctx.GetHeader(SessionKey)
	if sessionID == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "sessionId is null")
		return
	}
	authKey := utils.GetAuthKey(sessionID)
	loginTime, err := s.rdb.Get(ctx, authKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "session auth error")
		return
	}
	if loginTime == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "session auth fail")
		return
	}
	// ctx.Next() 只应该在中间件中使用
	ctx.Next()
}

func connRdb(s *SessionAuth) {
	// redis-cli
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	s.rdb = rdb
}
