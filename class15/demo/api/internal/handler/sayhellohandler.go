package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lanshan/Courseware-Backend-Go-2022/class15/demo/api/internal/logic"
	"lanshan/Courseware-Backend-Go-2022/class15/demo/api/internal/svc"
	"lanshan/Courseware-Backend-Go-2022/class15/demo/api/internal/types"
)

func SayHelloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SayHelloReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSayHelloLogic(r.Context(), svcCtx)
		resp, err := l.SayHello(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
