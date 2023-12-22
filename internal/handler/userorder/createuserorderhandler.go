package userorder

import (
	"net/http"

	"carservice/internal/logic/userorder"
	"carservice/internal/pkg/common/errcode"
	stdresponse "carservice/internal/pkg/httper/response"
	"carservice/internal/svc"
	"carservice/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateUserOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			// stdresponse.ResponseWithCtx(r.Context(), w, errcode.New(http.StatusBadRequest, "feature.", err.Error()))
			stdresponse.ResponseWithCtx(r.Context(), w, errcode.InvalidParamsError.Lazy(err.Error()))
			return
		}

		l := userorder.NewCreateUserOrderLogic(r.Context(), svcCtx)
		err := l.CreateUserOrder(&req)
		stdresponse.Response(w, nil, err)
	}
}
