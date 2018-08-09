package main

import "reflect"

func main() {
	test_001()
}

type Msg struct {
	name string
}

func test_001() {

	var m interface{}
	m = &Msg{"aaa"}
	mt := reflect.TypeOf(m)
	msgID := mt.Elem().Name()
	println("mt:", mt)
	println("msgID:", msgID)

	var m2 interface{}
	var m3 interface{}
	m2 = Msg{"aaa"}
	mt2 := reflect.TypeOf(m2)
	// msgID2 := mt2.Elem().Name()
	println("mt2:", mt2)
	// println("msgID2:", msgID2)

	m3 = &m2
	mt3 := reflect.TypeOf(m3)
	msgID3 := mt3.Elem().Name()
	println("mt3:", mt3)
	println("msgID3:", msgID3)
}
