/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package file handles file operations.
package file

import (
	"github.com/pkg/errors"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/logging"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Download downloads the given file.
func Download(url *url.URL) ([]byte, error) {
	response, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			logging.Error("Unable to close download file response body")
		}
	}()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to read download file response body")
	}
	return responseData, nil
}

// Read reads the given file.
func Read(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
