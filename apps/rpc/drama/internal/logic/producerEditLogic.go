package logic

import (
	"audio/apps/model/producer"
	"context"
	"errors"
	"time"

	"audio/apps/rpc/drama/drama"
	"audio/apps/rpc/drama/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProducerEditLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProducerEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProducerEditLogic {
	return &ProducerEditLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProducerEditLogic) ProducerEdit(in *drama.ProducerEditReq) (*drama.ProducerEditResp, error) {

	if in.Id < 0 {
		return nil, errors.New("INVALID_ARGUMENT")
	}

	producer := new(producer.Producer)
	if in.Id > 0 {
		_, err := l.svcCtx.ProducerModel.FindOne(l.ctx, in.Id)
		if err != nil {
			return nil, errors.New("INVALID_ARGUMENT")
		}

		producer.Id = in.Id
		producer.Name = in.Name
		producer.Uid = in.Uid
		producer.Mark = in.Mark
		producer.Status = in.Status

		err = l.svcCtx.ProducerModel.Update(l.ctx, producer)
		if err != nil {
			return nil, errors.New("INVALID_ARGUMENT")
		}
	}

	if in.Id == 0 {
		producer.Name = in.Name
		producer.Uid = in.Uid
		producer.Mark = in.Mark
		producer.Status = in.Status
		producer.CreatedAt = time.Now().Unix()
		resp, err := l.svcCtx.ProducerModel.Insert(l.ctx, producer)
		if err != nil {
			return nil, errors.New("UNKNOWN")
		}

		producer.Id, _ = resp.LastInsertId()

	}

	return &drama.ProducerEditResp{
		Id: int64(producer.Id),
	}, nil

}
