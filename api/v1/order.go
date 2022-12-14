package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderApi struct{}

func (b *OrderApi) OrderDetail(c *gin.Context) {
	fmt.Println("now in api: order detail api")
	c.JSON(http.StatusOK, gin.H{
		"message": "查询order信息成功",
	})
}

func (b *OrderApi) OrderInfo(c *gin.Context) {
	c.HTML(http.StatusOK, "detail.html", gin.H{
		"name": "admin",
		"pwd":  "123456",
	})
}

func (b *OrderApi) OrderIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "order index",
	})
}
