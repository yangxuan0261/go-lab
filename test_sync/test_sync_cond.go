// package test_sync_cond

package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Cond 实现一个条件变量，即等待或宣布事件发生的 goroutines 的会合点，它会保存一个通知列表。基本思想是当某中状态达成，goroutine 将会等待（Wait）在那里，当某个时刻状态改变时通过通知的方式（Broadcast，Signal）的方式通知等待的 goroutine。这样，不满足条件的 goroutine 唤醒继续向下执行，满足条件的重新进入等待序列。
*/

var count int = 4

func main() {
	ch := make(chan struct{}, 5)

	// 新建 cond
	var l sync.Mutex
	cond := sync.NewCond(&l)

	for i := 0; i < 5; i++ {
		go func(i int) {
			// 争抢互斥锁的锁定
			cond.L.Lock()
			defer func() {
				cond.L.Unlock()
				ch <- struct{}{}
			}()

			// 条件是否达成
			for count > i {
				cond.Wait()
				fmt.Printf("收到一个通知 goroutine%d\n", i)
			}

			fmt.Printf("goroutine%d 执行结束\n", i)
		}(i)
	}

	// 确保所有 goroutine 启动完成
	time.Sleep(time.Second)

	// 锁定一下
	fmt.Println("broadcast...")
	cond.L.Lock()
	count -= 1
	cond.Broadcast()
	cond.L.Unlock()

	time.Sleep(time.Second)
	fmt.Println("signal...")
	cond.L.Lock()
	count -= 2
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(time.Second)
	fmt.Println("broadcast...")
	cond.L.Lock()
	count -= 1
	cond.Broadcast()
	cond.L.Unlock()

	for i := 0; i < 5; i++ {
		<-ch
	}
}
