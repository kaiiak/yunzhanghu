package yunzhanghu

import "context"

const (
	registerURI = "/api/aic/new-economy/api-h5/v1/collect"
)

type RegisterInfo struct {
	DealerId            string `json:"dealer_id"`              // 商户代码(必填)
	BrokerId            string `json:"broker_id"`              // 经纪公司(必填)
	DealerUserId        string `json:"dealer_user_id"`         // 平台企业端的用户ID，在平台企业系统唯一且不变 (必填)
	PhoneNo             string `json:"phone_no"`               // 手机号 注意：区号和手机号以“-”连接，示例：+86-18618880001，身份信息为空时，此字段为必填项
	IdCard              string `json:"id_card"`                // 身份证号码（如果包含字母，建议字母大写。如果手机号为空，此字段为必填项）
	RealName            string `json:"real_name"`              // 身份证姓名（如果手机号为空，此字段为必填项）
	IdCardAddress       string `json:"id_card_address"`        // 身份证住址
	IdCardAgency        string `json:"id_card_agency"`         // 身份证签发机关
	IdCardNation        string `json:"id_card_nation"`         // 身份证民族（需要传编码，详见附录4.4民族代码表）
	IdCardValidityStart string `json:"id_card_validity_start"` //身份证有效期开始时间 （时间格式：yyyy-MM-dd）
	IdCardValidityEnd   string `json:"id_card_validity_end"`   //身份证有效期开始时间 （时间格式：yyyy-MM-dd）
}

type RegisterInfoResponseDealerId struct {
	DealerUserId string `json:"dealer_user_id"`
}
type RegisterInfoResponse struct {
	CommonResponse
	Data RegisterInfoResponseDealerId `json:"data"`
}

//Register 发起工商注册
func (y *Yunzhanghu) Register(ctx context.Context, DealerUserId, PhoneNo, IdCard, RealName, IdCardAddress, IdCardAgency, IdCardNation, IdCardValidityStart, IdCardValidityEnd string) (dealerId string, err error) {
	var (
		apiName = "发起工商注册"
		req     = &RegisterInfo{
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
		ret = new(RegisterInfoResponse)
	)
	responseBytes, err := y.postJSON(ctx, registerURI, apiName, req)
	if err != nil {
		return dealerId, err
	}

	if err = y.decodeWithError(responseBytes, ret, apiName); err != nil {
		return dealerId, err
	}

	return ret.Data.DealerUserId, nil
}
