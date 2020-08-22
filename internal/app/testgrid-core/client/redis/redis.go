/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package redis handles interaction with Redis server.
package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/constant"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/json"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/log"
	"time"
)

var (
	ctx      = context.Background()
	addr     string
	password string
	db       int
	rdb      *redis.Client
)

// Init initializes the redis client.
func Init() error {
	setConfigs()
	return connect()
}

func setConfigs() {
	addr = viper.GetString(constant.RedisServerAddressKey)
	password = viper.GetString(constant.RedisServerPasswordKey)
	db = viper.GetInt(constant.RedisServerDBKey)
}

func connect() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	log.Debug(fmt.Sprintf("Connected to the redis server running at '%v'", addr))
	return nil
}

// Set function sets the given key, value  without an expiration.
func Set(key string, val interface{}) error {
	d, err := json.ToString(val)
	if err != nil {
		return errors.Wrap(err, "Unable to convert given val to string")
	}
	err = rdb.Set(ctx, key, string(d), 0).Err()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Unable to set key: %s", key))
	}
	return nil
}

// Set function sets the given key, value  with the given expiration.
func SetWithExp(key string, val interface{}, exp time.Duration) error {
	d, err := json.ToString(val)
	if err != nil {
		return errors.Wrap(err, "Unable to convert given val to string")
	}
	err = rdb.Set(ctx, key, string(d), exp).Err()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Unable to set key: %s with exipration: %v", key, exp))
	}
	return nil
}

// Get function gets the value for a given key.
func Get(key string, t interface{}) error {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Unable to get value for key: %s", key))
	}
	err = json.Parse([]byte(val), t)
	if err != nil {
		return errors.Wrap(err, "Unable to parse to given struct value")
	}
	return nil
}

func SetBulk(m map[string]interface{}) error {
	pipe := rdb.Pipeline()
	for key, val := range m {
		d, err := json.ToString(val)
		if err != nil {
			return errors.Wrap(err, "Unable to convert given val to string")
		}
		pipe.Set(ctx, key, string(d), 0)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "Unable to insert bulk")
	}
	return nil
}
