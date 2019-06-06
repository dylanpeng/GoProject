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
		fmt.Printf("beanstalk init failed err:%s\n", err)
		return
	}

	defer c.Close()

	//设置tube队列
	c.Tube.Name = "test"
	c.TubeSet.Name["test"] = true

	//发送消息
	id, err := c.Put([]byte("new message"), 30, 0, 120*time.Second)
	fmt.Printf("put message success id:%s\n", id)
}
