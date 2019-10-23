package test_reflect

import (
	"reflect"
	"testing"
)

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
	var m3 interface{}
	m2 = Msg{"aaa"}
	mt2 := reflect.TypeOf(m2)
	msgID2 := mt2.Name() // 值类型直接 name, 如果使用 .Elem() 会报错
	println("mt2:", mt2)
	println("msgID2:", msgID2) // msgID2: Msg

	m3 = &m2
	mt3 := reflect.TypeOf(m3)
	msgID3 := mt3.Elem().Name()
	println("mt3:", mt3)
	println("msgID3:", msgID3) // msgID3: // m3 是任意指针, 无法识别到具体类型
}

func Test_str2func(t *testing.T) {
	// https://blog.csdn.net/wowzai/article/details/9327405
}
