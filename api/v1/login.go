package v1

import "github.com/gin-gonic/gin"

type LoginApi struct{}

func (b *LoginApi) Login(c *gin.Context) {
	// u := &system.SysUser{Username: l.Username, Password: l.Password}
	userService.Login()

	// response.FailWithMessage("验证码错误", c)
}
