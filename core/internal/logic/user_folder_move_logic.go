package logic

import (
	"cloud-disk/core/models"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderMoveLogic {
	return &UserFolderMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderMoveLogic) UserFolderMove(req *types.UserFolderMoveRequest, userIdentity string) (resp *types.UserFolderMoveReply, err error) {
	// 判断父级文件是否存在
	parentData := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity=? AND user_dentity=?", req.Identity, userIdentity).Get(parentData)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("文件夹不存在")
	}
	//更新ParentID
	_, err = l.svcCtx.Engine.Where("identity=?", req.Identity).Update(models.UserRepository{ParentId: int64(parentData.Id)})
	return
}
