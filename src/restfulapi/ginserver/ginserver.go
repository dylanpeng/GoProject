package main

import (
	"entities"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	r.GET("/getuser", func(c *gin.Context) {
		rid := c.DefaultQuery("id", "1")
		id, _ := strconv.ParseInt(rid, 10, 64)
		user := entities.User{ID: id, Age: 32, CreatedTime: time.Now(), UpdatedTime: time.Now(), IsDeleted: true}
		c.JSON(http.StatusOK, user)
	})

	r.POST("/adduser", func(c *gin.Context) {
		var user entities.User
		err := c.ShouldBind(&user)
		if err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.String(http.StatusBadRequest, "请求参数错误", err)
		}
	})

	r.Run(":9002") // listen and serve on 0.0.0.0:8080
}
