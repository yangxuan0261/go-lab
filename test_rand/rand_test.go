package test_rand

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_001(t *testing.T) {

	// 例如`rand.Intn`返回一个整型随机数n，0<=n<100
	fmt.Println("--- 111")
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))

	// `rand.Float64` 返回一个`float64` `f`,
	// `0.0 <= f < 1.0`
	fmt.Println("--- 222")
	fmt.Println(rand.Float64())

	// 这个方法可以用来生成其他数值范围内的随机数，
	// 例如`5.0 <= f < 10.0`
	fmt.Println("--- 333")
	fmt.Println((rand.Float64() * 5) + 5)
	fmt.Println((rand.Float64() * 5) + 5)

	// 为了使随机数生成器具有确定性，可以给它一个seed
	fmt.Println("--- 444")
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)

	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Intn(100))

	// 如果源使用一个和上面相同的seed，将生成一样的随机数
	fmt.Println("--- 555")
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Intn(100))
	fmt.Println()
}
