package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type User struct {
	Name        string    `json:"Name"`
	Age         int       `json:"Age"`
	CreatedTime time.Time `json:"CreatedTime"`
	IsDeleted   bool      `json:"IsDeleted"`
}

func main() {
	getClient()
	postClient()
}

func getClient() {
	u, _ := url.Parse("http://localhost:9001/GetUser")
	q := u.Query()
	q.Set("name", "dylan")
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", result)
}

func postClient() {
	var user User
	user.Name = "dylan"
	user.Age = 32
	user.CreatedTime = time.Now()
	user.IsDeleted = true
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json error:", err)
	}

	body := bytes.NewBuffer([]byte(b))
	res, err := http.Post("http://localhost:9001/PostUser", "application/json;charset=utf-8", body)
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", result)
}
