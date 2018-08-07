package test_sync_waitgroup

// package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		// 计数加 1
		wg.Add(1) // 注意，wg.Add() 方法一定要在 goroutine 开始前执行
		go func(i int) {
			// 计数减 1
			defer wg.Done()
			time.Sleep(time.Second * time.Duration(i))
			fmt.Printf("goroutine%d 结束\n", i)
		}(i)
	}

	// 等待执行结束
	wg.Wait()
	fmt.Println("所有 goroutine 执行结束")
}
