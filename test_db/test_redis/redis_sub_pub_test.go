package test_redis

import (
	"GoLab/common"
	"fmt"
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
)

// 参考:
// - https://blog.csdn.net/qq_17308321/article/details/89417493

func fn1(tp string, msg []byte) {
	fmt.Printf("--- fn1, tp:%s, msg:%s\n", tp, string(msg))
}

func fn2(tp string, msg []byte) {
	fmt.Printf("--- fn2, tp:%s, msg:%s\n", tp, string(msg))
}

func fn3(tp string, msg []byte) {
	fmt.Printf("--- fn3, tp:%s, msg:%s\n", tp, string(msg))
}

func Test_001(t *testing.T) {
	addr := "127.0.0.1:7379"

	suber := new(Subscriber)
	suber.Connect(addr)
	suber.Sub("hello", fn1)
	suber.Sub("hello", fn2)
	suber.Sub("world", fn3)

	c, err := redis.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	res, err := c.Do("PUBLISH", "hello", "大王叫我来巡山")
	if err == nil {
		fmt.Println("res:", res)
	}

	time.Sleep(time.Second * 1)
	println()
	suber.Unsub("hello", fn1)

	res, err = c.Do("PUBLISH", "hello", "大王叫我来巡山")
	if err == nil {
		fmt.Println("res:", res)
	}

	time.Sleep(time.Second * 1)
	println("")
	println("--- close sub")
	suber.Close()

	common.WaitSignal()
	fmt.Println("--- exit test")
}

//func Test_subscribe(t *testing.T) {
//	client := redis.NewClient(&redis.Options{
//		Addr:     "127.0.0.1:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//	_ = client
//}
