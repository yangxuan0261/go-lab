package main

import (
	"google.golang.org/grpc"

	pro "go_lab/test_grpc/grpc_stream_putget/proto"
	"context"
	"log"
	"time"

	_ "google.golang.org/grpc/balancer/grpclb"
)

const (
	ADDRESS = "localhost:50051"
)

func main() {
	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := pro.NewGreeterClient(conn)

	//调用服务端推送流
	reqstreamData := &pro.StreamReqData{Data: "--- cli data 333"}
	res, _ := c.GetStream(context.Background(), reqstreamData)
	for {
		data, err := res.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("--- cli GetStream recv:", data)
	}

	//客户端 推送 流
	putRes, _ := c.PutStream(context.Background())
	i := 1
	for {
		i++
		putRes.Send(&pro.StreamReqData{Data: "cli data 222"})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	//服务端 客户端 双向流
	allStr, _ := c.AllStream(context.Background())
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Println("--- cli AllStream recv:", data)
		}
	}()

	go func() {
		for {
			allStr.Send(&pro.StreamReqData{Data: "--- cli data 111"})
			time.Sleep(time.Second)
		}
	}()

	select {}

}
