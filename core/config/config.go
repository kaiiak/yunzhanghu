package config

// Config for 云账户
type Config struct {
	DesKey     string `json:"des_key" yaml:"des_key" mapstructure:"des_key"`
	Appkey     string `json:"app_key" yaml:"app_key" mapstructure:"app_key"`
	Dealer     string `json:"dealer" yaml:"dealer" mapstructure:"dealer"`
	Broker     string `json:"broker" yaml:"broker" mapstructure:"broker"`
	PrivateKey string `json:"private_key" yaml:"private_key" mapstructure:"private"`
}

func NewConfig() {

}
