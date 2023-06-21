package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"time"

	"audio/apps/rpc/drama/drama"
	"audio/apps/rpc/drama/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DramaInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDramaInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DramaInfoLogic {
	return &DramaInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DramaInfo 后台-有声作品详情
func (l *DramaInfoLogic) DramaInfo(in *drama.DramaInfoReq) (*drama.DramaInfoResp, error) {
	var dramaId = in.Id

	if dramaId <= 0 {
		return nil, errors.New("INVALID_ARGUMENT")
	}

	Drama, err := l.svcCtx.DramaModel.FindOne(l.ctx, dramaId)
	if err != nil {
		return nil, errors.New("UNKNOWN")
	}

	if Drama == nil {
		return nil, errors.New("NOT_FOUND")
	}

	var respDrama drama.DramaInfo
	_ = copier.Copy(&respDrama, Drama)

	AllEtimeStr := time.Unix(int64(Drama.DiscountAllEtime), 0).Format("2006-01-02 15:04:05")
	SingleEtimeStr := time.Unix(int64(Drama.DiscountSingleEtime), 0).Format("2006-01-02 15:04:05")
	respDrama.DiscountAllEtime = AllEtimeStr
	respDrama.DiscountSingleEtime = SingleEtimeStr

	return &drama.DramaInfoResp{
		Drama: &respDrama,
	}, nil
}
