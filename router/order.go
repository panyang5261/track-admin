package router

import (
	"github.com/gin-gonic/gin"
	v1 "track-admin/api/v1"
)

type OrderRouter struct{}

func (s *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	orderRouter := Router.Group("order")
	orderApi := v1.ApiGroupApp.OrderApi
	{
		orderRouter.GET("detail", orderApi.OrderDetail)
		orderRouter.GET("info", orderApi.OrderInfo)
		orderRouter.GET("index", orderApi.OrderIndex)

	}
	return orderRouter
}
