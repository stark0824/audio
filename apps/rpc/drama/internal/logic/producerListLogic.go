package logic

import (
	"audio/apps/model/producer"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"time"

	"audio/apps/rpc/drama/drama"
	"audio/apps/rpc/drama/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProducerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProducerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProducerListLogic {
	return &ProducerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProducerList 后台 - 全部工作室列表
func (l *ProducerListLogic) ProducerList(in *drama.ProducerAllReq) (*drama.ProducerAllResp, error) {
	whereBuilder := l.svcCtx.ProducerModel.RowBuilder()

	//Order 1 id desc
	list, err := l.svcCtx.ProducerModel.FindAll(l.ctx, whereBuilder, int(in.Order))

	if err != nil && err != producer.ErrNotFound {
		return nil, errors.New("UNKNOWN")
	}

	var resp []*drama.Producer
	if len(list) > 0 {
		for _, Producer := range list {
			var pbProducer drama.Producer
			_ = copier.Copy(&pbProducer, Producer)
			CreatedAt := time.Unix(int64(Producer.CreatedAt), 0).Format("2006-01-02 15:04:05")
			pbProducer.CreatedAt = CreatedAt
			resp = append(resp, &pbProducer)
		}
	}

	return &drama.ProducerAllResp{
		List: resp,
	}, nil
}
