package logic

import (
	"audio/common/globalkey"
	"context"
	"errors"
	"time"

	"audio/apps/rpc/drama/drama"
	"audio/apps/rpc/drama/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecPositionEditLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecPositionEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecPositionEditLogic {
	return &RecPositionEditLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RecPositionEditLogic) RecPositionEdit(in *drama.RecPositionDataReq) (*drama.RecPositionOpsResp, error) {

	if in.Id <= 0 {
		return nil, errors.New("UNKNOWN")
	}

	recDetail, err := l.svcCtx.RecommendModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("INVALID_ARGUMENT")
	}

	if in.Name == "" || in.Sort < 0 {
		return nil, errors.New("UNKNOWN")
	}

	v, _ := globalkey.TempType[int(in.TempType)]
	if v == "" {
		return nil, errors.New("UNKNOWN")
	}

	recDetail.Name = in.Name
	recDetail.NovelIds = in.NovelIds
	recDetail.Sort = in.Sort
	recDetail.TempType = in.TempType
	recDetail.UpdatedAt = time.Now().Unix()
	err = l.svcCtx.RecommendModel.Update(l.ctx, recDetail)

	if err != nil {
		return nil, errors.New("NOT_FOUND")
	}
	return &drama.RecPositionOpsResp{}, nil
}
