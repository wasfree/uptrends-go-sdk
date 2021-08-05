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

// ApiHttpAuthenticationType 
type ApiHttpAuthenticationType string

// List of ApiHttpAuthenticationType
const (
	APIHTTPAUTHENTICATIONTYPE_NONE ApiHttpAuthenticationType = "None"
	APIHTTPAUTHENTICATIONTYPE_BASIC ApiHttpAuthenticationType = "Basic"
	APIHTTPAUTHENTICATIONTYPE_NTLM ApiHttpAuthenticationType = "NTLM"
	APIHTTPAUTHENTICATIONTYPE_DIGEST ApiHttpAuthenticationType = "Digest"
)

var allowedApiHttpAuthenticationTypeEnumValues = []ApiHttpAuthenticationType{
	"None",
	"Basic",
	"NTLM",
	"Digest",
}

func (v *ApiHttpAuthenticationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApiHttpAuthenticationType(value)
	for _, existing := range allowedApiHttpAuthenticationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApiHttpAuthenticationType", value)
}

// NewApiHttpAuthenticationTypeFromValue returns a pointer to a valid ApiHttpAuthenticationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApiHttpAuthenticationTypeFromValue(v string) (*ApiHttpAuthenticationType, error) {
	ev := ApiHttpAuthenticationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApiHttpAuthenticationType: valid values are %v", v, allowedApiHttpAuthenticationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApiHttpAuthenticationType) IsValid() bool {
	for _, existing := range allowedApiHttpAuthenticationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApiHttpAuthenticationType value
func (v ApiHttpAuthenticationType) Ptr() *ApiHttpAuthenticationType {
	return &v
}

type NullableApiHttpAuthenticationType struct {
	value *ApiHttpAuthenticationType
	isSet bool
}

func (v NullableApiHttpAuthenticationType) Get() *ApiHttpAuthenticationType {
	return v.value
}

func (v *NullableApiHttpAuthenticationType) Set(val *ApiHttpAuthenticationType) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHttpAuthenticationType) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHttpAuthenticationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHttpAuthenticationType(val *ApiHttpAuthenticationType) *NullableApiHttpAuthenticationType {
	return &NullableApiHttpAuthenticationType{value: val, isSet: true}
}

func (v NullableApiHttpAuthenticationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHttpAuthenticationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

