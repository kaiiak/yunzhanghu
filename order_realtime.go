package yunzhanghu

import (
	"context"
)

const (
	paymentOrderRealtimeURI = "https://api-jiesuan.yunzhanghu.com/api/payment/v1/order-realtime"
)

type (
	reqOrderRealtime struct {
		OrderId   string `json:"order_id"`             // 商户订单号，由商户保持唯⼀一性(必填)
		DealerId  string `json:"dealer_id"`            // 商户代码(必填)
		BrokerId  string `json:"broker_id"`            // 经纪公司(必填)
		RealName  string `json:"real_name"`            // 银⾏行行开户姓名(必填)
		CardNo    string `json:"card_no"`              // 银⾏行行开户卡号(必填)
		PhoneNo   string `json:"phone_no"`             // ⽤用户或联系⼈人⼿手机号
		IdCard    string `json:"id_card"`              // 银⾏行行开户身份证号
		Pay       string `json:"pay"`                  // 打款⾦金金额(必填)
		AnchorId  string `json:"anchor_id,omitempty"`  // 打款⼈人id(选填)
		Notes     string `json:"notes,omitempty"`      // 备注信息(选填)
		PayRemark string `json:"pay_remark,omitempty"` // 打款备注(选填，最⼤大20个字符，⼀个汉字占2个字符，不不允许特殊字符:' " & | @ % * ( ) - : # ¥)
	}
	retOrderRealtime struct {
		CommonResponse
		Data struct {
			OrderId string `json:"order_id"`
			Ref     string `json:"ref"`
			Pay     string `json:"pay"`
		} `json:"data"`
	}
)

func (y *Yunzhanghu) OrderRealTime(ctx context.Context) error {
	var (
		apiName = "银行卡实时下单"
		req     = &reqOrderRealtime{}
		ret     = new(retOrderRealtime)
	)
	responseBytes, err := y.postJSON(paymentOrderRealtimeURI, apiName, req)
	if err != nil {
		return err
	}
	if err = y.decodeWithError(responseBytes, ret, apiName); err != nil {
		return err
	}
	return nil
}
