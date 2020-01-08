package cat

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

func Run(speed int) {
	log.Printf("--- Run, speed:%+v\n", speed)
}

func Test_Run(t *testing.T) {
	Run(111)
}

func Division(a int, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("b cant be zero")
	} else {
		return a / b, nil
	}
}

func Test_Division(t *testing.T) {
	if _, e := Division(6, 0); e != nil { //try a unit test on function
		t.Error("--- Division did not work as expected.") // 带行号的错误日志, 如果不是如预期的那么就报错,
	}
	t.Log("--- bbb") // 带行号的日志, 记录一些你期望记录的信息

	fmt.Println("--- aaa") // 无行号日志
}
