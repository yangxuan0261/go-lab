package test_sort

import (
	"fmt"
	"sort"
	"testing"
)

// 为了能够使用自定义函数来排序，我们需要一个
// 对应的排序类型，比如这里我们为内置的字符串
// 数组定义了一个别名ByLength
type ByLength []string

// 我们实现了sort接口的Len，Less和Swap方法, 必须实现的三个方法
// 这样我们就可以使用sort包的通用方法Sort
// Len和Swap方法的实现在不同的类型之间大致
// 都是相同的，只有Less方法包含了自定义的排序
// 逻辑，这里我们希望以字符串长度升序排序
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// 一切就绪之后，我们就可以把需要进行自定义排序
// 的字符串类型fruits转换为ByLength类型，然后使用
// sort包的Sort方法来排序
func Test_001(t *testing.T) {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

//------------ 自定义数据结构
type Actor struct {
	age  int32
	name string
}

type ByActorAge []*Actor

func (s ByActorAge) Len() int {
	return len(s)
}
func (s ByActorAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByActorAge) Less(i, j int) bool {
	return s[i].age < s[j].age
}

func Test_002(t *testing.T) {
	a1 := &Actor{name: "aaa", age: 55}
	a2 := &Actor{name: "bbc", age: 33}
	a3 := &Actor{name: "ccc", age: 66}
	a4 := &Actor{name: "ddd", age: 22}

	// // make 的话必须这样指定索引对应值, 用 append 会出错
	// slice1 := make(ByActorAge, 4, 4)
	// slice1[0] = a1
	// slice1[1] = a2
	// slice1[2] = a3
	// slice1[3] = a4
	// fmt.Println("slice1:", slice1)
	// fmt.Println("value:", reflect.TypeOf(slice1))

	// slice1 := ByActorAge{a1, a2, a3, a4} // 这种方式或下面一种都行
	slice1 := ByActorAge{}
	slice1 = append(slice1, a1)
	slice1 = append(slice1, a2)
	slice1 = append(slice1, a3)
	slice1 = append(slice1, a4)
	fmt.Println("slice1:", slice1)
	for _, val := range slice1 {
		fmt.Println(val.age, val.name)
	}

	fmt.Println("")
	sort.Sort(slice1)
	for _, val := range slice1 {
		fmt.Println(val.age, val.name)
	}
}
