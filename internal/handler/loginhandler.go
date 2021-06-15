package handler

import (
	"go-zero-token/utils"
	"net/http"

	"go-zero-token/internal/logic"
	"go-zero-token/internal/svc"
	"go-zero-token/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func LoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), ctx)
		resp, err := l.Login(req)
		if err != nil {
			utils.NewErrorJson(w, err)
		} else {
			httpx.OkJson(w,utils.NewOkBaseRep(resp))
		}
	}
}
