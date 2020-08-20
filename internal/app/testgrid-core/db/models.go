/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package db

import "github.com/wso2/testgrid-core/internal/app/testgrid-core/util/constant"

type InfraParam struct {
	ID    int    `gorm:"primary_key"`
	Param string `gorm:"column:param"`
	Type  string `gorm:"column:type"`
}

type HelmParam struct {
	ID           int    `gorm:"primary_key"`
	InfraParamID int    `gorm:"column:infra_param_id"`
	ParamKey     string `gorm:"column:param_key"`
	ParamVal     string `gorm:"column:param_value"`
}

func (HelmParam) TableName() string {
	return constant.HelmParamTable
}

func (InfraParam) TableName() string {
	return constant.InfraParamTable
}
