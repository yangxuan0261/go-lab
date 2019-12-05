package test_http

import (
	goprotobuf "GoLab/test_protobuf/proto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	proto "github.com/golang/protobuf/proto"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	reqBytes, _ := ioutil.ReadAll(req.Body)
	fmt.Fprintf(os.Stderr, "\n--- %s", reqBytes)
	// w.Write([]byte("Hello world"))

	// msg := &goprotobuf.HelloWorld{
	// 	Id:  proto.Int32(996),
	// 	Str: proto.String("what the fuck"),
	// }

	// buffer, _ := proto.Marshal(msg)
	// w.Write(buffer)

	// 上行解包
	pbReq := &goprotobuf.PBMessageRequest{}
	_ = proto.Unmarshal(reqBytes, pbReq)

	fmt.Printf("\n --- read Type: %d", *pbReq.Type)
	fmt.Printf("\n --- read Version: %s", *pbReq.Version)
	fmt.Printf("\n --- read Token: %s", *pbReq.Token)

	pbReq2 := &goprotobuf.PBStudentListReq{}
	_ = proto.Unmarshal(pbReq.MessageData, pbReq2)
	fmt.Printf("\n --- read Offset: %d", *pbReq2.Offset)
	fmt.Printf("\n --- read Limit: %d", *pbReq2.Limit)

	// 下行数据
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

func SayWorld(w http.ResponseWriter, req *http.Request) {
	reqBytes, _ := ioutil.ReadAll(req.Body)
	req.Header.Get("aaa")
	fmt.Printf("req body:%s, path:%s\n", string(reqBytes), req.URL.Path)

	w.Write([]byte("Hello world"))
}

func Test_Srv001(t *testing.T) {
	http.HandleFunc("/hello", SayHello)
	http.HandleFunc("/world", SayWorld)
	// http.HandleFunc("/hello_json", SayHello)
	http.ListenAndServe(":8001", nil)
	fmt.Println("---------------")
	// fmt.Fprintf("%s", "http://127.0.0.1:8001/hello")
}

func Test_SrvMux(t *testing.T) {
	m := &http.ServeMux{}
	m.HandleFunc("/world", SayWorld)
	http.ListenAndServe(":8001", m)
}

func Test_ReqGet(t *testing.T) {
	Url, err := url.Parse("http://baidu.com?fd=fdsf")
	if err != nil {
		panic(err.Error())
	}

	params := url.Values{}
	params.Set("a", "fdfds")
	params.Set("id", string("1"))
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	rsp, err := http.Get(urlPath)
	defer rsp.Body.Close()
	s, err := ioutil.ReadAll(rsp.Body)
	fmt.Println("--- rsp:", string(s))
}

type Server struct {
	ServerName string
	ServerIp   string
}

type ServerSlice struct {
	Server    []Server
	ServersID string
}

func Test_ReqPost(t *testing.T) {
	//post 第三个参数是io.reader interface
	//strings.NewReader  byte.NewReader bytes.NewBuffer  实现了read 方法
	s := ServerSlice{ServersID: "tearm", Server: []Server{{"beijing", "127.0.0.1"}, {"shanghai", "127.0.0.1"}}}
	b, _ := json.Marshal(s)
	fmt.Println(string(b))
	resp, _ := http.Post("http://baidu.com", "application/x-www-form-urlencoded", strings.NewReader("heel="+string(b)))
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("--- rsp:", string(body))

}

//func Test_ReqPost(t *testing.T) {
//
//	params := url.Values{}
//	params.Set("hello", "fdsfs") //这两种都可以
//	params = url.Values{"key": {"Value"}, "id": {"123"}}
//	resp, _ := http.PostForm("http://baidu.com",
//		body)
//
//	defer resp.Body.Close()
//	body, _ := ioutil.ReadAll(resp.Body)
//
//	fmt.Println(string(body))
//
//}

func Test_ReqPost002(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://baidu.com", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	//defer resp.Body.Close() // resp 可能为 nil

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println("--- rsp:", string(body))

}
