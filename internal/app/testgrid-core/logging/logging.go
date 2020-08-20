/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package logging handles the logging functionality.
package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util"
	"io"
	"os"
	"path/filepath"
	"time"
)

var logger logrus.Logger

// Init function initializes the logging package.
// Configuration initialization is a prerequisite.
func Init() {
	date := time.Now().Format("02-01-2006")
	logfile := filepath.Join(getLogsDir(), fmt.Sprintf("testgrid-core-%v.log", date))
	file, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0655)
	if err != nil {
		panic(err)
	}

	ioWriter := io.MultiWriter(os.Stdout, file)
	logLevel := getLogLevel()
	logger = logrus.Logger{
		Out:   ioWriter,
		Level: logLevel,
		Formatter: &Formatter{
			TimestampFormat: "02-01-2006 15:04:05",
			LogFormat:       "[%time%] %lvl% - %msg%",
		},
	}
}

func getLogLevel() logrus.Level {
	logLevel := viper.GetString(util.LogLevelKey)
	parsedLogLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}
	return parsedLogLevel
}

func getLogsDir() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
}

func Info(msg string) {
	logger.Info(msg)
}

func Error(msg string) {
	logger.Error(msg)
}

func Debug(msg string) {
	logger.Debug(msg)
}
