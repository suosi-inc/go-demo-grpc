package interceptor

import (
	"context"

	"github.com/suosi-inc/go-demo/grpc/internal/pkg/log"
	"github.com/suosi-inc/go-demo/grpc/internal/server/msg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var appList = map[string]string{
	"appId":   "appKey",
	"myAppId": "myAppKey",
}

func validAuth(appId, appKey string) error {
	if val, ok := appList[appId]; !ok || val != appKey {
		return msg.UserNotLogin.Error()
	}
	return nil
}

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
	if val, ok := md["app-id"]; ok {
		appId = val[0]
	}

	if val, ok := md["app-key"]; ok {
		appKey = val[0]
	}

	log.Infof("身份校验信息 appId: %s, appKey: %s", appId, appKey)

	err := validAuth(appId, appKey)
	if err != nil {
		return nil, err
	}

	// 继续处理请求
	v, err := handler(ctx, req)

	log.Info("退出身份校验拦截器")
	return v, err
}
