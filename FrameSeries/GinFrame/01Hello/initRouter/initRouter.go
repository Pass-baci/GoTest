package initRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/pass_baci/GoTest/FrameSeries/GinFrame/01Hello/middleware"
	"net/http"
	"strings"
)


func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Logger(), gin.Recovery())
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	//router.Static("/statics", "./statics") //添加静态资源
	//router.StaticFile("/favicon.ico","./favicon.ico") //添加网站图标
	//Test := router.Group("/")
	//{
	//	Test.GET("/", handler.Index)
	//	Test.GET("/login", handler.Login)
	//	Test.PUT("/", retHelloGiAndMethod)
	//	Test.POST("/", retHelloGiAndMethod)
	//	Test.DELETE("/", retHelloGiAndMethod)
	//	Test.PATCH("/", retHelloGiAndMethod)
	//	Test.HEAD("/", retHelloGiAndMethod)
	//	Test.OPTIONS("/", retHelloGiAndMethod)
	//}
	//UserTest := router.Group("/user")
	//{
	//	UserTest.GET("/:name", handler.UserSave)
	//}
	return router
}

func retHelloGiAndMethod(c *gin.Context) {
	c.String(http.StatusOK, "Hello Gin!"+strings.ToLower(c.Request.Method)+" method")
}