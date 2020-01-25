package test_string

import (
	"bytes"
	"strings"
	"testing"
)

// 命令行执行: go test -run=xxx -bench=. -benchtime=3s

const (
	sss = "https://mojotv.cn"
	cnt = 10000
)

var (
	bbb      = []byte(sss)
	expected = strings.Repeat(sss, cnt)
)

//使用 提前初始化  内置 copy函数
func BenchmarkCopyPreAllocate(b *testing.B) {
	b.ReportAllocs()

	var result string
	for n := 0; n < b.N; n++ {
		bs := make([]byte, cnt*len(sss))
		bl := 0
		for i := 0; i < cnt; i++ {
			bl += copy(bs[bl:], sss)
		}
		result = string(bs)
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

//使用 提前初始化  内置append 函数
func BenchmarkAppendPreAllocate(b *testing.B) {
	b.ReportAllocs()

	var result string
	for n := 0; n < b.N; n++ {
		data := make([]byte, 0, cnt*len(sss))
		for i := 0; i < cnt; i++ {
			data = append(data, sss...)
		}
		result = string(data)
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

//使用 提前初始化 bytes.Buffer
func BenchmarkBufferPreAllocate(b *testing.B) {
	b.ReportAllocs()

	var result string
	for n := 0; n < b.N; n++ {
		buf := bytes.NewBuffer(make([]byte, 0, cnt*len(sss)))
		for i := 0; i < cnt; i++ {
			buf.WriteString(sss)
		}
		result = buf.String()
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

//使用 strings.Repeat 本质是pre allocate + strings.Builder
func BenchmarkStringRepeat(b *testing.B) {
	b.ReportAllocs()

	var result string
	for n := 0; n < b.N; n++ {
		result = strings.Repeat(sss, cnt)
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

//使用 内置copy
func BenchmarkCopy(b *testing.B) {
	b.ReportAllocs()

	var result string
	for n := 0; n < b.N; n++ {
		data := make([]byte, 0, 64) // same size as bootstrap array of bytes.Buffer
		for i := 0; i < cnt; i++ {
			off := len(data)
			if off+len(sss) > cap(data) {
				temp := make([]byte, 2*cap(data)+len(sss))
				copy(temp, data)
				data = temp
			}
			data = data[0 : off+len(sss)]
			copy(data[off:], sss)
		}
		result = string(data)
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

//使用 内置append
func BenchmarkAppend(b *testing.B) {
	b.ReportAllocs()

	var result string
	for n := 0; n < b.N; n++ {
		data := make([]byte, 0, 64)
		for i := 0; i < cnt; i++ {
			data = append(data, sss...)
		}
		result = string(data)
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

//使用 bytes.Buffer
func BenchmarkBufferWriteBytes(b *testing.B) {
	b.ReportAllocs()

	var result string
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		for i := 0; i < cnt; i++ {
			buf.Write(bbb)
		}
		result = buf.String()
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

//使用 strings.Builder write bytes
func BenchmarkStringBuilderWriteBytes(b *testing.B) {
	b.ReportAllocs()

	var result string
	for n := 0; n < b.N; n++ {
		var buf strings.Builder
		for i := 0; i < cnt; i++ {
			buf.Write(bbb)
		}
		result = buf.String()
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

//使用 string buffer write string
func BenchmarkBufferWriteString(b *testing.B) {
	b.ReportAllocs()

	var result string
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		for i := 0; i < cnt; i++ {
			buf.WriteString(sss)
		}
		result = buf.String()
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

// 使用string 加号
func BenchmarkStringPlusOperator(b *testing.B) {
	b.ReportAllocs()

	var result string
	for n := 0; n < b.N; n++ {
		var str string
		for i := 0; i < cnt; i++ {
			str += sss
		}
		result = str
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}

/*
结果:
F:\a_link_workspace\go\GoWinEnv_new\src\go_lab\test_string>go test -run=xxx -bench=. -benchtime=3s
goos: windows
goarch: amd64
pkg: go_lab/test_string
BenchmarkCopyPreAllocate-8                 49311             62112 ns/op          344065 B/op          2 allocs/op
BenchmarkAppendPreAllocate-8               56575             63727 ns/op          344065 B/op          2 allocs/op
BenchmarkBufferPreAllocate-8               41299             87250 ns/op          344065 B/op          2 allocs/op
BenchmarkStringRepeat-8                   257836             14030 ns/op          172032 B/op          1 allocs/op
BenchmarkCopy-8                            31524            110825 ns/op          862306 B/op         13 allocs/op
BenchmarkAppend-8                          30850            120036 ns/op         1046403 B/op         23 allocs/op
BenchmarkBufferWriteBytes-8                23904            149534 ns/op          862371 B/op         14 allocs/op
BenchmarkStringBuilderWriteBytes-8         26404            139114 ns/op          874467 B/op         24 allocs/op
BenchmarkBufferWriteString-8               25032            156461 ns/op          862371 B/op         14 allocs/op
BenchmarkStringPlusOperator-8                 51          68346584 ns/op        885204123 B/op     10032 allocs/op
PASS
ok      go_lab/test_string       45.104s

*/