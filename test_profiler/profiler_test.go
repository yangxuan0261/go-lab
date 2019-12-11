package profiler_test

import (
	"GoLab/common/pprof"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"testing"
	"time"
)

func Test_openPprof(t *testing.T) {
	// 开启pprof，监听请求
	ip := "0.0.0.0:6060"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}

	// 然后打开网页: http://localhost:6060/debug/pprof/
}

func Test_leak01(t *testing.T) {
	// 开启pprof
	pprof.StartPprof("0.0.0.0:6060")

	var wg sync.WaitGroup
	wg.Add(1)
	tick := time.Tick(time.Second * 1)
	var buf []byte
	cnt := 0
	for range tick {
		buf = append(buf, make([]byte, 1024*1024*8)...)

		cnt++
		if cnt == 7 {
			break
		}
	}

	fmt.Println("--- tick over, wait")
	wg.Wait()
}

func Test_leak02(t *testing.T) {
	// 开启pprof
	pprof.StartPprof("0.0.0.0:6060")

	var wg sync.WaitGroup
	wg.Add(1)

	outCh := make(chan int)
	// 死代码，永不读取
	go func() {
		if false {
			<-outCh
		}
		select {}
	}()

	// 每s起100个goroutine，goroutine会阻塞，不释放内存
	tick := time.Tick(time.Second / 100)
	i := 0
	for range tick {
		i++
		//fmt.Println(i)
		alloc1(outCh)

		if i == 500 { // 3s
			break
		}
	}

	fmt.Println("--- tick over, wait")
	wg.Wait()
}

func alloc1(outCh chan<- int) {
	go alloc2(outCh)
}

func alloc2(outCh chan<- int) {
	func() {
		defer fmt.Println("alloc-fm exit")
		// 分配内存，假用一下
		buf := make([]byte, 1024*1024*10)
		_ = len(buf)
		//fmt.Println("alloc done")

		outCh <- 0 // 53行
	}()
}
