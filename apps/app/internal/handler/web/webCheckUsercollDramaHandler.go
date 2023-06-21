package web

import (
	"audio/common/result"
	"net/http"

	"audio/apps/app/internal/logic/web"
	"audio/apps/app/internal/svc"
	"audio/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func WebCheckUsercollDramaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckCollDramaReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := web.NewWebCheckUsercollDramaLogic(r.Context(), svcCtx)
		resp, err := l.WebCheckUsercollDrama(&req)
		result.HttpResult(r, w, resp, err)
	}
}
