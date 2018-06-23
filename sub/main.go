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

func init() {
	grpc.RegisterChannelz()
}

func main() {
	server := newServer()
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, server)
	reflection.Register(s)
	channelzsvc.RegisterChannelzServiceToServer(s)

	lis, err := net.Listen("tcp", ":9000")
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
	time.Sleep(time.Second)
	return in, nil
}
