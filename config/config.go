package config

import "github.com/go-ini/ini"

var (
	RedisConfig RedisConfigInfo
)

// RedisConfigInfo redis配置
type RedisConfigInfo struct {
	Host     string
	User     string
	Password string
	Database int
	PoolSize int
}

// 初始化配置
func InitConfig(file string) (*ini.File, error) {
	cfg, err := ini.Load(file)
	if err != nil {
		return nil, nil
	}

	section := cfg.Section("redis")
	RedisConfig.Host = section.Key("REDIS_HOST").String()
	return cfg, err
}
