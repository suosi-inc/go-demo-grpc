package setup

import (
	"context"

	"github.com/suosi-inc/go-demo/grpc/internal/client/config"
	"github.com/suosi-inc/go-demo/grpc/internal/pkg/di"
	"github.com/suosi-inc/go-demo/grpc/internal/pkg/log"
	pb "github.com/suosi-inc/go-demo/grpc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// commonCredential 自定义认证
type commonCredential struct{}

// GetRequestMetadata 实现自定义认证接口
func (c commonCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "myappid",
		"appkey": "myappkey",
	}, nil
}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c commonCredential) RequireTransportSecurity() bool {
	return false
}

// ClientInterceptor 客户端拦截器
func ClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

func UserClient(cfg config.ClientCfg) {
	conn, err := grpc.Dial(
		cfg.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(new(commonCredential)),
		// grpc.WithUnaryInterceptor(ClientInterceptor),
	)

	// defer func() { _ = conn.Close() }()
	if err != nil {
		log.Errorf("did not connect userService: %v", err)
	} else {
		di.SetUserClient(pb.NewUserServiceClient(conn))
	}
}
