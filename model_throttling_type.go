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

// ThrottlingType 
type ThrottlingType string

// List of ThrottlingType
const (
	THROTTLINGTYPE_INACTIVE ThrottlingType = "Inactive"
	THROTTLINGTYPE_BROWSER ThrottlingType = "Browser"
	THROTTLINGTYPE_SIMULATED ThrottlingType = "Simulated"
)

var allowedThrottlingTypeEnumValues = []ThrottlingType{
	"Inactive",
	"Browser",
	"Simulated",
}

func (v *ThrottlingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ThrottlingType(value)
	for _, existing := range allowedThrottlingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ThrottlingType", value)
}

// NewThrottlingTypeFromValue returns a pointer to a valid ThrottlingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewThrottlingTypeFromValue(v string) (*ThrottlingType, error) {
	ev := ThrottlingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ThrottlingType: valid values are %v", v, allowedThrottlingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ThrottlingType) IsValid() bool {
	for _, existing := range allowedThrottlingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ThrottlingType value
func (v ThrottlingType) Ptr() *ThrottlingType {
	return &v
}

type NullableThrottlingType struct {
	value *ThrottlingType
	isSet bool
}

func (v NullableThrottlingType) Get() *ThrottlingType {
	return v.value
}

func (v *NullableThrottlingType) Set(val *ThrottlingType) {
	v.value = val
	v.isSet = true
}

func (v NullableThrottlingType) IsSet() bool {
	return v.isSet
}

func (v *NullableThrottlingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThrottlingType(val *ThrottlingType) *NullableThrottlingType {
	return &NullableThrottlingType{value: val, isSet: true}
}

func (v NullableThrottlingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThrottlingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

