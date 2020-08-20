/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package file handles file operations.
package file

import (
	"github.com/pkg/errors"
	"github.com/wso2/testgrid-core/internal/app/testgrid-core/util/log"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Download downloads the file from the given file url.
func Download(url *url.URL) ([]byte, error) {
	response, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Error("Unable to close download file response body")
		}
	}()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to read download file response body")
	}
	return responseData, nil
}

// Read reads file from the given file path.
func Read(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
