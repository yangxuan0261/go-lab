package test_io

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func Test_write01(t *testing.T) {
	txt := "hello\n"
	path := "./temp_aaa.json"
	if err := ioutil.WriteFile(path, []byte(txt), os.ModePerm); err != nil {
		panic(err)
	} else {
		fmt.Println("--- success")
	}
	// ModePerm 0777 覆盖写入, ioutil 貌似没有追加, 追加参考: Test_writeAppend
	// 0644 也可以
}

func Test_read01(t *testing.T) {
	path := "./temp_aaa.json"
	if bts, err := ioutil.ReadFile(path); err != nil {
		panic(err)
	} else {
		fmt.Println("--- success, txt:", string(bts))
	}
}

func Test_writeAppend(t *testing.T) {
	path := "./temp_aaa.json"
	data := []byte("world\n")
	fl, err1 := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0644)
	defer fl.Close()

	if err1 != nil {
		panic(err1)
	}

	n, err2 := fl.Write(data)
	if err2 == nil && n < len(data) {
		err2 = io.ErrShortWrite
		panic(err2)
	}
}
