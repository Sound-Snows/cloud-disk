package models

import "time"

type UserRepository struct {
	Id                 int
	Identity           string
	UserIdentity       string
	ParentId           int64
	RepositoryIdentity string
	Ext                string
	Name               string
	CreateAt           time.Time `xorm:"created"`
	UpdateAt           time.Time `xorm:"updated"`
	DeleteAt           time.Time `xorm:"deleted"`
}

func (table UserRepository) TableName() string {
	return "user_repository"
}
