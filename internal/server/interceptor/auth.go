package interceptor

import (
	"context"

	"github.com/suosi-inc/go-demo/grpc/internal/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// Auth 身份校验
func Auth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Info("进入身份校验拦截器")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "no token ")
	}

	var (
		appId  string
		appKey string
	)

	// md 是一个 map[string][]string 类型的
	if val, ok := md["appid"]; ok {
		appId = val[0]
	}

	if val, ok := md["appkey"]; ok {
		appKey = val[0]
	}

	if appId != "myappid" || appKey != "myappkey" {
		return nil, grpc.Errorf(codes.Unauthenticated, "token invalid: appid=%s, appkey=%s", appId, appKey)
	}

	// 继续处理请求
	return handler(ctx, req)
}
