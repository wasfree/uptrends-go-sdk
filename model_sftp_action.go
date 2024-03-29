/*
Uptrends API v4

This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package uptrends

import (
	"encoding/json"
	"fmt"
)

// SftpAction 
type SftpAction string

// List of SftpAction
const (
	SFTPACTION_CONNECT_ONLY SftpAction = "ConnectOnly"
	SFTPACTION_TEST_FILE_EXISTS SftpAction = "TestFileExists"
	SFTPACTION_TEST_FILE_DOES_NOT_EXIST SftpAction = "TestFileDoesNotExist"
	SFTPACTION_DOWN_LOAD_FILE SftpAction = "DownLoadFile"
)

// All allowed values of SftpAction enum
var AllowedSftpActionEnumValues = []SftpAction{
	"ConnectOnly",
	"TestFileExists",
	"TestFileDoesNotExist",
	"DownLoadFile",
}

func (v *SftpAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SftpAction(value)
	for _, existing := range AllowedSftpActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SftpAction", value)
}

// NewSftpActionFromValue returns a pointer to a valid SftpAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSftpActionFromValue(v string) (*SftpAction, error) {
	ev := SftpAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SftpAction: valid values are %v", v, AllowedSftpActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SftpAction) IsValid() bool {
	for _, existing := range AllowedSftpActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SftpAction value
func (v SftpAction) Ptr() *SftpAction {
	return &v
}

type NullableSftpAction struct {
	value *SftpAction
	isSet bool
}

func (v NullableSftpAction) Get() *SftpAction {
	return v.value
}

func (v *NullableSftpAction) Set(val *SftpAction) {
	v.value = val
	v.isSet = true
}

func (v NullableSftpAction) IsSet() bool {
	return v.isSet
}

func (v *NullableSftpAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSftpAction(val *SftpAction) *NullableSftpAction {
	return &NullableSftpAction{value: val, isSet: true}
}

func (v NullableSftpAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSftpAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

