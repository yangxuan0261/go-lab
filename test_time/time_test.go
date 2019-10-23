package test_time

import (
	"log"
	"testing"
	"time"
)

func ticker() {
	t := time.NewTicker(time.Second * 3)
	defer t.Stop()
	cnt := 1
	for _ = range t.C { // chan 阻塞
		log.Println("--- cnt:", cnt)
		if cnt == 2 {
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
