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
const InfraCombinationGeneratorIDKey = "infraCombinationGeneratorID"
const InfraCombinationGeneratorIDDefault = "dummy-generator"
const DummyGeneratorID = "dummy-generator"

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
const HelmParamTblParamKeyCol = "param_key"
const HelmParamTblParamKeyQuery = HelmParamTblParamKeyCol + " = ?"

// Redis client constants.
const RedisServerAddressKey = "redis.address"
const RedisServerAddressDefault = "localhost:6379"
const RedisServerPasswordKey = "redis.password"
const RedisServerPasswordDefault = ""
const RedisServerDBKey = "redis.db"
const RedisServerDBDefault = 0

// Execution Plan constants.
const ExecutionPlanIDFmt = "%s-%s-%s-%s"
const DeployCMDRedisID = "cmd-deploy:%s-%s-deployment-%d"
const DeployCMDExePlanID = "test-plan-%s"
