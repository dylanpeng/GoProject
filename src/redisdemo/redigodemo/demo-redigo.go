package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
	}
	defer c.Close()

	//读写
	_, err = c.Do("SET", "testkey", "HelloWorld")
	if err != nil {
		fmt.Println("redis set testkey failed:", err)
	}

	text, err := redis.String(c.Do("GET", "testkey"))
	if err != nil {
		fmt.Println("redis get testkey failed:", err)
	} else {
		fmt.Printf("redis get testkey success:%v\n", text)
	}

	//设置过期时间
	_, err = c.Do("SET", "expirekey", "expirevalue", "EX", "3")
	if err != nil {
		fmt.Println("redis set expirekey failed:", err)
	}

	expirevalue, err := redis.String(c.Do("GET", "expirekey"))
	if err != nil {
		fmt.Println("redis get expirekey failed:", err)
	} else {
		fmt.Printf("get expirekey:%v\n", expirevalue)
	}

	time.Sleep(5 * time.Second)
	expirevalue, err = redis.String(c.Do("GET", "expirekey"))
	if err != nil {
		fmt.Println("redis get expirekey failed:", err)
	} else {
		fmt.Printf("get expirekey:%v\n", expirevalue)
	}

	//检查是否存在并删除
	is_key_exit, err := redis.Bool(c.Do("EXISTS", "testkey"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("exists or not: %v \n", is_key_exit)
	}

	_, err = c.Do("DEL", "testkey")

	is_key_exit, err = redis.Bool(c.Do("EXISTS", "testkey"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("exists or not: %v \n", is_key_exit)
	}

	//读写json到redis
	key := "profile"
	imap := map[string]string{"username": "666", "phonenumber": "888"}
	value, _ := json.Marshal(imap)
	n, err := c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("success")
	}
	var imapGet map[string]string
	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}
	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phonenumber"])

}
