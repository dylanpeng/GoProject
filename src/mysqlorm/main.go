package main

import (
	"entities"
	"fmt"
	"log"
	"mysqlorm/mysqlgorm"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal("Open mysql failed:", err.Error())
	}
	defer db.Close()

	user := entities.User{Name: "bbb", Age: 3, CreatedTime: time.Now(), UpdatedTime: time.Now(), IsDeleted: false}
	mysqlgorm.CreateUser(&user, db)
	fmt.Println(user)

	user2 := mysqlgorm.QueryUser(11, db)
	fmt.Println(user2)

	users := mysqlgorm.QueryUsers(db)
	fmt.Println(users)

	user2.Age = 32
	mysqlgorm.UpdateUser(&user2, db)
	fmt.Println(user2)

}
