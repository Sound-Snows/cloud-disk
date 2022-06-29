package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateReply, err error) {
	// 判断用户存储池中文件是否存在
	ur := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity=?", req.UserRepositoryIdentity).Get(ur)
	if err != nil {
		return
	}
	if !has {
		return nil, errors.New("user repository not found")
	}
	// 生成分享记录
	uuid := helper.GetUUID()
	data := &models.ShareBasic{
		Identity:               uuid,
		UserIdentity:           userIdentity,
		RepositoryIdentity:     ur.UserIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		ExpiredTime:            int64(req.ExpiredTime),
	}
	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return
	}
	resp = &types.ShareBasicCreateReply{Identity: uuid}
	return
}
