package logic

import (
	"audio/apps/model/recommend"
	"audio/apps/rpc/drama/drama"
	"audio/apps/rpc/drama/internal/svc"
	"audio/common/globalkey"
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecPositionAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecPositionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecPositionAddLogic {
	return &RecPositionAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RecPositionAddLogic) RecPositionAdd(in *drama.RecPositionAddReq) (*drama.RecPositionOpsResp, error) {

	if in.Name == "" || in.Sort < 0 {
		return nil, errors.New("INVALID_ARGUMENT")
	}

	v, _ := globalkey.TempType[int(in.TempType)]
	if v == "" {
		return nil, errors.New("INVALID_ARGUMENT")
	}

	rec := new(recommend.RecPosition)
	rec.Name = in.Name
	rec.Sort = in.Sort
	rec.TempType = in.TempType
	rec.NovelIds = in.NovelIds
	rec.UpdatedAt = time.Now().Unix()
	rec.CreatedAt = time.Now().Unix()

	_, err := l.svcCtx.RecommendModel.Insert(l.ctx, rec)
	if err != nil {
		return nil, errors.New("UNKNOWN")
	}
	return &drama.RecPositionOpsResp{}, nil
}
