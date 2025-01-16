package api

import (
	"go-learn/internal/services"

	"github.com/gin-gonic/gin"
)

const (
	rootPath   = "api"
	noAuthPath = "out/api"
)

func CmsRouters(r *gin.Engine) {
	cmsApp := services.NewCmsApp()
	session := &SessionAuth{}
	root := r.Group(rootPath).Use(session.Auth)
	root.GET("cms/ping", cmsApp.Hello)
	noAuth := r.Group(noAuthPath)
	noAuth.POST("register", cmsApp.Register)
	noAuth.POST("login", cmsApp.Login)
}
