package web

import (
	"audio/apps/rpc/drama/dramaclient"
	"audio/common/xerr"
	"context"

	"audio/apps/app/internal/svc"
	"audio/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebUsercollDramaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWebUsercollDramaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WebUsercollDramaLogic {
	return &WebUsercollDramaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebUsercollDramaLogic) WebUsercollDrama(req *types.CollDramaReq) (resp *types.Resp, err error) {
	if req.Uid == "" || req.Status == 10 || req.Did == 0 {
		return nil, xerr.NewErrCode(xerr.INVALID_ARGUMENT)
	}

	l.svcCtx.DramaRpcClient.UserCollDrama(l.ctx, &dramaclient.CollDramaReq{
		Uid:    req.Uid,
		Did:    req.Did,
		Status: req.Status,
	})

	return &types.Resp{}, nil
}
