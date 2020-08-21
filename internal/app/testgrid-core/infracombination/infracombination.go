/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package infracombination holds logic for infracombination generation.
package infracombination

import (
	"github.com/spf13/viper"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/constant"
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

// Generator is the interface which wraps infra combination generation functionality.
type Generator interface {
	Generate(osList []OS, dbList []DB, jdkList []JDK) ([]Combination, error)
}

//todo
type dummyGenerator struct {
	ID string
}

func (g *dummyGenerator) Generate(osList []OS, dbList []DB, jdkList []JDK) ([]Combination, error) {
	return []Combination{
		{
			OS:  osList[0],
			DB:  dbList[0],
			JDK: jdkList[0],
		},
		{
			OS:  osList[1],
			DB:  dbList[0],
			JDK: jdkList[0],
		},
	}, nil
}

// Generate generates infra combinations.
func Generate(osList []OS, dbList []DB, jdkList []JDK) ([]Combination, error) {
	generatorID := viper.GetString(constant.InfraCombinationGeneratorIDKey)
	var generator Generator
	switch generatorID {
	case constant.DummyGeneratorID:
		generator = &dummyGenerator{ID: constant.DummyGeneratorID}
	default:
		generator = &dummyGenerator{ID: constant.DummyGeneratorID}
	}
	return generator.Generate(osList, dbList, jdkList)
}

//func ToHelmParameters(combination Combination) map[string]string {
//	//helmParams := make(map[string]string)
//	//switch combination.DB {
//	//case :
//	//
//	//}
//}
