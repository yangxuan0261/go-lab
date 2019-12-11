package pprof

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func StartPprof(addr string) {
	go func() {
		// 开启pprof，监听请求
		//ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(addr, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", addr)
		}
	}()

	// 然后打开网页: http://localhost:6060/debug/pprof/
}
