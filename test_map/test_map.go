package test_map

// package main

import "fmt"

func main() {
	// test_map01()
	// test_map02()
	test_map03()
}

func test_map01() {
	var countryCapitalMap map[string]string /*创建集合 map[KeyType]ValueType*/
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country, val := range countryCapitalMap {
		fmt.Println(country, countryCapitalMap[country])
		fmt.Println("val", val)
	}

	/*查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["美国"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(captial) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("美国的首都是", captial)
	} else {
		fmt.Println("美国的首都不存在")
	}
}

func test_map02() {
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

// 用于测试会不会 复制拷贝 元素
func test_map03() {

	fn := func(tm map[string]int) {
		tm["aaa"] = 666
	}

	cm := make(map[string]int)
	cm["aaa"] = 111
	cm["bbb"] = 222

	fn(cm)

	fmt.Println("cm", cm) // cm map[aaa:666 bbb:222] 说明是引用传递, 不会复制拷贝元素
	for key, val := range cm {
		fmt.Println(key, val)
	}
}
