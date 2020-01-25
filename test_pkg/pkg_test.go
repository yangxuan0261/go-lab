package test_pkg

import (
	"go_lab/test_pkg/pkg001"
	pkg002 "go_lab/test_pkg/pkg002"
	"fmt"
	"testing"
)

func init() {
	fmt.Println("--- init pkg_test")
}

/*
需要指定 launch.json 中的参数为

"program": "${workspaceRoot}/src/go_lab/test_pkg/test_pkg.go", // 指定入口文件

*/

func Test_main(t *testing.T) {
	fmt.Println("Test_main")

	// test_001()
	pkg001.SayHello()
	// pkg001.getArea() // 报错, 访问外部只能访问 首字母大写 的方法

	pkg002.SayHello222()

	d1 := &pkg002.Dog{Name: "123", Age: 21}
	d2 := *d1
	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)

	/*
		--- init testpkg01
		--- init testpkg02-2
		--- init testpkg02-3
		--- init testpkg02 // 同 package 下执行的 init 顺序不同, 如果有顺序要求, 正确的做法是同 package 下只有一个 init 函数
		--- init pkg_test
		=== RUN   Test_main // 先执行所有 go 文件的 init 方法
		Test_main
		hello world
		hello world222
		d1: &{123 21}
		d2: {123 21}
	*/
}
