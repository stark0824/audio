package logic

import (
	"audio/apps/rpc/drama/drama"
	"audio/apps/rpc/drama/internal/svc"
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecPositionDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecPositionDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecPositionDelLogic {
	return &RecPositionDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RecPositionDelLogic) RecPositionDel(in *drama.RecPositionOpsReq) (*drama.RecPositionOpsResp, error) {

	if in.Id <= 0 {
		return nil, errors.New("INVALID_ARGUMENT")
	}

	recDetail, err := l.svcCtx.RecommendModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("INVALID_ARGUMENT")
	}

	recDetail.IsDel = 1
	recDetail.UpdatedAt = time.Now().Unix()
	err = l.svcCtx.RecommendModel.Update(l.ctx, recDetail)

	if err != nil {
		return nil, errors.New("UNKNOWN")
	}

	return &drama.RecPositionOpsResp{}, nil
}
