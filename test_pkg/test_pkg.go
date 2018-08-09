// package test_pkg

package main

import (
	"GoLab/test_pkg/pkg001"
	pkg002 "GoLab/test_pkg/pkg002"
	"fmt"
)

func init() {
	fmt.Println("--- init")
}

func main() {
	// test_001()
	pkg001.SayHello()
	// pkg001.getArea() // 报错, 访问外部只能访问 首字母大写 的方法

	pkg002.SayHello222()

	d1 := &pkg002.Dog{"123", 21}
	d2 := *d1
	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
}
