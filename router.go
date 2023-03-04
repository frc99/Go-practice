package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	gr := router.Group("/goods")
	{
		gr.GET("", goodsLists)
		///*action取出后面所有路径
		gr.GET("/:id/:action", goodDetail)
		gr.POST("/add", create)

	}
	router.Run(":8083")
}

// http://localhost:8083/goods/2/eat
func goodDetail(c *gin.Context) {
	id := c.Param("id")
	action := c.Param("action")
	c.JSON(http.StatusOK, map[string]string{
		"id":     id,
		"action": action})
}
func goodsLists(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"name": "goodsList"})
}
func create(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"name": "goodsList"})
}
