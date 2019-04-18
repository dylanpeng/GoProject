package entities

import (
	"time"
)

type User struct {
	ID          int64     `gorm:"column:ID;primary_key;auto_increment" form:"ID"`
	Name        string    `gorm:"column:Name"`
	Age         int       `gorm:"column:Age"`
	CreatedTime time.Time `gorm:"column:CreatedTime"`
	UpdatedTime time.Time `gorm:"column:UpdatedTime"`
	IsDeleted   bool      `gorm:"column:IsDeleted"`
}

func (User) TableName() string {
	return "User"
}
