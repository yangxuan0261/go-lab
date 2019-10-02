// package test_go

package main

import (
	"fmt"
	"time"
)

func main() {
	// test_001()
	test_002()
}

func child() {
	for {
		fmt.Print("child ")
		time.Sleep(1000000000)
	}
}

func parent() {
	for {
		fmt.Print("parent ")
		time.Sleep(1000000000)
	}
}

func test_001() {
	go child()
	parent() // 如果这个方法也加个 go 开个 goroutine 执行, 那么 test_001方法 将瞬间结束
}

// 通过信号通知 程序结束
func test_002() {
	c1 := make(chan string)

	go func(flag chan string) {
		for index := 0; index < 3; index++ {
			fmt.Printf("--- counting %d\n", index)
			time.Sleep(time.Second * 1)
		}
		flag <- "hello"
	}(c1)

	select {
	case msg := <-c1:
		fmt.Printf("收到通知 %s, 1s后结束\n", msg)
		time.Sleep(time.Second * 1)
		break
	}
	fmt.Println("程序结束 666")
}

// //go 关键字放在方法调用前新建一个 goroutine 并让他执行方法体
// go GetThingDone(param1, param2);

// //上例的变种，新建一个匿名方法并执行
// go func(param1, param2) {
// }(val1, val2)
