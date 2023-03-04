package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	//required 必须存在，uuid必须是uuid型的
	id   int    `uri:"id" binding:"required,uuid"`
	Name string `uri:"name",binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.Status(404)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name": person.Name,
			"id":   person.id,
		})
	})
	r.Run(":8083")
}
