package setup

import (
	"context"

	"github.com/suosi-inc/go-demo/grpc/internal/client/config"
	"github.com/suosi-inc/go-demo/grpc/internal/client/interceptor"
	"github.com/suosi-inc/go-demo/grpc/internal/pkg/di"
	"github.com/suosi-inc/go-demo/grpc/internal/pkg/log"
	pb "github.com/suosi-inc/go-demo/grpc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// userCredential 自定义认证
type userCredential struct{}

// GetRequestMetadata 实现自定义认证接口
func (c userCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"app-id":  config.Cfg.UserClient.AppId,
		"app-key": config.Cfg.UserClient.AppKey,
	}, nil
}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c userCredential) RequireTransportSecurity() bool {
	return false
}

func UserClient(cfg config.ClientCfg) {
	conn, err := grpc.Dial(
		cfg.Addr,
		// 禁用安全证书
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 带入认证元信息
		grpc.WithPerRPCCredentials(new(userCredential)),
		// 客户端过滤链(拦截器链)
		grpc.WithChainUnaryInterceptor(
			interceptor.Prof,
			interceptor.Other,
		),
	)

	// defer func() { _ = conn.Close() }()
	if err != nil {
		log.Errorf("did not connect userService: %v", err)
	} else {
		di.SetUserClient(pb.NewUserServiceClient(conn))
	}
}
