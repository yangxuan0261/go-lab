// package main

package test_base

import (
	syserr "GoLab/common/error"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"syscall"
	"testing"
	"time"
	"unsafe"
)

// Golang常用包有哪些？- https://www.zhihu.com/question/22009370
// https://godoc.org/

func Test_default(t *testing.T) {
	var a int32
	var b bool
	var c float32
	var d string
	var e1 *CPig
	var e2 CPig

	fmt.Printf("--- a:%+v\n", a)                 // 0
	fmt.Printf("--- b:%+v\n", b)                 // false
	fmt.Printf("--- c:%+v\n", c)                 // 0
	fmt.Printf("--- d:%+v, len:%d\n", d, len(d)) // "", 0
	fmt.Printf("--- e1:%+v\n", e1)               // nil, 空指针
	fmt.Printf("--- e2:%+v\n", e2)               // {name:}

	fmt.Println("--- NumCPU:", runtime.NumCPU()) // 8, cpu 线程数量
	fmt.Println("--- NumGoroutine:", runtime.NumGoroutine())
	runtime.GOMAXPROCS(runtime.NumCPU()) // 正确姿势, 指定为 cpu 的线程数量
}

func Test_string_int_float(t *testing.T) {
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

func Test_type(t *testing.T) {
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

type CBigPig struct {
	CPig
	isMale bool
}

func (self *CPig) Run(speed int) {
	fmt.Printf("--- CPig.Run, name:%s, speed:%d\n", self.name, speed)
}

func Test_dynamicCast(t *testing.T) {
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

	fmt.Println("--- aaa")
	bg1 := new(CBigPig)
	var bg2 interface{}
	bg2 = bg1
	bg3, ok := bg2.(*CPig)
	fmt.Println("--- bg3 res:", bg3 == nil, ok) // true false, 父类不能匹配成子类

	var bg4 interface{}
	bg4 = bg1.CPig                              // 是 对象
	bg5, ok := bg4.(*CPig)                      // 匹配成 指针, 匹配失败
	fmt.Println("--- bg5 res:", bg5 == nil, ok) // true false, 父类不能匹配成子类

	// 子类 匹配成 父类 的正确姿势
	var bg7 interface{}
	bg7 = &bg1.CPig                             // 是 指针
	bg8, ok := bg7.(*CPig)                      // 匹配成 指针, 匹配成功
	fmt.Println("--- bg8 res:", bg8 == nil, ok) // false true

	var bg9 interface{}
	pIns, ok := bg9.(*CPig)
	fmt.Println("--- try cast nil:", pIns, ok)
}

// https://studygolang.com/articles/5769
func TestString(t *testing.T) {
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

func TestLambda(t *testing.T) {
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

func TestFor(t *testing.T) {

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

func TestFuncPtr(t *testing.T) {
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

func TestPtr(t *testing.T) {
	ins1 := CBall{name: "hello"}
	ins2 := &ins1
	fmt.Printf("--- ins1 addr:%p\n", &ins1) // 打印指针地址
	fmt.Printf("--- ins2 addr:%p\n", ins2)
}

func TestPrintStack(t *testing.T) {
	func1 := func() {
		log.Println("--- fucn1")
		debug.PrintStack()
		log.Printf("--- stackInfo:%s\n", string(debug.Stack()))
	}
	func2 := func() {
		func1()
		log.Println("--- fucn2")
	}
	func2()
}

func TestOsInterrupt(t *testing.T) {

	go func() {
		time.Sleep(time.Second * 3)
		os.Exit(0)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	log.Println("--- testOsInterrupt")

	s := <-c
	log.Println("--- exist, signal:", s)
}

func TestOsInterrupt22(t *testing.T) {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		time.Sleep(time.Second * 2)
		// fmt.Println("--- try exit")
		// os.Exit(1)
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

func TestSyserr(t *testing.T) {
	defer func() { // 即使 panic 了也是可以在调用到 defer
		fmt.Println("--- 333")
		syserr.Recover()
	}()

	fmt.Println("--- 111")
	panic("--- wolegequ")
	fmt.Println("--- 222")
}

func TestReturn(t *testing.T) {
	fn := func(num int) (str string, err error) {
		if num > 5 {
			str = "aaa"
			err = errors.New("111")
			return
		} else {
			str = "ccc" // 无效
			return "bbb", errors.New("222")
		}
	}

	str, err := fn(1)
	fmt.Println("--- ret:", str, err)
}

func TestDefer(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("--- err:%+v", err)
		}
	}()
	defer log.Println("aaa")
	defer log.Println("bbb")

	log.Println("--- test 111")
	/*
		--- test 111
		bbb
		aaa
		--- err:hello
		// 后进先出
	*/
	panic("hello")
	log.Println("--- test 222")
}

func TestDefer02(t *testing.T) {
	ok := true
	if ok { //会根据运行时调用不同的 defer
		defer log.Println("bbb")
	} else {
		defer log.Println("ccc")
	}
	log.Println("aaa")
	/*
		2019/12/07 14:39:54 aaa
		2019/12/07 14:39:54 bbb
	*/
}

func TestEmptyStruct(t *testing.T) {
	s1 := struct{}{}
	a := 1
	_ = a
	s2 := struct{}{}
	log.Printf("--- s1, len:%v, addr:%p\n", unsafe.Sizeof(s1), &s1)
	log.Printf("--- s2, len:%v, addr:%p\n", unsafe.Sizeof(s2), &s2)
	/* 长度为 0, 地址一样
	2019/10/30 16:49:41 --- s1, len:0, addr:0x6777e8
	2019/10/30 16:49:41 --- s2, len:0, addr:0x6777e8
	*/
}

func TestCurrDir(t *testing.T) {
	str, err := os.Getwd()
	fmt.Println("--- pwd:", str, err)
}

func Test_FnEqual(t *testing.T) {

	fn1 := func() {

	}

	fn2 := func() {

	}

	fn3 := fn1

	fmt.Printf("--- fn1:%p\n", fn1)
	fmt.Printf("--- fn2:%p\n", fn2)
	fmt.Printf("--- fn3:%p\n", fn3)
	/*
		--- fn1:0x50bf60
		--- fn2:0x50bf70
		--- fn3:0x50bf60 // 与 fn1 地址相同
	*/

	//fmt.Printf("--- fn3 == fn1::%v\n", fn3 == fn1) // 编译报错, 方法只能与 nil 比较 (func can only be compared to nil)

	// 利用反射获取函数地址
	sf1 := reflect.ValueOf(fn1)
	sf2 := reflect.ValueOf(fn1)
	fmt.Println("--- aaa:", sf1.Pointer() == sf2.Pointer()) // true

	sf3 := reflect.ValueOf(fn2)
	fmt.Println("--- bbb:", sf1.Pointer() == sf3.Pointer()) // false

	// 直接使用函数地址
	ptr1 := &fn1
	ptr2 := &fn1
	fmt.Println("--- ccc:", ptr1 == ptr2) // true

	ptr3 := &fn2
	fmt.Println("--- ddd:", ptr1 == ptr3) // false
}

func Test_Ptr(t *testing.T) {

	type CPhone struct {
		Num int
	}

	type CPack struct {
		CPhone
		Name string
	}

	ins := &CPack{CPhone: CPhone{Num: 123}, Name: "123"}
	fmt.Printf("--- ptr1:%p\n", &ins.CPhone.Num)
	fmt.Printf("--- ptr2:%p\n", &(ins.CPhone.Num))
	/*
	   --- ptr1:0xc000004520
	   --- ptr2:0xc000004520 // 地址一样, 说明不用加 () 也可以去到最后一个字段的地址
	*/
}

func Test_copyStruct(t *testing.T) {
	var bArr1 []*CBall
	b1 := &CBall{name: "hello"}
	bArr1 = append(bArr1, b1)

	var bArr2 []*CBall
	b2 := *bArr1[0] // * 取到对象, 然后 值的复制拷贝, b2 新的一个对象
	bArr2 = append(bArr2, &(b2))
	bArr2[0].name = "world"

	fmt.Printf("--- bArr1:%+v\n", bArr1[0]) // bArr1:&{name:hello}
	fmt.Printf("--- bArr2:%+v\n", bArr2[0]) // bArr2:&{name:world} // 可以看到修改 bArr2[0] 不会影响到 bArr1[0]
}
