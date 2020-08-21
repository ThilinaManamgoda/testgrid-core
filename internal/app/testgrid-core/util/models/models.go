/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package models contains the structs to hold data.
package models

import (
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/deploy/helm/docker"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/infracombination"
)

// TestGridParentPlan represents the TestGrid core parent plan configuration.
type TestGridParentPlan struct {
	InfraParams     InfraParams     `yaml:"infraParams"`
	HelmDeployments HelmDeployments `yaml:"helmDeployments"`
}

// InfraParams represents the TestGrid core parent plan Infra Parameters.
type InfraParams struct {
	OS  []infracombination.OS  `yaml:"os"`
	JDK []infracombination.JDK `yaml:"jdk"`
	DB  []infracombination.DB  `yaml:"db"`
}

// HelmDeployments represents the TestGrid core parent plan Helm deployments.
type HelmDeployments struct {
	Chart        string         `yaml:"chart"`
	Version      string         `yaml:"version"`
	DockerBuilds []docker.Build `yaml:"version"`
}
