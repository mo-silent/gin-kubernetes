package conf

type System struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`    // 环境值
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值
}
