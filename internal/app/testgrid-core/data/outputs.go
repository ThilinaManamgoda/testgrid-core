/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package data

// MainDataFlowOutput represents a main data flow Output.
type MainDataFlowOutput struct {
	Name               string `yaml:"name" json:"name"`
	Key                string `yaml:"key" json:"key"`
	RuntimeDataFlowKey string `yaml:"runtimeDataFlowKey" json:"runtimeDataFlowKey"`
}
