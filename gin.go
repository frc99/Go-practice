package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	panic("异常")
	c.JSON(http.StatusOK, map[string]string{
		"message": "Hello world!"})
}
func main() {
	// 创建一个默认的路由引擎
	//default会默认开启logger和recovery
	r := gin.Default()
	//r := gin.New()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/ping", pong)
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(":8083")
}
