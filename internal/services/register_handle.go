package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-learn/internal/dao"
	"go-learn/internal/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type RegisterReq struct {
	UserID   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

type RegisterRsp struct {
	Message string `json:"message" binding:"required"`
}

func (c *CmsApp) Register(ctx *gin.Context) {
	var req RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("req:", req)
	// 密码加密
	hashedPassword, err := encryptPassword(req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
	}
	fmt.Println("hashedPassword:0", hashedPassword)
	accountDao := dao.NewAccountDao(c.db)
	// 检查账号是否已存在
	isExist, err := accountDao.IsExist(req.UserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
		return
	}
	if isExist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "账号已存在"})
		return
	}
	nowTime := time.Now()
	if err := accountDao.Create(model.Account{
		UserID:   req.UserID,
		Nickname: req.Nickname,
		Password: hashedPassword,
		Ct:       nowTime,
		Ut:       nowTime,
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &RegisterRsp{
			Message: fmt.Sprintf("注册成功"),
		},
	})
}

func encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("bcrypt generate from password error  = %v", err)
		return "", err
	}
	return string(hashedPassword), err
}
