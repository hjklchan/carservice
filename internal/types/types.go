// Code generated by goctl. DO NOT EDIT.
package types

type PingReq struct {
	HttpCode int `form:"httpCode, default=200"`
}

type PingRep struct {
	Result string `json:"result"`
}

type PingDbReq struct {
	AsResult string `form:"asResult, default=carservice"`
}

type PingDbRep struct {
	Result string `json:"result"`
}

type GetUserPhoneNumberReq struct {
	Code string `form:"code"`
}

type GetUserPhoneNumberRep struct {
	PhoneNumber string `json:"phoneNumber"`
}

type UploadFileRep struct {
	AbsolutePath string `json:"absolutePath"`
	RelativePath string `json:"relativePath"`
}

type UploadMultipleFilesRep struct {
	AbsolutePaths []string `json:"absolutePaths"`
	RelativePaths []string `json:"relativePaths"`
}

type WebsocketTestReq struct {
	Arg string `form:"arg"`
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

type GetUserByPhoneNumberReq struct {
	PhoneNumber string `form:"phoneNumber"`
}

type GetUserByPhoneNumberRep struct {
	Username  string `json:"username"`
	AvatarUrl string `json:"avatarUrl"`
}

type WechatAuthorizationReq struct {
	Code string `form:"code"`
}

type WechatAuthorizationRep struct {
	Token string `json:"token"`
}

type MockLoginReq struct {
	Token string `json:"token"`
}

type BrandOptionListItem struct {
	Id     uint   `json:"id"`
	Label  string `json:"label"`
	Pinyin string `json:"pinyin"`
}

type CarBrandOptionListReq struct {
	Page int `form:"page,optional"`
}

type CarBrandOptionListRep struct {
}

type BrandSeriesOptionListItem struct {
	Id     uint   `json:"id"`
	Label  string `json:"label"`
	Pinyin string `json:"pinyin"`
}

type BrandSeriesOptionListReq struct {
	BrandId uint `form:"brandId"`
	Page    int  `form:"page,optional"`
}

type BrandSeriesOptionListRep struct {
}

type CheckEmptyListReq struct {
}

type CheckEmptyListRep struct {
	Listable bool `json:"listable"`
}

type CreateCarOwnerInfoReq struct {
	Name              string  `json:"name"`
	PhoneNumber       string  `json:"phoneNumber"`
	MultilevelAddress string  `json:"multilevelAddress"`
	FullAddress       string  `json:"fullAddress"`
	Longitude         float64 `json:"longitude"`
	Latitude          float64 `json:"latitude"`
}

type CreateCarOwnerInfoRep struct {
}

type UpdateCarOwnerInfoReq struct {
	Id                uint    `path:"id"`
	Name              string  `json:"name"`
	PhoneNumber       string  `json:"phoneNumber"`
	MultilevelAddress string  `json:"multilevelAddress"`
	FullAddress       string  `json:"fullAddress"`
	Longitude         float64 `json:"longitude"`
	Latitude          float64 `json:"latitude"`
}

type UpdateCarOwnerInfoRep struct {
}

type CarOwnerInfoListItem struct {
	Id                uint   `json:"id"`
	Name              string `json:"name"`
	PhoneNumber       string `json:"phoneNumber"`
	MultilevelAddress string `json:"multilevelAddress"`
	FullAddress       string `json:"fullAddress"`
}

type GetCarOwnerInfoListReq struct {
	Id uint `path:"id"`
}

type GetCarOwnerInfoListRep struct {
}

type GetCarOwnerInfoReq struct {
	Id uint `path:"id"`
}

type GetCarOwnerInfoRep struct {
	Id                uint   `json:"id"`
	Name              string `json:"name"`
	PhoneNumber       string `json:"phoneNumber"`
	MultilevelAddress string `json:"multilevelAddress"`
	FullAddress       string `json:"fullAddress"`
}

type DeleteCarOwnerInfoReq struct {
	Id uint `path:"id"`
}

type CreateUserOrderReq struct {
	CarOwnerName              string  `json:"carOwnerName"`
	CarOwnerPhoneNumber       string  `json:"carOwnerPhoneNumber"`
	CarOwnerLongitude         float64 `json:"carOwnerLongitude"`
	CarOwnerLatitude          float64 `json:"carOwnerLatitude"`
	CarOwnerMultilevelAddress string  `json:"carOwnerMultilevelAddress"`
	CarOwnerFullAddress       string  `json:"carOwnerFullAddress"`
	CarBrandId                uint    `json:"carBrandId"`
	CarBrandSeriesId          uint    `json:"carBrandSeriesId"`
	PartnerStoreId            uint    `json:"partnerStoreId"`
	Requirements              string  `json:"requirements"`
}

type CreateUserOrderRep struct {
}
