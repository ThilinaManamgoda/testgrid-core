/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package helm handles Helm operations.
package helm

import (
	"fmt"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/data"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/deploy/helm/docker"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/infracombination"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/log"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/msg"
)

// Inputs represents inputs for a Helm deployment.
type Inputs struct {
	Static        []data.StaticInput   `yaml:"static" json:"static"`
	Dynamic       []data.DynamicInput  `yaml:"dynamic" json:"dynamic"`
	DataFlowInput []data.DataFlowInput `yaml:"dataFlow" json:"dataFlow"`
}

// Outputs represents outputs for a Helm deployment.
type Outputs struct {
	MainDataFlow []data.MainDataFlowOutput `yaml:"mainDataFlow" json:"mainDataFlow"`
}

// DockerBuilds represents docker image builds for a Helm deployment.
type DockerBuilds struct {
	Combination []docker.CombinationBuild `yaml:"combination" json:"combination"`
	Custom      []docker.Build            `yaml:"custom" json:"custom"`
}

// Deployment represents a Helm deployment.
type Deployment struct {
	Name         string       `yaml:"name" json:"name"`
	Chart        string       `yaml:"chart" json:"chart"`
	Version      string       `yaml:"version" json:"version"`
	DockerBuilds DockerBuilds `yaml:"dockerBuilds" json:"dockerBuilds"`
	Inputs       Inputs       `yaml:"inputs" json:"inputs"`
	Outputs      Outputs      `yaml:"outputs" json:"outputs"`
}

// ConfigureForCombination returns, a configured Helm deployment for the given Infra combination.
func (d Deployment) ConfigureForCombination(combination infracombination.Combination) Deployment {
	configuredCombinationDockerBuilds := make([]docker.CombinationBuild, len(d.DockerBuilds.Combination))
	for i, c := range d.DockerBuilds.Combination {
		configuredCombinationDockerBuilds[i] = c.ConfigureForCombination(combination)
		log.Debug(fmt.Sprintf(msg.CombinationDockerBuildInfo, c.Name, d.Name, combination.String()))
	}
	d.DockerBuilds.Combination = configuredCombinationDockerBuilds
	return d
}
