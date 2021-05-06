package yunzhanghu

import "context"

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
		CommonResponse
		Data Order `json:"data"`
	}
)

func (y *Yunzhanghu) QueryRealtimeOrder(ctx context.Context, orderId, channel, dataType string) (*Order, error) {
	var (
		apiName = "查询⼀个订单"
		resp    = new(retQueryRealtimeOrder)
	)
	respnseBytes, err := y.getJson(queryRealtimeOrderURI, apiName, resp)
	if err != nil {
		return nil, err
	}
	if err = y.decodeWithError(respnseBytes, resp, apiName); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
