package main

// package test_base

import (
	"fmt"
	"reflect"
	"strconv"
)

func test_string_int_float() {
	// #string到int
	num1, _ := strconv.Atoi("41")
	fmt.Printf("--- num:%d\n", num1)
	// #string到int64
	num2, _ := strconv.ParseInt("35", 10, 64)
	fmt.Printf("--- num:%d\n", num2)

	// #int到string
	str1 := strconv.Itoa(43)
	fmt.Printf("--- str1:%s\n", str1)
	// #int64到string
	str2 := strconv.FormatInt(int64(37), 10)
	fmt.Printf("--- str2:%s\n", str2)

	// float转string
	v := 3.1415926535
	str3 := strconv.FormatFloat(v, 'E', -1, 32) //float32s2 := strconv.FormatFloat(v, 'E', -1, 64)//float64
	fmt.Printf("--- str3:%s\n", str3)

	// string转float
	str4 := "3.1415926535"
	flt1, _ := strconv.ParseFloat(str4, 32)
	flt2, _ := strconv.ParseFloat(str4, 64)
	fmt.Printf("--- flt1:%f\n", flt1)
	fmt.Printf("--- flt2:%f\n", flt2)
}

func test_type() {
	// 参考: test_reflect.go

	num1 := 123
	mt := reflect.TypeOf(num1)
	fmt.Printf("--- num type:%s\n", mt.Name()) // --- num type:int

	bVar := false
	mt2 := reflect.TypeOf(bVar)
	fmt.Printf("--- bVar type:%s\n", mt2.Name()) // --- bVar type:bool

	str1 := "hello"
	mt3 := reflect.TypeOf(str1)
	fmt.Printf("--- str1 type:%s\n", mt3.Name()) // --- str1 type:string
}

func test_dynamicCast() {
	type CDog struct {
	}
	type CCat struct {
	}

	var ptr1 interface{}
	ptr1 = &CDog{}
	if myDog, ok := ptr1.(*CDog); ok { // 动态匹配 CDog
		fmt.Printf("--- ptr1 is CDog \n") // --- ptr1 is CDog
		_ = myDog
	} else {
		fmt.Printf("--- ptr1 is not CDog \n")
	}
}

func main() {
	test_string_int_float()
	// test_type()
	// test_dynamicCast()
}
