// server
package main

import (
	"fmt"
	"net"

	pb "github.com/thurt/simple-chat/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	//"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	//"github.com/golang/protobuf/proto"
)

const port = 10000

type echoServer struct{}

func (s *echoServer) Send(ctx context.Context, msg *pb.Msg) (*pb.Msg, error) {
	return &pb.Msg{"> " + msg.Text}, nil
}

func newServer() *echoServer {
	s := new(echoServer)
	return s
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterEchoServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
