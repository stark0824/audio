package logic

import (
	"audio/apps/model/follow"
	"audio/apps/rpc/drama/drama"
	"audio/apps/rpc/drama/internal/svc"
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCollDramaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCollDramaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCollDramaLogic {
	return &UserCollDramaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCollDramaLogic) UserCollDrama(in *drama.CollDramaReq) (*drama.Resp, error) {
	uid, _ := strconv.Atoi(in.Uid)
	if uid <= 0 || in.Did <= 0 {
		return nil, errors.New("INVALID_ARGUMENT")
	}

	whereBuilder := l.svcCtx.FollowModel.RowBuilder()
	followList, _ := l.svcCtx.FollowModel.FindOneByUidAndDid(l.ctx, whereBuilder, int64(uid), in.Did)

	if len(followList) > 0 && in.Status == 0 {
		l.svcCtx.FollowModel.Delete(l.ctx, followList[0].Id)
	} else if len(followList) == 0 && in.Status == 1 {
		follow := new(follow.Follow)
		follow.Uid = int64(uid)
		follow.Did = in.Did
		follow.UpdatedAt = time.Now().Unix()
		_, err := l.svcCtx.FollowModel.Insert(l.ctx, follow)
		if err != nil {
			return nil, errors.New("UNKNOWN")
		}
	}
	return &drama.Resp{}, nil
}
