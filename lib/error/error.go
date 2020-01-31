package syserr

import (
	"fmt"
	"log"
	"runtime/debug"
)

var ErrFn func(msg string)

func Recover() {
	if err := recover(); err != nil {
		msgerr := fmt.Sprintf("%v\nmystack:%s\n", err, string(debug.Stack()))
		if ErrFn == nil {
			ErrFn = func(s string) {
				log.Printf("msgerr:%s", s)
			}
		}
		ErrFn(msgerr)
	}
}
