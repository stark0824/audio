package logic

import (
	"audio/apps/model/producer"
	"audio/apps/rpc/drama/drama"
	"audio/apps/rpc/drama/internal/svc"
	"context"
	"errors"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type DramaShortListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDramaShortListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DramaShortListLogic {
	return &DramaShortListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DramaShortList 公共部分
func (l *DramaShortListLogic) DramaShortList(in *drama.DramaShortListReq) (*drama.DramaShortListResp, error) {

	whereBuilder := l.svcCtx.DramaModel.RowBuilder()
	var ids string = in.Ids
	list, err := l.svcCtx.DramaModel.FindListByIds(l.ctx, whereBuilder, ids)

	if err != nil && err != producer.ErrNotFound {
		return nil, errors.New("UNKNOWN")
	}

	var resp []*drama.DramaShort
	if len(list) > 0 {
		for _, dramaItem := range list {
			var pbDramaCommonInfo drama.DramaShort
			_ = copier.Copy(&pbDramaCommonInfo, dramaItem)
			resp = append(resp, &pbDramaCommonInfo)
		}
	}

	return &drama.DramaShortListResp{
		List: resp,
	}, nil
}
