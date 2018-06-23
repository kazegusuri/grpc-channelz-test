package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kazegusuri/grpc-channelz-test/pb"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "127.0.0.1:8000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("grpc.DialContext failed: %v", err)
	}
	defer func() { _ = conn.Close() }()

	cli := pb.NewEchoClient(conn)
	res, err := cli.Echo(ctx, &pb.EchoMessage{Message: "Hi"})
	if err != nil {
		log.Fatalf("req error %v\n", err)
	}
	fmt.Printf("%#v\n", res)
}
