package yunzhanghu

import "context"

const (
	authenticationVerifyIdURI = "/authentication/verify-id"
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
	responseBytes, err := y.postJSON(ctx, authenticationVerifyIdURI, apiName, input)
	if err != nil {
		return err
	}
	if err = y.decodeWithError(responseBytes, output, apiName); err != nil {
		return err
	}
	return nil
}

const (
	authenticationVerifyBankcardThreeFactorURI = "/authentication/verify-bankcard-three-factor"
	authenticationVerifyBankcardFourFactorURI  = "/authentication/verify-bankcard-four-factor"
)

type (
	reqVerifyBankcardFourFactor struct {
		CardNo   string `json:"card_no"`
		IdCard   string `json:"id_card"`
		RealName string `json:"real_name"`
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
	bs, err := y.postJSON(ctx, authenticationVerifyBankcardThreeFactorURI, apiName, req)
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
	bs, err := y.postJSON(ctx, authenticationVerifyBankcardFourFactorURI, apiName, req)
	if err != nil {
		return err
	}
	if err = y.decodeWithError(bs, ret, apiName); err != nil {
		return err
	}
	return nil
}
