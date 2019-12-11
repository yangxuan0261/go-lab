package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/goinggo/mapstructure"
)

func main() {
	// test_json_struct()
	// test_json_map()
	// test_map_struct()
}

// 结构体字段必须 首字母大写 (public)
type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

// 继承扩展字段
type StudentExt struct {
	Student
	Speed int
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
		fmt.Printf("--- stb string:%+v\n", string(strData))
	}

	var stb Student
	err = json.Unmarshal([]byte(strData), &stb)
	if err == nil {
		fmt.Printf("--- stb Struct:%+v\n", stb)
	}

	println()
	var stbExt StudentExt
	extStr := `{"Name":"Xiao Ming","Age":16,"Guake":true,"Classes":["Math","English","Chinese"],"Price":9.99,"Speed":123}`
	err = json.Unmarshal([]byte(extStr), &stbExt)
	if err == nil {
		fmt.Printf("--- stbExt Struct:%+v\n", stbExt)
	}

	strData, err = json.Marshal(stbExt)
	if err == nil {
		fmt.Printf("--- stbExt string:%+v\n", string(strData))
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
	Name string `json:"user_name"` // 这样可以映射为 Marshal 后的 string 的 key 值, 反之可以 Unmarshal 成对应的属性字段
	ID   int32  `json:"user_id"`
	Age  uint32 `json:"user_age"`
	Flag bool   `json:"user_age"`
	Arr  []string
}

func Test_map_struct(t *testing.T) {
	mapInstances := make(map[string]interface{})
	mapInstances["Name"] = "amy"
	mapInstances["ID"] = 7
	mapInstances["Age"] = 18
	mapInstances["Flag"] = true
	mapInstances["Arr"] = []string{"hello", "world"}

	var account1 Account
	err := mapstructure.Decode(mapInstances, &account1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("--- map2struct, account1:%+v\n", account1)

	account2 := Account{
		Name: "amy",
		ID:   007,
		Age:  18,
		Flag: false,
	}

	obj1 := reflect.TypeOf(account2)
	obj2 := reflect.ValueOf(account2)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	fmt.Printf("--- struct2map, data:%+v\n", data)
}

type Node struct {
	Id   uint32 // 1001001, 节点唯一id 命名, 1[类型][机子序列]
	Name string
	Meta map[string]string
}

type ImNode struct {
	Node
	Descr string
}

func Test_beautifyJson(t *testing.T) {
	jfile := "./temp_aaa.json"

	in := &ImNode{
		Node: Node{
			Id:   uint32(123),
			Name: "imnode",
			Meta: map[string]string{
				"key-111": "val-111",
				"key-222": "val-222",
			},
		},
		Descr: "hello",
	}

	bytes, err := json.MarshalIndent(in, "", "    ") // 带缩进
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(jfile, bytes, os.ModePerm)
}

func Test_parseArr(t *testing.T) {
	var arr1 []*Account
	arr1 = append(arr1, &Account{ID: 1}, &Account{ID: 2}, &Account{ID: 3})

	//bts, _ := json.Marshal(arr1) // 这里可以存 数组 或者 数组的地址
	bts, _ := json.Marshal(&arr1) // 这里可以存 数组 或者 数组的地址
	fmt.Printf("--- str:%s\n", string(bts))

	var arr2 []*Account                // 可以直接用声明的数组
	json.Unmarshal([]byte(bts), &arr2) // 这里一定要传数组的地址, 否则解码失败
	for k, v := range arr2 {
		fmt.Printf("--- k:%+v, v:%+v\n", k, v)
	}

	/*
		--- str:[{"user_name":"","user_id":1,"Arr":null},{"user_name":"","user_id":2,"Arr":null},{"user_name":"","user_id":3,"Arr":null}]
		--- k:0, v:&{Name: ID:1 Age:0 Flag:false Arr:[]}
		--- k:1, v:&{Name: ID:2 Age:0 Flag:false Arr:[]}
		--- k:2, v:&{Name: ID:3 Age:0 Flag:false Arr:[]}
	*/
}
