package main

import "fmt"

func main() {
	// test_001()
	// test_002()
	test_003()
}

func test_001() {
	fn1 := func(nums ...int) { // 形参输一个 slice 切片
		fmt.Print(nums, "\n")
	}

	fn1(1, 2)    // [1 2]
	fn1(3, 4, 5) // [3 4 5]
	nums := []int{1, 2, 3, 4}
	fn1(nums...) // [1 2 3 4], 把切片打散传入, 语法糖
}

func test_002() {
	fn1 := func(num int, args ...interface{}) {
		fmt.Printf("addr2:%p\n", args)
		if args == nil {
			fmt.Println("args is nil")
		}
		fmt.Print(num, args, "\n")
	}

	fn1(1, "123", true) // 1 [123 true]
	fn1(1)              // 1 [] // 切片时是空地址

	// nums := []int{1, 2, 3, 4} // 错误写法
	nums := []interface{}{1, 2, 3, 4} // 正确写法
	fmt.Printf("addr1:%p\n", nums)
	fn1(1, nums...) // 1 [1 2 3 4]

	/*
		addr1:0xc0420600c0
		addr2:0xc0420600c0 // 地址相同
	*/
}

type Actor struct {
	name string
}

type Animal struct {
	name string
}

func test_003() {
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
