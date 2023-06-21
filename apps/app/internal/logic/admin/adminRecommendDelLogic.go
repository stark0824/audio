package admin

import (
	"audio/apps/rpc/drama/dramaclient"
	"audio/common/xerr"
	"context"

	"audio/apps/app/internal/svc"
	"audio/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminRecommendDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminRecommendDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminRecommendDelLogic {
	return &AdminRecommendDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminRecommendDelLogic) AdminRecommendDel(req *types.RecommendOpsReq) (resp *types.Resp, err error) {

	if req.Id <= 0 {
		return nil, xerr.NewErrCode(xerr.INVALID_ARGUMENT)
	}
	_, err = l.svcCtx.DramaRpcClient.RecPositionDel(l.ctx, &dramaclient.RecPositionOpsReq{
		Id: int64(req.Id),
	})

	if err != nil {
		logx.Info(err)
		return nil, xerr.NewErrCode(xerr.UNKNOWN)
	}
	return &types.Resp{}, nil
}
