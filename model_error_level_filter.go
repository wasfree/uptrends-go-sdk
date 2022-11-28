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

// ErrorLevelFilter 
type ErrorLevelFilter string

// List of ErrorLevelFilter
const (
	ERRORLEVELFILTER_NO_ERROR ErrorLevelFilter = "NoError"
	ERRORLEVELFILTER_UNCONFIRMED ErrorLevelFilter = "Unconfirmed"
	ERRORLEVELFILTER_CONFIRMED ErrorLevelFilter = "Confirmed"
)

// All allowed values of ErrorLevelFilter enum
var AllowedErrorLevelFilterEnumValues = []ErrorLevelFilter{
	"NoError",
	"Unconfirmed",
	"Confirmed",
}

func (v *ErrorLevelFilter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorLevelFilter(value)
	for _, existing := range AllowedErrorLevelFilterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorLevelFilter", value)
}

// NewErrorLevelFilterFromValue returns a pointer to a valid ErrorLevelFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewErrorLevelFilterFromValue(v string) (*ErrorLevelFilter, error) {
	ev := ErrorLevelFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ErrorLevelFilter: valid values are %v", v, AllowedErrorLevelFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ErrorLevelFilter) IsValid() bool {
	for _, existing := range AllowedErrorLevelFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorLevelFilter value
func (v ErrorLevelFilter) Ptr() *ErrorLevelFilter {
	return &v
}

type NullableErrorLevelFilter struct {
	value *ErrorLevelFilter
	isSet bool
}

func (v NullableErrorLevelFilter) Get() *ErrorLevelFilter {
	return v.value
}

func (v *NullableErrorLevelFilter) Set(val *ErrorLevelFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorLevelFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorLevelFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorLevelFilter(val *ErrorLevelFilter) *NullableErrorLevelFilter {
	return &NullableErrorLevelFilter{value: val, isSet: true}
}

func (v NullableErrorLevelFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorLevelFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

