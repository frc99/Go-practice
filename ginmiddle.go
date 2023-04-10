package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HandlerFunc
func indexHandler(c *gin.Context) {
	fmt.Println("index")

	name, ok := c.Get("name") //取值 实现了跨中间件取值
	if !ok {
		name = "default user"
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 定义一个中间件
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	// 计时
	start := time.Now()

	//gin中间件中使用goroutine
	//当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）
	//go otherFunc(c.Copy()) // 在otherFunc中只能使用c的拷贝

	c.Next() //调用后续的处理函数

	//c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Printf("m1 out ...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ...")

	c.Set("name", "tony") //可以在请求上下文里面设置一些值，然后其他地方取值

	c.Abort() //阻止调用后续的处理函数 也就是 m2它自己走完就行
	//return
	fmt.Println("m2 out ...")
}

// 自定义认证中间件 通过这种方式实现一些灵活的控制
func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	// 或其他一些准备工作
	return func(c *gin.Context) {
		if doCheck {
			//这里存放具体的逻辑
			// 是否登录的判断
			// if 是登录用户
			c.Next()
			// else
			//   c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()

	// 点开Use 查看源码发现 Use(middleware ...HandlerFunc) IRoutes
	r.Use(m1, m2, authMiddleware(true)) //全局注册中间件函数m1 m2 authMiddleware

	/*访问/index
	执行indexHandler之前 去执行注册的中间件  总的执行打印顺序是：m1 in  -> m2 in -> index -> m2 out -> m1 out
	*/
	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})

	/* 路由组注册中间件方法1
	routeGroup1 := r.Group("/xx", authMiddleware(true)){
		routeGroup1.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "routeGroup1"})
		})
	}
	路由组注册中间件方法2
	routeGroup2 := r.Group("/xx2")
	routeGroup2.Use(authMiddleware(true)){
		routeGroup2.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "routeGroup2"})
		})
	}
	*/

	r.Run(":9090")
}
