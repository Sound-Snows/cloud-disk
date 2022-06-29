package models

import "time"

type UserBasic struct {
	Id       int
	Identity string
	Name     string
	Password string
	Email    string
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
	DeleteAt time.Time `xorm:"deleted"`
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
