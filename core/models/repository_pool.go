package models

import "time"

type RepositoryPool struct {
	Id       int
	Identity string
	Hash     string
	Name     string
	Ext      string
	Size     int64
	Path     string
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
	DeleteAt time.Time `xorm:"deleted"`
}

func (table RepositoryPool) TableName() string {
	return "repository_pool"
}
