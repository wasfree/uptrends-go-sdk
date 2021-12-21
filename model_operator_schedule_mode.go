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

// OperatorScheduleMode 
type OperatorScheduleMode string

// List of OperatorScheduleMode
const (
	OPERATORSCHEDULEMODE_ONE_TIME OperatorScheduleMode = "OneTime"
	OPERATORSCHEDULEMODE_DAILY OperatorScheduleMode = "Daily"
	OPERATORSCHEDULEMODE_WEEKLY OperatorScheduleMode = "Weekly"
	OPERATORSCHEDULEMODE_MONTHLY OperatorScheduleMode = "Monthly"
)

var allowedOperatorScheduleModeEnumValues = []OperatorScheduleMode{
	"OneTime",
	"Daily",
	"Weekly",
	"Monthly",
}

func (v *OperatorScheduleMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OperatorScheduleMode(value)
	for _, existing := range allowedOperatorScheduleModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OperatorScheduleMode", value)
}

// NewOperatorScheduleModeFromValue returns a pointer to a valid OperatorScheduleMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperatorScheduleModeFromValue(v string) (*OperatorScheduleMode, error) {
	ev := OperatorScheduleMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperatorScheduleMode: valid values are %v", v, allowedOperatorScheduleModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OperatorScheduleMode) IsValid() bool {
	for _, existing := range allowedOperatorScheduleModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OperatorScheduleMode value
func (v OperatorScheduleMode) Ptr() *OperatorScheduleMode {
	return &v
}

type NullableOperatorScheduleMode struct {
	value *OperatorScheduleMode
	isSet bool
}

func (v NullableOperatorScheduleMode) Get() *OperatorScheduleMode {
	return v.value
}

func (v *NullableOperatorScheduleMode) Set(val *OperatorScheduleMode) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorScheduleMode) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorScheduleMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorScheduleMode(val *OperatorScheduleMode) *NullableOperatorScheduleMode {
	return &NullableOperatorScheduleMode{value: val, isSet: true}
}

func (v NullableOperatorScheduleMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorScheduleMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

