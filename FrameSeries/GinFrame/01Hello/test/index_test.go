package test

import (
	"github.com/gin-gonic/gin"
	"github.com/pass_baci/GoTest/FrameSeries/GinFrame/01Hello/initRouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *gin.Engine


//初始化函数
func init() {
	gin.SetMode(gin.TestMode)
	router = initRouter.SetupRouter()
}

// TestIndexGetRouter测试SetupRouter函数的可用性
//func TestIndexGetRouter(t *testing.T) {
//	w := httptest.NewRecorder()  //创建一个 http 响应体
//	req, _ := http.NewRequest(http.MethodGet, "/", nil) //创建一个 http 请求体
//	router.ServeHTTP(w, req)
//	assert.Equal(t, http.StatusOK, w.Code) //判断状态码是否相同
//	assert.Equal(t, "Hello Gin!get method", w.Body.String()) //判断内容是否相同
//}

// TestIndexPutRouter测试Put请求的可用性
func TestIndexPutRouter(t *testing.T) {
	w := httptest.NewRecorder()  //创建一个 http 响应体
	req, _ := http.NewRequest(http.MethodPut, "/", nil) //创建一个 http 请求体
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code) //判断状态码是否相同
	assert.Equal(t, "Hello Gin!put method", w.Body.String()) //判断内容是否相同
}

//TestIndexGetParamRouter测试参数路由的可用性
func TestIndexGetParamRouter(t *testing.T) {
	username := "Test"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Test已保存", w.Body.String())
}

func TestIndexHtmlRouter(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}