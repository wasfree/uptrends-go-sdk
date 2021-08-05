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

// RegisterStatus 
type RegisterStatus string

// List of RegisterStatus
const (
	OK RegisterStatus = "OK"
	UNEXPECTED_ERROR RegisterStatus = "UnexpectedError"
)

var allowedRegisterStatusEnumValues = []RegisterStatus{
	"OK",
	"UnexpectedError",
}

func (v *RegisterStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RegisterStatus(value)
	for _, existing := range allowedRegisterStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RegisterStatus", value)
}

// NewRegisterStatusFromValue returns a pointer to a valid RegisterStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRegisterStatusFromValue(v string) (*RegisterStatus, error) {
	ev := RegisterStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RegisterStatus: valid values are %v", v, allowedRegisterStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RegisterStatus) IsValid() bool {
	for _, existing := range allowedRegisterStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RegisterStatus value
func (v RegisterStatus) Ptr() *RegisterStatus {
	return &v
}

type NullableRegisterStatus struct {
	value *RegisterStatus
	isSet bool
}

func (v NullableRegisterStatus) Get() *RegisterStatus {
	return v.value
}

func (v *NullableRegisterStatus) Set(val *RegisterStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterStatus(val *RegisterStatus) *NullableRegisterStatus {
	return &NullableRegisterStatus{value: val, isSet: true}
}

func (v NullableRegisterStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

