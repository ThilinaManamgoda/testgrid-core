/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package cmd

import (
	"fmt"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/logging"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	TestGridCoreVersion string
	BuildDate           string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   util.VersionCmdShortMsg,
	Long:    util.VersionCmdLongDesc[1:len(util.VersionCmdLongDesc)],
	Example: util.VersionCmdExamples[1:len(util.VersionCmdExamples)],
	Run:     versionCommand,
}

func versionCommand(cmd *cobra.Command, args []string) {
	logging.Info(fmt.Sprintf("TestGrid Core version: %v", TestGridCoreVersion))
	logging.Info(fmt.Sprintf("Release date: %v", BuildDate))
	logging.Info(fmt.Sprintf("OS\\Arch: %v\\%v", runtime.GOOS, runtime.GOARCH))
	logging.Info(fmt.Sprintf("Go version: %v", runtime.Version()))
}
