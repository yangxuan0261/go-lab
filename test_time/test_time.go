package main

import (
	"log"
	"time"
)

func test_ticker() {
	t := time.NewTicker(time.Second)
	defer t.Stop()
	cnt := 1
	for _ = range t.C { // chan 阻塞
		log.Println("--- cnt:", cnt)
		if cnt == 3 {
			log.Println("--- end ticker")
			return // 正确结束 ticker 的姿势, return 后会调用 t.Stop()
		}
		cnt += 1
	}
}

func main() {
	test_ticker()
	log.Println("--- done")
}
