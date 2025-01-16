package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginReq struct {
	UserID   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRsp struct {
	SessionID string `json:"session_id"`
	UserID    string `json:"user_id"`
	Nickname  string `json:"nickname"`
}

func (c *CmsApp) Login(ctx *gin.Context) {
	var req LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var (
		userID   = req.UserID
		password = req.Password
	)
	fmt.Println(userID, password)
}
