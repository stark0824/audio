package api

import (
	"audio/apps/rpc/drama/dramaclient"
	"context"
	"github.com/jinzhu/copier"
	"strings"

	"audio/apps/app/internal/svc"
	"audio/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiSoundSpecialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiSoundSpecialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiSoundSpecialLogic {
	return &ApiSoundSpecialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApiSoundSpecialLogic) ApiSoundSpecial(req *types.Req) (resp *types.ShortRecPositionResp, err error) {
	recResp, err := l.svcCtx.DramaRpcClient.RecPositionList(l.ctx, &dramaclient.RecPositionReq{
		Limit:    1,
		PageSize: 20,
	})

	if err != nil {
		return nil, err
	}

	var typesRecLists []types.ShortRecPosition

	if len(recResp.List) > 0 {
		for _, Recommend := range recResp.List {
			var typeRecommend types.ShortRecPosition
			_ = copier.Copy(&typeRecommend, Recommend)

			// -drama_list数据
			ids := strings.Replace(Recommend.NovelIds, "|", ",", -1)
			drama, err := l.svcCtx.DramaRpcClient.DramaShortList(l.ctx, &dramaclient.DramaShortListReq{
				Ids: ids,
			})

			var SubList []types.ShortInfo
			if err != nil || len(drama.List) == 0 {
				SubList = nil
			}

			if len(drama.List) > 0 {
				for _, Item := range drama.List {
					var subRecommend types.ShortInfo
					_ = copier.Copy(&subRecommend, Item)
					SubList = append(SubList, subRecommend)
				}

			}
			typeRecommend.List = SubList
			//drama_list数据
			typesRecLists = append(typesRecLists, typeRecommend)
		}
	}

	return &types.ShortRecPositionResp{
		List: typesRecLists,
	}, nil
}
