package service

import (
	"context"

	"github.com/suosi-inc/go-demo/grpc/internal/rpc/msg"
	pb "github.com/suosi-inc/go-demo/grpc/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userService struct {
	pb.UnimplementedUserServiceServer
}

var UserService = &userService{}

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

// Get 获取用户信息
func (s *userService) Get(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	id := in.GetId()
	if user, ok := userMap[id]; !ok {
		return nil, msg.UserEmpty.Error()
	} else {
		return &pb.GetUserResponse{Data: user}, nil
	}
}
