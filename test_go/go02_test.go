package test_go

import (
	"fmt"
	"testing"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func Test_02_001(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c // 这里会阻塞两次等待两个信号 c 都到齐了才往下走
	fmt.Println(x, y, x+y)
}

func Test_02_002(t *testing.T) {
	go func() {
		defer func() { // 这个是在 func 结束时调用, 而不是在 test_002 结束时调用
			fmt.Println("--- hello")
		}()

	}()

	time.Sleep(time.Second * 2)
	fmt.Println("--- world")
}
