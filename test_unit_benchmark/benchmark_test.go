package cat

import (
	"fmt"
	"sync/atomic"
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
	//fmt.Printf("--- CDog.Walk, name:%s, age:%d\n", d.Name, d.Age)
}

var dg IActor = &CDog{
	Name: "Tom",
	Age:  123,
}

func Benchmark_Cast(b *testing.B) {
	b.ReportAllocs() // 在report中包含内存分配信息，例如结果是:
	// Benchmark_Cast-3        1000000000               0.567 ns/op           0 B/op          0 allocs/op

	for i := 0; i < b.N; i++ { // b.N, 次数

		// 需要测试性能的接口
		d := CDog{}
		d.Walk(123)

		var dgIns *CDog
		dgIns = dg.(*CDog)
		_ = dgIns

	}
}

// 测试并发效率
func Benchmark_Parallel(b *testing.B) {
	b.ReportAllocs()

	cnt := uint64(1)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fmt.Println("--- cnt:", atomic.AddUint64(&cnt, 1))
			// 需要测试性能的接口
			dgIns := dg.(*CDog)
			//dgIns.Walk(123)
			_ = dgIns

		}
	})
}

func Benchmark_ResetTimer(b *testing.B) {
	b.ReportAllocs()

	d := CDog{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Walk(123)
	}
}
