package test_oop

// 参考:
// 面向对象 - http://wiki.jikexueyuan.com/project/magical-go/object-oriented.html
// Go 面向对象编程（译） - https://juejin.im/post/5d065cad51882523be6a92f2

import (
	"go-lab/test_oop/oop2"
	"fmt"
	"testing"
)

type Actor struct {
	name string
	age  int8
}

func (a *Actor) walk(dis float32) {
	fmt.Printf("%s walk meter %f", a.name, dis)
}

func Test_001(t *testing.T) {
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

func Test_002(t *testing.T) {
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

func (self *parent) Run(speed int) {
	fmt.Printf("--- parent, val:%d, speed:%d\n", self.val, speed)
}

type child struct {
	*parent // 继承 parent
	num     int
}

type child2 struct {
	parent // 继承 parent, 一般集成父类还是用 值类型,
	num    int
}

func (self *child2) Run(speed int) {
	self.parent.Run(speed) // 调用父类方法, 等价于其他语言的 super.xxx()
	fmt.Printf("--- child2, val:%d, speed:%d\n", self.val, speed)
}

func Test_003(t *testing.T) {
	var c child
	c = child{parent: &parent{1}, num: 2} // 这样初始化值必须按顺序赋值
	fmt.Println(c.num)
	fmt.Println(c.val)

	c.Run(111)        // 直接调用父类方法
	c.parent.Run(222) // 通过父类调用父类方法

	c2 := child2{parent: parent{33}, num: 3} // 这样初始化值必须按顺序赋值
	c2.Run(333)
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

func Test_004(t *testing.T) {
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

func Test_005(t *testing.T) {
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

func Test_private(t *testing.T) {
	d := oop2.DIns
	fmt.Printf("--- d:%+v\n", d)
}
