/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package main

import "github.com/wso2/testgrid-core/internal/app/testgrid-core/cmd"

var version string
var buildDate string

func main() {
	cmd.TestGridCoreVersion = version
	cmd.BuildDate = buildDate
	cmd.Execute()
}
