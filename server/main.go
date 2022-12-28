package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/suosi-inc/go-demo/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedUserServiceServer
}

// region 模拟的假数据
var admin = &pb.User{
	Id:           1,
	Name:         "admin",
	Weight:       62.85,
	HasSuper:     true,
	Status:       pb.UserStatus_ENABLE,
	Role:         &pb.Role{Id: 1, Name: "admin"},
	CreatedBy:    nil,
	MenuList:     []string{"admin.*", "feature1.*", "feature2.*"},
	CustomConfig: map[string]string{"color": "white", "background": "black"},
	CreatedTime:  &timestamppb.Timestamp{Seconds: 1664582400},
}
var user = &pb.User{
	Id:           2,
	Name:         "user",
	Weight:       74.21,
	HasSuper:     false,
	Status:       pb.UserStatus_ENABLE,
	Role:         &pb.Role{Id: 2, Name: "user"},
	CreatedBy:    nil,
	MenuList:     []string{"feature1.*", "feature2.*"},
	CustomConfig: map[string]string{"color": "black", "background": "white"},
	CreatedTime:  &timestamppb.Timestamp{Seconds: 1667260800},
}
var guest = &pb.User{
	Id:           3,
	Name:         "guest",
	Weight:       85.42,
	HasSuper:     false,
	Status:       pb.UserStatus_DISABLE,
	Role:         &pb.Role{Id: 3, Name: "guest"},
	CreatedBy:    nil,
	MenuList:     []string{"empty.*"},
	CustomConfig: map[string]string{"color": "red", "background": "green"},
	CreatedTime:  &timestamppb.Timestamp{Seconds: 1667260800},
}
var userMap = map[int64]*pb.User{admin.Id: admin, user.Id: user, guest.Id: guest}

// endregion

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	id := in.GetId()
	if user, ok := userMap[id]; !ok {
		return nil, errors.New("用户不存在")
	} else {
		return &pb.GetResponse{State: true, Message: "success", Data: user}, nil
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
