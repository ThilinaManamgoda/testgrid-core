/*
 * Copyright (c) 2020, WSO2 Inc. All Rights Reserved.
 */

// Package json holds the implementations for JSON data handling.
package json

import "encoding/json"

// ToString converts the given type to JSON string byte array.
func ToString(t interface{}) ([]byte, error) {
	d, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// Parse parses the given JSON string byte array to given type.
func Parse(d []byte, t interface{}) error {
	err := json.Unmarshal(d, t)
	if err != nil {
		return err
	}
	return nil
}
