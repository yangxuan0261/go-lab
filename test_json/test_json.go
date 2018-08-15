package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	test_001()
}

type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

func test_001() {
	st := &Student{
		Name:    "Xiao Ming",
		Age:     16,
		Guake:   true,
		Classes: []string{"Math", "English", "Chinese"},
		Price:   9.99,
	}

	strData, err := json.Marshal(st)
	if err == nil {
		fmt.Println("strData:", string(strData))
	}

	stb := &Student{}
	err = json.Unmarshal([]byte(strData), &stb)
	if err == nil {
		fmt.Println("stb:", stb)
	}
}
