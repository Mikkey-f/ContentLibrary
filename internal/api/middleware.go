package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const SessionKey = "session_id"

type SessionAuth struct {
}

func (s *SessionAuth) Auth(ctx *gin.Context) {
	sessionID := ctx.GetHeader(SessionKey)
	// TODO : imp auth
	if sessionID == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "sessionId is null")
	}
	// ctx.Next() 只应该在中间件中使用
	ctx.Next()
}
