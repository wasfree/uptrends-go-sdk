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

// TlsVersion 
type TlsVersion string

// List of TlsVersion
const (
	TLSVERSION_TLS12 TlsVersion = "Tls12"
	TLSVERSION_TLS11 TlsVersion = "Tls11"
	TLSVERSION_TLS10 TlsVersion = "Tls10"
	TLSVERSION_TLS12_TLS11 TlsVersion = "Tls12_Tls11"
	TLSVERSION_TLS11_TLS10 TlsVersion = "Tls11_Tls10"
	TLSVERSION_TLS12_TLS11_TLS10 TlsVersion = "Tls12_Tls11_Tls10"
	TLSVERSION_TLS12_TLS11_TLS10_WITH_FALLBACK TlsVersion = "Tls12_Tls11_Tls10_WithFallback"
)

var allowedTlsVersionEnumValues = []TlsVersion{
	"Tls12",
	"Tls11",
	"Tls10",
	"Tls12_Tls11",
	"Tls11_Tls10",
	"Tls12_Tls11_Tls10",
	"Tls12_Tls11_Tls10_WithFallback",
}

func (v *TlsVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TlsVersion(value)
	for _, existing := range allowedTlsVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TlsVersion", value)
}

// NewTlsVersionFromValue returns a pointer to a valid TlsVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTlsVersionFromValue(v string) (*TlsVersion, error) {
	ev := TlsVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TlsVersion: valid values are %v", v, allowedTlsVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TlsVersion) IsValid() bool {
	for _, existing := range allowedTlsVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TlsVersion value
func (v TlsVersion) Ptr() *TlsVersion {
	return &v
}

type NullableTlsVersion struct {
	value *TlsVersion
	isSet bool
}

func (v NullableTlsVersion) Get() *TlsVersion {
	return v.value
}

func (v *NullableTlsVersion) Set(val *TlsVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableTlsVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableTlsVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTlsVersion(val *TlsVersion) *NullableTlsVersion {
	return &NullableTlsVersion{value: val, isSet: true}
}

func (v NullableTlsVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTlsVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
