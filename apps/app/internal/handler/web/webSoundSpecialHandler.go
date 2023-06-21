package web

import (
	"audio/apps/app/internal/logic/web"
	"audio/common/result"
	"net/http"

	"audio/apps/app/internal/svc"
	"audio/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func WebSoundSpecialHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Req
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := web.NewWebSoundSpecialLogic(r.Context(), svcCtx)
		resp, err := l.WebSoundSpecial(&req)
		result.HttpResult(r, w, resp, err)
	}
}
