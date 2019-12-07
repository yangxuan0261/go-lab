package syserr

import (
	"fmt"
	"log"
	"runtime/debug"
)

var ErrFn func(msg string)
var tempFn func(msg string)

func ShowStack(b bool) {
	if ErrFn == nil {
		return
	}

	if b {
		if tempFn != nil {
			return
		}
		tempFn = ErrFn // 保存 原函数指针
		ErrFn = func(msg string) {
			debug.PrintStack()
			tempFn(msg)
		}
	} else {
		ErrFn = tempFn
		tempFn = nil
	}
}

func Recover() {
	if err := recover(); err != nil {
		msgerr := fmt.Sprintf("%v", err)
		if ErrFn != nil {
			ErrFn(msgerr)
		} else {
			ErrFn = func(s string) {
				log.Printf("msgerr:%s", s)
			}
			ErrFn(msgerr)
		}
	}
}
