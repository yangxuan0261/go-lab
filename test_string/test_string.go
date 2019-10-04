package main

import (
	"fmt"
	"strings"
)

// https://studygolang.com/articles/5769

func test_001() {
	s := "Hello,世界!!!!!"
	n := strings.Count(s, "!")
	fmt.Println(n) // 5
	n = strings.Count(s, "!!!")
	fmt.Println(n) // 2
}

func test_002() {
	s := "Hello,世界!!!!!"
	b := strings.Contains(s, "!!")
	fmt.Println(b) // true
	b = strings.Contains(s, "!?")
	fmt.Println(b) // false
	b = strings.Contains(s, "")
	fmt.Println(b) // true
}

func test_003() {
	s := "Hello,世界!"
	b := strings.ContainsRune(s, '\n')
	fmt.Println(b) // false
	b = strings.ContainsRune(s, '界')
	fmt.Println(b) // true
	b = strings.ContainsRune(s, 0)
	fmt.Println(b) // false
}

func test_004() {
	s := "Hello,世界!"
	i := strings.Index(s, "h")
	fmt.Println(i) // -1
	i = strings.Index(s, "!")
	fmt.Println(i) // 12
	i = strings.Index(s, "")
	fmt.Println(i) // 0
}

func main() {
	// test_001()
	// test_002()
	// test_003()
	test_004()
}
