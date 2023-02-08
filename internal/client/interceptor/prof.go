package interceptor

import (
	"context"
	"time"

	"github.com/suosi-inc/go-demo/grpc/internal/pkg/log"
	"google.golang.org/grpc"
)

// Prof 时间分析拦截器-客户端
func Prof(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Info("时间分析拦截器开始")
	var start = time.Now().UnixMicro()

	err := invoker(ctx, method, req, reply, cc, opts...)

	var end = time.Now().UnixMicro()
	log.Infof("耗时 %d 微秒", end-start)
	log.Info("时间分析拦截器结束")
	return err
}
