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

// SelectedPeriodType 
type SelectedPeriodType string

// List of SelectedPeriodType
const (
	PRESET_PERIOD SelectedPeriodType = "PresetPeriod"
	SPECIFIC_DATES SelectedPeriodType = "SpecificDates"
)

var allowedSelectedPeriodTypeEnumValues = []SelectedPeriodType{
	"PresetPeriod",
	"SpecificDates",
}

func (v *SelectedPeriodType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SelectedPeriodType(value)
	for _, existing := range allowedSelectedPeriodTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SelectedPeriodType", value)
}

// NewSelectedPeriodTypeFromValue returns a pointer to a valid SelectedPeriodType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSelectedPeriodTypeFromValue(v string) (*SelectedPeriodType, error) {
	ev := SelectedPeriodType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SelectedPeriodType: valid values are %v", v, allowedSelectedPeriodTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SelectedPeriodType) IsValid() bool {
	for _, existing := range allowedSelectedPeriodTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SelectedPeriodType value
func (v SelectedPeriodType) Ptr() *SelectedPeriodType {
	return &v
}

type NullableSelectedPeriodType struct {
	value *SelectedPeriodType
	isSet bool
}

func (v NullableSelectedPeriodType) Get() *SelectedPeriodType {
	return v.value
}

func (v *NullableSelectedPeriodType) Set(val *SelectedPeriodType) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectedPeriodType) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectedPeriodType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectedPeriodType(val *SelectedPeriodType) *NullableSelectedPeriodType {
	return &NullableSelectedPeriodType{value: val, isSet: true}
}

func (v NullableSelectedPeriodType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectedPeriodType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

