package main

import "fmt"

/*
从 slice或数组 中创建slice, 都是共享底层数组, 如果不同享数据, 得使用 copy 函数拷贝数据
*/

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func test_slice01() {
	var numbers1 = make([]int, 3, 5) //创建 slice 1
	printSlice(numbers1)             //len=3 cap=5 slice=[0 0 0]

	var numbers2 []int   //创建 slice 2
	printSlice(numbers2) //len=0 cap=0 slice=[]
	if numbers2 == nil {
		fmt.Println("切片是空的")
	}

}

func test_slice02() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8} //创建 slice 3
	printSlice(numbers)                         // len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]

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

// 删除测试
func test_slice_delete() {
	var removeByEle func(slice []int, elem int) []int
	removeByEle = func(slice []int, elem int) []int {
		if len(slice) == 0 {
			return slice
		}
		for i, v := range slice {
			if v == elem {
				slice = append(slice[:i], slice[i+1:]...)
				return removeByEle(slice, elem)
				break
			}
		}
		return slice
	}

	var removeByIndex func(slice []int, index int) []int
	removeByIndex = func(slice []int, index int) []int {
		if len(slice) == 0 {
			return slice
		}
		for i := range slice {
			if i == index {
				slice = append(slice[:i], slice[i+1:]...)
				break
			}
		}
		return slice
	}
	slice := []int{111, 222, 333, 444, 555, 333}
	for _, val := range slice {
		fmt.Print(val, " ")
	}

	fmt.Println("\n--- 按 元素 333 删除")
	ns1 := removeByEle(slice, 333)
	for _, val := range ns1 {
		fmt.Print(val, " ")
	}

	fmt.Println("\n--- 按 索引 1 删除")
	ns2 := removeByIndex(slice, 1)
	for _, val := range ns2 {
		fmt.Print(val, " ")
	}

}

// 从 slice 中创建 slice 的注意事项
func test_slice04() {
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	fmt.Println("slice:", len(slice), cap(slice), slice)             // slice: 5 5 [10 20 30 40 50]
	fmt.Println("newSlice:", len(newSlice), cap(newSlice), newSlice) // newSlice: 2 4 [20 30]
	/* 计算方式
	对底层数组容量是 k 的切片, slice[i:j]来说
	长度: j - i,	也就是 3 - 1
	容量: k - i,	也就是 5 - 1
	*/

	// 往 新slice 中 append 数据, 不超过容量的情况下, 会影响到底层数据, 因为此时是共享底层数据的
	fmt.Println("--- 分割线 1")
	newSlice2 := append(newSlice, 66)
	fmt.Println("newSlice2:", len(newSlice2), cap(newSlice2), newSlice2) // newSlice2: 3 4 [20 30 66]
	fmt.Println("slice:", len(slice), cap(slice), slice)                 // slice: 5 5 [10 20 30 66 50]

	// 往 新slice 中 append 的数量, 超过容量的部分, 不会影响原来的底层数据, 因为此时是不共享底层数据的, 从 newSlice 复制拷贝出了新的一份新内存数据
	newSlice21 := append(newSlice, 77, 88, 99)
	fmt.Println("newSlice21:", len(newSlice21), cap(newSlice21), newSlice21) // newSlice21: 5 8 [20 30 77 88 99]
	fmt.Println("slice:", len(slice), cap(slice), slice)                     // slice: 5 5 [10 20 30 66 50]
	// 尝试修改新复制拷贝出来的 newSlice21
	newSlice21[0] = 123
	fmt.Println("newSlice21 222 :", len(newSlice21), cap(newSlice21), newSlice21) // newSlice21 222 : 5 8 [123 30 77 88 99]
	fmt.Println("slice 222:", len(slice), cap(slice), slice)                      // slice 222: 5 5 [10 20 30 66 50] // slice 并没有被修改到, 所以已经和 newSlice21 不共享底层数组内存

	fmt.Println("--- 分割线 2")
	newSlice2[0] = 123                                               // 修改底层数组
	fmt.Println("newSlice:", len(newSlice), cap(newSlice), newSlice) // newSlice: 2 4 [123 30]
	fmt.Println("slice:", len(slice), cap(slice), slice)             // slice: 5 5 [10 123 30 66 50]

	fmt.Println("--- 分割线3")
	// 从数组中创建slice
	arr := [5]int{1, 2, 3}
	fmt.Println("arr:", len(arr), cap(arr), arr) // arr: 5 5 [1 2 3 0 0]
	newSlice3 := arr[1:3]
	fmt.Println("newSlice3:", len(newSlice3), cap(newSlice3), newSlice3) // newSlice3: 2 4 [2 3]

	fmt.Println("--- 分割线4")
	//如果不需要额外的容量, 用一下方式创建更加节省
	newSlice4 := slice[1:3:4]
	fmt.Println("newSlice4:", len(newSlice4), cap(newSlice4), newSlice4) // newSlice4: 2 3 [123 30]
	/* 计算方式
	对于 slice[i:j:k] 或 [1:3:4]]
	长度: j – i 或 3 - 1 = 2
	容量: k – i 或 4 - 1 = 3
	*/
	newSlice5 := slice[1:3:3]                                            // 不需要额外的容量
	fmt.Println("newSlice5:", len(newSlice5), cap(newSlice5), newSlice5) // newSlice5: 2 2 [123 30]

}

// 防止切除来的切片修改 原有切片的数据, 需要指定第三个参数
func test_slice05() {
	slice := []int{10, 20, 30, 40, 50}
	newSlice1 := slice[1:2:2]
	newSlice1 = append(newSlice1, 666)                                   // 复制拷贝到新的内存块, 就不会修改到 slice 的数据
	fmt.Println("slice:", len(slice), cap(slice), slice)                 // slice: 5 5 [10 20 30 40 50]
	fmt.Println("newSlice1:", len(newSlice1), cap(newSlice1), newSlice1) // newSlice1: 2 2 [20 666]

	// copy(dst, src) // 或者是使用系统 copy 函数
}

type CDog struct {
	name string
	age  int
}

func test_emptySlice() {

	var dogArr []*CDog
	if dogArr == nil {
		fmt.Println("--- is nil") // is nil, dogArr 所指向的对象 为 nil, 用 len 判断即可
	}
	fmt.Printf("--- dogArr:%p\n", &dogArr)         // 0xc000064440 有地址
	fmt.Printf("--- dogArr len:%d\n", len(dogArr)) // len:0

	dogArr2 := []*CDog{}                             // 空数组
	fmt.Printf("--- dogArr2 len:%d\n", len(dogArr2)) // len:0
}

func test_copy() {
	arr1 := []*CDog{
		&CDog{name: "aaa"},
		&CDog{name: "bbb"},
	}

	arr2 := make([]*CDog, len(arr1))
	copy(arr2, arr1)

	for _, val := range arr2 {
		fmt.Printf("--- name:%s\n", val.name)
	}
}

// 类似 c++ stl 中的 vector, 动态增长数组
func main() {
	// test_slice01()
	// test_slice02()
	// test_slice03()
	// test_slice04()
	// test_slice05()
	// test_slice_delete()

	// test_emptySlice()
	test_copy()
}
