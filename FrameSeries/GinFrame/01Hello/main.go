package main

import (
	"github.com/pass_baci/GoTest/FrameSeries/GinFrame/01Hello/initRouter"
)

//go get github.com/gin-gonic/gin
// @title Gin swagger
// @version 1.0
// @description Gin swagger 示例项目

// @contact.name youngxhu
// @contact.url https://youngxhui.top
// @contact.email youngxhui@g mail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	router := initRouter.SetupRouter()
	router.Run(":8080")
}
