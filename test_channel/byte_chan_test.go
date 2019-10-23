package test_chan

import (
	"fmt"
	"testing"
	"time"
)

func Test_byteChan(t *testing.T) {
	writeChan := make(chan []byte, 100)

	go func() {
		for b := range writeChan { // 切片可以这样使用, 也是可以阻塞等带数据
			fmt.Println("--- rec msg:", string(b))
		}
	}()

	time.Sleep(time.Second * 1)
	fmt.Println("--- 开始发送")
	go func() {
		for cnt := 1; ; cnt++ {
			writeChan <- []byte(fmt.Sprintf("hello cnt:%d", cnt))
			time.Sleep(time.Second * 1)
		}
	}()

	select {}
}

/*
--- 开始发送
--- rec msg: hello cnt:1
--- rec msg: hello cnt:2
--- rec msg: hello cnt:3
*/
