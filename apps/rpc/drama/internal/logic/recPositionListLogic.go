package logic

import (
	"audio/apps/model/producer"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"time"

	"audio/apps/rpc/drama/drama"
	"audio/apps/rpc/drama/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecPositionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecPositionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecPositionListLogic {
	return &RecPositionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RecPositionList 后台 - 有声推荐列表
func (l *RecPositionListLogic) RecPositionList(in *drama.RecPositionReq) (*drama.RecPositionResp, error) {
	whereBuilder := l.svcCtx.RecommendModel.RowBuilder().Where(squirrel.Eq{"is_del": 0})

	list, err := l.svcCtx.RecommendModel.FindAll(l.ctx, whereBuilder)

	if err != nil && err != producer.ErrNotFound {
		return nil, errors.New("UNKNOWN")
	}

	var resp []*drama.RecPosition
	if len(list) > 0 {
		for _, Producer := range list {
			var pbRecPosition drama.RecPosition
			_ = copier.Copy(&pbRecPosition, Producer)
			pbRecPosition.CreatedAt = time.Unix(Producer.CreatedAt, 0).Format("2006-01-02 15:04:05")
			pbRecPosition.UpdatedAt = time.Unix(Producer.UpdatedAt, 0).Format("2006-01-02 15:04:05")
			resp = append(resp, &pbRecPosition)
		}
	}

	return &drama.RecPositionResp{
		List: resp,
	}, nil
}
