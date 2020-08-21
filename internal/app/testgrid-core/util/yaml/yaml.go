/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package yaml holds the implementations for YAML data handling.
package yaml

import (
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/file"
	"gopkg.in/yaml.v2"
)

// Parse parses the given YAML data to the given struct.
func Parse(data []byte, t interface{}) error {
	err := yaml.Unmarshal(data, t)
	if err != nil {
		return err
	}
	return nil
}

// ToFile serialize the given struct and writes to the file.
func ToFile(filePath string, t interface{}) error {
	serializedData, err := yaml.Marshal(t)
	if err != nil {
		return err
	}
	err = file.Write(filePath, serializedData)
	if err != nil {
		return err
	}
	return nil
}
