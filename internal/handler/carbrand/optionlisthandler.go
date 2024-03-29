package carbrand

import (
	"net/http"

	"carservice/internal/logic/carbrand"
	"carservice/internal/pkg/common/errcode"
	"carservice/internal/pkg/httper/api"
	"carservice/internal/svc"
	"carservice/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func OptionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CarBrandOptionListReq
		if err := httpx.Parse(r, &req); err != nil {
			api.ResponseWithCtx(r.Context(), w, nil, errcode.New(http.StatusBadRequest, "feature.", err.Error()))
			return
		}
		l := carbrand.NewOptionListLogic(r.Context(), svcCtx)
		resp, err := l.OptionList()
		api.Response(w, resp, err)
	}
}
