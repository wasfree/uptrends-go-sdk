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

// MsaBodyType 
type MsaBodyType string

// List of MsaBodyType
const (
	MSABODYTYPE_RAW MsaBodyType = "Raw"
	MSABODYTYPE_VAULT_FILES MsaBodyType = "VaultFiles"
)

var allowedMsaBodyTypeEnumValues = []MsaBodyType{
	"Raw",
	"VaultFiles",
}

func (v *MsaBodyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MsaBodyType(value)
	for _, existing := range allowedMsaBodyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MsaBodyType", value)
}

// NewMsaBodyTypeFromValue returns a pointer to a valid MsaBodyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMsaBodyTypeFromValue(v string) (*MsaBodyType, error) {
	ev := MsaBodyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MsaBodyType: valid values are %v", v, allowedMsaBodyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MsaBodyType) IsValid() bool {
	for _, existing := range allowedMsaBodyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MsaBodyType value
func (v MsaBodyType) Ptr() *MsaBodyType {
	return &v
}

type NullableMsaBodyType struct {
	value *MsaBodyType
	isSet bool
}

func (v NullableMsaBodyType) Get() *MsaBodyType {
	return v.value
}

func (v *NullableMsaBodyType) Set(val *MsaBodyType) {
	v.value = val
	v.isSet = true
}

func (v NullableMsaBodyType) IsSet() bool {
	return v.isSet
}

func (v *NullableMsaBodyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsaBodyType(val *MsaBodyType) *NullableMsaBodyType {
	return &NullableMsaBodyType{value: val, isSet: true}
}

func (v NullableMsaBodyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsaBodyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

