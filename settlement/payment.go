package settlement

import (
	"context"
	"io"

	"github.com/kaiiak/yunzhanghu/core"
	"github.com/kaiiak/yunzhanghu/core/httpclient"
)

var ApiAddr = "https://api-jiesuan.yunzhanghu.com"

type Settlement struct {
	core.Core
}

func NewSettlement(ctx context.Context) {

}

const (
	paymentOrderAlipayURI = "/api/payment/v1/order-alipay"
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
		httpclient.CommonResponse
		Data struct {
			OrderId string `json:"order_id"`
			Ref     string `json:"ref"`
			Pay     string `json:"pay"`
		} `json:"data"`
	}
)

func (s *Settlement) OrderAlipay(ctx context.Context, orderId, realName, idCard, cardNo, pay, notes, payRemark string) error {
	var (
		apiName = "支付宝实时下单"
		req     = &reqOrderAlipay{
			DealerId:  s.Dealer,
			BrokerId:  s.Broker,
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
	responseBytes, err := s.PostJSON(ctx, paymentOrderAlipayURI, req)
	if err != nil {
		return err
	}
	if err = s.DecodeWithError(responseBytes, ret, apiName); err != nil {
		return err
	}
	return nil
}

const (
	paymentOrderRealtimeURI = "/api/payment/v1/order-realtime"
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
		httpclient.CommonResponse
		Data struct {
			OrderId string `json:"order_id"`
			Ref     string `json:"ref"`
			Pay     string `json:"pay"`
		} `json:"data"`
	}
)

func (s *Settlement) OrderRealTime(ctx context.Context) error {
	var (
		apiName = "银行卡实时下单"
		req     = &reqOrderRealtime{}
		ret     = new(retOrderRealtime)
	)
	responseBytes, err := s.PostJSON(ctx, paymentOrderRealtimeURI, req)
	if err != nil {
		return err
	}
	if err = s.DecodeWithError(responseBytes, ret, apiName); err != nil {
		return err
	}
	return nil
}

const (
	queryRealtimeOrderURI = "/api/payment/v1/query-realtime-order"
)

type (
	reqQueryRealtimeOrder struct {
		OrderId  string `json:"order_id"`  // 商户订单号
		Channel  string `json:"channel"`   // 银⾏行行卡，⽀支付宝，微信(不不填时为银⾏行行卡订单查询)(选填)
		DataType string `json:"data_type"` // 如果为encryption，则对返回的data进⾏行行加密(选填)
	}
	retQueryRealtimeOrder struct {
		httpclient.CommonResponse
		Data Order `json:"data"`
	}
)

func (s *Settlement) QueryRealtimeOrder(ctx context.Context, orderId, channel string, encrypted bool) (*Order, error) {
	var (
		apiName = "查询⼀个订单"
		resp    = new(retQueryRealtimeOrder)
		req     = &reqQueryRealtimeOrder{
			OrderId: orderId,
			Channel: channel,
		}
	)
	if encrypted {
		req.DataType = "encryption"
	}
	respnseBytes, err := s.GetJson(ctx, queryRealtimeOrderURI, req)
	if err != nil {
		return nil, err
	}
	if err = s.DecodeWithError(respnseBytes, resp, apiName); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

const (
	uploadIdCardImageURI = "/api/payment/v1/sign/idcard/image"
)

type (
	reqUploadIdCardImage struct {
		DealerId string `json:"dealer_id"` // 商户代码(必填)
		BrokerId string `json:"broker_id"` // 经纪公司(必填)
		RealName string `json:"real_name"` // ⽤用户姓名(必填)
		IdCard   string `json:"id_card"`   // ⽤用户身份证号(必填)
	}
	retUploadIdCardImage struct {
		httpclient.CommonResponse
		Data interface{} `json:"data"`
	}
)

func (s *Settlement) UploadIdCardImage(ctx context.Context, realName, idCard string, image, backgroud io.Reader) error {
	var (
		apiName = "身份证信息上传"
		req     = &reqUploadIdCardImage{
			DealerId: s.Dealer,
			BrokerId: s.Broker,
			RealName: realName,
			IdCard:   idCard,
		}
		resp = new(retUploadIdCardImage)
	)
	responsesBytes, err := s.PostForm(ctx, uploadIdCardImageURI, apiName, req, map[string]io.Reader{
		"id_card_image":           image,
		"id_card_image_backgroud": backgroud,
	})
	if err != nil {
		return err
	}
	if err = s.DecodeWithError(responsesBytes, resp, apiName); err != nil {
		return err
	}
	return nil
}

const (
	authAlipayURI = "/api/payment/v1/auth-alipay"
)

type (
	reqAuthAlipay struct {
		BrokerId string `json:"broker_id"`
	}
	retAuthAlipay struct {
		httpclient.CommonResponse
		Data struct {
			Info string `json:"info"`
		} `json:"data"`
	}
)

func (s *Settlement) AuthAlipay(ctx context.Context) (string, error) {
	var (
		apiName = "⽀付宝授权"
		req     = &reqAuthAlipay{BrokerId: s.Broker}
		resp    = new(retAuthAlipay)
	)
	responseBytes, err := s.GetJson(ctx, authAlipayURI, req)
	if err != nil {
		return "", err
	}
	if err = s.DecodeWithError(responseBytes, resp, apiName); err != nil {
		return "", err
	}
	return resp.Data.Info, nil
}
