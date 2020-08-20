/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util"
	"os"
	"strings"
)

// TestGridCore represents the TestGrid core configurations.
type TestGridCore struct {
	Log                       Log    `json:"log"`
	InfraCombinationGenerator string `json:"infraCombinationGenerator"`
}

// Log represents the TestGrid core log configurations.
type Log struct {
	Level string `json:"level"`
}

func setDefaults() {
	viper.SetDefault(util.LogLevelKey, util.LogLevelDefault)
	//todo
	viper.SetDefault(util.InfraCombinationGeneratorKey, util.InfraCombinationGeneratorDefault)
}

// Init reads in config file and ENV variables if set.
func Init(cfgFile string) {
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
	setDefaults()
	viper.SetEnvPrefix(util.EnvPreFix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
