package base

// package main

import (
	"fmt"
	"math"
	"unsafe"
)

// 下面两种做法都只是为了引入包, 防止自动格式化时 引入包被删除
// import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

var _ = math.Inf

func main() {
	// test_default()
	// test001()
	// test_const_iota()
	// test_size()
	// test_defer()
	// test_switch()
	test_for()
	// test_arr()
	// test_funcPoint()
	// test_struct()
	// test_string()
}

func test_default() {
	var a int     // 0
	var b float32 // +0.000000e+000
	var c string  // ""
	var d []int   // [0/0]0x0
	var e *int    // 0x0

	if c == "" { // string 默认是空串
		println("c=\"\"")
	}

	if d == nil {
		println("d=nil")
	}

	if e == nil { // 指针默认是nil
		println("e=nil")
	}

	println(a, b, c, d, e)
}

// eee := 456 // 且只能在函数体中使用

//变量声明
func test001() {
	// 记住make只用于map，slice和channel，并且不返回指针。要获得一个显式的指针，使用new进行分配，或者显式地使用一个变量的地址。
	tm := make(map[string]int)
	println(tm)

	var aaa int32 // 第一种，指定变量类型，声明后若不赋值，使用默认值。
	aaa = 123

	var bbb = 456 // 第二种，根据值自行判定变量类型。

	ccc := 789 // 第三种，省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误。且只能在函数体中使用

	// var a string = "abc" // 尝试编译这段代码将得到错误 a declared and not used, 必须被使用

	fmt.Println(aaa, bbb, ccc)

	// 多变量声明
	var x, y int
	var c, d int = 1, 2
	var e, f = 123, "hello"
	a, b := 7, "abc"
	println(x, y, a, b, c, d, e, f)

	// 常量, 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
	const b001 string = "abc"                // 显式类型定义：
	const b002 = "abc"                       // 隐式类型定义：
	const b003, b004, b005 = 1, false, "str" //多重赋值
	println(b001, b002, b003, b004, b005)

	const ( // 常量还可以用作枚举：
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	println(Unknown, Female, Male)

	// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：
	const (
		Ea = "abc"
		Eb = len(Ea)
		Ec = unsafe.Sizeof(Ea) //字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。
	)
	println(Ea, Eb, Ec) //abc 3 16

}

func test_const_iota() {
	// iota，特殊常量，可以认为是一个可以被编译器修改的常量。

	// const (
	// 	a1 = iota
	// 	b1 = iota
	// 	c1 = iota
	// )

	// 等价于
	const (
		a1 = iota
		b1
		c1
	)

	const d1 = iota         // 0
	println(a1, b1, c1, d1) //0 1 2 0

	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		w        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, w) // 0 1 2 ha ha 100 100 7 8

	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
	// iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3。
}

func test_size() {
	var x = 123           // 隐式声明用 int64, 8
	var y int = 123       // 8
	var a int32 = 123     // 4
	var a1 *int32 = &a    // 8	// 指针的大小是 8
	var b = 123.3         // 隐式声明用 float64, 8
	var c float32 = 123.3 // 4
	var d = "asd"         // 16

	println("x:", unsafe.Sizeof(x))   // 8
	println("y:", unsafe.Sizeof(y))   // 8
	println("a:", unsafe.Sizeof(a))   // 4
	println("a1:", unsafe.Sizeof(a1)) // 8
	println("b:", unsafe.Sizeof(b))   // 8
	println("c:", unsafe.Sizeof(c))   // 4
	println("d:", unsafe.Sizeof(d))   // 16

	// println(" y == a", y == a) // 运行时报错,  (mismatched types int and int32)
}

func test_defer() {
	/*
		defer 的思想类似于C++中的析构函数，不过Go语言中“析构”的不是对象，而是函数，defer就是用来添加函数结束时执行的语句。注意这里强调的是添加，而不是指定，因为不同于C++中的析构函数是静态的，Go中的defer是动态的。
		defer 中使用匿名函数依然是一个闭包。
	*/
	x, y := 1, 2

	defer func(a int) {
		fmt.Printf("x:%d,y:%d\n", a, y) // y 为闭包引用
	}(x) // 复制 x 的值

	x += 100
	y += 100
	fmt.Println(x, y)

	// 101 102
	// x:1,y:102
}

func test_switch() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		fmt.Println("--- match 111")
		grade = "A"
	case 80: // 如果这里 再次 case 90 会编译报错, 所以 go 里面不需要 break
		fmt.Println("--- match 222")
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)

	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}

func test_for() {
	/* for 循环 */
	for index := 0; index < 3; index++ { // 和 C 语言的 for 一样：
		fmt.Printf("index=%d\n", index)
	}

	fmt.Println("\n--- 分割线 1")
	var b int = 4
	var a int = 0
	for a < b { // 和 C 的 while 一样：
		a++
		fmt.Printf("a=%d. b=%d\n", a, b)
	}

	fmt.Println("\n--- 分割线 2")
	// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下
	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		if i == 5 {
			break
			// return

		}
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	fmt.Println("\n--- 分割线 3")
	// for { // 和 C 的 for(;;) 一样,
	// 	fmt.Println("hello")
	// }
}

func test_arr() {
	balance1 := [7]float32{1000.0, 2.0} // 指定长度
	fmt.Println("len=", len(balance1))
	for i, v := range balance1 {
		fmt.Printf("balance1[%d]=%f\n", i, v)
	}

	fmt.Println("\n--- 分割线 1")
	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0} // 动态长度
	fmt.Println("len=", len(balance2))
	for i, v := range balance2 {
		fmt.Printf("balance2[%d]=%f\n", i, v)
	}
}

// 函数指针
type FnAdd func(int, int) (int, string)

func test_funcPoint() {
	fn := func(a int, b int) (int, string) {
		return a + b, "hello"
	}
	var myfunc FnAdd = fn
	total, flag := myfunc(2, 3)
	fmt.Print(total, " ", flag, "\n")

	fn2 := func(a int, b int) (r1 int, r2 string) {
		r1 = a + b
		r2 = "world"
		return
	}
	var myfunc2 FnAdd = fn2
	total2, flag2 := myfunc2(4, 5)
	fmt.Print(total2, " ", flag2, "\n")
	fmt.Print(fn2, "\n")
}

// 结构体
type User struct {
	name string `defaultName` // `` 里面的是注释内容, 并不是默认值
	age  int8
}

// 一般会专门起一个方法来New对象
func NewUser() *User {
	return &User{
		name: "defaultName",
		age:  10,
	}
}

func test_struct() {
	u1 := &User{
		name: "aaa",
		age:  12,
	}
	fmt.Print(u1, "\n")
	fmt.Printf("addr:%p\n", u1)

	var u2 *User = &User{
		name: "bbb",
		age:  12,
	}
	fmt.Print(u2, "\n")
	fmt.Printf("addr:%p\n", u2)

	u3 := NewUser()
	fmt.Print(u3, "\n")
	fmt.Printf("addr:%p\n", u3)
}

func test_string() {

}
