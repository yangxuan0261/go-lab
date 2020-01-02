package test_base

import (
	"fmt"
	"log"
	"testing"
)

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

func Test_DeferReturn(t *testing.T) {
	num := 1
	fn1 := func() int {
		defer func() {
			num++
		}()
		return num
	}

	n1 := fn1()
	fmt.Printf("--- n1:%d\n", n1) // n1:1
	fmt.Printf("--- num:%d\n", num)
}

func f1() (r int) {
	defer func() { r++ }()
	return 0
}

/* 结果: 1
func f1()(r int){ // 1.赋值r=0 // 2.闭包引用，返回值被修改 defer func(){r++}() // 3.空的 return}
*/

func f2() (r int) {
	t := 5
	defer func() { t = t + 5 }()
	return t
}

/* 结果: 5
func f2()(r int){t:=5 // 1.赋值r=t // 2.闭包引用，但是没有修改 返回值r defer func(){t=t+5}() // 3.空的 return }
*/

func f3() (r int) {
	defer func(r int) { r = r + 5 }(r)
	return 1
}

/* 结果: 1
func f3()(r int){ // 1.赋值r=1 // 2.r作为函数参数，不会修改要返回的那个r值 defer func(rint){r=r+5}(r) // 3.空的 return}*/
/*
采坑点
使用 defer 最容易采坑的地方是和带命名返回参数的函数一起使用时。

defer 语句定义时，对外部变量的引用是有两种方式的，分别是作为函数参数和作为闭包引用。作为函数参数，则在 defer 定义时就把值传递给 defer，并被缓存起来；作为闭包引用的话，则会在 defer 函数真正调用时根据整个上下文确定当前的值。

避免掉坑的关键是要理解这条语句：

return xxx
这条语句并不是一个原子指令，经过编译之后，变成了三条指令：

1.返回值=xxx
2.调用defer函数
3.空的return

1,3 步才是 return 语句真正的命令，第 2 步是 defer 定义的语句，这里就有可能会操作返回值。
*/

func Test_DeferReturn02(t *testing.T) {
	fmt.Printf("--- f1:%d\n", f1()) // f1:1
	fmt.Printf("--- f2:%d\n", f2()) // f2:5
	fmt.Printf("--- f3:%d\n", f3()) // f3:1
}
