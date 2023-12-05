// Code generated by goctl. DO NOT EDIT.
package types

type PingReq struct {
	HttpCode int `form:"httpCode, default=200"`
}

type PingRep struct {
	Result string `json:"result"`
}

type PingDbReq struct {
	Value string `json:"value, default=carservice"`
}

type PingDbRep struct {
	Result string `json:"result"`
}

type SendCaptchaReq struct {
	PhoneNumber string `form:"phoneNumber"`
}

type SendCaptchaRep struct {
}

type PhoneNumberLoginReq struct {
	PhoneNumber string `json:"phoneNumber"`
	Captcha     string `json:"captcha"`
}

type PhoneNumberLoginRep struct {
	Token string `json:"token"`
}
