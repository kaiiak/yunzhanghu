package yunzhanghu

import (
	"context"
)

const (
	authenticationVerifyIdURI = "https://api-jiesuan.yunzhanghu.com/authentication/verify-id"
)

type (
	reqVerifyid struct {
		IdCard   string `json:"id_card"`
		RealName string `json:"real_name"`
	}
	retVerifyid struct {
		CommonResponse
	}
)

func (y *Yunzhanghu) VerifyId(ctx context.Context, realName, idCard string) error {
	var (
		input = &reqVerifyid{
			IdCard:   idCard,
			RealName: realName,
		}
		output  = new(retVerifyid)
		apiName = "身份证实名验证"
	)
	responseBytes, err := y.postJSON(authenticationVerifyIdURI, apiName, input)
	if err != nil {
		return err
	}
	if err = y.decodeWithError(responseBytes, output, apiName); err != nil {
		return err
	}
	return nil
}
