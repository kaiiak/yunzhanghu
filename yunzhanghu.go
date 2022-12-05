package yunzhanghu

import (
	"github.com/kaiiak/yunzhanghu/core"
	"github.com/kaiiak/yunzhanghu/economy"
	"github.com/kaiiak/yunzhanghu/settlement"
)

type Yunzhanghu struct {
	*core.Config
}

func NewYunzhanghu(desKey string, appkey string, dealer string, broker string, privateKey string) *Yunzhanghu {
	return &Yunzhanghu{&core.Config{DesKey: desKey, Appkey: appkey, Dealer: dealer, Broker: broker, PrivateKey: privateKey}}
}

// 新经济个体工商户相关接口
func (y *Yunzhanghu) NewEconomy() *economy.Economy {
	return economy.NewEconomy(y.Config)
}

// 结算系统相关接口
func (y *Yunzhanghu) NewSettlement() *settlement.Settlement {
	return settlement.NewSettlement(y.Config)
}
