package main

import (
	goprotobuf "GoLab/test_protobuf/proto"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	proto "github.com/golang/protobuf/proto"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	s, _ := ioutil.ReadAll(req.Body)
	fmt.Fprintf(os.Stderr, "--- %s", s)
	// w.Write([]byte("Hello world"))

	// msg := &goprotobuf.HelloWorld{
	// 	Id:  proto.Int32(996),
	// 	Str: proto.String("what the fuck"),
	// }

	// buffer, _ := proto.Marshal(msg)
	// w.Write(buffer)

	msg := &goprotobuf.PBStudentListRsp{
		List: []uint32{1, 2, 3},
	}
	data, _ := proto.Marshal(msg)

	// bufMd5 := make([]byte, 9)
	// copy(bufMd5, []byte("Hello aaa"))

	msg2 := &goprotobuf.PBMessageResponse{
		Type2:       proto.Uint32(123),
		MessageData: data,
		ResultCode:  proto.Uint32(456),
		ResultInfo:  proto.String("Hello bbb"),
	}
	buffer2, _ := proto.Marshal(msg2)
	w.Write(buffer2)
}

func main() {
	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":8001", nil)
	fmt.Println("---------------")
	// fmt.Fprintf("%s", "http://127.0.0.1:8001/hello")
}
