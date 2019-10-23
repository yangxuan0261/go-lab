package test_oop

import (
	"fmt"
	"testing"
)

// 参考: http://wiki.jikexueyuan.com/project/magical-go/object-oriented.html

type GateHandler interface {
	Bind(addr string)
	UnBind()
}

type Gate struct {
	name    string
	handler GateHandler
}

// 只要实现了GateHandler 接口, 就可以赋值到这个 handler字段, 其实这里的实现是 handler 和 实例的 g 是同一个对象.
func (this *Gate) SetGateHandler(hdr GateHandler) {
	this.handler = hdr
}

func (this *Gate) GetGateHandler() GateHandler {
	return this.handler
}

// impl GateHandler
func (this *Gate) Bind(addr string) {
	fmt.Println("Bind:", addr)
	fmt.Printf("g:%p, g.handler:%p\n", this, this.handler)
}

func (this *Gate) UnBind() {
	fmt.Println("UnBind:")
}

func Test_101(t *testing.T) {
	g := &Gate{
		name: "hello",
	}
	g.SetGateHandler(g)
	g.GetGateHandler().Bind("world")
}
