package main

import (
	"entities"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化引擎
	r := gin.Default()

	//简单的Get请求
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//GET请求通过name获取参数  /user/test
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	//GET请求通过正常的URL获取参数  /getuser?id=2
	r.GET("/getuser", func(c *gin.Context) {
		rid := c.DefaultQuery("id", "1")
		id, _ := strconv.ParseInt(rid, 10, 64)
		user := entities.User{ID: id, Age: 32, CreatedTime: time.Now(), UpdatedTime: time.Now(), IsDeleted: true}
		c.JSON(http.StatusOK, user)
	})

	//POST请求通过绑定获取对象
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
