package yunzhanghu

import (
	"context"
)

const (
	authenticationVerifyBankcardThreeFactorURI = "/authentication/verify-bankcard-three-factor"
	authenticationVerifyBankcardFourFactorURI  = "/authentication/verify-bankcard-four-factor"
)

type (
	reqVerifyBankcardFourFactor struct {
		CardNo   string `json:"card_no,omitempty"`
		IdCard   string `json:"id_card,omitempty"`
		RealName string `json:"real_name,omitempty"`
		Mobile   string `json:"mobile,omitempty"`
	}
	retVerifyBankcardFourFactor struct {
		CommonResponse
	}
)

func (y *Yunzhanghu) VerifyBankcardThreeFactor(ctx context.Context, cardNo, idCard, realName string) error {
	var (
		req = &reqVerifyBankcardFourFactor{
			CardNo:   cardNo,
			IdCard:   idCard,
			RealName: realName,
		}
		ret     = new(retVerifyBankcardFourFactor)
		apiName = "银行卡三要素认证"
	)
	bs, err := y.postJSON(authenticationVerifyBankcardThreeFactorURI, apiName, req)
	if err != nil {
		return err
	}
	if err = y.decodeWithError(bs, ret, apiName); err != nil {
		return err
	}
	return nil
}

func (y *Yunzhanghu) VerifyBankcardFourFactor(ctx context.Context, cardNo, idCard, realName, mobile string) error {
	var (
		req = &reqVerifyBankcardFourFactor{
			CardNo:   cardNo,
			IdCard:   idCard,
			RealName: realName,
			Mobile:   mobile,
		}
		ret     = new(retVerifyBankcardFourFactor)
		apiName = "银行卡四要素认证"
	)
	bs, err := y.postJSON(authenticationVerifyBankcardFourFactorURI, apiName, req)
	if err != nil {
		return err
	}
	if err = y.decodeWithError(bs, ret, apiName); err != nil {
		return err
	}
	return nil
}
