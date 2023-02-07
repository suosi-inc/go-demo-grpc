package interceptor

import (
	"context"
	"time"

	"github.com/suosi-inc/go-demo/grpc/internal/pkg/log"
	"google.golang.org/grpc"
)

// Prof 记录时间
func Prof(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Info("进入时间拦截器")

	var start = time.Now().UnixMicro()

	// 继续处理请求
	val, err := handler(ctx, req)

	var end = time.Now().UnixMicro()
	log.Infof("耗时 %d 微秒", end-start)

	if err != nil {
		return val, err
	}

	return val, nil
}
