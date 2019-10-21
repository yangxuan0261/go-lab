package main

import (
	pb "GoLab/test_grpc/grpc_call/aaa"
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const port = ":50051"

//server 用于实现从proto 服务定义生成的 helloworld.GreeterServer接口.
type server struct{}

// SayHello 实现 helloworld.GreeterServer接口.
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("--- from cli, say:%s\n", in.Name)
	return &pb.HelloReply{Message: "hello " + in.Name}, nil
}

func (*server) SayBye(srv pb.Greeter_SayByeServer) error {
	log.Printf("-- srv SayBye begin\n")

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()

		for {
			data, err := srv.Recv()
			if err != nil {

				return
			}
			log.Printf("-- srv recv:%+v\n", data)
		}
	}()

	go func() {
		defer wg.Done()
		cnt := int32(1)
		for {
			srv.Send(&pb.HelloReply{Message: fmt.Sprintf("srv msg-%d", cnt)})
			cnt++
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("-- srv start:\n")

	//创建 gRPC 服务器，将我们实现的Greeter服务绑定到一个端口
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

// 生成 pb: protoc -I .\protos\ --go_out=plugins=grpc:./aaa .\protos/*.proto
// 参考
// - https://www.jianshu.com/p/85e9cfa16247
