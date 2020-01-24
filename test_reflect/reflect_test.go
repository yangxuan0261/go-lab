package test_reflect

import (
	"fmt"
	"reflect"
	"testing"
)

// Go Reflect 性能 - https://colobu.com/2019/01/29/go-reflect-performance/
// 优化, 可以通过 自定义的生成器脚本, 生成, 避免使用反射

type Msg struct {
	name string
}

func Test_001(t *testing.T) {

	var m interface{}
	m = &Msg{"aaa"}
	mt := reflect.TypeOf(m)
	msgID := mt.Elem().Name() // 指针类型需要 .Elem(), .Elem() 是指针所指向的 对象
	println("mt:", mt)
	println("msgID:", msgID) //msgID: Msg

	var m2 interface{}
	m2 = Msg{"aaa"}
	mt2 := reflect.TypeOf(m2)
	msgID2 := mt2.Name() // 值类型直接 name, 如果使用 .Elem() 会报错
	println("mt2:", mt2)
	println("msgID2:", msgID2) // msgID2: Msg

	var m3 interface{}
	m3 = &m2
	mt3 := reflect.TypeOf(m3)
	msgID3 := mt3.Elem().Name()
	println("mt3:", mt3)
	println("msgID3:", msgID3) // msgID3: // m3 是任意指针, 无法识别到具体类型
}

func Test_str2func(t *testing.T) {
	// https://blog.csdn.net/wowzai/article/details/9327405
}

type Foo struct {
	Name string
}

type Bar struct {
	Name string
}

func Test_type2Instance(t *testing.T) {
	var regStruct map[string]interface{}
	regStruct = make(map[string]interface{})
	regStruct["Foo"] = Foo{}
	regStruct["Bar"] = Bar{}
	for k, v := range regStruct {
		fmt.Printf("--- k:%v, v:%p\n", k, &v)
	}

	str := "Bar"
	if regStruct[str] != nil {
		t := reflect.ValueOf(regStruct[str]).Type()
		v := reflect.New(t).Elem()
		fmt.Printf("--- aaa:%p\n", &v)
	}
}

// 根据类型实例化对象, 需要浪费一个对象的空间, 错误姿势
func Test_type2Instance2(t *testing.T) {
	type T = Bar
	data := T{}
	ptr := reflect.New(reflect.TypeOf(data)).Interface() // 指针对象
	//obj := reflect.New(reflect.TypeOf(data)).Elem().Interface() // 指针 指向的对象

	if ins, ok := ptr.(*Bar); ok {
		ins.Name = "hello"
		fmt.Printf("--- ins:%+v\n", ins) // --- ins:&{Name:hello}
	} else {
		t.Error("--- type error")
	}
}

// 根据类型实例化对象, 不需要浪费一个对象的空间, 正确姿势
/*
灵感来自 protobuf 生产的代码. proto.RegisterType((*Skin)(nil), "datacfg.skin")
*/
func Test_type2Instance3(t *testing.T) {
	x := (*Bar)(nil)
	fmt.Printf("--- x:%+v\n", x) // --- x:<nil>

	t1 := reflect.TypeOf(x)
	fmt.Printf("--- t1:%+v\n", t1) // --- t1:*test_reflect.Bar

	v1 := reflect.ValueOf(x)
	fmt.Printf("--- v1:%+v\n", v1)          // --- v1:<nil>
	fmt.Printf("--- Kind:%+v\n", v1.Kind()) // --- Kind:ptr, 居然可以答应成字符串
	fmt.Printf("--- Pointer:%+v\n", v1.Pointer())
	fmt.Printf("--- t1:%+v\n", v1.Kind() == reflect.Ptr)

	println()
	name1 := reflect.Zero(t1).Interface().(*Bar)
	fmt.Printf("--- name1:%+v\n", name1) // --- name1:<nil>

	println()
	ptr := reflect.New(t1.Elem()).Interface() // 指针对象, t1 是个指针类型, 所以实例化时要用
	fmt.Printf("--- ptr:%+v\n", ptr)          // --- t2:*test_reflect.Bar

	t2 := reflect.TypeOf(ptr)
	fmt.Printf("--- t2:%+v\n", t2) // --- t2:*test_reflect.Bar

	if ins, ok := ptr.(*Bar); ok {
		ins.Name = "hello"
		fmt.Printf("--- ins:%+v\n", ins) // --- ins:&{Name:hello}
	} else {
		t.Error("--- type error")
	}

	println()
	// 分装成一个方法
	instanceFn := func(x interface{}) interface{} {
		return reflect.New(reflect.TypeOf(x).Elem()).Interface()
	}

	itf1 := instanceFn((*Bar)(nil))
	if ins, ok := itf1.(*Bar); ok {
		ins.Name = "world"
		fmt.Printf("--- ins:%+v\n", ins)
	} else {
	}
}
