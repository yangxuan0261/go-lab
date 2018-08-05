package oop01

// 参考: http://wiki.jikexueyuan.com/project/magical-go/object-oriented.html

import (
	"fmt"
)

func main() {
	// test_001()
	// test_002()
	test_003()
}

type Actor struct {
	name string
	age  int8
}

func (a *Actor) walk(dis float32) {
	fmt.Printf("%s walk meter %f", a.name, dis)
}

func test_001() {
	var m1 *Actor = new(Actor)
	m1.walk(12.333)
}

// 封装特性
type data struct {
	val int
}

func (p_data *data) set(num int) {
	p_data.val = num
}

func (p_data *data) get() int {
	return p_data.val
}

func (p_data *data) get22() (int, string) {
	return p_data.val, "hello"
}

func test_002() {
	p_data := &data{4}
	p_data.set(5)
	fmt.Println(p_data.get())

	_, dv := p_data.get22()
	fmt.Println(dv)
}

// 继承特性
type parent struct {
	val int
}

type child struct {
	parent
	num int
}

func test_003() {
	var c child
	c = child{parent{1}, 2}
	fmt.Println(c.num)
	fmt.Println(c.val)
}

// 继承特性
type act interface {
	write()
}

type xiaoming struct {
}

type xiaofang struct {
}

func (xm *xiaoming) write() {
	fmt.Println("xiaoming write")
}

func (xf *xiaofang) write() {
	fmt.Println("xiaofang write")
}

func test_004() {
	var w act

	xm := xiaoming{}
	xf := xiaofang{}

	w = &xm
	w.write()

	w = &xf
	w.write()
}
