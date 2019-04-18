package mysqlgorm

import (
	"entities"
	"fmt"
)

func CreateUser(user *entities.User) (int64, error) {
	db, err := GetDBInstance()
	if err != nil {
		fmt.Println("CreateUser failed error:" + err.Error())
		return -1, err
	}
	db.Create(user)
	return user.ID, err
}

func QueryUser(id int64) (entities.User, error) {
	db, err := GetDBInstance()
	if err != nil {
		fmt.Println("QueryUser failed error:" + err.Error())
		return entities.User{}, err
	}
	var user entities.User
	db.Where("ID = ?", id).First(&user)
	return user, nil
}

func QueryUsers() ([]entities.User, error) {
	db, err := GetDBInstance()
	if err != nil {
		fmt.Println("QueryUser failed error:" + err.Error())
		return nil, err
	}
	var users []entities.User
	db.Where("ID < ?", 6).Find(&users)
	return users, nil
}

func UpdateUser(user *entities.User) error {
	db, err := GetDBInstance()
	if err != nil {
		fmt.Println("CreateUser failed error:" + err.Error())
		return err
	}
	db.Save(user)
	return nil
}
