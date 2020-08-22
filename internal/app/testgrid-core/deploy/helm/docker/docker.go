/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package docker handles docker image building tasks.
package docker

import (
	"fmt"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/infracombination"
)

// DynamicArg represents a dynamic argument for a Docker build.
type DynamicArg struct {
	Name        string `yaml:"name" json:"name"`
	Type        string `yaml:"type" json:"type"`
	Map         string `yaml:"map" json:"map"`
	DataFlowKey string `yaml:"dataFlowKey" json:"dataFlowKey"`
}

// StaticArg represents a static argument for a Docker build.
type StaticArg struct {
	Name        string `yaml:"name" json:"name"`
	Map         string `yaml:"map" json:"map"`
	Value       string `yaml:"value" json:"value"`
	DataFlowKey string `yaml:"dataFlowKey" json:"dataFlowKey"`
}

// DataFlowArg represents a data flow argument for a Docker build.
type DataFlowArg struct {
	Name string `yaml:"name" json:"name"`
	Map  string `yaml:"map" json:"map"`
	Key  string `yaml:"key" json:"key"`
}

// Args represents arguments for a Docker build.
type Args struct {
	Static   []StaticArg   `yaml:"static" json:"static"`
	Dynamic  []DynamicArg  `yaml:"dynamic" json:"dynamic"`
	DataFlow []DataFlowArg `yaml:"dataFlow" json:"dataFlow"`
}

// Build represents the Docker image build configurations.
type Build struct {
	Name string `yaml:"name" json:"name"`
	Job  string `yaml:"job" json:"job"`
	Args Args   `yaml:"args" json:"args"`
}

// CombinationBuild represents the Docker image combination build configurations.
type CombinationBuild struct {
	Name           string `yaml:"name" json:"name"`
	Job            string `yaml:"job" json:"job"`
	Args           Args   `yaml:"args" json:"args"`
	Product        string `yaml:"product" json:"product"`
	ProductVersion string `yaml:"productVersion" json:"productVersion"`
}

// ConfigureForCombination returns, a configured Combination Build for the given Infra combination.
func (c CombinationBuild) ConfigureForCombination(combination infracombination.Combination) CombinationBuild {
	staticArgs := make([]StaticArg, len(c.Args.Static))
	copy(staticArgs, c.Args.Static)
	c.Args.Static = append(staticArgs,
		StaticArg{
			Name:        fmt.Sprintf("%s-os", c.Name),
			Map:         "OS",
			Value:       string(combination.OS),
			DataFlowKey: fmt.Sprintf("%s-arg-jdk", c.Name),
		},
		StaticArg{
			Name:        fmt.Sprintf("%s-jdk", c.Name),
			Map:         "JDK",
			Value:       string(combination.JDK),
			DataFlowKey: fmt.Sprintf("%s-arg-jdk", c.Name),
		})
	return c
}
