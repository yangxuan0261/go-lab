package main

import (
	"flag"
	"fmt"
	"testing"
)

/*
参考: http://qefee.com/2014/02/02/go%E8%AF%AD%E8%A8%80%E7%9A%84flag%E5%8C%85%E7%AE%80%E5%8D%95%E4%BD%BF%E7%94%A8%E6%95%99%E7%A8%8B/

需要指定 launch.json 中的参数为

   "program": "${workspaceRoot}/src/go_lab/test_flag/test_flag.go", // 指定入口文件
   "args": [
       "-conf=../../bin/conf/server.json",
       "--log",
       "../../bin/logs"
   ],
*/

func Test_001(t *testing.T) {
	confPath := flag.String("conf", "default path", "Server configuration file path")

	var log string
	flag.StringVar(&log, "log", "default name", "help msg for name") // 参数为指针

	flag.Parse() // 解析, 才能获取到参数

	fmt.Printf("confPath:%v, len:%v\n", *confPath, len(*confPath))
	fmt.Println("log:", log)

}
