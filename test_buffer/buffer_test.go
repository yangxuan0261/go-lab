package test_buffer

import (
	"bytes"
	"fmt"
	"testing"
)

// https://www.sunansheng.com/archives/25.html
// https://my.oschina.net/u/943306/blog/127981
// https://www.jianshu.com/p/92c174072c11

// 向Buffer中写入数据
func Test_001(t *testing.T) {
	newBytes := []byte(" go")
	//创建一个内容Learning的缓冲器
	buf := bytes.NewBuffer([]byte("Learning"))
	//将newBytes这个slice写到buf的尾部
	buf.Write(newBytes)
	fmt.Println(buf.String()) // Learning go
}

// 从Buffer中读取数据
func Test_002(t *testing.T) {
	bufs := bytes.NewBufferString("Learning swift.")
	fmt.Println("缓冲器：" + bufs.String()) // Learning swift.
	l := make([]byte, 5)
	bufs.Read(l) //把 bufs 的内容读入到l内,因为l容量为5,所以只读了5个过来
	fmt.Println("读取到的内容：" + string(l)) // Learn
	fmt.Println("缓冲器：" + bufs.String()) // ing swift. // 前面的被读走了
}

func Test_BufferString01(t *testing.T) {
	s := " world"
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String()) // hello
	buf.WriteString(s)        //将 s 这个 string 写到 buf 的尾部
	fmt.Println(buf.String()) // hello world
}
