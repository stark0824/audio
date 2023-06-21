package admin

import (
	"audio/apps/rpc/drama/dramaclient"
	"audio/common/globalkey"
	"audio/common/xerr"
	"context"
	"github.com/jinzhu/copier"

	"audio/apps/app/internal/svc"
	"audio/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminDramaInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminDramaInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminDramaInfoLogic {
	return &AdminDramaInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminDramaInfoLogic) AdminDramaInfo(req *types.DramaInfoRequest) (resp *types.DramaInfoResp, err error) {

	if req.DramaId <= 0 {
		return nil, xerr.NewErrCode(xerr.INVALID_ARGUMENT)
	}

	dramaInfoResp, err := l.svcCtx.DramaRpcClient.DramaInfo(l.ctx, &dramaclient.DramaInfoReq{
		Id: int64(req.DramaId),
	})

	if err != nil {
		return nil, xerr.NewErrCode(xerr.UNKNOWN)
	}

	var dramaInfo types.Info
	_ = copier.Copy(&dramaInfo, dramaInfoResp.Drama)

	Producer, err := l.svcCtx.DramaRpcClient.ProducerDict(l.ctx, &dramaclient.Req{})

	var typesProducerLists []types.Producer

	if len(Producer.List) > 0 {
		for _, ProducerList := range Producer.List {
			var typeProducer types.Producer
			_ = copier.Copy(&typeProducer, ProducerList)

			typesProducerLists = append(typesProducerLists, typeProducer)
		}
	}

	var PayTypeList = []types.PayType{
		{Value: 0, Label: "免费收听"},
		{Value: 1, Label: "全集购买"},
		{Value: 2, Label: "打包购买"},
	}

	return &types.DramaInfoResp{
		DramaInfo:    dramaInfo,
		Category:     globalkey.Category,
		Process:      globalkey.Process,
		Status:       globalkey.Status,
		CataLog:      globalkey.CataLog,
		ProducerList: typesProducerLists,
		PayType:      PayTypeList,
	}, nil
}
