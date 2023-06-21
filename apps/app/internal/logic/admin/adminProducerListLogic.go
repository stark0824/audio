package admin

import (
	"audio/apps/rpc/drama/dramaclient"
	"audio/common/globalkey"
	"context"
	"github.com/jinzhu/copier"

	"audio/apps/app/internal/svc"
	"audio/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminProducerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminProducerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminProducerListLogic {
	return &AdminProducerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminProducerListLogic) AdminProducerList(req *types.Req) (resp *types.ProducerListResp, err error) {
	Producer, err := l.svcCtx.DramaRpcClient.ProducerList(l.ctx, &dramaclient.ProducerAllReq{
		Order: 1,
	})

	var typesProducerLists []types.ProducerList

	if len(Producer.List) > 0 {
		for _, Item := range Producer.List {
			var typeProducer types.ProducerList
			_ = copier.Copy(&typeProducer, Item)

			v, _ := globalkey.ProducerStatus[int(Item.Status)]
			if v == "" {
				typeProducer.StatusVal = ""
			} else {
				typeProducer.StatusVal = v
			}

			typesProducerLists = append(typesProducerLists, typeProducer)
		}
	}

	return &types.ProducerListResp{
		List:   typesProducerLists,
		Status: globalkey.ProducerStatus,
	}, nil
}
