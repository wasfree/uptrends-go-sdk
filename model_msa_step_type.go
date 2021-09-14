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

// MsaStepType the model 'MsaStepType'
type MsaStepType string

// List of MsaStepType
const (
	MSASTEPTYPE_HTTP_REQUEST MsaStepType = "HttpRequest"
	MSASTEPTYPE_DELAY MsaStepType = "Delay"
)

var allowedMsaStepTypeEnumValues = []MsaStepType{
	"HttpRequest",
	"Delay",
}

func (v *MsaStepType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MsaStepType(value)
	for _, existing := range allowedMsaStepTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MsaStepType", value)
}

// NewMsaStepTypeFromValue returns a pointer to a valid MsaStepType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMsaStepTypeFromValue(v string) (*MsaStepType, error) {
	ev := MsaStepType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MsaStepType: valid values are %v", v, allowedMsaStepTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MsaStepType) IsValid() bool {
	for _, existing := range allowedMsaStepTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MsaStepType value
func (v MsaStepType) Ptr() *MsaStepType {
	return &v
}

type NullableMsaStepType struct {
	value *MsaStepType
	isSet bool
}

func (v NullableMsaStepType) Get() *MsaStepType {
	return v.value
}

func (v *NullableMsaStepType) Set(val *MsaStepType) {
	v.value = val
	v.isSet = true
}

func (v NullableMsaStepType) IsSet() bool {
	return v.isSet
}

func (v *NullableMsaStepType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsaStepType(val *MsaStepType) *NullableMsaStepType {
	return &NullableMsaStepType{value: val, isSet: true}
}

func (v NullableMsaStepType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsaStepType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

