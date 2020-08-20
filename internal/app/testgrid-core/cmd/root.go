/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/config"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/constant"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/log"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     "testgrid-core",
	Short:   util.RootCmdShortMsg,
	Long:    util.RootCmdLongDesc[1:len(util.RootCmdLongDesc)],
	Example: util.RootCmdExamples[1:len(util.RootCmdExamples)],
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(generateTestPlansCmd)

	if err := rootCmd.Execute(); err != nil {
		//todo child command
		log.ErrorAndExit(errors.Wrap(err, fmt.Sprintf("Error when executing the '%s' command", rootCmd.Context())), constant.OsExitCode_1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, log.Init)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.testgrid-core.yaml)")
}

func initConfig() {
	config.Init(cfgFile)
}
