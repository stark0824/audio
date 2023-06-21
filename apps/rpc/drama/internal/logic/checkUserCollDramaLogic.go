package logic

import (
	"context"
	"errors"
	"strconv"

	"audio/apps/rpc/drama/drama"
	"audio/apps/rpc/drama/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserCollDramaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserCollDramaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserCollDramaLogic {
	return &CheckUserCollDramaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserCollDramaLogic) CheckUserCollDrama(in *drama.CheckCollDramaReq) (*drama.CheckCollDramaResp, error) {

	uid, _ := strconv.Atoi(in.Uid)
	if uid <= 0 || in.Did <= 0 {
		return nil, errors.New("INVALID_ARGUMENT")
	}

	whereBuilder := l.svcCtx.FollowModel.RowBuilder()
	followList, _ := l.svcCtx.FollowModel.FindOneByUidAndDid(l.ctx, whereBuilder, int64(uid), in.Did)

	flag := 0
	if len(followList) > 0 {
		flag = 1
	}

	return &drama.CheckCollDramaResp{
		IsFollow: int64(flag),
	}, nil
}
