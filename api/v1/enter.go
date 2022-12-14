package v1

import (
	"track-admin/service"
)

type ApiGroup struct {
	LoginApi
	OrderApi
}

var (
	userService = new(service.UserService)
)

var ApiGroupApp = new(ApiGroup)
