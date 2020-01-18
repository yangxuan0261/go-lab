package main

// package test_file

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"testing"
)

func main() {
	// go Test_01(nil)
	go Test_02(nil)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	fmt.Printf("shutdown (signal: %v)\n", sig)
}

func Test_01(t *testing.T) {
	fmt.Println("hello world")
}

//获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			// ok := strings.HasSuffix(fi.Name(), ".go")
			// if ok {
			files = append(files, dirPth+PthSep+fi.Name())
			// }
		}
	}

	return files, dirs, nil
}

func Test_02(t *testing.T) {
	path := "C:/Users/Administrator/Desktop/testcpp"
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	path2 := "C:/Users/Administrator/Desktop/testcpp/hello.cpp"
	fmt.Println("--- filepath.Dir", filepath.Dir(path2))             // 获取该路径所在的父目录 C:\Users\Administrator\Desktop\testcpp
	fmt.Println("--- filepath.Ext", filepath.Ext(path2))             // 获取该路径的扩展名 .cpp
	fmt.Println("--- filepath.Base", filepath.Base(path2))           // 获取文件名 hello.cpp
	fmt.Println("--- filepath.FromSlash", filepath.FromSlash(path2)) // 转换路径 C:\Users\Administrator\Desktop\testcpp\hello.cpp
	fmt.Println("--- filepath.ToSlash", filepath.ToSlash(path2))     // 转换路径 C:/Users/Administrator/Desktop/testcpp/hello.cpp
}

//获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

func Test_03(t *testing.T) {
	dstDir := "C:/Users/Administrator/Desktop/testcpp"
	files, dirs, _ := GetFilesAndDirs(dstDir)

	fmt.Printf("获取的文件夹\n")
	for _, dir := range dirs {
		fmt.Printf("- %s\n", dir)
	}

	fmt.Printf("获取的文件\n")
	for _, file := range files {
		fmt.Printf("- %s\n", file)
	}

	for _, table := range dirs {
		temp, _, _ := GetFilesAndDirs(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	for _, table1 := range files {
		fmt.Printf("获取的文件为[%s]\n", table1)
	}

	fmt.Printf("=======================================\n")
	xfiles, _ := GetAllFiles("./simplemath")
	for _, file := range xfiles {
		fmt.Printf("获取的文件为[%s]\n", file)
	}
}

func Exists(path string) bool {
	info, err := os.Stat(path) //os.Stat获取文件信息
	fmt.Printf("--- isDir:%v, info:%+v\n", info.IsDir(), info)

	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func ExistsDir(path string) bool {
	info, err := os.Stat(path) //os.Stat获取文件信息
	if err == nil && info.IsDir() {
		return true
	}
	return false
}

func ExistsFile(path string) bool {
	info, err := os.Stat(path) //os.Stat获取文件信息
	if err == nil && !info.IsDir() {
		return true
	}
	return false
}

func Test_api(t *testing.T) {
	//fmt.Println("--- path:", os.Args[0]) // C:\Users\wolegequ\AppData\Local\Temp\___Test_api_in_GoLab_test_file.exe

	path1 := "F:/a_link_workspace/go/GoWinEnv_new/src/GoLab/test_file/"
	path2 := "F:/a_link_workspace/go/GoWinEnv_new/src/GoLab/test_file/file_test.go"
	fmt.Printf("--- exist: %v\n", ExistsDir(path1))
	fmt.Printf("--- filepath.Base: %v\n", filepath.Base(path2))
	fmt.Printf("--- filepath.Dir: %v\n", filepath.Dir(path2))
	fmt.Printf("--- filepath.Ext: %v\n", filepath.Ext(path2))
	fmt.Printf("--- filepath.Join: %v\n", filepath.Join(filepath.Dir(path2), filepath.Base(path2)))
}
