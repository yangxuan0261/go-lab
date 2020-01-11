package iterator_test

import (
	"fmt"
	"github.com/json-iterator/go"
	"testing"
)

// 高性能，100% 兼容的“encoding/json” 替代品 - https://github.com/json-iterator/go

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Student struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Guake   bool     `json:"guake"`
	Classes []string `json:"classes"`
	Price   float32  `json:"price"`
}

func Test_json_struct(t *testing.T) {
	st := &Student{
		Name:    "Xiao Ming",
		Age:     16,
		Guake:   true,
		Classes: []string{"Math", "English", "Chinese"},
		Price:   9.99,
	}

	strData, err := json.Marshal(st)
	strData2, err := json.MarshalIndent(st, "", "    ") // beauty
	if err == nil {
		fmt.Printf("--- stb string:%+v\n", string(strData))
		fmt.Printf("--- stb string22:%+v\n", string(strData2))
	}

	var stb Student
	err = json.Unmarshal(strData, &stb)
	if err == nil {
		fmt.Printf("--- stb Struct:%+v\n", stb)
	}
}
