/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package log handles the log functionality.
package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/constant"
	"io"
	"os"
	"path/filepath"
	"time"
)

var logger logrus.Logger
var logLevel logrus.Level

// Init function initializes the log package.
// Configuration initialization is a prerequisite.
func Init() {
	date := time.Now().Format("02-01-2006")
	logfile := filepath.Join(getLogsDir(), fmt.Sprintf("testgrid-core-%v.log", date))
	file, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0655)
	if err != nil {
		panic(err)
	}

	ioWriter := io.MultiWriter(os.Stdout, file)
	logLevel = getLogLevel()
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
	logLevel := viper.GetString(constant.LogLevelKey)
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

func ErrorAndExit(err error, exitCode int) {
	logger.Error(err)
	os.Exit(exitCode)
}

func Debug(msg string) {
	logger.Debug(msg)
}

func IsDebugEnabled() bool {
	return logger.IsLevelEnabled(logrus.DebugLevel)
}

type GormLogger struct{}

func (*GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		logger.WithFields(
			logrus.Fields{
				"module":        "gorm",
				"type":          "sql",
				"rows_returned": v[5],
				"src":           v[1],
				"values":        v[4],
				"duration":      v[2],
			},
		).Info(v[3])
	case "log":
		logger.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}
