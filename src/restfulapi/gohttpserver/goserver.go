package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type User struct {
	Name        string    `json:"Name"`
	Age         int       `json:"Age"`
	CreatedTime time.Time `json:"CreatedTime"`
	IsDeleted   bool      `json:"IsDeleted"`
}

func main() {
	http.HandleFunc("/PostUser", PostUserHandler)
	http.HandleFunc("/GetUser", GetUserHandler)
	http.ListenAndServe("localhost:9001", nil)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", result)

		var user User
		json.Unmarshal([]byte(result), &user)
		fmt.Println(user)
		fmt.Fprintf(w, "Add User %s", result)
	} else {
		w.Write([]byte("Only support post"))
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "GET" {
		if len(r.Form["name"]) > 0 {
			fmt.Println(r.Form["name"])
		}
		var user User
		user.Name = "dylan"
		user.Age = 32
		user.CreatedTime = time.Now()
		user.IsDeleted = true
		b, err := json.Marshal(user)
		if err != nil {
			fmt.Println("json error:", err)
		}
		fmt.Fprint(w, string(b))
	} else {
		w.Write([]byte("Only support get"))
	}
}
