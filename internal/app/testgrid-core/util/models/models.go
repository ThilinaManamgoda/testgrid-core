/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package models contains the structs to hold data.
package models

import (
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/deploy/helm"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/infracombination"
)

// TestGridParentPlan represents the TestGrid core parent plan configuration.
type TestGridParentPlan struct {
	InfraParams     InfraParams       `yaml:"infraParams"`
	HelmDeployments []helm.Deployment `yaml:"helmDeployments"`
}

// InfraParams represents the TestGrid core parent plan Infra Parameters.
type InfraParams struct {
	OS  []infracombination.OS  `yaml:"os"`
	JDK []infracombination.JDK `yaml:"jdk"`
	DB  []infracombination.DB  `yaml:"db"`
}

// ExecutionPlan represents the TestGrid core execution plan configuration.
type ExecutionPlan struct {
	ParentUUID      string            `yaml:"parentUuid"`
	UUID            string            `yaml:"uuid"`
	JobID           string            `yaml:"jobId"`
	HelmDeployments []helm.Deployment `yaml:"helmDeployments"`
}

// TestPlan represents the TestGrid core command execution plan configuration.
type TestPlan struct {
	ID      string   `yaml:"id"`
	DeployCMD []string `yaml:"deployCMD"`
}
