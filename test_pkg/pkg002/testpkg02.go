package pkg00234

import "fmt"

type Dog struct {
	Name string // 外部访问需要首字母大写
	Age  int8   // 首字母小写只能内部访问
}

func SayHello222() { // 外部访问需要首字母大写
	fmt.Println("hello world222")
}

func getArea222() { // 首字母小写只能内部访问
	const WIDTH int = 10
	const HEIGHT int = 20
	var area int
	area = WIDTH * HEIGHT
	fmt.Printf("面积为222 : %d\n", area)

}
