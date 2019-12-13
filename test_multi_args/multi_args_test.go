package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_001(t *testing.T) {
	fn1 := func(nums ...int) { // 形参输一个 slice 切片
		fmt.Print(nums, "\n")
	}

	fn1(1, 2)    // [1 2]
	fn1(3, 4, 5) // [3 4 5]
	nums := []int{1, 2, 3, 4}
	fn1(nums...) // [1 2 3 4], 把切片打散传入, 语法糖
}

func Test_002(t *testing.T) {
	fn1 := func(num int, args ...interface{}) {
		fmt.Printf("len:%d, addr2:%p, args:%+v, num:%d\n", len(args), args, args, num)
		if args == nil {
			fmt.Println("args is nil")
		}
		at := reflect.TypeOf(args)
		fmt.Printf("at:%+v\n", at, ) // at:[]interface {}, 数组
		if len(args) > 0 {
			fmt.Printf("[0]:%+v\n", args[0])
		}
	}

	fn1(1, "123", true) // len:2, addr2:0xc000004520, args:[123 true], num:1
	fn1(1)              // len:0, addr2:0x0, args:[], num:1, args is nil // 切片时是空地址

	println()
	// nums := []int{1, 2, 3, 4} // 错误写法,  必须是 interface{} 类型数组
	nums := []interface{}{1, 2, 3, 4} // 正确写法
	fmt.Printf("addr1:%p\n", nums)    // addr1:0xc00001a6c0
	fn1(1, nums...)                   // len:4, addr2:0xc00001a6c0, args:[1 2 3 4], num:1 // 地址相同, 数组必须要通过 ... 展开
}

type Actor struct {
	name string
}

type Animal struct {
	name string
}

func Test_003(t *testing.T) {
	fn1 := func(args []interface{}) {
		fmt.Printf("fn1 args addr:%p, len:%d\n", args, len(args))

		if args == nil {
			fmt.Println("fn1 args is nil")
		} else {
			fmt.Println("fn1 args:", args)
		}
	}

	fn2 := func(args ...interface{}) {
		fmt.Println("fn2 args:", args, " len:", len(args))
		fmt.Println("fn2 args[0]:", args[0]) // 如果 fn2() 这样不传参调用的话, 这里会空指针奔溃

		if a2, ok := args[2].(*Actor); ok { // 类型装换, 获取正确类型 *Actor
			fmt.Println("a2 name:", a2.name)
		} else {
			fmt.Println("a2 is not *Animal")
		}
	}

	a1 := Actor{"hello"}
	a2 := &Actor{"world"}
	fn2(1, a1, a2)

	// fn1() // 报错 not enough arguments in call to fn1, 必须得有参数
	pi1 := []interface{}{}
	pi2 := []interface{}{a1}
	pi3 := []interface{}{a2}
	fmt.Printf("pi1 addr:%p\n", pi1) // 地址和方法里的地址相同, 说明是个切片都是
	fmt.Printf("pi2 addr:%p\n", pi2)
	fmt.Printf("pi3 addr:%p\n", pi3)

	fn1(pi1)
	fn1(pi2)
	fn1(pi3)
}

func Test_004(t *testing.T) {
	func1 := func(nums ...int) {
		fmt.Printf("--- nums:%+v\n", nums)
	}

	arr1 := []int{
		4, 5, 6,
	}

	// func1(1, 2, 3, arr1...) // 错误

	arr0 := []int{
		1, 2, 3,
	}

	arr0 = append(arr0, arr1...) // 正确
	func1(arr0...)
}

func Test_parse(t *testing.T) {
	func1 := func(nums ...interface{}) {
		fmt.Printf("--- nums:%+v\n", nums)

		if len(nums) > 0 {
			if a, ok := nums[0].(*Actor); ok {
				fmt.Printf("--- a.name:%+v\n", a.name)
			}
		}
	}

	a1 := &Actor{name: "hello"}
	func1(a1, 222)
}
