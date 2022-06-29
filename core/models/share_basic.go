package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type ShareBasicDetails struct {
	RepositoryIdentity string `json:"repositoryIdentity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Size               int64  `json:"size"`
	Path               string `json:"path"`
}

type ShareBasic struct {
	Id                     int
	Identity               string
	UserIdentity           string
	RepositoryIdentity     string
	UserRepositoryIdentity string
	ExpiredTime            int64
	CreateAt               time.Time `xorm:"created"`
	UpdateAt               time.Time `xorm:"updated"`
	DeleteAt               time.Time `xorm:"deleted"`
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}

// Insert 保存用户分享
func (s ShareBasic) Insert(engine *xorm.Engine) (int64, error) {
	return engine.Insert(&s)
}

// CountUp 累加计数
func (s ShareBasic) CountUp(identity string, engine *xorm.Engine) error {
	_, err := engine.Exec("UPDATE share_basic SET click_num = click_num + 1 WHERE identity = ?", identity)
	if err != nil {
		return err
	}
	return nil
}

// GetBasicDetail 获取资源详情
func (s ShareBasic) GetBasicDetail(identity string, engine *xorm.Engine) (*ShareBasicDetails, error) {
	shareBasicDetails := new(ShareBasicDetails)
	_, err := engine.Table("share_basic sb").Where("sb.identity = ?", identity).
		Join("LEFT", "repository_pool rp", "sb.repository_identity = rp.identity").
		Join("LEFT", "user_repository ur", "ur.identity = sb.user_repository_identity").
		Select("sb.repository_identity,ur.name,rp.path,rp.size,rp.ext").Get(shareBasicDetails)
	if err != nil {
		return nil, err
	}
	return shareBasicDetails, nil
}
