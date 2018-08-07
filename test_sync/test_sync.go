package test_sync

// package main

// 参考: https://deepzz.com/post/golang-sync-package-usage.html

import (
	"fmt"
	"sync"
)

/*
一个互斥锁只能同时被一个 goroutine 锁定，其它 goroutine 将阻塞直到互斥锁被解锁（重新争抢对互斥锁的锁定）
*/

func main() {
	test001()
}

// ---------------
// 互斥锁
func test001() {
	quit := make(chan bool)

	var a = 0

	// 启动 100 个协程，需要足够大
	var mylock sync.Mutex
	for i := 0; i < 100; i++ {
		go func(idx int) {
			mylock.Lock() // 100 个协程使用同一个互斥锁mylock, 确保 a+1时, 只能有一个协程能操作. 和c++多线程的互斥锁是一个意思
			defer mylock.Unlock()
			a += 1
			fmt.Printf("goroutine %d, a=%d\n", idx, a)
			if a == 100 {
				quit <- true
			}
		}(i)
	}

	// 确保所有协程执行完
	select {
	case <-quit:
		fmt.Println("--- 程序结束666")
		break
	}
}
