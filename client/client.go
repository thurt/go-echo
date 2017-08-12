// client
package main

import (
	"bufio"
	"fmt"
	"os"

	pb "github.com/thurt/simple-chat/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const serverAddr = "127.0.0.1:10000"

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewEchoClient(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		res, err := client.Send(context.Background(), &pb.Msg{scanner.Text()})
		if err != nil {
			grpclog.Fatalf(err.Error())
		}
		fmt.Println(res.Text)
	}
}
