package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/k0kubun/pp"
	"github.com/kazegusuri/grpc-channelz-test/pb"
	"google.golang.org/grpc"
	channelzsvc "google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/reflection"
)

var (
	cli1 pb.EchoClient
	cli2 pb.EchoClient
)

func init() {
	grpc.RegisterChannelz()
}

func main() {
	ctx := context.Background()
	conn1, err := grpc.DialContext(ctx, "127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.DialContext failed: %v", err)
	}
	defer func() { _ = conn1.Close() }()

	conn2, err := grpc.DialContext(ctx, "127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.DialContext failed: %v", err)
	}
	defer func() { _ = conn2.Close() }()

	cli1 = pb.NewEchoClient(conn1)
	cli2 = pb.NewEchoClient(conn2)

	go func() {
		for {
			time.Sleep(3 * time.Second)
			_, err := cli1.Echo(ctx, &pb.EchoMessage{Message: "hello"})
			if err != nil {
				log.Printf("goroutine: echo failed: %v", err)
			} else {
				log.Printf("goroutine: echo succeeded")
			}
		}
	}()

	server := newServer()
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, server)
	reflection.Register(s)
	channelzsvc.RegisterChannelzServiceToServer(s)

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("err %v\n", err)
	}

	select {}
}

type server struct {
}

func newServer() *server {
	return &server{}
}

func (s *server) Echo(ctx context.Context, in *pb.EchoMessage) (*pb.EchoMessage, error) {
	pp.Printf("%v\n", in)
	return cli2.Echo(ctx, in)
}
