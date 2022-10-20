package yunzhanghu

import "context"

const (
	preStartURI = "/api/aic/new-economy/api-h5/v1/h5url"
)

type PreStart struct {
	DealerId     string `json:"dealer_id"`      // 商户代码(必填)
	BrokerId     string `json:"broker_id"`      // 经纪公司(必填)
	DealerUserId string `json:"dealer_user_id"` // 平台企业端的用户ID，在平台企业系统唯一且不变 (必填)
	ClientType   int    `json:"client_type"`    //客户端类型 2：H5+API(必填)
	NotifyUrl    string `json:"notify_url"`     //用户注册后注册结果回调通知URL 注意：回调会以平台企业维度回调不会以用户维度回调，即会统一回调平台企业一个接口再由平台企业决定针对该用户的后续处理(必填)
	Color        string `json:"color"`          //H5页面主体颜色 支持自定义，默认蓝色#007AFF
	ReturnUrl    string `json:"return_url"`     //退出H5页面时通过此URL跳转到指定页面，如果为空，将会执行jsBridge上注册的YZHJScloseView方法
	CustomTitle  string `json:"custom_title"`   //支持自定义，默认云账户1：个体工商户申请
}

type PreStartResponseBody struct {
	H5Url string `json:"h5_url"` //个体工商户注册H5页面URL
}
type PreStartResponse struct {
	CommonResponse
	Data PreStartResponseBody `json:"data"` //body

}

//GetPreStartUrl 预启动接口
func (y *Yunzhanghu) GetPreStartUrl(ctx context.Context, dealerUserId, notifyUrl, color, returnUrl, customTitle string, clientType int) (url string, err error) {
	var (
		apiName = "预启动接口"
		req     = &PreStart{
			DealerId:     y.Dealer,
			BrokerId:     y.Broker,
			DealerUserId: dealerUserId,
			ClientType:   clientType,
			NotifyUrl:    notifyUrl,
			Color:        color,
			ReturnUrl:    returnUrl,
			CustomTitle:  customTitle,
		}
		ret = new(PreStartResponse)
	)
	responseBytes, err := y.postJSON(ctx, preStartURI, apiName, req)
	if err != nil {
		return url, err
	}

	if err = y.decodeWithError(responseBytes, ret, apiName); err != nil {
		return url, err
	}

	url = ret.Data.H5Url
	return url, nil
}
