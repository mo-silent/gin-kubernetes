package conf

type Mysql struct {
	Path     string `mapstructure:"path" json:"path" yaml:"path"`             // 数据库路径
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             // 数据库端口
	Config   string `mapstructure:"config" json:"config" yaml:"config"`       // 数据库连接参数
	DBname   string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`    // 数据库名称
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.DBname + "?" + m.Config
}
