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

type AdminRecommendEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminRecommendEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminRecommendEditLogic {
	return &AdminRecommendEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminRecommendEditLogic) AdminRecommendEdit(req *types.RecommendDataReq) (resp *types.Resp, err error) {

	if req.Name == "" || req.NovelIds == "" || req.Sort < 0 || req.Id <= 0 {
		return nil, xerr.NewErrCode(xerr.INVALID_ARGUMENT)
	}

	v, _ := globalkey.TempType[int(req.TempType)]

	if v == "" {
		return nil, xerr.NewErrCode(xerr.INVALID_ARGUMENT)
	}

	_, err = l.svcCtx.DramaRpcClient.RecPositionEdit(l.ctx, &dramaclient.RecPositionDataReq{
		Id:       req.Id,
		Name:     req.Name,
		NovelIds: req.NovelIds,
		TempType: req.TempType,
		Sort:     req.Sort,
	})

	if err != nil {
		logx.Info(err)
		return nil, xerr.NewErrCode(xerr.UNKNOWN)
	}
	logx.Info(333)

	return &types.Resp{}, nil
}
