package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/goinggo/mapstructure"
)

func main() {
	// test_json_struct()
	// test_json_map()
	// test_map_struct()
}

type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
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
	if err == nil {
		fmt.Println("strData:", string(strData))
	}

	var stb Student // 只需要声明就可以, 并不需要初始化
	fmt.Println("stb:", stb)
	err = json.Unmarshal([]byte(strData), &stb)
	if err == nil {
		fmt.Println("stb:", stb)
	}
}

func Test_json_map(t *testing.T) {
	fmt.Println("--- json to map")

	jsonStr := `
	{
		"user_name":"amy",
		"user_id":7,
		"user_age":18,
		"student":{"Name":"world","Age":456}
	}
`
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(mapResult)

	fmt.Println("--- map to json")
	mapInstances := map[string]interface{}{"user_name": "amy", "user_id": 7, "user_age": 18}
	mapInstances["student"] = Student{Name: "hello", Age: 123}

	jsonBytes, err := json.Marshal(mapInstances)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonBytes))
}

type Account struct {
	Name string `json:"user_name"`
	ID   int32  `json:"user_id"`
	Age  uint32 `json:"user_age"`
}

func Test_map_struct(t *testing.T) {
	fmt.Println("--- map to struct")
	mapInstances := make(map[string]interface{})
	mapInstances["Name"] = "amy"
	mapInstances["ID"] = 7
	mapInstances["Age"] = 18

	var account1 Account
	err := mapstructure.Decode(mapInstances, &account1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account1)

	fmt.Println("--- struct to map")
	account2 := Account{
		Name: "amy",
		ID:   007,
		Age:  18,
	}

	obj1 := reflect.TypeOf(account2)
	obj2 := reflect.ValueOf(account2)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}

	fmt.Println(data)
}
