/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package models contains the structs to hold data.
package models

import "github.com/wso2/testgrid-core/internal/app/testgrid-core/infracombination"

// TestGridParentPlan represents the TestGrid core parent plan configuration.
type TestGridParentPlan struct {
	InfraParams InfraParams `yaml:"infraParams"`
}

// InfraParams represents the TestGrid core parent plan Infra Parameters.
type InfraParams struct {
	OS  []infracombination.OS  `yaml:"os"`
	JDK []infracombination.JDK `yaml:"jdk"`
	DB  []infracombination.DB  `yaml:"db"`
}
