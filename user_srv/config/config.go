package config

type Config struct {
	ProjectName string `mapstructure:"project_name" json:"project_name"`
	MD5Salt     string `mapstructure:"md5_salt" json:"md5_salt"`
	Uuid     string `mapstructure:"Uuid" json:"Uuid"`
	RunMod      string `mapstructure:"run_mod" json:"run_mod"`
	Mysql       Mysql  `mapstructure:"mysql" json:"mysql"`
	Logger      Logger `mapstructure:"logger" json:"logger"`
	Consul         Consul  `mapstructure:"consul" json:"consul"`
}

// Mysql 数据库配置
type Mysql struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     uint32 `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Database string `mapstructure:"database" json:"database"`
}

// Logger 日志配置
type Logger struct {
	Filename   string `mapstructure:"filename" json:"filename"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age"`
	Compress   bool   `mapstructure:"compress" json:"compress"`
}

type Consul struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     uint32 `mapstructure:"port" json:"port"`
	Tags []string `mapstructure:"tags" json:"tags"`
}
