/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package cmd

import (
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util"

	"github.com/spf13/cobra"
)

// generateTestPlansCmd represents the generateTestPlans command
var generateTestPlansCmd = &cobra.Command{
	Use:     "generate-test-plans",
	Short:   util.GenerateTestPlansCmdShortMsg,
	Long:    util.GenerateTestPlansCmdLongDesc[1:len(util.GenerateTestPlansCmdLongDesc)],
	Example: util.GenerateTestPlansCmdExamples[1:len(util.GenerateTestPlansCmdExamples)],
	Run:     generateTestPlansCommand,
}

func generateTestPlansCommand(cmd *cobra.Command, args []string) {

}

func init() {

}

