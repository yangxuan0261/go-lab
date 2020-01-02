package test_map

import (
	"fmt"
	"sort"
	"sync"
	"testing"
)

func Test_mapBase(t *testing.T) {
	var countryCapitalMap map[string]string                                                // 声明, 但还未分配内存
	fmt.Printf("--- addr 111:%p, isnil:%v\n", countryCapitalMap, countryCapitalMap == nil) // 0x0, isnil:true
	countryCapitalMap = make(map[string]string)                                            // 分配内存
	fmt.Printf("--- addr 222:%p, isnil:%v\n", countryCapitalMap, countryCapitalMap == nil) // 0xc000058510, isnil:false

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country, val := range countryCapitalMap { // 遍历
		fmt.Println(country, countryCapitalMap[country])
		fmt.Println("val", val)
	}

	/*查看元素在集合中是否存在 */
	if captial, ok := countryCapitalMap["美国"]; ok {
		fmt.Println("美国的首都是", captial)
	} else {
		fmt.Println("美国的首都不存在")
	}

	if captial, ok := countryCapitalMap["Japan"]; ok {
		fmt.Println("Japan的首都是", captial)
	} else {
		fmt.Println("Japan的首都不存在")
	}
}

func Test_mapDelete(t *testing.T) {
	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	fmt.Println("\n--- 分割线 2")
	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	fmt.Println("\n--- 分割线 3")
	countryCapitalMap["China"] = "guangzhou"
	countryCapitalMap["China"] = "shenzhen" // 重复设置使用最后一个

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}

func Test_map03(t *testing.T) {

	fn := func(tm map[string]int) {
		tm["aaa"] = 666
	}

	cm := make(map[string]int)
	cm["aaa"] = 111
	cm["bbb"] = 222

	fn(cm)

	fmt.Println("cm", cm) // cm map[aaa:666 bbb:222] 说明是引用传递
	for key, val := range cm {
		fmt.Println(key, val)
	}
}

func Test_mapClear(t *testing.T) {
	cm := make(map[string]int)
	cm["aaa"] = 111
	cm["bbb"] = 222

	cm = make(map[string]int) // 重新分配一个
	//cm = nil // 或者 置为 nil 都可以让 内存被回收
	//或者遍历一遍 delete

	query := map[string]string{}

	query["test0"] = "0"
	query["test1"] = "1"
	query["test2"] = "2"

	for k, v := range query {
		delete(query, "test1") // 可以在遍历中删除 还没遍历到的元素
		fmt.Println(query, k, v)
	}
}

func Test_mapCop(t *testing.T) {
	cm1 := make(map[string]int)
	cm1["aaa"] = 111
	cm1["bbb"] = 222
	cm1["ccc"] = 333

	// 错误的拷贝
	cm2 := cm1
	fmt.Printf("--- cm1 addr:%p\n", cm1) // 0xc00005e510
	fmt.Printf("--- cm2 addr:%p\n", cm2) // 0xc00005e510, 地址相同, 指向同一快内存

	cm1["aaa"] = 666                                  // 修改 cm1 会影响到 cm2
	fmt.Printf("--- cm2 addr:%p, map:%v\n", cm2, cm2) // 0xc00005e510, map[aaa:666 bbb:222 ccc:333]

	// 正确的拷贝
	cm3 := make(map[string]int) // 分配新地址
	for k, v := range cm1 {
		cm3[k] = v
	}

	cm1["aaa"] = 777                                  // 修改 cm1 不会影响到 cm3
	fmt.Printf("--- cm3 addr:%p, map:%v\n", cm3, cm3) // 0xc00005e570, map[aaa:666 bbb:222 ccc:333]
}

type CDog struct {
	name string
	age  int
}

func (self *CDog) Run(speed int) {
	fmt.Printf("--- CDog.Run, name:%s, age:%d, speed:%d\n", self.name, self.age, speed)
}

func Test_value(t *testing.T) {
	dogMap := map[string]*CDog{ // key: string, value:*CDog (CDog指针)
		"xxx": &CDog{name: "xxx", age: 111}, // 初始化 map
	}
	dogMap["bbb"] = &CDog{name: "bbb", age: 456}
	dogMap["aaa"] = &CDog{name: "aaa", age: 123}
	dogMap["ccc"] = &CDog{name: "ccc", age: 789}
	for k, v := range dogMap {
		fmt.Println("------ key:", k)
		dogMap[k].Run(666)
		v.Run(777)
	}

	fmt.Println("--- len:", len(dogMap)) // --- len: 4
}

func Test_sort(t *testing.T) {
	dogMap := map[string]*CDog{ // key: string, value:*CDog (CDog指针)
		"xxx": &CDog{name: "xxx", age: 111}, // 初始化 map
	}
	dogMap["bbb"] = &CDog{name: "bbb", age: 456}
	dogMap["aaa"] = &CDog{name: "aaa", age: 123}
	dogMap["ccc"] = &CDog{name: "ccc", age: 789}

	keys := make([]string, len(dogMap)) // 遍历 map, 用数组装起来排序
	i := 0
	for k, _ := range dogMap {
		keys[i] = k
		i++
	}
	fmt.Printf("--- keys:%+v\n", keys)

	sort.Sort(sort.StringSlice(keys))
	fmt.Printf("--- keys sort:%+v\n", keys)
}

func Test_foreachEmpty(t *testing.T) {
	var m1 map[int]string

	fmt.Printf("--- isnil:%t\n", m1 == nil)

	for k, v := range m1 { // nil 也是可以直接遍历
		fmt.Println("--- kv:", k, v)
	}
}

/*
参考:
- Go 1.9 sync.Map揭秘 - https://colobu.com/2017/07/11/dive-into-sync-Map/#sync-Map%E7%9A%84%E6%80%A7%E8%83%BD
- 由浅入深聊聊Golang的sync.Map - https://juejin.im/post/5d36a7cbf265da1bb47da444
线程安全的 map: sync.Map
*/
func Test_syncMap(t *testing.T) {
	type userInfo struct {
		Name string
		Age  int
	}

	var m sync.Map

	vv, ok := m.LoadOrStore("1", "one")
	fmt.Println(vv, ok) //one false, ok 表示是否已经存在这个 key

	vv, ok = m.Load("1")
	fmt.Println(vv, ok) //one true

	vv, ok = m.LoadOrStore("1", "oneone")
	fmt.Println(vv, ok) //one true

	vv, ok = m.Load("1")
	fmt.Println(vv, ok) //one true

	m.Store("1", "oneone")
	vv, ok = m.Load("1")
	fmt.Println(vv, ok) // oneone true

	m.Store("2", "two")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

	m.Delete("1")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

	map1 := make(map[string]userInfo)
	var user1 userInfo
	user1.Name = "ChamPly"
	user1.Age = 24
	map1["user1"] = user1

	var user2 userInfo
	user2.Name = "Tom"
	user2.Age = 18
	m.Store("map_test", map1)

	mapValue, _ := m.Load("map_test")

	for k, v := range mapValue.(map[string]userInfo) {
		fmt.Println(k, v)
		fmt.Println("name:", v.Name)
	}
}

func Test_interfaceKey(t *testing.T) {
	aMap := make(map[interface{}]interface{})
	aMap[1] = "aaa"
	aMap[2] = "bbb"
	aMap["xxx"] = "ccc"

	var key1 interface{} = 2

	fmt.Printf("--- v:%v\n", aMap[2])     // v:bbb
	fmt.Printf("--- v:%v\n", aMap[key1])  // v:bbb
	fmt.Printf("--- v:%v\n", aMap["xxx"]) // v:ccc // 值类型的 key 可以直接使用

	d1 := new(CDog)
	d2 := new(CDog)
	aMap[d1] = "dog111"
	fmt.Printf("--- d1:%v\n", aMap[d1]) // d1:dog111 // 指针类型必须用对应的指针才能获取到
	fmt.Printf("--- d2:%v\n", aMap[d2]) // d2:<nil>
}
