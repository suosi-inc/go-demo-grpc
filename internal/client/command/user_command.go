package command

import (
	"context"
	"time"

	"github.com/suosi-inc/go-demo/grpc/internal/client/config"
	"github.com/suosi-inc/go-demo/grpc/internal/pkg/di"
	"github.com/suosi-inc/go-demo/grpc/internal/pkg/log"
	pb "github.com/suosi-inc/go-demo/grpc/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userCommand struct {
}

var UserCommand userCommand

// Search 搜索
func (userCommand) Search() {
	client := di.GetUserClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*config.Cfg.UserClient.Timeout)
	defer cancel()

	result, err := client.Search(ctx, &pb.SearchUserRequest{
		Page: 1,
		Size: 10,
		CreatedTimeRange: &pb.SearchUserRequest_CreatedTimeRange{
			Start: &timestamppb.Timestamp{Seconds: time.Now().Unix() - 60*86400},
			End:   &timestamppb.Timestamp{Seconds: time.Now().Unix()},
		},
	})
	if err != nil {
		log.Errorf("error: %s\n", err.Error())
	}
	log.Infof("result: %+v\n", result)
}

// Get 获取单条信息
func (userCommand) Get() {
	client := di.GetUserClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*config.Cfg.UserClient.Timeout)
	defer cancel()

	result, err := client.Get(ctx, &pb.GetUserRequest{Id: 1})
	if err != nil {
		log.Errorf("error: %s\n", err.Error())
	}
	log.Infof("result: %+v\n", result)
}

// Add 添加
func (userCommand) Add() {
	client := di.GetUserClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*config.Cfg.UserClient.Timeout)
	defer cancel()

	result, err := client.Add(ctx, &pb.AddUserRequest{
		Name:     "test",
		HasSuper: false,
		Status:   pb.UserStatus_DISABLE,
	})
	if err != nil {
		log.Errorf("error: %s\n", err.Error())
	}
	log.Infof("result: %+v\n", result)
}

// Edit 编辑
func (userCommand) Edit() {
	client := di.GetUserClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*config.Cfg.UserClient.Timeout)
	defer cancel()

	name := "test_demo"
	result, err := client.Edit(ctx, &pb.EditUserRequest{Id: 3, Name: &name})
	if err != nil {
		log.Errorf("error: %s\n", err.Error())
	}
	log.Infof("result: %+v\n", result)
}

// Remove 删除
func (userCommand) Remove() {
	client := di.GetUserClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*config.Cfg.UserClient.Timeout)
	defer cancel()

	result, err := client.Remove(ctx, &pb.RemoveUserRequest{Id: 1})
	if err != nil {
		log.Errorf("error: %s\n", err.Error())
	}
	log.Infof("result: %+v\n", result)
}
