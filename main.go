package main

import (
	"track-admin/initialize"
)

func main() {
	r := initialize.Routers()
	r.LoadHTMLGlob("./templates/*/*")
	r.Run("127.0.0.1:8080") // 监听并在 0.0.0.0:8080 上启动服务
}
