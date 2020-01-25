package main

import (
	pb "go-lab/test_grpc/grpc_call/aaa"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const port = ":50051"

//server 用于实现从proto 服务定义生成的 helloworld.GreeterServer接口.
type server struct{}

// SayHello 实现 helloworld.GreeterServer接口.
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("--- from cli, say:%s\n", in.Name)
	pr, _ := peer.FromContext(ctx)
	fmt.Printf("--- from cli, peer:%+v\n", pr)
	return &pb.HelloReply{Message: "hello " + in.Name}, nil
}

func (*server) SayBye(srv pb.Greeter_SayByeServer) error {
	ctx := srv.Context()

	// 设置头部信息, cli 可以获取到
	hdr := metadata.MD{}
	hdr["aaa"] = []string{"a-111", "a-222"}
	srv.SetHeader(hdr)

	log.Printf("--- srv SayBye begin\n") // 每一个连接进来都是一个独立的 connection

	// 获取客户端自定义信息
	headers, ok := metadata.FromIncomingContext(srv.Context())
	if ok {
		log.Printf("--- headers:%+v\n", headers)
	}

	// 获取客户端id
	pr, _ := peer.FromContext(ctx)
	addrSlice := strings.Split(pr.Addr.String(), ":")
	log.Printf("--- peer:%+v\n", pr)
	log.Printf("--- addSlice:%+v\n", addrSlice)
	if pr.Addr == net.Addr(nil) {
		log.Printf("--- pr.Addr is nil\n")
		return fmt.Errorf("getClientIP, peer.Addr is nil")
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()

		for {
			data, err := srv.Recv()
			if err != nil {
				// io.EOF
				log.Printf("-- srv Recv err:%+v\n", err)
				return
			}
			log.Printf("-- srv Recv:%+v\n", data)
		}
	}()

	go func() {
		defer wg.Done()
		cnt := int32(1)
		for {
			err := srv.Send(&pb.HelloReply{Message: fmt.Sprintf("srv msg-%d", cnt)})
			if err != nil {
				// io.EOF
				log.Printf("-- srv Send err:%+v\n", err)
				return
			}
			cnt++
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	log.Printf("--- srv SayBye end\n")
	return nil
}

// TLS认证
func getCreds() grpc.ServerOption {
	cert, err := tls.LoadX509KeyPair("../conf/server.pem", "../conf/server.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../conf/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	tlsCfg := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		ClientAuth:         tls.RequireAndVerifyClientCert,
		ClientCAs:          certPool,
		InsecureSkipVerify: false,
	}
	creds := credentials.NewTLS(tlsCfg)
	return grpc.Creds(creds)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("-- srv start:\n")

	//创建 gRPC 服务器，将我们实现的Greeter服务绑定到一个端口

	s := grpc.NewServer(
		getCreds(),
	)
	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil { // 会阻塞
		log.Fatalf("failed to server: %v", err)
	}
	log.Printf("-- srv start 222:\n")
}

// 生成 pb: protoc -I .\protos\ --go_out=plugins=grpc:./aaa .\protos/*.proto
// 参考
// - https://www.jianshu.com/p/85e9cfa16247
// 证书生成 参考: https://blog.csdn.net/yangxuan0261/article/details/102508827
