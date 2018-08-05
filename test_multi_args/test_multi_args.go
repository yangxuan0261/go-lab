package test_multi_args

// package main

import "fmt"

func main() {
	// test_001()
	test_002()
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
