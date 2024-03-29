package sms

import (
	"net/http"

	"carservice/internal/logic/sms"
	"carservice/internal/pkg/common/errcode"
	stdresponse "carservice/internal/pkg/httper/response"
	smsutil "carservice/internal/pkg/sms"
	"carservice/internal/svc"
	"carservice/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			stdresponse.ResponseWithCtx(r.Context(), w, nil, errcode.New(http.StatusBadRequest, "feature.", err.Error()))
			return
		}

		// Customize validation.
		if !smsutil.CheckPhoneNumber(req.PhoneNumber) {
			stdresponse.ResponseWithCtx(r.Context(), w, nil, errcode.New(http.StatusBadRequest, "feature.", "无效的手机号码"))
			return
		}

		l := sms.NewSendCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.SendCaptcha(&req)
		stdresponse.Response(w, resp, err)
	}
}
