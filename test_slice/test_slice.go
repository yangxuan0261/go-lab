package test_slice

// package main

import "fmt"

// 类似 c++ stl 中的 vector, 动态增长数组
func main() {
	// test_slice01()
	test_slice02()
	// test_slice03()
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func test_slice01() {
	var numbers1 = make([]int, 3, 5)
	printSlice(numbers1) //len=3 cap=5 slice=[0 0 0]

	var numbers2 []int
	printSlice(numbers2) //len=0 cap=0 slice=[]
	if numbers2 == nil {
		fmt.Println("切片是空的")
	}

}

func test_slice02() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers) // len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers) //numbers == [0 1 2 3 4 5 6 7 8]

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4]) // numbers[1:4] == [1 2 3]

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3]) // numbers[:3] == [0 1 2]

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:]) // numbers[4:] == [4 5 6 7 8]

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1) // len=0 cap=5 slice=[]

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	numbers[1] = 100
	printSlice(number2) // len=2 cap=9 slice=[0 100] // number2 指向 numbers 切片的指针
	for i, v := range number2 {
		fmt.Printf("number2[%d]=%d\n", i, v)
	}
	// number2[0]=0
	// number2[1]=100

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3) // len=3 cap=7 slice=[2 3 4]
}

func test_slice03() {
	var numbers []int
	printSlice(numbers) // len=0 cap=0 slice=[]

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers) // len=1 cap=1 slice=[0]

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers) // len=2 cap=2 slice=[0 1]

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers) // len=5 cap=6 slice=[0 1 2 3 4]

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1) // len=5 cap=12 slice=[0 1 2 3 4]
}
