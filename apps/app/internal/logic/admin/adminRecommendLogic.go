package admin

import (
	"audio/apps/rpc/drama/dramaclient"
	"audio/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"strings"

	"audio/apps/app/internal/svc"
	"audio/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminRecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminRecommendLogic {
	return &AdminRecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminRecommendLogic) AdminRecommend(req *types.RecommendListReq) (resp *types.RecommendListResp, err error) {
	recResp, err := l.svcCtx.DramaRpcClient.RecPositionList(l.ctx, &dramaclient.RecPositionReq{})

	if err != nil {
		return nil, xerr.NewErrCode(xerr.UNKNOWN)
	}

	var (
		typesRecLists []types.RecPosition
	)

	if len(recResp.List) > 0 {
		for _, Recommend := range recResp.List {
			var typeRecommend types.RecPosition
			_ = copier.Copy(&typeRecommend, Recommend)

			if Recommend.TempType == 1 {
				typeRecommend.TempName = "单排"
			} else if Recommend.TempType == 2 {
				typeRecommend.TempName = "横排"
			} else if Recommend.TempType == 3 {
				typeRecommend.TempName = "竖排"
			}

			ids := strings.Replace(typeRecommend.NovelIds, "|", ",", -1)
			drama, err := l.svcCtx.DramaRpcClient.DramaShortList(l.ctx, &dramaclient.DramaShortListReq{
				Ids: ids,
			})

			if err != nil || len(drama.List) == 0 {
				typeRecommend.NovelIds = ""
			}

			if len(drama.List) > 0 {
				var novelName []string
				for _, Item := range drama.List {
					novelName = append(novelName, Item.Name)
				}
				typeRecommend.NovelName = strings.Join(novelName, "|")
			}
			typesRecLists = append(typesRecLists, typeRecommend)
		}
	}

	return &types.RecommendListResp{
		List:  typesRecLists,
		Count: int64(len(recResp.List)),
	}, nil
}
