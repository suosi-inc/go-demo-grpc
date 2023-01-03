package service

import (
	"context"

	"github.com/suosi-inc/go-demo/grpc/internal/server/msg"
	pb "github.com/suosi-inc/go-demo/grpc/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	CreatedBy:    admin,
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
	CreatedBy:    admin,
	MenuList:     []string{"empty.*"},
	CustomConfig: map[string]string{"color": "red", "background": "green"},
	CreatedTime:  &timestamppb.Timestamp{Seconds: 1667260800},
}
var userMap = map[int64]*pb.User{admin.Id: admin, user.Id: user, guest.Id: guest}

// endregion

// Get 获取用户信息
func (userService) Get(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if user, ok := userMap[req.GetId()]; !ok {
		return nil, msg.UserEmpty.Error()
	} else {
		return &pb.GetUserResponse{Data: user}, nil
	}
}

// Search 用户列表
func (userService) Search(ctx context.Context, req *pb.SearchUserRequest) (*pb.SearchUserResponse, error) {
	userList := make([]*pb.User, 0)
	total := int64(0)
	for _, user := range userMap {
		userList = append(userList, user)
		total++
	}
	return &pb.SearchUserResponse{List: userList, Total: total}, nil
}

// Add 添加用户
func (userService) Add(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

// Edit 编辑用户
func (userService) Edit(ctx context.Context, req *pb.EditUserRequest) (*pb.EditUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}

// Remove 删除用户
func (userService) Remove(ctx context.Context, req *pb.RemoveUserRequest) (*pb.RemoveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
