package pkg001

import "fmt"

func SayHello() { // 外部访问需要首字母大写
	fmt.Println("hello world")
}

func getArea() { // 首字母小写只能内部访问
	const WIDTH int = 10
	const HEIGHT int = 20
	var area int
	area = WIDTH * HEIGHT
	fmt.Printf("面积为 : %d\n", area)

}
