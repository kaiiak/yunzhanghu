package settlement

import (
	"context"

	"github.com/kaiiak/yunzhanghu/core"
)

const (
	authenticationVerifyIdURI = "/authentication/verify-id"
)

type (
	reqVerifyid struct {
		IdCard   string `json:"id_card"`
		RealName string `json:"real_name"`
	}
	retVerifyid struct {
		core.CommonResponse
	}
)

func (s *Settlement) VerifyId(ctx context.Context, realName, idCard string) error {
	var (
		input = &reqVerifyid{
			IdCard:   idCard,
			RealName: realName,
		}
		output  = new(retVerifyid)
		apiName = "身份证实名验证"
	)
	responseBytes, err := core.PostJSON(s.newContext(ctx, core.NewSHA256Sign(s.Appkey)), authenticationVerifyIdURI, input)
	if err != nil {
		return err
	}
	if err = core.DecodeWithError(responseBytes, output, apiName); err != nil {
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
		core.CommonResponse
	}
)

func (s *Settlement) VerifyBankcardThreeFactor(ctx context.Context, cardNo, idCard, realName string) error {
	var (
		req = &reqVerifyBankcardFourFactor{
			CardNo:   cardNo,
			IdCard:   idCard,
			RealName: realName,
		}
		ret     = new(retVerifyBankcardFourFactor)
		apiName = "银行卡三要素认证"
	)
	bs, err := core.PostJSON(s.newContext(ctx, core.NewSHA256Sign(s.Appkey)), authenticationVerifyBankcardThreeFactorURI, req)
	if err != nil {
		return err
	}
	if err = core.DecodeWithError(bs, ret, apiName); err != nil {
		return err
	}
	return nil
}

func (s *Settlement) VerifyBankcardFourFactor(ctx context.Context, cardNo, idCard, realName, mobile string) error {
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
	bs, err := core.PostJSON(s.newContext(ctx, core.NewSHA256Sign(s.Appkey)), authenticationVerifyBankcardFourFactorURI, req)
	if err != nil {
		return err
	}
	if err = core.DecodeWithError(bs, ret, apiName); err != nil {
		return err
	}
	return nil
}
