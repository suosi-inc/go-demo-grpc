package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/suosi-inc/go-demo/grpc/internal/client"
	"github.com/suosi-inc/go-demo/grpc/internal/pkg"
)

const (
	AppName      = "client"
	AppShortDesc = "grpc-client-demo"
	AppLongDesc  = "grpc-client-demo"
)

var (
	cfgFile string

	AppCmd = &cobra.Command{
		Use:          AppName,
		Short:        AppShortDesc,
		Long:         AppLongDesc,
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			initConfig()

			pkg.InitZapLogger()
		},
		Run: func(cmd *cobra.Command, args []string) {
			client.NewApp(args)
		},
	}
)

// initConfig Init config
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Load default config file
		viper.SetConfigName(AppName)
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AddConfigPath(".")
	}

	// Load environment variables
	viper.AutomaticEnv()

	// Default config
	defaultConfig()

	// Read config file into viper config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panic("Config file not found")
		} else {
			log.Panic(err)
		}
	}

	mergeEnvConfig()

	log.Println("App using config file:", viper.ConfigFileUsed())
}

// defaultConfig Default config
func defaultConfig() {
	viper.SetDefault("app.name", AppName)

	viper.SetDefault("logger.file", AppName+".log")
	viper.SetDefault("logger.format", "text")
	viper.SetDefault("logger.level", "debug")
	viper.SetDefault("logger.maxsize", 128)
	viper.SetDefault("logger.maxage", 3)
	viper.SetDefault("logger.maxbackups", 7)
}

// Merge in environment specific config
func mergeEnvConfig() {
	configFilePath := ""
	env := []string{"product", "develop"}

out:
	for _, e := range env {
		configName := AppName + "-" + e + ".yaml"
		configPaths := []string{configName, "./config/" + configName}

		for _, path := range configPaths {
			if _, err := os.Stat(path); err == nil {
				configFilePath = path
				break out
			}
		}
	}

	if configFilePath == "" {
		return
	}
	configBytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(fmt.Errorf("Could not read config file: %s \n", err))
	}
	err = viper.MergeConfig(bytes.NewBuffer(configBytes))
	if err != nil {
		panic(fmt.Errorf("Merge config file error: %s \n", err))
	}
}

func Execute() {
	if err := AppCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
