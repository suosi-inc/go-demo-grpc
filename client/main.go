package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/suosi-inc/go-demo/grpc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() { _ = conn.Close() }()

	userClient := protobuf.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := userClient.Get(ctx, &protobuf.GetUserRequest{Id: 4})
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	log.Printf("result: %+v\n", result)
}
