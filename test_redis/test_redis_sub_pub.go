package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 参考: https://blog.csdn.net/wangshubo1989/article/details/75050024

func main() {

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	// TODO: 暂时不知道怎么通过redis做管道
	c.Do("PUBLISH", "hello", "test send2")
}
