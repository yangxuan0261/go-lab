package test_protobuf

import (
	goprotobuf "go_lab/test_protobuf/proto"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	proto "github.com/golang/protobuf/proto"
)

func Test_main(t *testing.T) {
	pwrite()
	time.Sleep(time.Second * 1)
	pread()
}

func pwrite() {
	msg := &goprotobuf.HelloWorld{
		Id:  proto.Int32(996),
		Str: proto.String("what the fuck"),
	}

	path := string("d:/test.txt")
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}

	defer f.Close()
	buffer, err := proto.Marshal(msg) // 必须是
	f.Write(buffer)
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("--- has err:", err.Error())
		os.Exit(-1)
	}
}

func pread() {

	path := string("d:/test.txt")
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}

	defer file.Close()
	fi, err := file.Stat()
	CheckError(err)
	buffer := make([]byte, fi.Size())
	_, err = io.ReadFull(file, buffer)
	CheckError(err)

	msg := &goprotobuf.HelloWorld{}
	err = proto.Unmarshal(buffer, msg)
	CheckError(err)

	fmt.Printf("read: %s\n", msg.String())
}
