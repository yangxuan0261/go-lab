package main

import (
	pro "GoLab/test_grpc/grpc_stream_putget/proto"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type server struct {
}

//服务端 单向流
func (s *server) GetStream(req *pro.StreamReqData, res pro.Greeter_GetStreamServer) error {
	i := 0
	for {
		log.Println("--- srv GetStream req data:", req.Data) // req 没有定义为 stream, 可以直接 .xxx 获取属性值
		i++
		res.Send(&pro.StreamResData{Data: "--- srv GetStream"})
		time.Sleep(1 * time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

//客户端 单向流
func (this *server) PutStream(cliStr pro.Greeter_PutStreamServer) error {

	for {
		if data, err := cliStr.Recv(); err == nil {
			log.Println("--- srv PutStream recv:", data)
		} else {
			log.Println("break, err :", err)
			break
		}
	}

	return nil
}

//客户端服务端 双向流
func (this *server) AllStream(allStr pro.Greeter_AllStreamServer) error {

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Println("--- srv AllStream recv:", data.Data)
		}
		wg.Done()
	}()

	go func() {
		for {
			allStr.Send(&pro.StreamResData{Data: "--- srv data"})
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}

func main() {
	//监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return
	}
	//创建一个grpc 服务器
	s := grpc.NewServer()
	//注册事件
	pro.RegisterGreeterServer(s, &server{})
	//处理链接
	s.Serve(lis)
}
