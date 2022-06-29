package logic

import (
	"cloud-disk/core/models"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailReply, err error) {
	resp = &types.ShareBasicDetailReply{}
	//点击计数
	shareBasic := new(models.ShareBasic)
	err = shareBasic.CountUp(req.Identity, l.svcCtx.Engine)
	if err != nil {
		return resp, nil
	}
	//获取资源详情数据
	basicDetail, err := shareBasic.GetBasicDetail(req.Identity, l.svcCtx.Engine)
	if err != nil {
		return resp, nil
	}
	resp.RepositoryIdentity = basicDetail.RepositoryIdentity
	resp.Name = basicDetail.Name
	resp.Ext = basicDetail.Ext
	resp.Path = basicDetail.Path
	return
}
