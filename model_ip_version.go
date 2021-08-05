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

// IpVersion 
type IpVersion string

// List of IpVersion
const (
	IPVERSION_IP_V4 IpVersion = "IpV4"
	IPVERSION_IP_V6 IpVersion = "IpV6"
)

var allowedIpVersionEnumValues = []IpVersion{
	"IpV4",
	"IpV6",
}

func (v *IpVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IpVersion(value)
	for _, existing := range allowedIpVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IpVersion", value)
}

// NewIpVersionFromValue returns a pointer to a valid IpVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIpVersionFromValue(v string) (*IpVersion, error) {
	ev := IpVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IpVersion: valid values are %v", v, allowedIpVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IpVersion) IsValid() bool {
	for _, existing := range allowedIpVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IpVersion value
func (v IpVersion) Ptr() *IpVersion {
	return &v
}

type NullableIpVersion struct {
	value *IpVersion
	isSet bool
}

func (v NullableIpVersion) Get() *IpVersion {
	return v.value
}

func (v *NullableIpVersion) Set(val *IpVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableIpVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableIpVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpVersion(val *IpVersion) *NullableIpVersion {
	return &NullableIpVersion{value: val, isSet: true}
}

func (v NullableIpVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
