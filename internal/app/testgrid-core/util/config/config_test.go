/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package config

import (
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/constant"
	"os"
	"testing"
)

func TestSetDefaultConfig(t *testing.T) {
	setDefaults()
	assert.Equal(t, viper.GetString(constant.LogLevelKey), constant.LogLevelDefault)
	// Database defaults
	assert.Equal(t, viper.GetString(constant.DatabaseUserNameKey), constant.DatabaseUserNameDefault)
	assert.Equal(t, viper.GetString(constant.DatabasePasswordKey), constant.DatabasePasswordDefault)
	assert.Equal(t, viper.GetString(constant.DatabaseHostKey), constant.DatabaseHostDefault)
	assert.Equal(t, viper.GetString(constant.DatabasePortKey), constant.DatabasePortDefault)
	assert.Equal(t, viper.GetString(constant.DatabaseNameKey), constant.DatabaseNameDefault)
	assert.Equal(t, viper.GetString(constant.DatabaseLogLevelKey), constant.DatabaseLogLevelDefault)
	assert.Equal(t, viper.GetString(constant.DatabaseMaxRetriesKey), constant.DatabaseMaxRetriesDefault)
}

func TestInitConfig(t *testing.T) {
	Init("")
	configEnv := constant.EnvPreFix + "_LOG_LEVEL"
	configVal := "debug-test"
	err := os.Setenv(configEnv, configVal)
	if err != nil {
		t.Errorf("Failed to set configuration env %v with error: %v", configEnv, err)
	}
	assert.Equal(t, viper.GetString(constant.LogLevelKey), configVal)
}
