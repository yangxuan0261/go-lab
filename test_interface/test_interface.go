package test_interface

// package main

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
	test_003()
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
