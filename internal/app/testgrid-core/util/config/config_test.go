/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package config

import (
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util"
	"os"
	"testing"
)

func TestSetDefaultConfig(t *testing.T) {
	setDefaults()
	assert.Equal(t, viper.GetString(util.LogLevelKey), util.LogLevelDefault)
}

func TestInitConfig(t *testing.T) {
	Init("")
	configEnv := util.EnvPreFix + "_LOG_LEVEL"
	configVal := "debug-test"
	err := os.Setenv(configEnv, configVal)
	if err != nil {
		t.Errorf("Failed to set configuration env %v with error: %v", configEnv, err)
	}
	assert.Equal(t, viper.GetString(util.LogLevelKey), configVal)
}
