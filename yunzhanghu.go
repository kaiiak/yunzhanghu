package yunzhanghu

type (
	Yunzhanghu struct {
		DesKey string `json:"des_key" yaml:"des_key" mapstructure:"des_key"`
		Appkey string `json:"app_key" yaml:"app_key" mapstructure:"app_key"`
		Dealer string `json:"dealer" yaml:"dealer" mapstructure:"dealer"`
		Broker string `json:"broker" yaml:"broker" mapstructure:"broker"`
		// ApiAddr string `json:"api_addr" yaml:"api_addr" mapstructure:"api_addr"`
	}
)

func New(desKey, appkey, dealer, broker string) *Yunzhanghu {
	return &Yunzhanghu{desKey, appkey, dealer, broker}
}
