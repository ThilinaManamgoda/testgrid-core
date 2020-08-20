/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package infracombination holds logic for infracombination generation.
package infracombination

import (
	"github.com/spf13/viper"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util"
)

type OS string
type DB string
type JDK string

// Combination represents a infra combination.
type Combination struct {
	OS  OS
	DB  DB
	JDK JDK
}

// Generator is interface that wraps infra combination generation functionality.
type Generator interface {
	Generate(osList []OS, dbList []DB, jdkList []JDK) (Combination, error)
}

//todo
type dummyGenerator struct {
}

func (g *dummyGenerator) Generate(osList []OS, dbList []DB, jdkList []JDK) ([]Combination, error) {
	return []Combination{
		{
			OS:  osList[0],
			DB:  dbList[0],
			JDK: jdkList[0],
		},
	}, nil
}

// Generate generates a infra combination.
func Generate(osList []OS, dbList []DB, jdkList []JDK) ([]Combination, error) {
	generator := viper.GetString(util.InfraCombinationGeneratorKey)
	switch generator {
	case "dummy-generator":
		dummyGen := &dummyGenerator{}
		return dummyGen.Generate(osList, dbList, jdkList)
	default:
		dummyGen := &dummyGenerator{}
		return dummyGen.Generate(osList, dbList, jdkList)
	}
}
