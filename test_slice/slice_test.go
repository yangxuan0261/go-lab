package test_slice

import (
	"fmt"
	"sort"
	"testing"
)

/*
从 slice或数组 中创建slice, 都是共享底层数组, 如果不同享数据, 得使用 copy 函数拷贝数据
*/

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func Test_slice01(t *testing.T) {
	var numbers1 = make([]int, 3, 5) //创建 slice 1
	printSlice(numbers1)             //len=3 cap=5 slice=[0 0 0]

	var numbers2 []int   //创建 slice 2
	printSlice(numbers2) //len=0 cap=0 slice=[]
	if numbers2 == nil {
		fmt.Println("切片是空的")
	}

}

func Test_slice02(t *testing.T) {
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

func Test_slice03(t *testing.T) {
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
func Test_slice_delete(t *testing.T) {
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
func Test_slice04(t *testing.T) {
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

// 防止切出来的切片修改 原有切片的数据, 需要指定第三个参数
func Test_slice05(t *testing.T) {
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

func Test_emptySlice(t *testing.T) {
	var dogArr2 []*CDog                                                                                                 // 这样声明为 nil, 有地址, 可以 append, 正确的数组声明方式
	fmt.Printf("--- dogArr2 len:%d, cap:%d, addr:%p, isnil:%v\n", len(dogArr2), cap(dogArr2), &dogArr2, dogArr2 == nil) // len:0, cap:0, addr:0xc000004540, isnil:true
	dogArr2 = append(dogArr2, &CDog{}, &CDog{}, &CDog{})
	fmt.Printf("--- dogArr2 len:%d, cap:%d, addr:%p\n", len(dogArr2), cap(dogArr2), &dogArr2) // len:3, cap:4, addr:0xc000004540, 地址居然还没变
	fmt.Printf("--- dog:%v\n", dogArr2[0])

	println()
	var arr1 []int = nil                                                                                 // 这样声明为 nil, 有地址, 可以 append
	fmt.Printf("--- arr1 len:%d, cap:%d, addr:%p, isnil:%v\n", len(arr1), cap(arr1), &arr1, arr1 == nil) // len:0, cap:0, addr:0xc000004600, isnil:true
	var arr2 []int = make([]int, 0)                                                                      // make 生成的数组 不为 nil, 正确的数组声明方式
	fmt.Printf("--- arr2 len:%d, cap:%d, addr:%p, isnil:%v\n", len(arr2), cap(arr2), &arr2, arr2 == nil) // arr2 len:0, cap:0, addr:0xc0000046a0, isnil:false
	var arr3 = []int{}                                                                                   // {} 初始化的 不为 nil
	fmt.Printf("--- arr3 len:%d, cap:%d, addr:%p, isnil:%v\n", len(arr3), cap(arr3), &arr3, arr3 == nil) // arr2 len:0, cap:0, addr:0xc0000046c0, isnil:false
}

// 清空数组
func Test_clearSlice(t *testing.T) {
	// ------------ test2
	arr1 := make([]int, 0, 10)
	arr2 := arr1
	fmt.Printf("--- arr1, len:%d, cap:%d, addr:%p\n", len(arr1), cap(arr1), arr1) // len:0, cap:10, addr:0xc0000122d0
	fmt.Printf("--- arr2, len:%d, cap:%d, addr:%p\n", len(arr2), cap(arr2), arr2) // len:0, cap:10, addr:0xc0000122d0 // 地址未变

	println("aaa")
	arr1 = append(arr1, 1, 2, 3)
	fmt.Printf("--- arr1, len:%d, cap:%d, addr:%p\n", len(arr1), cap(arr1), arr1) // len:3, cap:10, addr:0xc0000122d0
	fmt.Printf("--- arr2, len:%d, cap:%d, addr:%p\n", len(arr2), cap(arr2), arr2) // len:0, cap:10, addr:0xc0000122d0 // arr2 不受 arr1 的影响, 依旧是 len:0 // TODO: 这个有点不明白为什么 不受 arr1 的影响

	println("bbb")
	//参考:
	// https://programming.guide/go/clear-slice.html
	// https://yourbasic.org/golang/clear-slice/
	// 错误 的清空方式
	arr1 = arr1[:0]                                                               // 清空数组, 但是是 错误 的, 元素会继续占据内存
	fmt.Printf("--- arr1, len:%d, cap:%d, addr:%p\n", len(arr1), cap(arr1), arr1) // len:0, cap:10, addr:0xc0000122d0 // 地址未变, 虽然长度 len 为 0, 但依旧可以切出 原来的数据
	fmt.Printf("--- arr1, reappears:%v\n", arr1[:2])                              // [1 2], 重新出现原来的数据, If the slice is extended again, the original data reappears.
	fmt.Printf("--- arr2, len:%d, cap:%d, addr:%p\n", len(arr2), cap(arr2), arr2) // len:0, cap:10, addr:0xc0000122d0
	arr1 = append(arr1, 4, 5, 6)
	fmt.Printf("--- arr1, len:%d, cap:%d, addr:%p\n", len(arr1), cap(arr1), arr1) // len:3, cap:10, addr:0xc0000122d0 // 地址未变
	arr6 := append(arr1, 4, 5, 6)                                                 // 新的数组变量, 地址改变 ****************
	fmt.Printf("--- arr6, len:%d, cap:%d, addr:%p\n", len(arr6), cap(arr6), arr6) // len:6, cap:10, addr:0xc0000122d0 // 地址未变
	arr6[1] = 999
	fmt.Printf("--- arr1[1]:%v\n", arr1) // [4 999 6] // 会修改到 arr1 的数据
	fmt.Printf("--- arr6[1]:%v\n", arr6) // [4 999 6 4 5 6]

	// 正确 的清空方式
	println("bbb 222")
	arr1 = nil                                                                    // 清空数组, 正确的姿势
	fmt.Printf("--- arr1, len:%d, cap:%d, addr:%p\n", len(arr1), cap(arr1), arr1) // len:0, cap:0, addr:0x0 // 空地址
	arr1 = append(arr1, 4, 5, 6)                                                  // 即使为 nil, 除了索引元素之外, 几乎其他 api 都可以使用, 如: append
	fmt.Printf("--- arr1, len:%d, cap:%d, addr:%p\n", len(arr1), cap(arr1), arr1) // len:3, cap:4, addr:0xc0000104e0 // 会分配到一个新内存

	//fmt.Printf("--- arr1, not reappears:%v\n", arr1[:2])                               // 不可用切出数据, 不然会数组越界崩溃

	// ------------ test2
	println("ccc")
	arr3 := make([]int, 0, 10)
	arr3 = append(arr3, 1, 2, 3)
	fmt.Printf("--- arr3, len:%d, cap:%d, addr:%p\n", len(arr3), cap(arr3), arr3) // len:3, cap:10, addr:0xc000012320

	arr3 = arr3[:cap(arr3)]                                                       // 填满数组
	fmt.Printf("--- arr3, len:%d, cap:%d, addr:%p\n", len(arr3), cap(arr3), arr3) // len:10, cap:10, addr:0xc000012320
	arr3 = append(arr3, 1, 2, 3)                                                  // 动态扩容, 导致地址改变
	fmt.Printf("--- arr3, len:%d, cap:%d, addr:%p\n", len(arr3), cap(arr3), arr3) // len:13, cap:20, addr:0xc0000a4000
}

func Test_copy(t *testing.T) {
	// 错误姿势
	arr3 := []int{1, 2, 3}
	arr4 := arr3
	arr4[2] = 777
	fmt.Printf("--- arr3, addr:%p, arr:%v\n", &arr3, arr3) // [1 2 777]
	fmt.Printf("--- arr4, addr:%p, arr:%v\n", &arr4, arr4) // [1 2 777] // 会修改到 arr3 的数据

	// 正确姿势
	arr1 := []int{1, 2, 3}

	arr2 := make([]int, len(arr1))
	copy(arr2, arr1)

	arr2[2] = 666
	fmt.Printf("--- arr1, addr:%p, arr:%v\n", &arr1, arr1) // [1 2 3]
	fmt.Printf("--- arr2, addr:%p, arr:%v\n", &arr2, arr2) // [1 2 666] // 不会修改到 arr1 的数据
}

func Test_arrAppendArr(t *testing.T) {
	arr1 := []int{
		4, 5, 6,
	}

	arr0 := []int{
		1, 2, 3,
	}

	arr0 = append(arr0, arr1...) // 正确
	fmt.Printf("--- arr0:%+v\n", arr0)
}

type Person struct {
	Name string
	Age  int
}

// 按照 Person.Age 从大到小排序
type PersonSlice []Person

func (a PersonSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a PersonSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a PersonSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Age < a[i].Age
}

// 封装成 wrap
type PersonWrapper struct { //注意此处
	people []Person
	by     func(p, q *Person) bool
}

func (pw PersonWrapper) Len() int { // 重写 Len() 方法
	return len(pw.people)
}
func (pw PersonWrapper) Swap(i, j int) { // 重写 Swap() 方法
	pw.people[i], pw.people[j] = pw.people[j], pw.people[i]
}
func (pw PersonWrapper) Less(i, j int) bool { // 重写 Less() 方法
	return pw.by(&pw.people[i], &pw.people[j])
}

func Test_sort(t *testing.T) {
	fmt.Println("--- 基础类型排序")
	// 基础类型
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Sort(sort.Reverse(sort.IntSlice(intList))) // 逆序
	sort.Sort(sort.Reverse(sort.Float64Slice(float8List)))
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)

	fmt.Println()
	// 数据结构
	// 参考: https://itimetraveler.github.io/2016/09/07/%E3%80%90Go%E8%AF%AD%E8%A8%80%E3%80%91%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B%E6%8E%92%E5%BA%8F%E5%92%8C%20slice%20%E6%8E%92%E5%BA%8F/
	people := []Person{
		{"shang san", 12},
		{"aaa", 12},
		{"zzz", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}

	fmt.Println(people)

	sort.Sort(PersonSlice(people)) // 按照 Age 的逆序排序
	fmt.Println(people)

	sort.Sort(sort.Reverse(PersonSlice(people))) // 按照 Age 的升序排序
	fmt.Println(people)

	fmt.Println("--- 字段优先级排序")
	sort.Sort(PersonWrapper{people, func(a, b *Person) bool {
		if a.Age == b.Age { // 排序优先级 Age > Name
			return a.Name > b.Name
		} else {
			return a.Age < b.Age // Age 递减排序
		}
	}})
	fmt.Println(people)
}

func Test_Append(t *testing.T) {
	var arr1 []*Person
	fmt.Printf("--- arr1, len:%d, isnil:%v\n", len(arr1), arr1 == nil) // arr1, len:0, isnil:true

	arr1 = append(arr1, nil)
	fmt.Printf("--- arr1, len:%d, isnil:%v\n", len(arr1), arr1 == nil) // arr1, len:1, isnil:false, 有一个 nil 在 0 位置

	var arr2 []*Person
	var arr3 []*Person
	arr3 = append(arr3, arr2...)
	fmt.Printf("--- arr3, len:%d, isnil:%v\n", len(arr3), arr3 == nil) // arr3, len:0, isnil:true, nil append nil 还是 nil
}

// 类似 c++ stl 中的 vector, 动态增长数组
