package main

import (
	"context"
	"net"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	pb "github.com/yuwe1/gopratice/proto-pratice/protodemo/protodemo1"
	"google.golang.org/grpc"
)

type CallServer struct {
}

func (c *CallServer) SayCall(ctx context.Context, re *pb.CallRequest) (res *pb.CallResponse, err error) {

	var an *any.Any
	if re.Infos["A"] == "B" {
		an, err = ptypes.MarshalAny(&pb.Res{Reply: "请求正确"})
	} else {
		an, err = ptypes.MarshalAny(&pb.Res{Reply: "请求出错"})
	}
	return &pb.CallResponse{
		Reply:   "Hello World !!",
		Details: []*any.Any{an},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// 新建一个grpc服务器
	grpcServer := grpc.NewServer()
	// 向grpc服务器注册SayHelloServer
	pb.RegisterCallServiceServer(grpcServer, &CallServer{})
	// 启动服务
	grpcServer.Serve(lis)
}
