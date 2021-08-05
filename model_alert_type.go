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

// AlertType 
type AlertType string

// List of AlertType
const (
	OK AlertType = "Ok"
	ERROR AlertType = "Error"
	REMINDER AlertType = "Reminder"
)

var allowedAlertTypeEnumValues = []AlertType{
	"Ok",
	"Error",
	"Reminder",
}

func (v *AlertType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlertType(value)
	for _, existing := range allowedAlertTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlertType", value)
}

// NewAlertTypeFromValue returns a pointer to a valid AlertType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlertTypeFromValue(v string) (*AlertType, error) {
	ev := AlertType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlertType: valid values are %v", v, allowedAlertTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlertType) IsValid() bool {
	for _, existing := range allowedAlertTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertType value
func (v AlertType) Ptr() *AlertType {
	return &v
}

type NullableAlertType struct {
	value *AlertType
	isSet bool
}

func (v NullableAlertType) Get() *AlertType {
	return v.value
}

func (v *NullableAlertType) Set(val *AlertType) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertType) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertType(val *AlertType) *NullableAlertType {
	return &NullableAlertType{value: val, isSet: true}
}

func (v NullableAlertType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

