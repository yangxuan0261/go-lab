package cat

import (
	"fmt"
	"testing"
	"time"
)

// 参考: https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter09/09.1.html#%E6%8A%A5%E5%91%8A%E6%96%B9%E6%B3%95

func Test_Log(t *testing.T) {
	fmt.Printf("Test_Log 1\n")
	t.Log("aaa")
	t.Logf("aaa:%d", 123)
	/* 带行号输出日志
		unit_test.go:12: aaa
	    unit_test.go:13: aaa:123
	*/
}

func Test_Fail(t *testing.T) {
	fmt.Printf("Test_Fail 1\n")
	t.Fail() // 会标记为 测试失败, 但不会中断测试
	fmt.Printf("Test_Fail 2\n")
	t.FailNow() // 会标记为 测试失败, 同时中断测试
	fmt.Printf("Test_Fail 3\n")

}

func Test_Error(t *testing.T) {
	fmt.Printf("Test_Error 1\n")
	t.Error("bbb") // 等价于 Log + Fail, 会标记为 测试失败, 但不会中断测试
	t.Errorf("bbb:%d", 123)
	fmt.Printf("Test_Error 2\n")
}

func Test_Fatal(t *testing.T) {
	fmt.Printf("Test_Fatal 1\n")
	t.Fatal("ccc") // 等价于 Log + FailNow, 会标记为 测试失败, 同时中断测试
	t.Fatalf("ccc:%d", 123)
	fmt.Printf("Test_Fatal 2\n")
}

func Test_SkipNow(t *testing.T) {
	fmt.Printf("Test_SkipNow 1\n")
	t.SkipNow() // 会中断测试, 但不会标记为 测试失败
	fmt.Printf("Test_SkipNow 2\n")
}

func Test_SubTest(t *testing.T) {

	// 多个子测试
	t.Run("subtest001", func(t *testing.T) { // 队列, 先见先出 执行
		fmt.Println("--- fn 111 - 1:", time.Now().Unix())
		time.Sleep(time.Second * 2)
		fmt.Println("--- fn 111 - 2:", time.Now().Unix())
	})

	t.Run("subtest002", func(t *testing.T) {
		fmt.Println("--- fn 222 - 1:", time.Now().Unix())
		time.Sleep(time.Second * 3)
		fmt.Println("--- fn 222 - 2:", time.Now().Unix())
	})

	t.Run("subtest003", Test_Log) // 也可以添加

	t.Cleanup(func() { // 栈, 后进先出 执行
		fmt.Println("Cleaning Up! 111")
	})

	t.Cleanup(func() {
		fmt.Println("Cleaning Up! 222")
	})

	/*
		Cleaning Up! 222
		Cleaning Up! 111
	*/
}
