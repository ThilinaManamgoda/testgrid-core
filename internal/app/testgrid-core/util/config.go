/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */


package util

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// TestGridCoreConfig represents the TestGrid core configurations.
type TestGridCoreConfig struct {
	Log Log `json:"log"`
}

// Log represents the TestGrid core logging configurations.
type Log struct {
	Level string `json:"level"`
}

func setDefaultConfig(){
	viper.SetDefault(LogLevelKey, LogLevelDefault)
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".testgrid-core" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".testgrid-core")
	}
	setDefaultConfig()
	viper.SetEnvPrefix(EnvPreFix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
