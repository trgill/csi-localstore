package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "localstore/proto/storage/api"

	"google.golang.org/grpc"
)

const (
	defaultCommand   = "list"
	defaultAddress   = "localhost:55056"
	defaulDeviceList = ""
)

var (
	addr    = flag.String("addr", defaultAddress, "Port Address")
	command = flag.String("command", defaultCommand, "Command")
	devices = flag.String("devlist", defaulDeviceList, "Device List")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewLocalStorageClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var r *pb.StorageReply
	var rc error

	switch {
	case *command == "create":
		r, rc = client.CreateVolume(ctx, &pb.StorageRequest{Name: *command})
	case *command == "delete":
		r, rc = client.RemoveVolume(ctx, &pb.StorageRequest{Name: *command})
	case *command == "resize":
		r, rc = client.ResizeVolume(ctx, &pb.StorageRequest{Name: *command})
	case *command == "list":
		r, rc = client.ListVolumes(ctx, &pb.StorageRequest{Name: *command})
	}

	if err != nil {
		log.Fatalf("could not send: %v", rc)
	}
	log.Printf("Request: %s", r.GetMessage())
}
