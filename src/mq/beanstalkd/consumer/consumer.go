package main

import (
	"fmt"
	"github.com/kr/beanstalk"
	"time"
)

func main(){
	//连接beanstalkd服务器
	c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")

	if err != nil{
		fmt.Printf("create beanstalkd failed err:%s\n", err.Error())
	}

	defer c.Close()

	//设置tube队列
	c.Tube.Name = "test"
	c.TubeSet.Name["test"] = true

	for {
		//队列取消息 10秒后没有取到返回超时error
		id, body, err := c.Reserve(10 * time.Second)

		if err != nil{
			fmt.Printf("read job failed err:%s\n", err)
			continue
		}

		fmt.Printf("tube:%s | body:%s | id:%s\n", c.Tube.Name, string(body), id)

		err = c.Delete(id)
	}
}
