package main

// package test_base

import (
	"fmt"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
)

// Golang常用包有哪些？- https://www.zhihu.com/question/22009370
// https://godoc.org/

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

type IActor interface {
	Run(speed int)
}

type CPig struct {
	name string
}

func (self *CPig) Run(speed int) {
	fmt.Printf("--- CPig.Run, name:%s, speed:%d\n", self.name, speed)
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

	Print := func(i interface{}) {
		switch i.(type) {
		case string:
			fmt.Printf("type is string,value is:%v\n", i.(string))
			break
		case float64:
			fmt.Printf("type is float32,value is:%v\n", i.(float64))
			break
		case int:
			fmt.Printf("type is int,value is:%v\n", i.(int))
			break
		default:
			fmt.Printf("type is unknown\n")
		}
	}
	var i interface{}
	i = "hello"
	Print(i)
	i = 100
	Print(i)
	i = 1.29
	Print(i)

	pig := &CPig{name: "hello"}
	var actor interface{}
	actor = pig
	if actIns, actOk := actor.(IActor); actOk { // 接口的匹配指针不需要 *, actor 必须是 指针才能匹配成功, 如果是 对象 话将匹配失败
		actIns.Run(666)
	} else {
		fmt.Println("--- cast IActor fail")
	}
}

// https://studygolang.com/articles/5769
func testString() {
	Slash := func(r rune) rune {
		if r == '\\' {
			return '/'
		}
		return r
	}

	s := "C:\\Windows\\System32\\FileName"
	ms := strings.Map(Slash, s)
	fmt.Printf("%q\n", ms) // "C:/Windows/System32/FileName"
}

func testLambda() {
	// https://blog.csdn.net/wangshubo1989/article/details/79217291
	text := "hello"
	foo := func(age11 int) (int, string) {
		fmt.Printf("--- text:%s, age11:%d\n", text, age11)
		return 666, "world"
	}

	// calling the closure
	age22, value := foo(123)
	fmt.Printf("--- value:%s, age22:%d\n", value, age22)
}

func testFor() {

	for i := 1; i < 10; i++ { // 和 C 语言的 for 一样:
		fmt.Printf("--- i:%d\n", i)
	}

	cnt := 1
	flag := true
	for flag { // 和 C 的 while 一样：
		if cnt == 5 {
			flag = false
		} else {
			fmt.Printf("--- cnt:%d\n", cnt)
			cnt++
		}
	}

	cnt222 := 1
	for { // 和 C 的 for(;;) 一样：
		if cnt222 == 5 {
			break
		} else {
			fmt.Printf("--- cnt222:%d\n", cnt222)
			cnt222++
		}
	}
}

type bbbFunc func(int, string)

func aaaFunc(arg1 int, arg2 string) {
	fmt.Printf("--- arg1:%d, arg2:%s\n", arg1, arg2)
}

type CBall struct {
	name string
}

func (self *CBall) Run(speed int) {
	fmt.Printf("--- CBall, name:%s, run speed:%d\n", self.name, speed)
}

func testFuncPtr() {
	var funcPtr1 func(int, string) // 函数指针
	funcPtr1 = aaaFunc
	funcPtr1(123, "hello")

	var funcPtr2 bbbFunc // 函数指针
	funcPtr2 = aaaFunc
	funcPtr2(456, "world")

	insBall := &CBall{name: "Tom"}
	insBall.Run(666)

	var funcPtr3 func(int)
	funcPtr3 = insBall.Run // 实例对象的函数
	funcPtr3(777)
}

func testPtr() {
	ins1 := CBall{name: "hello"}
	ins2 := &ins1
	fmt.Printf("--- ins1 addr:%p\n", &ins1) // 打印指针地址
	fmt.Printf("--- ins2 addr:%p\n", ins2)
}

func testPrintStack() {
	func1 := func() {
		fmt.Println("--- fucn1")
		debug.PrintStack()
	}
	func2 := func() {
		func1()
		fmt.Println("--- fucn2")
	}
	func2()
}

func main() {
	// test_string_int_float()
	// test_type()
	test_dynamicCast()
	// testLambda()
	// testString()
	// testFor()
	// testFuncPtr()
	// testPtr()
	// testPrintStack()
}
