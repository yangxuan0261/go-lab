// package test_pkg

package main

import (
	"fmt"
	"testing"
)

type EnumType int32

const (
	EPolicyMIN EnumType = iota
	EPolicyMAX
	EPolicyMID
	EPolicyAVG
)

var enumMap = map[EnumType]string{
	EPolicyMIN: "EPolicyMIN",
	EPolicyMAX: "EPolicyMAX",
	EPolicyMID: "EPolicyMID",
	EPolicyAVG: "EPolicyAVG",
}

// switch 与 map 的性能比较: https://segmentfault.com/a/1190000011361164
// func (p EnumType) String() string { // 重写 String() 方法
// 	switch p {
// 	case Policy_MIN:
// 		return "MIN"
// 	case Policy_MAX:
// 		return "MAX"
// 	case Policy_MID:
// 		return "MID"
// 	case Policy_AVG:
// 		return "AVG"
// 	default:
// 		return "UNKNOWN"
// 	}
// }

func (p EnumType) String() string {
	if val, ok := enumMap[p]; ok {
		return val
	} else {
		return "Unknown enum"
	}
}

func foo(p EnumType) {
	fmt.Printf("--- ccc: %v\n", p) // MAX
	fmt.Println("--- aaa:", p)     // MAX
}

func foo2(p string) {
	fmt.Printf("--- bbb: %s\n", p) // MAX
}

func Test_main(t *testing.T) {
	foo(EPolicyMAX)
	foo2(EPolicyMAX.String())
}
