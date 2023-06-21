package admin

import (
	"audio/apps/app/internal/logic/admin"
	"audio/apps/app/internal/types"
	"audio/common/result"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"audio/apps/app/internal/svc"
)

func AdminRecommendAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendDataReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := admin.NewAdminRecommendAddLogic(r.Context(), svcCtx)
		resp, err := l.AdminRecommendAdd(&req)
		result.HttpResult(r, w, resp, err)
	}
}
