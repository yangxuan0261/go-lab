package profiler_test

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
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
	go func() {
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
		}
	}()

	tick := time.Tick(time.Second * 1)
	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
	}
}
