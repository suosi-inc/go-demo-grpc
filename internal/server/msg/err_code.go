package msg

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	UserEmpty    = errCode{Code: 100001, Msg: "用户不存在"}
	UserNotLogin = errCode{Code: codes.Unauthenticated, Msg: "用户授权失败"}
)

type errCode struct {
	Code codes.Code
	Msg  string
}

func (c *errCode) Error() error {
	return status.Errorf(c.Code, c.Msg)
}
