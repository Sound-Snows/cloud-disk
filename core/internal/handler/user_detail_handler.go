package handler

import (
	"cloud-disk/core/internal/types"
	"net/http"

	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDatailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewUserDetailLogic(r.Context(), svcCtx)
		resp, err := l.UserDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}