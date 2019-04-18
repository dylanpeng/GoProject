package main

import (
	"entities"
	"fmt"
	"log"
	"mysqlorm/mysqlgorm"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := mysqlgorm.GetDBInstance()
	if err != nil {
		log.Fatal("Open mysql failed:", err.Error())
	}
	defer db.Close()

	user := entities.User{Name: "bbb", Age: 3, CreatedTime: time.Now(), UpdatedTime: time.Now(), IsDeleted: false}
	mysqlgorm.CreateUser(&user)
	fmt.Println(user)

	user2, _ := mysqlgorm.QueryUser(11)
	fmt.Println(user2)

	users, _ := mysqlgorm.QueryUsers()
	fmt.Println(users)

	user2.Age = 32
	mysqlgorm.UpdateUser(&user2)
	fmt.Println(user2)

}
