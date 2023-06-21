package api

import (
	"audio/apps/rpc/drama/dramaclient"
	"audio/common/xerr"
	"context"

	"audio/apps/app/internal/svc"
	"audio/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiUsercollDramaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiUsercollDramaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiUsercollDramaLogic {
	return &ApiUsercollDramaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApiUsercollDramaLogic) ApiUsercollDrama(req *types.CollDramaReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	if req.Uid == "" || req.Status == 10 || req.Did == 0 {
		return nil, xerr.NewErrCode(xerr.INVALID_ARGUMENT)
	}

	l.svcCtx.DramaRpcClient.UserCollDrama(l.ctx, &dramaclient.CollDramaReq{
		Uid:    req.Uid,
		Did:    req.Did,
		Status: req.Status,
	})
	return
}
