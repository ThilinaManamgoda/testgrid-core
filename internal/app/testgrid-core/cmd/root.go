/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/logging"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
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
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, logging.Init)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.testgrid-core.yaml)")
}

func initConfig() {
	util.InitConfig(cfgFile)
}
