package redis

import (
	"distributed-id/config"
	"github.com/gomodule/redigo/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	Pool *redis.Pool
)

func init() {
	config.InitConfig("../config.ini")
	redisHost := config.RedisConfig.Host
	if redisHost == "" {
		redisHost = "127.0.0.1:6379"
	}
	Pool = newPool(redisHost)
	//cleanupHook()
}

func newPool(server string) *redis.Pool {

	return &redis.Pool{ //实例化一个连接池

		MaxIdle:     3,                 //最初的连接数量
		IdleTimeout: 300 * time.Second, //连接关闭时间 300秒 （300秒不使用自动关闭）

		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}
