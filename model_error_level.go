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

// ErrorLevel the model 'ErrorLevel'
type ErrorLevel string

// List of ErrorLevel
const (
	ERRORLEVEL_NO_ERROR ErrorLevel = "NoError"
	ERRORLEVEL_UNCONFIRMED ErrorLevel = "Unconfirmed"
	ERRORLEVEL_CONFIRMED ErrorLevel = "Confirmed"
)

var allowedErrorLevelEnumValues = []ErrorLevel{
	"NoError",
	"Unconfirmed",
	"Confirmed",
}

func (v *ErrorLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorLevel(value)
	for _, existing := range allowedErrorLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorLevel", value)
}

// NewErrorLevelFromValue returns a pointer to a valid ErrorLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewErrorLevelFromValue(v string) (*ErrorLevel, error) {
	ev := ErrorLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ErrorLevel: valid values are %v", v, allowedErrorLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ErrorLevel) IsValid() bool {
	for _, existing := range allowedErrorLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorLevel value
func (v ErrorLevel) Ptr() *ErrorLevel {
	return &v
}

type NullableErrorLevel struct {
	value *ErrorLevel
	isSet bool
}

func (v NullableErrorLevel) Get() *ErrorLevel {
	return v.value
}

func (v *NullableErrorLevel) Set(val *ErrorLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorLevel(val *ErrorLevel) *NullableErrorLevel {
	return &NullableErrorLevel{value: val, isSet: true}
}

func (v NullableErrorLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

