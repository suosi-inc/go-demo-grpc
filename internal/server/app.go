package server

import (
	"errors"
	"fmt"
	"net"

	"github.com/spf13/viper"
	"github.com/suosi-inc/go-demo/grpc/internal/pkg/log"
	"github.com/suosi-inc/go-demo/grpc/internal/server/config"
	"github.com/suosi-inc/go-demo/grpc/internal/server/interceptor"
	"github.com/suosi-inc/go-demo/grpc/internal/server/service"
	pb "github.com/suosi-inc/go-demo/grpc/protobuf"
	"google.golang.org/grpc"
)

func NewServer() error {
	bootstrap()

	setupDi()

	return runServer()
}

// bootstrap Bootstrap app
func bootstrap() {
	// Config map into struct
	err := viper.Unmarshal(&config.Cfg)
	if err != nil {
		panic("Unable to decode config into struct: ")
	}
	log.Info("Config into struct: ", log.Any("cfg", config.Cfg))
}

func setupDi() {
}

func runServer() error {
	lis, err := net.Listen("tcp", config.Cfg.Server.Port)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to listen: %v", err))
	}

	server := grpc.NewServer(
		// 注册拦截器
		grpc.ChainUnaryInterceptor(
			interceptor.Prof,
			interceptor.Auth,
		),
	)
	pb.RegisterUserServiceServer(server, service.UserService)
	log.Info(fmt.Sprintf("start server: %v", lis.Addr()))
	return server.Serve(lis)
}
