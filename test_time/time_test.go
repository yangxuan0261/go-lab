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

func ticker2() {
	t := time.NewTicker(time.Second * 1)
	defer t.Stop()
	cnt := 1

	for {
		log.Println("--- cnt:", cnt)
		if cnt == 5 {
			log.Println("--- end ticker")
			return // 正确结束 ticker 的姿势, return 后会调用 t.Stop()
		}
		select {
		case <-t.C:
			fmt.Printf("Ticker running..., time:%v\n", time.Now().Unix())
			//case stop := <-ch: // 可以加入 context 取消定时器
			//	if stop {
			//		fmt.Println("Ticker Stop")
			//		return
			//	}
		}
		cnt += 1
	}
}

func Test_main(t *testing.T) {
	log.Println("--- start")
	//ticker()
	ticker2()
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

func Test_Unix(t *testing.T) {
	tm := time.Now()
	utime := tm.Unix()
	fmt.Printf("--- utime:%+v\n", utime) // --- utime:1575288231

	untime := tm.UnixNano()
	fmt.Printf("--- untime:%+v\n", untime) // --- untime:1575288231 652637100

	println()
	fmt.Printf("时间戳（秒）：\t\t\t%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（纳秒）：\t\t%v;\n", time.Now().UnixNano())
	fmt.Printf("时间戳（毫秒）1：\t\t%v;\n", time.Now().UnixNano()/1e6)
	fmt.Printf("时间戳（毫秒）2：\t\t%v;\n", time.Now().UnixNano()/int64(time.Millisecond))
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n", time.Now().UnixNano()/1e9)

	println()
	fmt.Printf("time.Second: \t\t%v\n", int64(time.Second))
	fmt.Printf("time.Millisecond: \t%v\n", int64(time.Millisecond))
	fmt.Printf("time.Microsecond: \t%v\n", int64(time.Microsecond))
	fmt.Printf("time.Nanosecond: \t%v\n", int64(time.Nanosecond))
}
