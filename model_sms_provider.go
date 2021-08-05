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

// SmsProvider 
type SmsProvider string

// List of SmsProvider
const (
	USE_ACCOUNT_SETTING SmsProvider = "UseAccountSetting"
	SMS_PROVIDER_EUROPE SmsProvider = "SmsProviderEurope"
	SMS_PROVIDER_INTERNATIONAL SmsProvider = "SmsProviderInternational"
	SMS_PROVIDER_EUROPE2 SmsProvider = "SmsProviderEurope2"
	SMS_PROVIDER_USA SmsProvider = "SmsProviderUSA"
)

var allowedSmsProviderEnumValues = []SmsProvider{
	"UseAccountSetting",
	"SmsProviderEurope",
	"SmsProviderInternational",
	"SmsProviderEurope2",
	"SmsProviderUSA",
}

func (v *SmsProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SmsProvider(value)
	for _, existing := range allowedSmsProviderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SmsProvider", value)
}

// NewSmsProviderFromValue returns a pointer to a valid SmsProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSmsProviderFromValue(v string) (*SmsProvider, error) {
	ev := SmsProvider(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SmsProvider: valid values are %v", v, allowedSmsProviderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SmsProvider) IsValid() bool {
	for _, existing := range allowedSmsProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SmsProvider value
func (v SmsProvider) Ptr() *SmsProvider {
	return &v
}

type NullableSmsProvider struct {
	value *SmsProvider
	isSet bool
}

func (v NullableSmsProvider) Get() *SmsProvider {
	return v.value
}

func (v *NullableSmsProvider) Set(val *SmsProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsProvider(val *SmsProvider) *NullableSmsProvider {
	return &NullableSmsProvider{value: val, isSet: true}
}

func (v NullableSmsProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

