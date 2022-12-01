package httpclient

import "fmt"

type StatusCode string

var (
	responseCodes = map[StatusCode]string{
		"0000": "成功",
		"1001": "签名已过期",
		"1002": "请求参数格式不正确",
		"1003": "签名错误",
		"1004": "加密错误",
		"1005": "商户未设置3deskey或没有设置appkey",
		"2001": "上传数据有误",
		"2002": "已上传过该笔流水",
		"2003": "实名认证失败",
		"2006": "银⾏卡号错误",
		"2011": "订单不存在",
		"2018": "该笔订单不存在",
		"2016": "错误的打款金额",
		"2024": "订单⾦额小于0",
		"2027": "账户余额不足",
		"2038": "该商户不属于该经纪公司",
		"3000": "银行系统连接失败，请联系我们修复",
		"5000": "不存在的账单",
	}
)

func (s StatusCode) Message() string {
	if msg, ok := responseCodes[s]; ok {
		return msg
	}
	return fmt.Sprintf("%v", s)
}
