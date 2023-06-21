package logic

import (
	"audio/apps/model/producer"
	"audio/apps/rpc/drama/drama"
	"audio/apps/rpc/drama/internal/svc"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProducerDictLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProducerDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProducerDictLogic {
	return &ProducerDictLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProducerDictLogic) ProducerDict(in *drama.Req) (*drama.ProducerDictResp, error) {
	whereBuilder := l.svcCtx.ProducerModel.RowBuilder()
	list, err := l.svcCtx.ProducerModel.FindAllByNormal(l.ctx, whereBuilder)

	if err != nil && err != producer.ErrNotFound {
		logx.Info(err)
		return nil, errors.New("UNKNOWN")
	}

	var resp []*drama.ProducerDict
	if len(list) > 0 {
		for _, Producer := range list {
			var pbProducer drama.ProducerDict
			_ = copier.Copy(&pbProducer, Producer)
			resp = append(resp, &pbProducer)
		}
	}

	return &drama.ProducerDictResp{
		List: resp,
	}, nil

}
