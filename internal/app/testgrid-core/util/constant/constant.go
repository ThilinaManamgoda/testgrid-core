/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

package constant

const OsExitCode_1 = 1

const HTTPSScheme = "https"
const HTTPScheme = "http"
const FileScheme = "file"

const EnvPreFix = "TESTGRID_CORE"
const LogLevelKey = "log.level"
const LogLevelDefault = "info"

//todo
const InfraCombinationGeneratorKey = "infraCombinationGenerator"
const InfraCombinationGeneratorDefault = "dummy-generator"

const FileFlag = "file"

// Database constants.
const MySQL = "mysql"
const DatabaseUserNameKey = "db.username"
const DatabaseUserNameDefault = "root"
const DatabasePasswordKey = "db.password"
const DatabasePasswordDefault = "root123"
const DatabaseHostKey = "db.host"
const DatabaseHostDefault = "localhost"
const DatabasePortKey = "db.port"
const DatabasePortDefault = 3306
const DatabaseNameKey = "db.name"
const DatabaseNameDefault = "testgrid"
const DatabaseLogLevelKey = "db.logLevel"
const DatabaseLogLevelDefault = true
const DatabaseMaxRetriesKey = "db.maxRetries"
const DatabaseMaxRetriesDefault = "3"
const HelmParamTable = "helm_params"
const InfraParamTable = "infra_params"
