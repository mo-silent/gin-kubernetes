package conf

type Server struct {
	System     System `mapstructure:"system" json:"system" yaml:"system"`
	Kubeconfig string `mapstructure:"kubeconfig" json:"kubeconfig" yaml:"kubeconfig"`
	Mysql      Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
