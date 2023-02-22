package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/suosi-inc/go-demo/grpc/cmd/client"
	"github.com/suosi-inc/go-demo/grpc/cmd/server"
)

var rootCmd = &cobra.Command{
	Use:          "grpc-demo",
	Short:        "args: `server`/`client`",
	Long:         `server: open server, client: request service data`,
	SilenceUsage: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("must provide at least one entry")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
}

func init() {
	rootCmd.AddCommand(server.AppCmd)
	rootCmd.AddCommand(client.AppCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
