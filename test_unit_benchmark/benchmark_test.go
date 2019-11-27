package cat

import (
	"fmt"
	"testing"
)

type IActor interface {
	Walk(speed int32)
}

type CDog struct {
	Name string
	Age  int32
}

func (d *CDog) Walk(speed int32) {
	fmt.Printf("--- CDog.Walk, name:%s, age:%d\n", d.Name, d.Age)
}

var dg IActor = &CDog{
	Name: "Tom",
	Age:  123,
}

func Benchmark_Cast(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N, 次数

		// 需要测试性能的接口
		var dgIns *CDog
		dgIns = dg.(*CDog)
		_ = dgIns

	}
}

// 测试并发效率
func Benchmark_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

			// 需要测试性能的接口
			dgIns := dg.(*CDog)
			//dgIns.Walk(123)
			_ = dgIns

		}
	})
}
