package cat

import (
	"sync"
	"testing"
)

/*
参考: https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter09/09.1.html#parallel-%E6%B5%8B%E8%AF%95
并行测试, 只要测试方法里加上 t.Parallel()
*/

var pairs = []struct {
	k string
	v string
}{
	{"polaris", " 徐新华 "},
	{"studygolang", "Go 语言中文网 "},
	{"stdlib", "Go 语言标准库 "},
	{"polaris1", " 徐新华 1"},
	{"studygolang1", "Go 语言中文网 1"},
	{"stdlib1", "Go 语言标准库 1"},
	{"polaris2", " 徐新华 2"},
	{"studygolang2", "Go 语言中文网 2"},
	{"stdlib2", "Go 语言标准库 2"},
	{"polaris3", " 徐新华 3"},
	{"studygolang3", "Go 语言中文网 3"},
	{"stdlib3", "Go 语言标准库 3"},
	{"polaris4", " 徐新华 4"},
	{"studygolang4", "Go 语言中文网 4"},
	{"stdlib4", "Go 语言标准库 4"},
}

var (
	data   = make(map[string]string)
	locker sync.RWMutex
)

func WriteToMap(k, v string) {
	locker.Lock()
	defer locker.Unlock()
	data[k] = v
}

func ReadFromMap(k string) string {
	locker.RLock()
	defer locker.RUnlock()
	return data[k]
}

func TestWriteToMap(t *testing.T) {
	t.Parallel()
	for i := 1; i < 10000; i++ {
		for _, tt := range pairs {
			WriteToMap(tt.k, tt.v)
		}
	}
}

func TestReadFromMap(t *testing.T) {
	t.Parallel()
	for i := 1; i < 10000; i++ {
		for _, tt := range pairs {
			actual := ReadFromMap(tt.k)
			if actual != tt.v {
				t.Errorf("the value of key(%s) is %s, expected: %s", tt.k, actual, tt.v)
			}
		}
	}
}

// 如果把 锁 相关的代码注释掉, 会报错: fatal error: concurrent map read and map write
