// Code generated by goctl. DO NOT EDIT.
package types

type ServerPingReq struct {
	HttpCode int `form:"httpCode, default=200"`
}

type ServerPingRep struct {
	Result string `json:"result"`
}

type ServerPingDbReq struct {
	AsResult string `form:"asResult, default=carservice"`
}

type ServerPingDbRep struct {
	Result string `json:"result"`
}

type GetAppsetRep struct {
	AppName string `json:"appName"`
	AppDesc string `json:"appDesc"`
	AppVers string `json:"appVers"`
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

type UploadImageReq struct {
}

type UploadImageRep struct {
	Url string `json:"url"`
}

type WebsocketTestReq struct {
	Arg string `form:"arg"`
}

type WebsocketTestRep struct {
}

type WebsocketServiceListItem struct {
	ServiceId   uint8  `json:"serviceId"`
	ServiceDesc string `json:"serviceDesc"`
}

type WebsocketServiceReq struct {
}

type WebsocketServiceRep struct {
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
	UserId int64 `form:"userId,optional"`
}

type MockLoginRep struct {
	Token string `json:"token"`
}

type GetUserProfileReq struct {
}

type GetUserProfileRep struct {
	Id          uint   `json:"id"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	AvatarUrl   string `json:"avatarUrl"`
}

type UpdateUserProfileReq struct {
	Username  string `json:"username"`
	AvatarUrl string `json:"avatarUrl"`
}

type UpdateUserProfileRep struct {
}

type CarBrandOptionListItem struct {
	Id     uint   `json:"id"`
	Label  string `json:"label"`
	Pinyin string `json:"pinyin"`
}

type CarBrandOptionListReq struct {
	Page int `form:"page,optional"`
}

type CarBrandOptionListRep struct {
}

type CarBrandSeriesOptionListItem struct {
	Id     uint   `json:"id"`
	Label  string `json:"label"`
	Pinyin string `json:"pinyin"`
}

type GetCarBrandSeriesOptionListReq struct {
	BrandId uint `form:"brandId"`
	Page    int  `form:"page,optional"`
}

type GetCarBrandSeriesOptionListRep struct {
}

type CarOwnerInfoCheckEmptyListRep struct {
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

type UserOrderStatusListItem struct {
	Status string `json:"status"`
	Label  string `json:"label"`
	GoTag  string `json:"goTag"`
}

type UserOrderListItem struct {
	Id           uint   `json:"id"`
	Deletable    bool   `json:"deletable"`
	OrderNumber  string `json:"orderNumber"`
	PartnerStore string `json:"partnerStore"`
	Requirements string `json:"requirements"`
	OrderStatus  string `json:"orderStatus"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type CarReplacementItem struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	EstF32Price float32 `json:"estF32Price"`
	EstU64Price uint64  `json:"estU64Price"`
	Counter     uint    `json:"counter"`
}

type CreateUserOrderReq struct {
	CarOwnerName        string               `json:"carOwnerName"`
	CarOwnerPhoneNumber string               `json:"carOwnerPhoneNumber"`
	CarOwnerLongitude   float64              `json:"carOwnerLongitude"`
	CarOwnerLatitude    float64              `json:"carOwnerLatitude"`
	CarOwnerMultiLvAddr string               `json:"carOwnerMultiLvAddr"`
	CarOwnerFullAddress string               `json:"carOwnerFullAddress"`
	CarBrandId          int64                `json:"carBrandId"`
	CarSeriesId         int64                `json:"carSeriesId"`
	PartnerStoreId      int64                `json:"partnerStoreId,optional"`
	Requirements        string               `json:"requirements"`
	AgreeToTerms        uint8                `json:"agreeToTerms"`
	CarReplacements     []CarReplacementItem `json:"carReplacements"`
}

type CreateUserOrderRep struct {
	NewId uint `json:"newId"`
}

type UpdateUserOrderReq struct {
	CarOwnerName        string  `json:"carOwnerName"`
	CarOwnerPhoneNumber string  `json:"carOwnerPhoneNumber"`
	CarOwnerLongitude   float64 `json:"carOwnerLongitude"`
	CarOwnerLatitude    float64 `json:"carOwnerLatitude"`
	CarOwnerMultLvAddr  string  `json:"carOwnerMultLvAddr"`
	CarOwnerFullAddress string  `json:"carOwnerFullAddress"`
	CarBrandId          uint    `json:"carBrandId"`
	CarSeriesId         uint    `json:"carSeriesId"`
	PartnerStoreId      uint    `json:"partnerStoreId"`
	Requirements        string  `json:"requirements"`
}

type UpdateUserOrderRep struct {
}

type GetUserOrderReq struct {
	Id uint `path:"id"`
}

type GetUserOrderRep struct {
	Id                  uint   `json:"id"`
	OrderNumber         string `json:"orderNumber"`
	CarOwnerName        string `json:"carOwnerName"`
	CarOwnerMultiLvAddr string `json:"carOwnerMultiLvAddr"`
	CarOwnerFullAddr    string `json:"carOwnerFullAddress"`
	CarBrandName        string `json:"carBrandName"`
	CarSeriesName       string `json:"carSeriesName"`
	PartnerStore        string `json:"partnerStore"`
	PartnerStoreAddr    string `json:"partnerStoreAddr"`
	Requirements        string `json:"requirements"`
	OrderStatus         string `json:"orderStatus"`
	CreatedAt           string `json:"createdAt"`
	UpdatedAt           string `json:"updatedAt"`
	CarOwnerPhoneNumber string `json:"carOwnerPhoneNumber"`
	CarBrandId          uint   `json:"carBrandId"`
	CarSeriesId         uint   `json:"carSeriesId"`
}

type GetUserOrderListReq struct {
	Page      int    `form:"page,optional"`
	Status    string `form:"status,optional"`
	DateStart string `form:"dateStart,optional"`
	DateEnd   string `form:"dateEnd,optional"`
}

type GetUserOrderListRep struct {
}

type DeleteUserOrderReq struct {
	Id uint `path:"id"`
}

type CancelUserOrderReq struct {
	Id int64 `path:"id"`
}

type PaymentOrderReq struct {
	Id              int64 `path:"id"`
	CarReplacements []int `json:"carReplacements"`
}

type PaymentOrderRep struct {
	GoSdkVersion string `json:"goSdkVersion"`
	PrepayId     string `json:"prepayId"`
	TimeStamp    string `json:"timeStamp"`
	NonceStr     string `json:"nonceStr"`
	Package      string `json:"package"`
	SignType     string `json:"signType"`
	PaySign      string `json:"paySign"`
}

type PaymentOrderCallbackReq struct {
}

type PaymentOrderCallbackRep struct {
}

type RefundOrderReq struct {
	Id int64 `path:"id"`
}

type RefundOrderRep struct {
}

type RefundOrderCallbackReq struct {
}

type RefundOrderCallbackRep struct {
}

type PartnerStoreListItem struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	FullAddress string  `json:"fullAddress"`
	Gap         float32 `json:"gap"`
	Unit        string  `json:"unit"`
}

type GetPartnerStoreListReq struct {
	Address  string  `form:"address"`
	LimitGap float32 `form:"limitGap,optional"`
}

type GetPartnerStoreListRep struct {
}

type BulletinListItem struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

type GetBulletinListReq struct {
	Limit int32 `form:"limit,optional"`
}

type CreateOrderCommentReq struct {
	UserOrderId uint   `json:"userOrderId"`
	Title       string `json:"title,optional"`
	Rate        int8   `json:"rate"`
	Content     string `json:"content"`
}

type CreateOrderCommentRep struct {
}

type OrderCommentListItem struct {
}

type GetOrderCommentListReq struct {
}

type DeleteOrderCommentReq struct {
	Id uint `path:"id"`
}

type CarReplacement struct {
	Id          uint             `json:"id"`
	Title       string           `json:"title"`
	EstF32Price float32          `json:"estF32Price"`
	EstU64Price uint64           `json:"estU64Price"`
	Counter     uint             `json:"counter"`
	Children    []CarReplacement `json:"children"`
}

type CarReplacementReq struct {
	CarSeriesId uint `form:"carSeriesId"`
}
