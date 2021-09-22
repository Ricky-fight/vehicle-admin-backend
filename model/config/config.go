package config

import "time"

type Config struct {
	Mysql  Mysql  `mapstructure:"mysql"`  // mysql配置
	System System `mapstructure:"system"` // 服务器配置
	Jwt    Jwt    `mapstructure:"jwt"`    // jwt配置
}

type Mysql struct {
	User     string `mapstructure:"user"`     // 用户名
	Password string `mapstructure:"password"` // 密码
	Host     string `mapstructure:"host"`     // 地址
	Port     string `mapstructure:"port"`     // 端口
	Dbname   string `mapstructure:"dbname"`   // 库名
}

type System struct {
	Port string `mapstructure:"port"` // 服务器运行端口
}

type Jwt struct {
	Domain              string        `mapstructure:"domain"`              // 域名
	TokenExpireDuration time.Duration `mapstructure:"tokenExpireDuration"` // token有效时间
	Secret              string        `mapstructure:"secret"`              // 服务器私有密钥
}
