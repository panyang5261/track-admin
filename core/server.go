package core

import (
	"fmt"
	"track-admin/initialize"
)

func RunWindowsServer() {

	initialize.Routers()

	fmt.Println("welcome to track-admin, core server")

}
