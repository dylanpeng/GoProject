package main

import (
	"entities"
	"fmt"
	"mysqlorm/mysqlgorm"
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
		id, err := strconv.ParseInt(rid, 10, 64)
		if err != nil {
			fmt.Println("getuser badrequest error:" + err.Error())
			c.String(http.StatusBadRequest, "bad request param")
			return
		}
		user, err := mysqlgorm.QueryUser(id)
		if err != nil {
			fmt.Println("QueryUser error:" + err.Error())
			c.String(http.StatusBadRequest, "QueryUser error, please contact administrator")
			return
		}
		c.JSON(http.StatusOK, user)
	})

	//POST请求通过绑定获取对象
	r.POST("/adduser", func(c *gin.Context) {
		var user entities.User
		err := c.ShouldBind(&user)
		user.ID = 0
		user.CreatedTime = time.Now()
		user.UpdatedTime = time.Now()
		if err != nil {
			c.String(http.StatusBadRequest, "Bad Requst", err)
			return
		}
		mysqlgorm.CreateUser(&user)
		c.JSON(http.StatusOK, user)
	})

	r.Run(":9002") // listen and serve on 0.0.0.0:8080
}
