/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/constant"
	"os"
	"strings"
)

// TestGridCore represents the TestGrid core configurations.
type TestGridCore struct {
	Log                       Log    `json:"log"`
	InfraCombinationGenerator string `json:"infraCombinationGenerator"`
	DB                        DB     `json:"db"`
	Redis                     Redis  `json:"redis"`
}

// Log represents the TestGrid core log configurations.
type Log struct {
	Level string `json:"level"`
}

// Redis represents the TestGrid core Redis Client configurations.
type Redis struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// DB represents the TestGrid core database configurations.
type DB struct {
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Name       int    `json:"name"`
	LogLevel   string `json:"logLevel"`
	MaxRetries int    `json:"maxRetries"`
}

func setDefaults() {
	viper.SetDefault(constant.LogLevelKey, constant.LogLevelDefault)
	//todo
	viper.SetDefault(constant.InfraCombinationGeneratorIDKey, constant.InfraCombinationGeneratorIDDefault)

	// Database defaults.
	viper.SetDefault(constant.DatabaseUserNameKey, constant.DatabaseUserNameDefault)
	viper.SetDefault(constant.DatabasePasswordKey, constant.DatabasePasswordDefault)
	viper.SetDefault(constant.DatabaseHostKey, constant.DatabaseHostDefault)
	viper.SetDefault(constant.DatabasePortKey, constant.DatabasePortDefault)
	viper.SetDefault(constant.DatabaseNameKey, constant.DatabaseNameDefault)
	viper.SetDefault(constant.DatabaseLogLevelKey, constant.DatabaseLogLevelDefault)
	viper.SetDefault(constant.DatabaseMaxRetriesKey, constant.DatabaseMaxRetriesDefault)

	// Redis client defaults.
	viper.SetDefault(constant.RedisServerAddressKey, constant.RedisServerAddressDefault)
	viper.SetDefault(constant.RedisServerPasswordKey, constant.RedisServerPasswordDefault)
	viper.SetDefault(constant.RedisServerDBKey, constant.RedisServerDBDefault)
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
	viper.SetEnvPrefix(constant.EnvPreFix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
