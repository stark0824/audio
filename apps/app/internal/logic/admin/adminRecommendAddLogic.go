package admin

import (
	"audio/apps/rpc/drama/dramaclient"
	"audio/common/globalkey"
	"audio/common/xerr"
	"context"

	"audio/apps/app/internal/svc"
	"audio/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminRecommendAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminRecommendAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminRecommendAddLogic {
	return &AdminRecommendAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminRecommendAddLogic) AdminRecommendAdd(req *types.RecommendDataReq) (resp *types.Resp, err error) {

	if req.Name == "" || req.NovelIds == "" || req.Sort < 0 {
		return nil, xerr.NewErrCode(xerr.INVALID_ARGUMENT)
	}

	v, _ := globalkey.TempType[int(req.TempType)]
	if v == "" {
		return nil, xerr.NewErrCode(xerr.INVALID_ARGUMENT)
	}

	_, err = l.svcCtx.DramaRpcClient.RecPositionAdd(l.ctx, &dramaclient.RecPositionAddReq{
		Name:     req.Name,
		NovelIds: req.NovelIds,
		TempType: req.TempType,
		Sort:     req.Sort,
	})

	if err != nil {
		logx.Info(err)
		return nil, xerr.NewErrCode(xerr.UNKNOWN)
	}
	return &types.Resp{}, nil
}
