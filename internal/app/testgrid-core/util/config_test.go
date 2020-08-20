/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package util

import (
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestSetDefaultConfig(t *testing.T) {
	setDefaultConfig()
	assert.Equal(t, viper.GetString(LogLevelKey), LogLevelDefault)
}

func TestInitConfig(t *testing.T) {
	InitConfig("")
	configEnv := EnvPreFix + "_LOG_LEVEL"
	configVal := "debug-test"
	err := os.Setenv(configEnv, configVal)
	if err != nil {
		t.Errorf("Failed to set configuration env %v with error: %v", configEnv, err)
	}
	assert.Equal(t, viper.GetString(LogLevelKey), configVal)
}
