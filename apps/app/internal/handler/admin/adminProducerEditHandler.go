package admin

import (
	"audio/apps/app/internal/logic/admin"
	"audio/common/result"
	"net/http"

	"audio/apps/app/internal/svc"
	"audio/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AdminProducerEditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProducerEditReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := admin.NewAdminProducerEditLogic(r.Context(), svcCtx)
		resp, err := l.AdminProducerEdit(&req)
		result.HttpResult(r, w, resp, err)
	}
}
