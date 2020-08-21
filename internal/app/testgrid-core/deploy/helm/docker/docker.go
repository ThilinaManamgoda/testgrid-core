/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package docker handles docker building tasks.
package docker

// DynamicArg represents a dynamic argument for a Docker build.
type DynamicArg struct {
	Name       string `yaml:"name"`
	Type       string `yaml:"type"`
	Key        string `yaml:"key"`
	MappingKey string `yaml:"type"`
}

// StaticArg represents a static argument for a Docker build.
type StaticArg struct {
	Name  string `yaml:"name"`
	Key   string `yaml:"key"`
	Value string `yaml:"type"`
}

// Build represents the Docker build configurations.
type Build struct {
	Name        string       `yaml:"name"`
	Type        string       `yaml:"type"`
	Job         string       `yaml:"job"`
	StaticArgs  []StaticArg  `yaml:"staticArgs"`
	DynamicArgs []DynamicArg `yaml:"dynamicArgs"`
}
