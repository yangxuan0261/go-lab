package test_pkg

import (
	"fmt"
	"test/test_pkg/pkg001"
	"test/test_pkg/pkg001/pkg002"
)

func main() {
	// test_001()
	pkg001.SayHello()
	// pkg001.getArea() // 报错, 访问外部只能访问 首字母大写 的方法

	pkg00234.SayHello222()

	d1 := &pkg00234.Dog{"123", 21}
	d2 := *d1
	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
}
