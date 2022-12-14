package initialize

import (
	"github.com/gin-gonic/gin"
	"track-admin/router"
	//"track-admin/router"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()

	//PublicGroup := Router.Group("")
	//{
	//	// 健康监测
	//	PublicGroup.GET("/health", func(c *gin.Context) {
	//		c.JSON(http.StatusOK, "ok")
	//	})
	//}

	ApiGroup := Router.Group("api")
	{
		new(router.LoginRouter).InitLoginRouter(ApiGroup)
		new(router.OrderRouter).InitOrderRouter(ApiGroup)
	}

	return Router
}
