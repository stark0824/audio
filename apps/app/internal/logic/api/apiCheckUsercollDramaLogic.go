package api

import (
	"audio/apps/rpc/drama/dramaclient"
	"audio/common/xerr"
	"context"

	"audio/apps/app/internal/svc"
	"audio/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiCheckUsercollDramaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiCheckUsercollDramaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiCheckUsercollDramaLogic {
	return &ApiCheckUsercollDramaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApiCheckUsercollDramaLogic) ApiCheckUsercollDrama(req *types.CheckCollDramaReq) (resp *types.CheckCollDramaResp, err error) {
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
