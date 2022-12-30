package setup

import (
	"github.com/suosi-inc/go-demo/grpc/internal/client/config"
	"github.com/suosi-inc/go-demo/grpc/internal/pkg/di"
	"github.com/suosi-inc/go-demo/grpc/internal/pkg/log"
	pb "github.com/suosi-inc/go-demo/grpc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClient(cfg config.ClientCfg) {
	conn, err := grpc.Dial(cfg.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// defer func() { _ = conn.Close() }()
	if err != nil {
		log.Errorf("did not connect userService: %v", err)
	} else {
		di.SetUserClient(pb.NewUserServiceClient(conn))
	}
}
