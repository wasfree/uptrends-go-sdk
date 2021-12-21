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

// MonitorGroupAuthorizationType 
type MonitorGroupAuthorizationType string

// List of MonitorGroupAuthorizationType
const (
	MONITORGROUPAUTHORIZATIONTYPE_VIEW_MONITOR_DATA_IN_GROUP MonitorGroupAuthorizationType = "ViewMonitorDataInGroup"
	MONITORGROUPAUTHORIZATIONTYPE_VIEW_MONITORS_IN_GROUP MonitorGroupAuthorizationType = "ViewMonitorsInGroup"
	MONITORGROUPAUTHORIZATIONTYPE_EDIT_MONITORS_IN_GROUP MonitorGroupAuthorizationType = "EditMonitorsInGroup"
	MONITORGROUPAUTHORIZATIONTYPE_CREATE_AND_DELETE_MONITORS_IN_GROUP MonitorGroupAuthorizationType = "CreateAndDeleteMonitorsInGroup"
)

var allowedMonitorGroupAuthorizationTypeEnumValues = []MonitorGroupAuthorizationType{
	"ViewMonitorDataInGroup",
	"ViewMonitorsInGroup",
	"EditMonitorsInGroup",
	"CreateAndDeleteMonitorsInGroup",
}

func (v *MonitorGroupAuthorizationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MonitorGroupAuthorizationType(value)
	for _, existing := range allowedMonitorGroupAuthorizationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MonitorGroupAuthorizationType", value)
}

// NewMonitorGroupAuthorizationTypeFromValue returns a pointer to a valid MonitorGroupAuthorizationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMonitorGroupAuthorizationTypeFromValue(v string) (*MonitorGroupAuthorizationType, error) {
	ev := MonitorGroupAuthorizationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MonitorGroupAuthorizationType: valid values are %v", v, allowedMonitorGroupAuthorizationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MonitorGroupAuthorizationType) IsValid() bool {
	for _, existing := range allowedMonitorGroupAuthorizationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorGroupAuthorizationType value
func (v MonitorGroupAuthorizationType) Ptr() *MonitorGroupAuthorizationType {
	return &v
}

type NullableMonitorGroupAuthorizationType struct {
	value *MonitorGroupAuthorizationType
	isSet bool
}

func (v NullableMonitorGroupAuthorizationType) Get() *MonitorGroupAuthorizationType {
	return v.value
}

func (v *NullableMonitorGroupAuthorizationType) Set(val *MonitorGroupAuthorizationType) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorGroupAuthorizationType) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorGroupAuthorizationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorGroupAuthorizationType(val *MonitorGroupAuthorizationType) *NullableMonitorGroupAuthorizationType {
	return &NullableMonitorGroupAuthorizationType{value: val, isSet: true}
}

func (v NullableMonitorGroupAuthorizationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorGroupAuthorizationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

