// package test_error

package main

// http://www.runoob.com/go/go-error-handling.html

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

func main() {
	// test_001()
	// test_0012()
	// test_002()
	// test_003()
	test_004()
}

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

// 普通错误处理
func test_001() {
	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

}

//----------------
/*
// Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
type error interface {
    Error() string
}

Go提供了两种创建一个实现了error interface的类型的变量实例的方法
errors.New("your first error code")
fmt.Errorf("error value is %d\n", 123)
*/

func test_0012() {
	fn := func(a int) (int, error) {
		return a + 10, errors.New("your first error code") // 内置第一种返回 error 实例
	}

	fn2 := func(a int) (int, error) {
		return a + 10, fmt.Errorf("error value is %d", a) // 内置第二种返回 error 实例
	}

	fn3 := func(a int) (int, error) {
		return a + 10, nil
	}

	if val, err := fn(10); err != nil {
		fmt.Print("error111: ", err, val, "\n")
	}
	if val, err := fn2(20); err != nil {
		fmt.Print("error222: ", err, val, "\n")
	}
	if val, err := fn3(30); err != nil {
		fmt.Print("error333: ", err, val, "\n")
	}
}

// ---------------
/*
正如名字一样，这个（recover）内建函数被用于从 panic 或 错误场景中恢复：让程序可以从 panicking 重新获得控制权，停止终止过程进而恢复正常执行。
recover 只能在 defer 修饰的函数（参见 6.4 节）中使用：用于取得 panic 调用中传递过来的错误值，如果是正常执行，调用 recover 会返回 nil，且没有其它效果。
总结：panic 会导致栈被展开直到 defer 修饰的 recover() 被调用或者程序中止。
下面例子中的 protect 函数调用函数参数 g 来保护调用者防止从 g 中抛出的运行时 panic，并展示 panic 中的信息：
*/

func protectExec(g func()) {
	defer func() {
		log.Println("done")
		// Println executes normally even if there is a panic
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	log.Println("start")
	g() //   possible runtime-error
}

// ---------------
//这是一个展示 panic，defer 和 recover 怎么结合使用的完整例子：
// 参考:http://wiki.jikexueyuan.com/project/the-way-to-go/13.3.html

// defer panic 相当于 try catch

func badCall() {
	// 1.
	// panic("bad end")

	// 2.
	b := 0
	a := 10 / b // 除 0
	_ = a
}

func myfunc() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panicing %s\r\n", err)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n") // 不会运行
}

func test_002() {
	fmt.Printf("Calling test\r\n")
	myfunc()
	fmt.Printf("Test completed\r\n")
}

/*
Calling test
Panicing bad end
Test completed
*/

// --------------- 自定义错误

// 自定义包中的错误处理和 panicking http://wiki.jikexueyuan.com/project/the-way-to-go/13.4.html

type IMyErr interface {
	Log(args ...interface{}) string
}

type CMyError struct {
	Age  int
	Name string
}

func (self *CMyError) Log(args ...interface{}) string { // 实现 IMyErr 的所有接口
	return fmt.Sprintf("--- is CMyError, name:%s\n", self.Name)
}

func myFunc222() {
	defer func() {
		if err := recover(); err != nil {
			myErr, ok := err.(IMyErr) // 动态匹配 IMyErr
			if !ok {
				fmt.Printf("--- unknown error\n")
			} else {
				fmt.Println("--- myErr:", myErr.Log())
			}
		}
	}()
	fmt.Printf("myFunc222 111\r\n")
	num, err := strconv.Atoi("aaa") // 报错
	if err != nil {
		panic(&CMyError{Age: 12, Name: "wolegequ"}) // 实例化 自定义错误 CMyError, 将在 recover() 中获取到这个 interface{} 参数
	}
	_ = num
	fmt.Printf("myFunc222 222\r\n")
}

func test_003() {
	fmt.Printf("Calling test\r\n")
	myFunc222()
	fmt.Printf("Test completed\r\n")
}

func test_004() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("--- error 111, msg:%v\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer func() { // panic 只能在对应的 goroutine 中捕获
			if err := recover(); err != nil {
				log.Printf("--- error 222, msg:%v\n", err)
			}
			wg.Done()
		}()

		time.Sleep(time.Second * 3)
		panic("--- err")
	}()

	go func() {
		log.Println("--- hello")
		time.Sleep(time.Second)
	}()

	log.Println("--- start")
	// wg.Wait()
	for {
	}
	log.Println("--- end")
}

/*
Calling test
myFunc222 111
--- is CMyError, name:wolegequ, err:strconv.Atoi: parsing "aaa": invalid syntax
Test completed
*/
