package test_time

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func ticker() {
	t := time.NewTicker(time.Second * 1)
	defer t.Stop()
	cnt := 1
	for _ = range t.C { // chan 阻塞
		log.Println("--- cnt:", cnt)
		if cnt == 5 {
			log.Println("--- end ticker")
			return // 正确结束 ticker 的姿势, return 后会调用 t.Stop()
		}
		cnt += 1
	}
}

func Test_main(t *testing.T) {
	log.Println("--- start")
	ticker()
	log.Println("--- done")
}

func Test_timeout(t *testing.T) {
	fmt.Println("--- begin")
	in := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		select {
		case v := <-in:
			fmt.Println(v)
		case <-time.After(time.Second * 3):
			return // 超时
		}
	}()
	wg.Wait()
	fmt.Println("--- over")
}
