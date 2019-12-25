package test_tail

import (
	"GoLab/common"
	syslog "GoLab/common/log"
	"fmt"
	"github.com/hpcloud/tail"
	"testing"
	"time"
)

var logfile = "./temp_access.json"

func writeLog() {
	defer fmt.Println("--- defer writeLog")
	syslog.Init(logfile, "./temp_error.json", true)

	t := time.NewTicker(time.Second * 1)
	defer t.Stop()
	cnt := 1

	for {
		select {
		case <-t.C:
			syslog.Access.Sugar().Infof("hello-%d", cnt)
			cnt++
		}
	}
}

func listenLog() {
	defer fmt.Println("--- defer listenLog")

	tl, err := tail.TailFile(logfile, tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}

	fmt.Printf("--- start listenLog\n")

	for line := range tl.Lines {
		fmt.Printf("--- line:%s\n", line.Text)
	}
}

func Test_tail(t *testing.T) {
	go writeLog()
	go listenLog()

	common.WaitSignal()
	fmt.Println("--- exit Test_tail")
}
