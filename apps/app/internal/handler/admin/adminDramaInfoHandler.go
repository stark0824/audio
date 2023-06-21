package admin

import (
	"audio/apps/app/internal/logic/admin"
	"audio/common/result"
	"net/http"

	"audio/apps/app/internal/svc"
	"audio/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AdminDramaInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DramaInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := admin.NewAdminDramaInfoLogic(r.Context(), svcCtx)
		resp, err := l.AdminDramaInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
