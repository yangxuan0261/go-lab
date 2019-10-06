package main

import (
	"fmt"
	"time"
)

// 除 default 外，如果只有一个 case 语句评估通过，那么就执行这个case里的语句；
// 除 default 外，如果有多个 case 语句评估通过，那么通过伪随机的方式随机选一个；
// 如果 default 外的 case 语句都没有通过评估，那么执行 default 里的语句；
// 如果没有 default，那么 代码块会被阻塞，指导有一个 case 通过评估；否则一直阻塞

func main() {
	// test_chan01()
	// test_chan012()
	test_chan013_readwrite()
	// test_chan02()
	// test_chan03()
	// test_chan04()
}

func test_chan01() {
	show := func(c chan int) {
		for {
			data := <-c // 阻塞 c, 直至有东西发送到 c
			fmt.Println("receive:", data)
		}
	}

	c := make(chan int)
	go show(c)
	for {
		num := 6
		fmt.Println("send:", num)
		c <- num // 将 num 发送到 c
		time.Sleep(time.Second * 3)
	}
}

// 使用 for range 阻塞 chan, 效果等同 test_chan01
func test_chan012() {
	show := func(c chan int) {
		for b := range c {
			fmt.Println("receive:", b)
		}
	}

	c := make(chan int)
	go show(c)
	for {
		num := 6
		fmt.Println("send:", num)
		c <- num
		time.Sleep(time.Second * 3)
	}
}

// chan 的只读,只写,读写
func test_chan013_readwrite() {
	fnRW := func(c chan int) { // c可以读写
		c <- 6
		val := <-c
		fmt.Println("val:", val)
	}

	fnR := func(c <-chan int) { // c只读
		// c <- 6 // 报错: send to receive-only type <-chan int
		val := <-c
		fmt.Println("val:", val)
	}

	fnW := func(c chan<- int) { // c只写
		// <-c // 报错: receive from send-only type chan<- int
		c <- 6
	}

	ci := make(chan int)
	go fnRW(ci)
	go fnR(ci)
	go fnW(ci)

	select {}
}

// 多 channel 访问
func test_chan02() {
	fibonacci := func(c, quit chan int) {
		x, y := 1, 1

		cnt := 0
		for {
			cnt++
			fmt.Printf("--- for cnt=%d\n ", cnt)
			select {
			case c <- x:
				x, y = y, x+y
				fmt.Printf("x=%d, y=%d\n", x, y)
			case <-quit:
				println("--- 收到 结束信号 111, 1s后停止")
				return
			}
		}
	}

	show := func(c, quit chan int) {
		for i := 0; i < 10; i++ {
			fmt.Println("c", <-c)
			time.Sleep(time.Second * 1)
		}

		fmt.Println("--- 3s后发送 结束信号")
		time.Sleep(time.Second * 3)
		quit <- 0
	}

	data := make(chan int)
	leave := make(chan int)

	go fibonacci(data, leave)
	go show(data, leave)

	// for { // 阻塞, 不然 test_chan02 方法将瞬间结束
	// 	time.Sleep(1000000000)
	// }
	select {
	case <-leave: // 收到
		println("--- 收到 结束信号 222, 1s后停止")
		time.Sleep(time.Second * 1)
		return
	}
}

func test_chan03() {
	c1 := make(chan string) // 声明 信号
	c2 := make(chan string)

	go func() { // go 关键字, 新建一个协程跑这个方法
		time.Sleep(time.Second * 1)
		c1 <- "one" // 往 c1 信号中丢数据, 也就是通知 c1 阻塞的地方可以继续跑了
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1: // 阻塞, 等待 c1 信号通知, 如果收到通知, 这跑这个case, 并把数据丢该 msg1
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	fmt.Println("程序结束666")
	/*
		msg1 := <-c1 表示 阻塞, 等待c1信号通知, 收到通知后把数据 赋值给 msg1. 如果不需要信号中的数据, 可以可以这样写 <-c1

		c1 <- "one" 表示 通知 c1 信号阻塞的地方可以继续运行了, 并往里面丢了一个数据 "one"
	*/
}

//-------------
// 初始化多个 chan, 等待全部返回
func test_chan04() {
	num := 3
	c1 := make(chan int)

	go func() {
		time.Sleep(time.Second * 1)
		for index := 0; index < num; index++ {
			fmt.Println("---index:", index)
			c1 <- index
		}
	}()

	fmt.Println("开始阻塞等待")
	for index := 0; index < num; index++ {
		v := <-c1
		fmt.Println("---v:", v)
	}

	fmt.Println("程序结束 666")
}

/*
num := 3
c1 := make(chan int, num) // 缓冲区可以存储3个int类型的整数, 一次性将3个整数存入channel，在读取的时候，也是一次性读取.
							如果要求必须全部存入后才能读取的话, 必须指定缓冲区长度
开始阻塞等待
---index: 0
---index: 1
---index: 2
---v: 0
---v: 1
---v: 2
程序结束 666


c1 := make(chan int) // 缓冲区默认为1个, 存入和读取也就是混乱的.
开始阻塞等待
---index: 0
---index: 1
---v: 0
---v: 1
---index: 2
---v: 2
程序结束 666
*/
