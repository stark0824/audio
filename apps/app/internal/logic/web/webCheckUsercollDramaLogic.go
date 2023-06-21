package web

import (
	"audio/apps/rpc/drama/dramaclient"
	"audio/common/xerr"
	"context"

	"audio/apps/app/internal/svc"
	"audio/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebCheckUsercollDramaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWebCheckUsercollDramaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WebCheckUsercollDramaLogic {
	return &WebCheckUsercollDramaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebCheckUsercollDramaLogic) WebCheckUsercollDrama(req *types.CheckCollDramaReq) (resp *types.CheckCollDramaResp, err error) {

	if req.Uid == "" || req.Did == 0 {
		return nil, xerr.NewErrCode(xerr.INVALID_ARGUMENT)
	}

	res, err := l.svcCtx.DramaRpcClient.CheckUserCollDrama(l.ctx, &dramaclient.CheckCollDramaReq{
		Uid: req.Uid,
		Did: req.Did,
	})

	return &types.CheckCollDramaResp{
		Follow: res.IsFollow,
	}, nil
}
