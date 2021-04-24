package yunzhanghu

import (
	"context"
)

const (
	paymentOrderAlipayURI = "https://api-jiesuan.yunzhanghu.com//api/payment/v1/order-alipay"
)

type (
	reqOrderAlipay struct {
		OrderId   string `json:"order_id"`             // 商户订单号，由商户保持唯⼀一性(必填)
		DealerId  string `json:"dealer_id"`            // 商户代码(必填)
		BrokerId  string `json:"broker_id"`            // 经纪公司(必填)
		RealName  string `json:"real_name"`            // 姓名(必填)
		IdCard    string `json:"id_card"`              // 身份证(必填)
		CardNo    string `json:"card_no"`              // 收款⼈人⽀支付宝账户(必填)
		Pay       string `json:"pay"`                  // 打款⾦金金额(单位为元, 必填)
		Notes     string `json:"notes,omitempty"`      // 备注信息(选填)
		PayRemark string `json:"pay_remark,omitempty"` // 打款备注(选填，最⼤大20个字符，⼀一个汉字占2个字符，不不允许特殊字符:' " & | @ % * ( ) - : // ¥)
	}
	retOrderAlipay struct {
		CommonResponse
		Data struct {
			OrderId string `json:"order_id"`
			Ref     string `json:"ref"`
			Pay     string `json:"pay"`
		} `json:"data"`
	}
)

func (y *Yunzhanghu) OrderAlipay(ctx context.Context, orderId, realName, idCard, cardNo, pay, notes, payRemark string) error {
	var (
		apiName = "支付宝实时下单"
		req     = &reqOrderAlipay{
			DealerId:  y.Dealer,
			BrokerId:  y.Broker,
			OrderId:   orderId,
			RealName:  realName,
			IdCard:    idCard,
			CardNo:    cardNo,
			Pay:       pay,
			Notes:     notes,
			PayRemark: payRemark,
		}
		ret = new(retOrderAlipay)
	)
	responseBytes, err := y.postJSON(paymentOrderAlipayURI, apiName, req)
	if err != nil {
		return err
	}
	if err = y.decodeWithError(responseBytes, ret, apiName); err != nil {
		return err
	}
	return nil
}
