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
// - https://www.tizi365.com/archives/306.html

func fn1(tp string, msg []byte) {
	fmt.Printf("--- fn1, tp:%s, msg:%s\n", tp, string(msg))
}

func fn2(tp string, msg []byte) {
	fmt.Printf("--- fn2, tp:%s, msg:%s\n", tp, string(msg))
}

func fn3(tp string, msg []byte) {
	fmt.Printf("--- fn3, tp:%s, msg:%s\n", tp, string(msg))
}

func Test_AnotherConn(t *testing.T) {
	addr := "127.0.0.1:7379"

	suber, serr := NewSubscriber(addr)
	if serr != nil {
		panic(serr)
	}
	suber.Sub("hello", fn1)
	suber.Sub("hello", fn2)
	suber.Sub("world", fn3)

	c, err := redis.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	res, err := c.Do("PUBLISH", "hello", "大王叫我来巡山 111")
	if err == nil {
		fmt.Println("res 111:", res)
	}

	time.Sleep(time.Second * 1)
	println()
	suber.Unsub("hello", fn1)

	res, err = c.Do("PUBLISH", "hello", "大王叫我来巡山 222")
	if err == nil {
		fmt.Println("res 222:", res)
	}

	time.Sleep(time.Second * 1)
	println("")
	suber.Unsub("hello", fn2) // 测试所有回调取消, 取消订阅 topic
	res, err = c.Do("PUBLISH", "hello", "大王叫我来巡山 333")
	if err == nil {
		fmt.Println("res 333:", res) // res 333: 0, 因为 s.client.Unsubscribe 了, 所有没有被消费, 会返回 0
	}

	time.Sleep(time.Second * 1)
	println("")
	println("--- close sub")
	suber.Close()

	common.WaitSignal()
	fmt.Println("--- exit Test_AnotherConn")
}

//
//func Test_SameConn(t *testing.T) {
//	addr := "127.0.0.1:7379"
//
//	suber := new(Subscriber)
//	suber.Connect(addr)
//	suber.Sub("hello", fn1)
//	suber.Sub("hello", fn2)
//	suber.Sub("world", fn3)
//
//	res, err := suber.Pub("hello", "大王叫我来巡山 111")
//	if err == nil {
//		fmt.Println("res:", res)
//	}
//
//	time.Sleep(time.Second * 1)
//	println()
//	suber.Unsub("hello", fn1)
//
//	res, err = suber.Pub("hello", "大王叫我来巡山 111")
//	if err == nil {
//		fmt.Println("res:", res)
//	}
//
//	time.Sleep(time.Second * 1)
//	println("")
//	println("--- close sub")
//	suber.Close()
//
//	common.WaitSignal()
//	fmt.Println("--- exit Test_SameConn")
//}
