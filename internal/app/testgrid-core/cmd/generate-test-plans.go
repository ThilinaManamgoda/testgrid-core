/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/db"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/infracombination"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/constant"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/file"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/log"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/models"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/yaml"
	"net/url"

	"github.com/spf13/cobra"
)

var testGridParentPlanFile string

// generateTestPlansCmd represents the generateTestPlans command.
var generateTestPlansCmd = &cobra.Command{
	Use:     "generate-test-plans",
	Short:   util.GenerateTestPlansCmdShortMsg,
	Long:    util.GenerateTestPlansCmdLongDesc[1:len(util.GenerateTestPlansCmdLongDesc)],
	Example: util.GenerateTestPlansCmdExamples[1:len(util.GenerateTestPlansCmdExamples)],
	Run:     generateTestPlansCommand,
}

func generateTestPlansCommand(cmd *cobra.Command, args []string) {
	testGridParentPlan := getTestGridParentPlan()
	i := testGridParentPlan.InfraParams
	c, err := infracombination.Generate(i.OS, i.DB, i.JDK)
	if err != nil {
		log.ErrorAndExit(errors.Wrap(err, "Unable to generate combinations"), constant.OsExitCode_1)
	}

	log.Info(fmt.Sprintf("OS %s, DB %s, JDK %s", c[0].OS, c[0].DB, c[0].JDK))
	err = db.Init()
	if err != nil {
		log.ErrorAndExit(errors.Wrap(err, "Unable initiate database connection"), constant.OsExitCode_1)
	}
	hList := make([]db.HelmParam, 0)
	_, err = db.RetrieveHelmParamListForParamKey( "wso2.mysql.enabled" ,&hList)
	if err != nil {
		log.ErrorAndExit(errors.Wrap(err, "Unable to database"), constant.OsExitCode_1)
	}
	log.Info(hList[0].ParamVal)
}

func init() {
	generateTestPlansCmd.Flags().StringVarP(&testGridParentPlanFile, constant.FileFlag, "f", "", "TestGrid parent plan location")
	err := generateTestPlansCmd.MarkFlagRequired(constant.FileFlag)
	if err != nil {
		log.ErrorAndExit(errors.Wrap(err, fmt.Sprintf("Error when making the '%s' required", constant.FileFlag)), constant.OsExitCode_1)
	}
}

func getTestGridParentPlan() models.TestGridParentPlan {
	u, err := url.Parse(testGridParentPlanFile)
	if err != nil {
		log.ErrorAndExit(errors.Wrap(err, "Unable to parse the Test Grid parent plan location"), constant.OsExitCode_1)
	}
	var testGridParentPlanData []byte
	switch u.Scheme {
	case constant.HTTPSScheme, constant.HTTPScheme:
		testGridParentPlanData, err = file.Download(u)
		if err != nil {
			log.ErrorAndExit(errors.Wrap(err, "Unable to download the Test Grid parent plan"), constant.OsExitCode_1)
		}
	case constant.FileScheme:
		testGridParentPlanData, err = file.Read(u.Path)
		if err != nil {
			log.ErrorAndExit(errors.Wrap(err, "Unable to read the Test Grid parent plan"), constant.OsExitCode_1)
		}
	default:
		log.ErrorAndExit(errors.New("Unable to identify the scheme for Test Grid parent plan location"), constant.OsExitCode_1)
	}

	var testGridParentPlan models.TestGridParentPlan
	err = yaml.Parse(testGridParentPlanData, &testGridParentPlan)
	if err != nil {
		log.ErrorAndExit(errors.Wrap(err, "Unable to parse the Test Grid parent plan"), constant.OsExitCode_1)
	}
	return testGridParentPlan
}
