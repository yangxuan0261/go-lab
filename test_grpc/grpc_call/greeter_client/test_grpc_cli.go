package main

import (
	pb "GoLab/test_grpc/grpc_call/aaa"
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

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
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cl := pb.NewGreeterClient(conn)

	stream, err := cl.SayBye(context.Background())
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
