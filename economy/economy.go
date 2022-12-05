package economy

import (
	"context"

	"github.com/kaiiak/yunzhanghu/core"
)

const (
	economyApiAddr = "https://api-aic.yunzhanghu.com"
)

type Economy struct {
	*core.Config
}

func NewEconomy(cnf *core.Config) *Economy {
	return &Economy{cnf}
}

func (e *Economy) newContext(ctx context.Context, sign core.Signer) *core.Context {
	return core.NewContext(ctx, e.Config, economyApiAddr, sign)
}
