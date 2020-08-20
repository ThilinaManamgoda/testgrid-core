/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/logging"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/file"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/models"
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
	fmt.Println(testGridParentPlan.Combinations.OS[1])
}

func init() {
	generateTestPlansCmd.Flags().StringVarP(&testGridParentPlanFile, util.FileFlag, "f", "", "TestGrid parent plan location")
	err := generateTestPlansCmd.MarkFlagRequired(util.FileFlag)
	if err != nil {
		logging.ErrorAndExit(errors.Wrap(err, fmt.Sprintf("Error when making the '%s' required", util.FileFlag)), util.OsExitCode_1)
	}
}

func getTestGridParentPlan() models.TestGridParentPlan {
	u, err := url.Parse(testGridParentPlanFile)
	if err != nil {
		logging.ErrorAndExit(errors.Wrap(err, "Unable to parse the Test Grid parent plan location"), util.OsExitCode_1)
	}
	var testGridParentPlanData []byte
	switch u.Scheme {
	case util.HTTPSSchema, util.HTTPSchema:
		testGridParentPlanData, err = file.Download(u)
		if err != nil {
			logging.ErrorAndExit(errors.Wrap(err, "Unable to download the Test Grid parent plan"), util.OsExitCode_1)
		}
	case util.FileSchema:
		testGridParentPlanData, err = file.Read(u.Path)
		if err != nil {
			logging.ErrorAndExit(errors.Wrap(err, "Unable to read the Test Grid parent plan"), util.OsExitCode_1)
		}
	default:
		logging.ErrorAndExit(errors.New("Unable to identify the schema for Test Grid parent plan location"), util.OsExitCode_1)
	}

	var testGridParentPlan models.TestGridParentPlan
	err = util.ParseYaml(testGridParentPlanData, &testGridParentPlan)
	if err != nil {
		logging.ErrorAndExit(errors.Wrap(err, "Unable to parse the Test Grid parent plan"), util.OsExitCode_1)
	}
	return testGridParentPlan
}
