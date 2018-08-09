package main

import (
	"fmt"
	"reflect"
)

/*
interface{} 像 csharp 中的 Object, 所有类型的基类, 有装箱拆箱操作, 也就是有点性能上的消耗.
*/

func main() {
	// test_001()
	// test_002()
	// test_003()
	// test_004()
	// test_005()
	test_006()
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone *NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone *IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func test_001() {
	var phone1, phone2 Phone // 接口是一个指针

	phone1 = new(NokiaPhone)
	phone1.call()

	phone2 = new(IPhone)
	phone2.call()

	println(phone1, phone2) // (0x4ce940,0x54ee08) (0x4ce920,0x54ee08), 接口是一个指针
}

// -------------

func test_002() {
	fn1 := func(val interface{}) {
		v := reflect.ValueOf(val) // 使用 reflect 库
		fmt.Print(v.Kind(), "\n")

		if v.Kind() == reflect.Int {
			fmt.Print(v, val, "\n")
		}
		if v.Kind() == reflect.Bool {
			fmt.Print(v, val, "\n")
		}
		if v.Kind() == reflect.Float64 {
			fmt.Print(v, val, "\n")
		}
	}
	fn1(123)
	fn1(true)
	fn1(123.2)
}

type Element interface{}

func test_003() {
	var e Element = 100
	switch value := e.(type) { //type是一个关键字
	case int:
		fmt.Println("int", value)
	case string:
		fmt.Println("string", value)
	default:
		fmt.Println("unknown", value)
	}
}

// --------------
// 通过继承来实现接口
type Module interface {
	OnInit()
	OnDestroy()
}

type Actor struct {
}

func (a *Actor) OnInit() {
	fmt.Println("--- Actor OnInit")
}

type Cat struct {
	*Actor
}

func (c *Cat) OnDestroy() {
	fmt.Println("--- Cat OnDestroy")
}

func test_004() {
	var tor Module
	tor = new(Cat)
	tor.OnInit()
	tor.OnDestroy()
}

// --------------
// 类型转换
type Human struct {
	name string
}

type Animal struct {
	name string
}

func test_005() {
	var a1 interface{}
	a1 = &Human{"aaa"}
	println("a1:", a1)

	a2 := a1.(*Human)
	println("a2:", a2)

	// a3 := a1.(*Animal)
	// println("a3:", a3)

	switch x := a1.(type) {
	case *Human:
		println("is *Human")
		println("x:", x.name)
	case *Animal:
		println("is *Animal")
		println("x:", x)
	default:
		println("unknown types")
		println("x:", x)
	}

}

// --------------
// 返回值检查 居然还有这种功效
func test_006() {
	var f interface{}
	f = func([]interface{}) {

	}

	checkFn := func(n int8) {
		var ok bool
		switch n {
		case 0:
			_, ok = f.(func([]interface{}))
		case 1:
			_, ok = f.(func([]interface{}) interface{})
		case 2:
			_, ok = f.(func([]interface{}) []interface{})
		default:
			panic("bug")
		}

		if !ok {
			panic("bug")
		}
		return
	}

	checkFn(0)
}
