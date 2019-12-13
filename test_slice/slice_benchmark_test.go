package test_slice

import (
	"strconv"
	"testing"
)

var length = 1000
var maps map[string]string
var slices []string
var arrays [1000]string

func init() {
	maps = make(map[string]string, length)
	slices = make([]string, length)
	for i := 0; i < length; i++ {
		maps[strconv.Itoa(i)] = "abc"
		slices[i] = "abc"
		arrays[i] = "abc"
	}

}

func BenchmarkIterateMap(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, _ = range maps {
		}
	}
}

func BenchmarkIterateSlices(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, _ = range slices {
		}
	}
}

func BenchmarkIterateArrays(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, _ = range arrays {
		}
	}
}

func BenchmarkIterateMapF(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, k := range maps {
			_ = k
		}
	}
}

func BenchmarkIterateSlicesF(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, k := range slices {
			_ = k
		}
	}
}

func BenchmarkIterateArraysF(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, k := range arrays {
			_ = k
		}
	}
}

func BenchmarkIterateSlicesFor(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			_ = slices[j]
		}
	}
}

func BenchmarkIterateArraysFor(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for j := 0; j < length; j++ {
			_ = arrays[j]
		}
	}
}

/*
BenchmarkIterateMap-8         	   97801	     12156 ns/op	       0 B/op	       0 allocs/op
BenchmarkIterateSlices-8      	 4557112	       264 ns/op	       0 B/op	       0 allocs/op
BenchmarkIterateArrays-8      	 3893553	       265 ns/op	       0 B/op	       0 allocs/op
BenchmarkIterateMapF-8        	   98648	     12132 ns/op	       0 B/op	       0 allocs/op
BenchmarkIterateSlicesF-8     	 4540880	       263 ns/op	       0 B/op	       0 allocs/op
BenchmarkIterateArraysF-8     	 2554587	       501 ns/op	       0 B/op	       0 allocs/op
BenchmarkIterateSlicesFor-8   	 2804695	       451 ns/op	       0 B/op	       0 allocs/op
BenchmarkIterateArraysFor-8   	 2278812	       549 ns/op	       0 B/op	       0 allocs/op
*/
