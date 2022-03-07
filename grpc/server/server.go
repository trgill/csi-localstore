package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "localstore/proto/storage/api"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 55056, "server port")
)

type server struct {
	pb.UnimplementedLocalStorageServer
}

func (s *server) CreateVolume(ctx context.Context, in *pb.StorageRequest) (*pb.StorageReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.StorageReply{Message: "StorageRequest " + in.GetName()}, nil
}

func (s *server) RemoveVolume(ctx context.Context, in *pb.StorageRequest) (*pb.StorageReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.StorageReply{Message: "StorageRequest " + in.GetName()}, nil
}

func (s *server) ResizeVolume(ctx context.Context, in *pb.StorageRequest) (*pb.StorageReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.StorageReply{Message: "StorageRequest " + in.GetName()}, nil
}

func (s *server) ListVolumes(ctx context.Context, in *pb.StorageRequest) (*pb.StorageReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.StorageReply{Message: "StorageRequest " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLocalStorageServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
