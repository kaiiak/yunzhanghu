package economy

import (
	"context"

	"github.com/kaiiak/yunzhanghu/core"
	"github.com/kaiiak/yunzhanghu/core/httpclient"
)

type Economy struct {
	core.Core
}

const (
	preStartURI = "/api/aic/new-economy/api-h5/v1/h5url"
)

type (
	reqH5URL struct {
		DealerId     string `json:"dealer_id"`      // 商户代码(必填)
		BrokerId     string `json:"broker_id"`      // 经纪公司(必填)
		DealerUserId string `json:"dealer_user_id"` // 平台企业端的用户ID，在平台企业系统唯一且不变 (必填)
		ClientType   int    `json:"client_type"`    // 客户端类型 2：H5+API(必填)
		NotifyUrl    string `json:"notify_url"`     // 用户注册后注册结果回调通知URL 注意：回调会以平台企业维度回调不会以用户维度回调，即会统一回调平台企业一个接口再由平台企业决定针对该用户的后续处理(必填)
		Color        string `json:"color"`          // H5页面主体颜色 支持自定义，默认蓝色#007AFF
		ReturnUrl    string `json:"return_url"`     // 退出H5页面时通过此URL跳转到指定页面，如果为空，将会执行jsBridge上注册的YZHJScloseView方法
		CustomTitle  string `json:"custom_title"`   // 支持自定义，默认云账户1：个体工商户申请
	}
	retH5URL struct {
		httpclient.CommonResponse
		Data struct {
			H5Url string `json:"h5_url"` //个体工商户注册H5页面URL
		} `json:"data"` //body

	}
)

// GetEconomyH5URL 预启动接口
func (y *Economy) GetEconomyH5URL(ctx context.Context, dealerUserId, notifyUrl, color, returnUrl, customTitle string, clientType int) (h5url string, err error) {
	var (
		apiName = "预启动接口"
		req     = &reqH5URL{
			DealerId:     y.Dealer,
			BrokerId:     y.Broker,
			DealerUserId: dealerUserId,
			ClientType:   clientType,
			NotifyUrl:    notifyUrl,
			Color:        color,
			ReturnUrl:    returnUrl,
			CustomTitle:  customTitle,
		}
		ret = new(retH5URL)
	)
	responseBytes, err := y.PostJSON(ctx, preStartURI, req)
	if err != nil {
		return
	}

	if err = y.DecodeWithError(responseBytes, ret, apiName); err != nil {
		return
	}

	h5url = ret.Data.H5Url
	return
}

const (
	collectURI = "/api/aic/new-economy/api-h5/v1/collect"
)

type (
	reqCollect struct {
		DealerId            string `json:"dealer_id"`              // 商户代码(必填)
		BrokerId            string `json:"broker_id"`              // 经纪公司(必填)
		DealerUserId        string `json:"dealer_user_id"`         // 平台企业端的用户ID，在平台企业系统唯一且不变 (必填)
		PhoneNo             string `json:"phone_no"`               // 手机号 注意：区号和手机号以“-”连接，示例：+86-18618880001，身份信息为空时，此字段为必填项
		IdCard              string `json:"id_card"`                // 身份证号码（如果包含字母，建议字母大写。如果手机号为空，此字段为必填项）
		RealName            string `json:"real_name"`              // 身份证姓名（如果手机号为空，此字段为必填项）
		IdCardAddress       string `json:"id_card_address"`        // 身份证住址
		IdCardAgency        string `json:"id_card_agency"`         // 身份证签发机关
		IdCardNation        string `json:"id_card_nation"`         // 身份证民族（需要传编码，详见附录4.4民族代码表）
		IdCardValidityStart string `json:"id_card_validity_start"` // 身份证有效期开始时间 （时间格式：yyyy-MM-dd）
		IdCardValidityEnd   string `json:"id_card_validity_end"`   // 身份证有效期开始时间 （时间格式：yyyy-MM-dd）
	}
	retCollect struct {
		httpclient.CommonResponse
		Data struct {
			DealerUserId string `json:"dealer_user_id"`
		} `json:"data"`
	}
)

// Register 工商实名信息录入
func (y *Economy) Collect(ctx context.Context, DealerUserId, PhoneNo, IdCard, RealName, IdCardAddress, IdCardAgency, IdCardNation, IdCardValidityStart, IdCardValidityEnd string) (dealerId string, err error) {
	var (
		apiName = "工商实名信息录入"
		req     = &reqCollect{
			DealerId:            y.Dealer,
			BrokerId:            y.Broker,
			DealerUserId:        DealerUserId,
			PhoneNo:             PhoneNo,
			IdCard:              IdCard,
			RealName:            RealName,
			IdCardAddress:       IdCardAddress,
			IdCardAgency:        IdCardAgency,
			IdCardNation:        IdCardNation,
			IdCardValidityStart: IdCardValidityStart,
			IdCardValidityEnd:   IdCardValidityEnd,
		}
		ret = new(retCollect)
	)
	responseBytes, err := y.PostJSON(ctx, collectURI, req)
	if err != nil {
		return
	}

	if err = y.DecodeWithError(responseBytes, ret, apiName); err != nil {
		return
	}
	dealerId = ret.Data.DealerUserId
	return
}

const (
	economyStatusURI = "/api/aic/new-economy/api-h5/v1/status"
)

type EconomyRegisterStatus struct {
	DealerId     string `json:"dealer_id"`      // 商户代码(必填)
	BrokerId     string `json:"broker_id"`      // 经纪公司(必填)
	OpenId       string `json:"open_id"`        // 用户唯一标识不上传平台企业端的用户ID和身份证二要素时，用户唯一标识为必传字段
	IdCard       string `json:"id_card"`        // 身份证号码 不上传用户唯一标识和平台企业端的用户ID时，身份证姓名和证件号码同为必传字段
	RealName     string `json:"real_name"`      // 身份证姓名 不上传用户唯一标识和平台企业端的用户ID时，身份证姓名和证件号码同为必传字段
	DealerUserId string `json:"dealer_user_id"` // 平台企业端的用户ID，在平台企业系统唯一且不变 (必填)
}

type EconomyRegisterStatusResponseBody struct {
	Status              int    `json:"status"`                //注册状态
	StatusMessage       string `json:"status_message"`        //注册状态描述
	StatusDetail        int    `json:"status_detail"`         //注册详情状态码
	StatusDetailMessage string `json:"status_detail_message"` //注册详情状态码描述
	ApplyedAt           string `json:"applyed_at"`            //注册发起时间 格式：yyyy-MM-dd
	RegistedAt          string `json:"registed_at"`           //注册完成时间 格式：yyyy-MM-dd
	Uscc                string `json:"uscc"`                  //统一社会信用代码 个体工商户唯一标识，注册成功返回
}
type EconomyRegisterStatusResponse struct {
	httpclient.CommonResponse
	Data EconomyRegisterStatusResponseBody `json:"data"` //body

}

type (
	reqEconomyStatus struct {
		DealerId     string `json:"dealer_id"`      // 商户ID(必填)
		BrokerId     string `json:"broker_id"`      // 综合服务主体ID(必填)
		OpenId       string `json:"open_id"`        // 用户唯一标识 不上传平台企业端的用户ID和身份证二要素时， 用户唯一标识为必传字段
		RealName     string `json:"real_name"`      // 身份证姓名 不上传平台企业端的用户ID和身份证二要素时， 用户姓名为必传字段
		IdCard       string `json:"id_card"`        // 身份证号码 不上传平台企业端的用户ID和身份证二要素时， 身份证号码为必传字段
		DealerUserId string `json:"dealer_user_id"` // 平台企业端的用户ID 不上传用户唯一标识、身份证姓名、身份证号码时，平台企业端的用户ID为必传字段
	}
	retEconomyStatus struct {
		httpclient.CommonResponse
		Data EconomyStatus `json:"data"`
	}
	EconomyStatus struct {
		Status              int64  `json:"status"`
		StatusMessage       string `json:"status_message"`
		StatusDetail        int64  `json:"status_detail"`
		StatusDetailMessage string `json:"status_detail_message"`
		ApplyedAt           string `json:"applyed_at"`
		RegistedAt          string `json:"registed_at"`
		USCC                string `json:"uscc"`
		RealName            string `json:"real_name"`
		IdCard              string `json:"id_card"`
	}
)

func (y *Economy) GetEconomyStatusWithOpenId(ctx context.Context, OpenId string) (status EconomyStatus, err error) {
	var (
		apiName = "查询个体工商户状态"
		req     = &reqEconomyStatus{
			DealerId: y.Dealer,
			BrokerId: y.Broker,
			OpenId:   OpenId,
		}
		ret = new(retEconomyStatus)
	)
	responseBytes, err := y.PostJSON(ctx, economyStatusURI, req)
	if err != nil {
		return status, err
	}

	if err = y.DecodeWithError(responseBytes, ret, apiName); err != nil {
		return status, err
	}

	status = ret.Data
	return status, nil
}

func (y *Economy) GetEconomyStatusWithIdCard(ctx context.Context, readName, idCard, dealerUserId string) (status EconomyStatus, err error) {
	var (
		apiName = "查询个体工商户状态"
		req     = &reqEconomyStatus{
			DealerId:     y.Dealer,
			BrokerId:     y.Broker,
			RealName:     readName,
			IdCard:       idCard,
			DealerUserId: dealerUserId,
		}
		ret = new(retEconomyStatus)
	)
	responseBytes, err := y.GetJson(ctx, economyStatusURI, req)
	if err != nil {
		return status, err
	}

	if err = y.DecodeWithError(responseBytes, ret, apiName); err != nil {
		return status, err
	}

	status = ret.Data
	return status, nil
}
