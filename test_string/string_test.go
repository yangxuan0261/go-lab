package test_string

import (
	"fmt"
	"strings"
	"testing"
)

// https://studygolang.com/articles/5769

func Test_001(t *testing.T) {
	s := "Hello,世界!!!!!"
	n := strings.Count(s, "!")
	fmt.Println(n) // 5
	n = strings.Count(s, "!!!")
	fmt.Println(n) // 2
}

func Test_002(t *testing.T) {
	s := "Hello,世界!!!!!"
	b := strings.Contains(s, "!!")
	fmt.Println(b) // true
	b = strings.Contains(s, "!?")
	fmt.Println(b) // false
	b = strings.Contains(s, "")
	fmt.Println(b) // true
}

func Test_003(t *testing.T) {
	s := "Hello,世界!"
	b := strings.ContainsRune(s, '\n')
	fmt.Println(b) // false
	b = strings.ContainsRune(s, '界')
	fmt.Println(b) // true
	b = strings.ContainsRune(s, 0)
	fmt.Println(b) // false
}

func Test_004(t *testing.T) {
	s := "Hello,世界!"
	i := strings.Index(s, "h")
	fmt.Println(i) // -1
	i = strings.Index(s, "!")
	fmt.Println(i) // 12
	i = strings.Index(s, "")
	fmt.Println(i) // 0
}
