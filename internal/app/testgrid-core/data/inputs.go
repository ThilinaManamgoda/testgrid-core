/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package data handles data flow operations.
package data

// DynamicInput represents a Dynamic input.
type DynamicInput struct {
	Name        string `yaml:"name" json:"name"`
	Type        string `yaml:"type" json:"type"`
	Map         string `yaml:"map" json:"map"`
	DataFlowKey string `yaml:"dataFlowKey" json:"dataFlowKey"`
}

// DataFlowInput represents a data flow input.
type DataFlowInput struct {
	Name string `yaml:"name" json:"name"`
	Map  string `yaml:"map" json:"map"`
	Key  string `yaml:"key" json:"key"`
}

// StaticInput represents a static input.
type StaticInput struct {
	Name        string `yaml:"name" json:"name"`
	Map         string `yaml:"map" json:"map"`
	Value       string `yaml:"value" json:"value"`
	DataFlowKey string `yaml:"dataFlowKey" json:"dataFlowKey"`
}
