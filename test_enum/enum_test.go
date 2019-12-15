// package test_pkg

package main

import (
	"fmt"
	"testing"
)

type EnumType int32

const (
	EPolicyMIN EnumType = iota // 从 0 开始递增
	EPolicyMAX
	EPolicyMID
	EPolicyAVG
)

type EnumActor int32

const (
	EAHello EnumActor = iota + 1000 // 从 1000 开始递增
	EAWorld
	EARun

	EAWalk EnumActor = iota + 2000 // *** 这里并不是 2000 开头, 而是 iota = 3 了, 所以这里是 2003
	EAFly
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
//
//func (p EnumType) String() string {
//	if val, ok := enumMap[p]; ok {
//		return val
//	} else {
//		return "Unknown enum"
//	}
//}

func foo(p EnumType) {
	fmt.Printf("--- ccc: %v\n", p) // MAX
	fmt.Println("--- aaa:", p)     // MAX
}

func foo2(p string) {
	fmt.Printf("--- bbb: %s\n", p) // MAX
}

func Test_String(t *testing.T) {
	foo(EPolicyMAX)
	//foo2(EPolicyMAX.String())
}

func Test_Print(t *testing.T) {
	fmt.Printf("--- EPolicyMIN:%+v\n", EPolicyMIN)
	fmt.Printf("--- EPolicyMAX:%+v\n", EPolicyMAX)
	fmt.Printf("--- EPolicyMID:%+v\n", EPolicyMID)
	fmt.Printf("--- EPolicyAVG:%+v\n", EPolicyAVG)

	println()
	fmt.Printf("--- EAHello:%+v\n", EAHello)
	fmt.Printf("--- EAWorld:%+v\n", EAWorld)
	fmt.Printf("--- EARun:%+v\n", EARun)
	fmt.Printf("--- EAWalk:%+v\n", EAWalk)
	fmt.Printf("--- EAFly:%+v\n", EAFly)

	/* 只要不重写 String() 方法即可
	--- EPolicyMIN:0
	--- EPolicyMAX:1
	--- EPolicyMID:2
	--- EPolicyAVG:3

	--- EAHello:1000
	--- EAWorld:1001
	--- EARun:1002
	--- EAWalk:2003
	--- EAFly:2004
	*/
}
