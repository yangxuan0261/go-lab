package main

import (
	pb "GoLab/test_grpc/grpc_call/aaa"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func getTls() grpc.DialOption {
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
		RootCAs:            certPool,
		InsecureSkipVerify: true,
	}

	creds := credentials.NewTLS(tlsCfg)
	return grpc.WithTransportCredentials(creds)
}

func testCall() {
	//创建一个gRPC频道，指定连接的主机名和服务器端口
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("--- from srv, Greeting %s", r.Message)

}

func testStream() {
	// conn, err := grpc.Dial(address, grpc.WithInsecure()) // 无 tls 证书
	conn, err := grpc.Dial(address, getTls()) // 有 tls 证书

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cl := pb.NewGreeterClient(conn)

	// 设置自定义信息, srv 可以获取到
	header := metadata.New(map[string]string{"qqq": "q-111", "www": "w-222"})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	stream, err := cl.SayBye(ctx)
	if err != nil {
		log.Printf("--- cl.SayBye, err:%+v\n", err)
		return
	}
	md, _ := stream.Header() // 获取 srv SetHeader 值
	log.Printf("--- srv hdr:%+v\n", md)

	if err != nil {
		log.Println("err:", err)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 5)

		for {
			data, err := stream.Recv()
			if err != nil {
				log.Printf("-- cli Recv err:%+v\n", err)
				return
			}
			log.Printf("-- cli recv:%+v\n", data)
		}
	}()

	go func() {
		defer wg.Done()

		cnt := int32(1)
		for {
			msg := &pb.HelloRequest{Name: fmt.Sprintf("john-%d", cnt)}
			log.Printf("-- cli Send:%+v\n", msg)
			err := stream.Send(msg)
			if err != nil {
				log.Printf("-- cli Send err:%+v\n", err)
				return
			}
			cnt++
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}

func main() {
	// testCall()
	testStream()
}

// 踩坑
// 1. 客户端 tls 报错: x509: certificate is not valid for any names, but wanted to match localhost
// > 将 InsecureSkipVerify: false, 改为 InsecureSkipVerify: true,

// 设置自定义信息, 参考: http://ralphbupt.github.io/2017/05/27/gRPC%E4%B9%8Bmetadata/
