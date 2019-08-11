package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/yuwe1/gopratice/proto-pratice/protodemo/protodemo1"
)

func main() {
	// 创建一个 gRPC channel 和服务器交互
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed:%v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := pb.NewCallServiceClient(conn)

	// 直接调用
	resp1, err := client.SayCall(context.Background(), &pb.CallRequest{
		Greeting: "Hello Server 1 !!",
		Infos:    map[string]string{"A": "B"},
	})

	log.Printf("Resp1:%+v", resp1)

	resp2, err := client.SayCall(context.Background(), &pb.CallRequest{
		Greeting: "Hello Server 2 !!",
	})

	log.Printf("Resp2:%s", resp2.GetDetails())
}
