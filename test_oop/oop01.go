package main

// package oop01

// 参考:
// 面向对象 - http://wiki.jikexueyuan.com/project/magical-go/object-oriented.html
// Go 面向对象编程（译） - https://juejin.im/post/5d065cad51882523be6a92f2

import (
	"fmt"
)

func main() {
	// test_001()
	// test_002()
	// test_003()
	test_005()
	// test_101()
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
	val int // 开头字母 大写 为 public, 小写 为 private, 别的 package 不能调用 private
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
	parent // 继承 parent
	num    int
}

func test_003() {
	var c child
	c = child{parent{1}, 2} // 这样初始化值必须按顺序赋值
	fmt.Println(c.num)
	fmt.Println(c.val)
}

// 多态特性, 实现接口
type act interface {
	write()
	read()
}

type xiaoming struct {
}

type xiaofang struct {
}

func (xm *xiaoming) write() { // 只要实现了 act 的接口, 就可以将 xiaofang 的地址赋值给 act (interface) 变量.
	fmt.Println("xiaoming write")
}

func (xm *xiaoming) read() {
	fmt.Println("xiaofang write")
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

	_ = xf
	// w = &xf // 编译报错, 因为 xiaofang 没有实现 read() 方法
}

// struct 里面嵌入 interface
type xiaoli struct {
	age int
	act // 接口成员
}

func (this *xiaoli) write() { // 实现接口
	fmt.Printf("\nxiaoli write, age:%d", this.age)
}

func test_005() {
	xl := &xiaoli{}
	xl.age = 123
	xl.write()

	var xl2 act // 只要实现了接口, 就可以赋值调用
	xl2 = xl
	xl2.write()
	// xl2.read() // 执行时会闪退, 因为 xiaoli 没有实现 read() 方法, 所以不建议使用这种方式, 因为编译期不能提示, 运行期有问题才闪退
	fmt.Printf("\nxl:%v", xl)
	fmt.Printf("\nxl2:%v", xl2)
}
