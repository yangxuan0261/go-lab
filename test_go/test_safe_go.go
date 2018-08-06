// package test_go

package main

import (
	"fmt"
	"runtime"
)

const LenStackBuf int32 = 4096

// one Go per goroutine (goroutine not safe)
type Go struct {
	ChanCb    chan func()
	pendingGo int
}

type LinearGo struct {
	f  func()
	cb func()
}

func New(l int) *Go {
	g := new(Go)
	g.ChanCb = make(chan func(), l)
	return g
}

func (g *Go) Go(f func(), cb func()) {
	g.pendingGo++
	fmt.Println("pendingGo++:", g.pendingGo)

	go func() {
		defer func() {
			g.ChanCb <- cb
			if r := recover(); r != nil {
				if LenStackBuf > 0 {
					buf := make([]byte, LenStackBuf)
					l := runtime.Stack(buf, false)
					fmt.Printf("%v: %s", r, buf[:l])
				} else {
					fmt.Printf("%v", r)
				}
			}
		}()

		f()
	}()
}

func (g *Go) Cb(cb func()) {
	defer func() {
		fmt.Println("pendingGo--:", g.pendingGo)
		g.pendingGo--
		if r := recover(); r != nil {
			if LenStackBuf > 0 {
				buf := make([]byte, LenStackBuf)
				l := runtime.Stack(buf, false)
				fmt.Printf("%v: %s", r, buf[:l])
			} else {
				fmt.Printf("%v", r)
			}
		}
	}()

	if cb != nil {
		cb()
	}
}

func (g *Go) Close() {
	for g.pendingGo > 0 {
		g.Cb(<-g.ChanCb)
	}
}

func (g *Go) Idle() bool {
	return g.pendingGo == 0
}

//----------------------
// 安全的 goroutine 调用
func main() {
	d := New(10)

	// go 1
	var res int
	d.Go(func() {
		fmt.Println("1 + 1 = ?")
		res = 1 + 1
	}, func() {
		// panic("bad thing happen!!!") // 尝试报错
		fmt.Println(res)
	})

	d.Cb(<-d.ChanCb)

	// go 2
	d.Go(func() {
		fmt.Print("My name is ")
	}, func() {
		fmt.Println("Leaf")
	})

	d.Close()
}
