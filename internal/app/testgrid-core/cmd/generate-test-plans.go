/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package cmd

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/client/redis"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/deploy/helm"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/infracombination"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/constant"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/file"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/log"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/models"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/msg"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/yaml"
	"net/url"
	"os"
	"path"
)

var testGridParentPlanFile string
var testGridParentPlan models.TestGridParentPlan
var infraCombinations []infracombination.Combination
var executionPlans []models.ExecutionPlan
var testPlan []models.TestPlan

var parentUUID string

// generateTestPlansCmd represents the generateTestPlans command.
var generateTestPlansCmd = &cobra.Command{
	Use:     "generate-test-plans",
	Short:   util.GenerateTestPlansCmdShortMsg,
	Long:    util.GenerateTestPlansCmdLongDesc[1:len(util.GenerateTestPlansCmdLongDesc)],
	Example: util.GenerateTestPlansCmdExamples[1:len(util.GenerateTestPlansCmdExamples)],
	Run:     generateTestPlansCommand,
}

func generateTestPlansCommand(cmd *cobra.Command, args []string) {
	populateTestGridParentPlan()
	generateInfraCombinations()
	generateExecutionPlans()
	generateTestPlan()
	persistTestPlan()
}

func persistTestPlan() {
	d, err := os.Getwd()
	if err != nil {
		log.ErrorAndExit(errors.Wrap(err, "Unable get working directory"), constant.OsExitCode_1)
	}
	for _, p := range testPlan {
		err = yaml.ToFile(path.Join(d, "test-plans", p.ID), p)
		if err != nil {
			log.ErrorAndExit(errors.Wrap(err, fmt.Sprintf("Unable persist test plan %s", p.ID)), constant.OsExitCode_1)
		}
	}
}

func generateTestPlan() {
	log.Info("Generate command execution plan")
	redisBulkInserts := make(map[string]interface{})
	for _, e := range executionPlans {
		var deployCMDRedisIDs []string
		for i, h := range e.HelmDeployments {
			deployCMDRedisID := fmt.Sprintf(constant.DeployCMDRedisID, e.ParentUUID, e.UUID, i)
			redisBulkInserts[deployCMDRedisID] = h
			log.Debug(fmt.Sprintf(msg.DeployCMDConfigRedisIDInfo, h.Name, deployCMDRedisID))
			deployCMDRedisIDs = append(deployCMDRedisIDs, deployCMDRedisID)
		}
		testPlan = append(testPlan, models.TestPlan{
			ID:        fmt.Sprintf(constant.DeployCMDExePlanID, e.JobID),
			DeployCMD: deployCMDRedisIDs})
	}
	err := redis.Init()
	if err != nil {
		log.ErrorAndExit(errors.Wrap(err, "Unable init Redis client"), constant.OsExitCode_1)
	}
	err = redis.SetBulk(redisBulkInserts)
	if err != nil {
		log.ErrorAndExit(errors.Wrap(err, "Unable set execution commands configs in Redis"), constant.OsExitCode_1)
	}
}

func generateExecutionPlans() {
	log.Info("Generating Execution plans")
	parentUUID = getUUID()
	for _, combination := range infraCombinations {
		helmDeployments, err := configureHelmDeploymentsForCombination(combination)
		if err != nil {
			log.ErrorAndExit(errors.Wrap(err, "Unable copy Helm deployments"), constant.OsExitCode_1)
		}
		uID := getUUID()
		e := models.ExecutionPlan{
			ParentUUID:      parentUUID,
			UUID:            uID,
			JobID:           fmt.Sprintf(constant.ExecutionPlanIDFmt, uID, combination.OS, combination.JDK, combination.DB),
			HelmDeployments: helmDeployments,
		}
		log.Debug(fmt.Sprintf(msg.ExecutionPlanInfo, parentUUID, uID))
		executionPlans = append(executionPlans, e)
	}
}

func configureHelmDeploymentsForCombination(combination infracombination.Combination) ([]helm.Deployment, error) {
	helmDeployments := make([]helm.Deployment, len(testGridParentPlan.HelmDeployments))
	for i, h := range testGridParentPlan.HelmDeployments {
		helmDeployments[i] = h.ConfigureForCombination(combination)
	}
	return helmDeployments, nil
}

func getUUID() string {
	return uuid.New().String()
}

func generateInfraCombinations() {
	log.Info("Generating Infra combinations")
	var err error
	i := testGridParentPlan.InfraParams
	infraCombinations, err = infracombination.Generate(i.OS, i.DB, i.JDK)
	if err != nil {
		log.ErrorAndExit(errors.Wrap(err, "Unable to generate combinations"), constant.OsExitCode_1)
	}
	if log.IsDebugEnabled() {
		for i, c := range infraCombinations {
			log.Debug(fmt.Sprintf("Combination-%d: %s", i, c.String()))
		}
	}
}

func init() {
	generateTestPlansCmd.Flags().StringVarP(&testGridParentPlanFile, constant.FileFlag, "f", "", "TestGrid parent plan location")
	err := generateTestPlansCmd.MarkFlagRequired(constant.FileFlag)
	if err != nil {
		log.ErrorAndExit(errors.Wrap(err, fmt.Sprintf("Error when making the '%s' required", constant.FileFlag)), constant.OsExitCode_1)
	}
}

func populateTestGridParentPlan() {
	u, err := url.Parse(testGridParentPlanFile)
	if err != nil {
		log.ErrorAndExit(errors.Wrap(err, "Unable to parse the Test Grid parent plan location"), constant.OsExitCode_1)
	}
	var testGridParentPlanData []byte
	switch u.Scheme {
	case constant.HTTPSScheme, constant.HTTPScheme:
		log.Info(fmt.Sprintf("Downloading Test Grid parent plan from: %s", u.String()))
		testGridParentPlanData, err = file.Download(u)
		if err != nil {
			log.ErrorAndExit(errors.Wrap(err, "Unable to download the Test Grid parent plan"), constant.OsExitCode_1)
		}
	case constant.FileScheme:
		log.Info(fmt.Sprintf("Reading Test Grid parent plan from file: %s", u.Path))
		testGridParentPlanData, err = file.Read(u.Path)
		if err != nil {
			log.ErrorAndExit(errors.Wrap(err, "Unable to read the Test Grid parent plan"), constant.OsExitCode_1)
		}
	default:
		log.ErrorAndExit(errors.New("Unable to identify the scheme for Test Grid parent plan location"), constant.OsExitCode_1)
	}
	err = yaml.Parse(testGridParentPlanData, &testGridParentPlan)
	if err != nil {
		log.ErrorAndExit(errors.Wrap(err, "Unable to parse the Test Grid parent plan"), constant.OsExitCode_1)
	}
}
