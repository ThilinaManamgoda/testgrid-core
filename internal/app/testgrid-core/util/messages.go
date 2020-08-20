/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package util contains utilities.
package util

import (
	"github.com/lithammer/dedent"
	"runtime"
)

const ExecutableName = "testgrid-core_" + runtime.GOOS

const RootCmdShortMsg string = "Update tool keeps WSO2 products up-to-date."

var RootCmdLongDesc = dedent.Dedent(`
TestGrid Core is a simple command-line tool that connects to the  
WSO2 Update service, determines which updates are new and relevant, and 
downloads them and updates the product.

To find out the latest on WSO2 Update, visit https://wso2.com/updates
`)

var RootCmdExamples = dedent.Dedent(`
	# Check the current version of TestGrid Core in use on your system
  	` + ExecutableName + ` version
`)

const VersionCmdShortMsg string = "Print the TestGrid Core version information."

var VersionCmdLongDesc = dedent.Dedent(`
Version command prints the TestGrid Core distribution version information
such as TestGrid Core version, release date, operating system
Architecture and Go version.
		`)

var VersionCmdExamples = dedent.Dedent(`
		# Get update tool version information
		  ` + ExecutableName + ` version
		`)

const GenerateTestPlansCmdShortMsg string = "Generates TestGrid plans."

var GenerateTestPlansCmdLongDesc = dedent.Dedent(`
GenerateTestPlans command generates test plans for a given parent plan definition.
		`)

var GenerateTestPlansCmdExamples = dedent.Dedent(`
		# Generate test plans
		  ` + ExecutableName + ` generate-test-plans --file <TESTGRID_YAML>
		`)
