package interceptor

import (
	"context"

	"github.com/suosi-inc/go-demo/grpc/internal/pkg/log"
	"google.golang.org/grpc"
)

// Other 客户端其他拦截器
func Other(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Infof("客户端其他拦截器开始")

	err := invoker(ctx, method, req, reply, cc, opts...)

	log.Info("客户端其他拦截器结束")
	return err
}
