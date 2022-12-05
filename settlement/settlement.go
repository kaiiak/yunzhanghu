package settlement

import (
	"context"

	"github.com/kaiiak/yunzhanghu/core"
)

const settlementApiAddr = "https://api-jiesuan.yunzhanghu.com"

type Settlement struct {
	*core.Config
}

func NewSettlement(cnf *core.Config) *Settlement {
	return &Settlement{cnf}
}

func (s *Settlement) newContext(ctx context.Context, sign core.Signer) *core.Context {
	return core.NewContext(ctx, s.Config, settlementApiAddr, sign)
}
