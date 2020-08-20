/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package yaml holds the implementations for YAML data handling.
package yaml

import (
	"gopkg.in/yaml.v2"
)

func Parse(data []byte, t interface{}) error {
	err := yaml.Unmarshal(data, t)
	if err != nil {
		return err
	}
	return nil
}
