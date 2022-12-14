package router

import (
	"github.com/gin-gonic/gin"
	v1 "track-admin/api/v1"
)

type LoginRouter struct{}

func (s *LoginRouter) InitLoginRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	loginRouter := Router.Group("base")
	loginApi := v1.ApiGroupApp.LoginApi
	{
		loginRouter.POST("login", loginApi.Login)
	}
	return loginRouter

}
