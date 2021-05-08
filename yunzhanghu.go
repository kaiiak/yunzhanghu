package yunzhanghu

import (
	"math/rand"
	"time"
)

const (
	defaultApiAddr = "https://api-jiesuan.yunzhanghu.com"
)

var (
	random *rand.Rand
)

type (
	Yunzhanghu struct {
		DesKey  string `json:"des_key" yaml:"des_key" mapstructure:"des_key"`
		Appkey  string `json:"app_key" yaml:"app_key" mapstructure:"app_key"`
		Dealer  string `json:"dealer" yaml:"dealer" mapstructure:"dealer"`
		Broker  string `json:"broker" yaml:"broker" mapstructure:"broker"`
		ApiAddr string `json:"api_addr" yaml:"api_addr" mapstructure:"api_addr"`
	}
)

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func New(desKey, appkey, dealer, broker, apiAddr string) *Yunzhanghu {
	if apiAddr == "" {
		apiAddr = defaultApiAddr
	}
	return &Yunzhanghu{desKey, appkey, dealer, broker, apiAddr}
}
