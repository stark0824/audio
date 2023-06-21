package admin

import (
	"audio/apps/rpc/drama/dramaclient"
	"audio/common/xerr"
	"context"

	"audio/apps/app/internal/svc"
	"audio/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminProducerEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminProducerEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminProducerEditLogic {
	return &AdminProducerEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminProducerEditLogic) AdminProducerEdit(req *types.ProducerEditReq) (resp *types.ProducerEditResp, err error) {
	if req.Name == "" || req.Mark == "" || req.Uid == 0 {
		return nil, xerr.NewErrCode(xerr.INVALID_ARGUMENT)
	}

	result, err := l.svcCtx.DramaRpcClient.ProducerEdit(l.ctx, &dramaclient.ProducerEditReq{
		Id:     req.Id,
		Name:   req.Name,
		Mark:   req.Mark,
		Status: req.Status,
		Uid:    req.Uid,
	})

	if err != nil {
		return nil, xerr.NewErrCode(xerr.UNKNOWN)
	}

	return &types.ProducerEditResp{
		Id: result.Id,
	}, nil
}
