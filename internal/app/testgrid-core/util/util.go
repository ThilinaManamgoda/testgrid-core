/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package util

import (
	"gopkg.in/yaml.v2"
)

func ParseYaml(data []byte, t interface{}) error {
	err := yaml.Unmarshal(data, t)
	if err != nil {
		return err
	}
	return nil
}
