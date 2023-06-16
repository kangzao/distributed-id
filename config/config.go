package config

// ConsulConfig 默认情况下，mapstructure使用结构体中字段的名称做这个映射，例如我们的结构体有一个Name字段，
// mapstructure解码时会在map[string]interface{}中查找键名name。注意，这里的name是大小写不敏感的！
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name"` //通过name 进行服务发现
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
}
