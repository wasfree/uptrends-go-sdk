/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package uptrends

import (
	"encoding/json"
	"fmt"
)

// ScheduleMode 
type ScheduleMode string

// List of ScheduleMode
const (
	ONE_TIME ScheduleMode = "OneTime"
	DAILY ScheduleMode = "Daily"
	WEEKLY ScheduleMode = "Weekly"
	MONTHLY ScheduleMode = "Monthly"
)

var allowedScheduleModeEnumValues = []ScheduleMode{
	"OneTime",
	"Daily",
	"Weekly",
	"Monthly",
}

func (v *ScheduleMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScheduleMode(value)
	for _, existing := range allowedScheduleModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScheduleMode", value)
}

// NewScheduleModeFromValue returns a pointer to a valid ScheduleMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScheduleModeFromValue(v string) (*ScheduleMode, error) {
	ev := ScheduleMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScheduleMode: valid values are %v", v, allowedScheduleModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScheduleMode) IsValid() bool {
	for _, existing := range allowedScheduleModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleMode value
func (v ScheduleMode) Ptr() *ScheduleMode {
	return &v
}

type NullableScheduleMode struct {
	value *ScheduleMode
	isSet bool
}

func (v NullableScheduleMode) Get() *ScheduleMode {
	return v.value
}

func (v *NullableScheduleMode) Set(val *ScheduleMode) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleMode) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleMode(val *ScheduleMode) *NullableScheduleMode {
	return &NullableScheduleMode{value: val, isSet: true}
}

func (v NullableScheduleMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

