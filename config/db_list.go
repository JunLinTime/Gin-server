package config

type DB struct {
	Disable   bool   `mapstructure:"disable" json:"disable" yaml:"disable"`
	Type      string `mapstructure:"type" json:"type" yaml:"type"`
	AliasName string `mapstructure:"alias-name" json:"alias-name" yaml:"alias-name"`
	Path      string `mapstructure:"path" json:"path" yaml:"path"`             // 服务器地址:端口
	Port      string `mapstructure:"port" json:"port" yaml:"port"`             //:端口
	Config    string `mapstructure:"config" json:"config" yaml:"config"`       // 高级配置
	Dbname    string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`     // 数据库名
	Username  string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
	Password  string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
}

func (m *DB) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
