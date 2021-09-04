package yunzhanghu

type (
	Order struct {
		Pay                 string            `json:"pay"`
		AnchorId            string            `json:"anchor_id"`
		BrokerAmount        string            `json:"broker_amount"`
		BrokerId            string            `json:"broker_id"`
		CardNo              string            `json:"card_no"`
		DealerId            string            `json:"dealer_id"`
		IdCard              string            `json:"id_card"`
		OrderId             string            `json:"order_id"`
		PhoneNo             string            `json:"phone_no"`
		RealName            string            `json:"real_name"`
		Ref                 string            `json:"ref"`
		Notes               string            `json:"notes"`
		Status              OrderStatus       `json:"status"`                //  订单状态码，详⻅见:附录1订单状态码
		StatusDetail        OrderStatusDetail `json:"status_detail"`         //  订单详细状态码，详⻅见:附录2订单详细状态码
		StatusMessage       string            `json:"status_message"`        //  状态码说明，详⻅见:附录1订单状态码
		StatusDetailMessage string            `json:"status_detail_message"` //  状态详细状态码说明，详⻅见:附录2订单详细状态码
		SysAmount           string            `json:"sys_amount"`
		PayRemark           string            `json:"pay_remark"` //  打款备注(选填，最⼤大20个字符，⼀一个汉字占2个字符，不不允许特殊字符:' " & | @ % * ( ) - : # ¥)
		Tax                 string            `json:"tax"`
		TaxLevel            string            `json:"tax_level"`
		CreatedAt           Time              `json:"created_at"`    //  订单接收时间
		FinishedTime        Time              `json:"finished_time"` //  订单处理理时间
		EncryData           string            `json:"encry_data"`    //  当且仅当data_type为encryption时，返回且仅返回该加密数据字段
	}
)

func (o Order) GetOrderStatus() string {
	return o.Status.Message()
}

func (o Order) GetOrderStatusDetail() string {
	return o.StatusDetail.Message()
}
