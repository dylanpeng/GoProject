package mysqlgorm

import (
	"entities"

	"github.com/jinzhu/gorm"
)

func CreateUser(user *entities.User, db *gorm.DB) {
	db.Create(user)
}

func QueryUser(id int64, db *gorm.DB) entities.User {
	var user entities.User
	db.Where("ID = ?", id).First(&user)
	return user
}

func QueryUsers(db *gorm.DB) []entities.User {
	var users []entities.User
	db.Where("ID < ?", 6).Find(&users)
	return users
}

func UpdateUser(user *entities.User, db *gorm.DB) {
	db.Save(user)
}
