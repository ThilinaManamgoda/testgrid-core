/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package db handles database interactions.
package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/constant"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/log"
	"math"
	"strconv"
	"time"
)

var (
	url        string
	userName   string
	password   string
	host       string
	port       int
	dbName     string
	logMode    bool
	maxRetries int
	db         *gorm.DB
)

// Entity represents a table in the database.
type Entity interface {
	TableName() string
}

func backOff(min, max time.Duration, attempt int) time.Duration {
	du := math.Pow(2, float64(attempt))
	sleep := time.Duration(du) * time.Second
	if sleep < min {
		return min
	}
	if sleep > max {
		return max
	}
	return sleep
}

// Init initialize database parameters and open a DB connection.
func Init() error {
	setConfigs()
	return connect()
}

func setConfigs() {
	userName = viper.GetString(constant.DatabaseUserNameKey)
	password = viper.GetString(constant.DatabasePasswordKey)
	host = viper.GetString(constant.DatabaseHostKey)
	port = viper.GetInt(constant.DatabasePortKey)
	dbName = viper.GetString(constant.DatabaseNameKey)
	logMode = viper.GetBool(constant.DatabaseLogLevelKey)
	maxRetries = viper.GetInt(constant.DatabaseMaxRetriesKey)
}

// connect starts a DB connection and returns any error occurred.
func connect() error {
	url = userName + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + dbName + "?charset=utf8"
	var err error
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(constant.MySQL, url)
		if err == nil {
			break
		}
		bt := backOff(1*time.Second, 60*time.Second, i)

		log.Debug(fmt.Sprintf("Retrying the DB connection. err: %v", err))
		time.Sleep(bt)
	}
	if err != nil {
		return errors.Wrap(err, "Cannot initiate database connection")
	}
	if logMode {
		log.Debug("Debug logs are enabled for Database client")
		db.LogMode(logMode)
		db.SetLogger(&log.GormLogger{})
	}
	return nil
}

// CloseDBCon function closes the open DB connections.
func CloseDBCon() {
	log.Debug("closing DB connection")
	if err := db.Close(); err != nil {
		log.Error(fmt.Sprintf("unable to close the DB connection %v", err))
	}
}

// Store saves the given ServiceInstance in the Database.
// Returns any error encountered.
func Store(e Entity) error {
	return db.Table(e.TableName()).Create(e).Error
}

// Update updates the given ServiceInstance in the Database.
// Returns any error encountered.
func Update(e Entity) error {
	return db.Table(e.TableName()).Save(e).Error
}

// Delete deletes the given entity from the Database.
// Returns any error encountered.
func Delete(e Entity) error {
	return db.Table(e.TableName()).Delete(e).Error
}

// Retrieve function initialize the given entity from the database if exists.
// Returns true if the entity exists and any error encountered.
func Retrieve(e Entity) (bool, error) {
	result := db.Table(e.TableName()).Where(e).Find(e)
	if result.Error != nil {
		if result.RecordNotFound() {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func RetrieveHelmParamListForParamKey(paramKey string, helmParams interface{}) (bool, error) {
	result := db.Where(constant.HelmParamTblParamKeyQuery, paramKey).Find(helmParams)
	if result.Error != nil {
		if result.RecordNotFound() {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}
