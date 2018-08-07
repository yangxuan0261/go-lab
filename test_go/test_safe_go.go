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
		g.pendingGo--
		fmt.Println("pendingGo--:", g.pendingGo)
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

	d.Cb(<-d.ChanCb) // 这里用 ChanCb (chan func() 函数信号) 阻塞, 等 Go() 中执行完 f 后, 函数结束时 (defer) 将发送 回调函数 发送给 ChanCb信号 (g.ChanCb <- cb), 这里接收到信号后 获取到 回调函数 cb 作为 Cb 的形参传进去, 在里面执行 cb()

	// go 2
	d.Go(func() {
		fmt.Print("My name is ")
	}, func() {
		fmt.Println("Leaf")
	})

	d.Close()
}
