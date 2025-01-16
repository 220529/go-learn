package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const SessionKey = "session_id"

type SessionAuth struct{}

func (s *SessionAuth) Auth(ctx *gin.Context) {
	sessionId := ctx.Request.Header.Get(SessionKey)
	if sessionId == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "session is id null")
	}
	fmt.Println(sessionId)
	ctx.Next()
}
