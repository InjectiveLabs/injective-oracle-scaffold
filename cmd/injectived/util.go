package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	tmcfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/InjectiveLabs/injective-oracle-scaffold/cmd/injectived/config"
	"github.com/InjectiveLabs/injective-oracle-scaffold/logging"
)

// InterceptConfigsPreRunHandler performs a pre-run function for the root daemon
// application command. It will create a Viper literal and a default server
// Context. The server Tendermint configuration will either be read and parsed
// or created and saved to disk, where the server Context is updated to reflect
// the Tendermint configuration. The Viper literal is used to read and parse
// the application configuration. Command handlers can fetch the server Context
// to get the Tendermint configuration or to get access to Viper.
func InterceptConfigsPreRunHandler(cmd *cobra.Command) error {
	rootViper := viper.New()
	rootViper.BindPFlags(cmd.Flags())
	rootViper.BindPFlags(cmd.PersistentFlags())

	serverCtx := server.NewDefaultContext()
	config, err := interceptConfigs(serverCtx, rootViper)
	if err != nil {
		return err
	}

	if ll := rootViper.GetString("log-level"); len(ll) > 0 {
		config.LogLevel = ll
	} else if len(config.LogLevel) == 0 {
		config.LogLevel = "main:info,state:info,statesync:info,*:error"
	}

	logger := logging.NewWrappedSuplog("info", config.LogLevel, false)

	if rootViper.GetBool(server.FlagTrace) {
		logger = log.NewTracingLogger(logger)
	}

	serverCtx.Config = config
	serverCtx.Logger = logger.With("module", "main")

	return server.SetCmdServerContext(cmd, serverCtx)
}

// interceptConfigs parses and updates a Tendermint configuration file or
// creates a new one and saves it. It also parses and saves the application
// configuration file. The Tendermint configuration file is parsed given a root
// Viper object, whereas the application is parsed with the private package-aware
// viperCfg object.
func interceptConfigs(ctx *server.Context, rootViper *viper.Viper) (*tmcfg.Config, error) {
	rootDir := rootViper.GetString(flags.FlagHome)
	configPath := filepath.Join(rootDir, "config")
	configFile := filepath.Join(configPath, "config.toml")

	conf := tmcfg.DefaultConfig()

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		tmcfg.EnsureRoot(rootDir)

		if err = conf.ValidateBasic(); err != nil {
			return nil, fmt.Errorf("error in config file: %v", err)
		}

		conf.RPC.PprofListenAddress = "localhost:6060"
		conf.P2P.RecvRate = 5120000
		conf.P2P.SendRate = 5120000
		conf.Consensus.TimeoutCommit = 5 * time.Second
		tmcfg.WriteConfigFile(configFile, conf)
	} else {
		rootViper.SetConfigType("toml")
		rootViper.SetConfigName("config")
		rootViper.AddConfigPath(configPath)
		if err := rootViper.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("failed to read in app.toml: %w", err)
		}

		if err := rootViper.Unmarshal(conf); err != nil {
			return nil, err
		}
	}

	conf.SetRoot(rootDir)

	appConfigFilePath := filepath.Join(configPath, "app.toml")
	if _, err := os.Stat(appConfigFilePath); os.IsNotExist(err) {
		appConf, err := config.ParseConfig(ctx.Viper)
		if err != nil {
			return nil, fmt.Errorf("failed to parse app.toml: %w", err)
		}

		config.WriteConfigFile(appConfigFilePath, appConf)
	}

	ctx.Viper.SetConfigType("toml")
	ctx.Viper.SetConfigName("app")
	ctx.Viper.AddConfigPath(configPath)
	if err := ctx.Viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read in app.toml: %w", err)
	}

	return conf, nil
}
