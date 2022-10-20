package yunzhanghu

import "context"

const (
	registerStatusURI = "/api/aic/new-economy/api-h5/v1/status"
)

type RegisterStatus struct {
	DealerId     string `json:"dealer_id"`      // 商户代码(必填)
	BrokerId     string `json:"broker_id"`      // 经纪公司(必填)
	OpenId       string `json:"open_id"`        // 用户唯一标识不上传平台企业端的用户ID和身份证二要素时，用户唯一标识为必传字段
	IdCard       string `json:"id_card"`        // 身份证号码 不上传用户唯一标识和平台企业端的用户ID时，身份证姓名和证件号码同为必传字段
	RealName     string `json:"real_name"`      // 身份证姓名 不上传用户唯一标识和平台企业端的用户ID时，身份证姓名和证件号码同为必传字段
	DealerUserId string `json:"dealer_user_id"` // 平台企业端的用户ID，在平台企业系统唯一且不变 (必填)
}

type RegisterStatusResponseBody struct {
	Status              int    `json:"status"`                //注册状态
	StatusMessage       string `json:"status_message"`        //注册状态描述
	StatusDetail        int    `json:"status_detail"`         //注册详情状态码
	StatusDetailMessage string `json:"status_detail_message"` //注册详情状态码描述
	ApplyedAt           string `json:"applyed_at"`            //注册发起时间 格式：yyyy-MM-dd
	RegistedAt          string `json:"registed_at"`           //注册完成时间 格式：yyyy-MM-dd
	Uscc                string `json:"uscc"`                  //统一社会信用代码 个体工商户唯一标识，注册成功返回
}
type RegisterStatusResponse struct {
	CommonResponse
	Data RegisterStatusResponseBody `json:"data"` //body

}

//GetRegisterStatus 查询个体工商户状态
func (y *Yunzhanghu) GetRegisterStatus(ctx context.Context, dealerUserId, openId, realName, idCard string) (data RegisterStatusResponseBody, err error) {
	var (
		apiName = "查询个体工商户状态"
		req     = &RegisterStatus{
			DealerId:     y.Dealer,
			BrokerId:     y.Broker,
			DealerUserId: dealerUserId,
			OpenId:       openId,
			RealName:     realName,
			IdCard:       idCard,
		}
		ret = new(RegisterStatusResponse)
	)
	responseBytes, err := y.getJson(ctx, registerStatusURI, apiName, req)
	if err != nil {
		return data, err
	}

	if err = y.decodeWithError(responseBytes, ret, apiName); err != nil {
		return data, err
	}

	data = ret.Data
	return data, nil
}
