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

// LastErrorLevel 
type LastErrorLevel string

// List of LastErrorLevel
const (
	LASTERRORLEVEL_NO_ERROR LastErrorLevel = "NoError"
	LASTERRORLEVEL_UNCONFIRMED LastErrorLevel = "Unconfirmed"
	LASTERRORLEVEL_CONFIRMED LastErrorLevel = "Confirmed"
	LASTERRORLEVEL_INACTIVE LastErrorLevel = "Inactive"
	LASTERRORLEVEL_INCONCLUSIVE LastErrorLevel = "Inconclusive"
)

// All allowed values of LastErrorLevel enum
var AllowedLastErrorLevelEnumValues = []LastErrorLevel{
	"NoError",
	"Unconfirmed",
	"Confirmed",
	"Inactive",
	"Inconclusive",
}

func (v *LastErrorLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LastErrorLevel(value)
	for _, existing := range AllowedLastErrorLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LastErrorLevel", value)
}

// NewLastErrorLevelFromValue returns a pointer to a valid LastErrorLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLastErrorLevelFromValue(v string) (*LastErrorLevel, error) {
	ev := LastErrorLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LastErrorLevel: valid values are %v", v, AllowedLastErrorLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LastErrorLevel) IsValid() bool {
	for _, existing := range AllowedLastErrorLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LastErrorLevel value
func (v LastErrorLevel) Ptr() *LastErrorLevel {
	return &v
}

type NullableLastErrorLevel struct {
	value *LastErrorLevel
	isSet bool
}

func (v NullableLastErrorLevel) Get() *LastErrorLevel {
	return v.value
}

func (v *NullableLastErrorLevel) Set(val *LastErrorLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableLastErrorLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableLastErrorLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastErrorLevel(val *LastErrorLevel) *NullableLastErrorLevel {
	return &NullableLastErrorLevel{value: val, isSet: true}
}

func (v NullableLastErrorLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastErrorLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

