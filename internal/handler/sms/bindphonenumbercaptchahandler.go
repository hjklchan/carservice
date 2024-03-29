package sms

import (
	"net/http"

	"carservice/internal/logic/sms"
	"carservice/internal/pkg/common/errcode"
	stdresponse "carservice/internal/pkg/httper/response"
	"carservice/internal/svc"
	"carservice/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func BindPhoneNumberCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			stdresponse.ResponseWithCtx(r.Context(), w, nil, errcode.New(http.StatusBadRequest, "feature.", err.Error()))
			return
		}

		l := sms.NewBindPhoneNumberCaptchaLogic(r.Context(), svcCtx)
		err := l.BindPhoneNumberCaptcha(&req)
		stdresponse.Response(w, nil, err)
	}
}
