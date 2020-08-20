/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package log

import (
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

const (
	defaultLogFormat       = "[%time%] %lvl% - %msg%"
	defaultTimestampFormat = time.RFC3339
)

// Formatter provides the custom formatting capabilities to the logger.
type Formatter struct {
	TimestampFormat string
	LogFormat       string
}

// Format formats and returns the given log entry.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	output = strings.Replace(output, "%msg%", entry.Message, 1)

	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%lvl%", level, 1)

	for k, val := range entry.Data {
		switch v := val.(type) {
		case string:
			output = strings.Replace(output, "%"+k+"%", v, 1)
		case int:
			s := strconv.Itoa(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		case bool:
			s := strconv.FormatBool(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}
	return []byte(output + "\n"), nil
}
