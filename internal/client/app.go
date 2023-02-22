package client

import (
	"github.com/spf13/viper"
	"github.com/suosi-inc/go-demo/grpc/internal/client/command"
	"github.com/suosi-inc/go-demo/grpc/internal/client/config"
	"github.com/suosi-inc/go-demo/grpc/internal/pkg/log"
	"github.com/suosi-inc/go-demo/grpc/internal/pkg/setup"
)

func NewApp(args []string) {
	bootstrap()

	setupDi()

	if len(args) > 0 {
		runCommand(args)
	}
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
	setup.UserClient(config.Cfg.UserClient)
}

func runCommand(args []string) {
	arg := args[0]
	switch arg {
	case "user:get":
		command.UserCommand.Get()
	case "user:search":
		command.UserCommand.Search()
	case "user:add":
		command.UserCommand.Add()
	case "user:edit":
		command.UserCommand.Edit()
	case "user:remove":
		command.UserCommand.Remove()
	default:
		log.Error("At least one parameter")
	}
}
