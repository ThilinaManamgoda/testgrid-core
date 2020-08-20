/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package models contains the structs to hold data.
package models

// TestGridParentPlan represents the TestGrid core parent plan configuration.
type TestGridParentPlan struct {
	Combinations Combinations `yaml:"combinations"`
}

// Combinations represents the TestGrid core parent plan combinations.
type Combinations struct {
	OS  []string `yaml:"os"`
	JDK []string `yaml:"jdk"`
	DB  []string `yaml:"db"`
}
