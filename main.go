package main

import (
	redisUtil "distributed-id/redis"
	"fmt"
)

func main() {
	redisUtil.Set("test", []byte("777"))
	result, _ := redisUtil.Get("test")

	fmt.Println(string(result))

}
