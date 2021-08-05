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

// VaultSectionAuthorizationType 
type VaultSectionAuthorizationType string

// List of VaultSectionAuthorizationType
const (
	VAULTSECTIONAUTHORIZATIONTYPE_VIEW_VAULT_SECTION VaultSectionAuthorizationType = "ViewVaultSection"
	VAULTSECTIONAUTHORIZATIONTYPE_CHANGE_VAULT_SECTION VaultSectionAuthorizationType = "ChangeVaultSection"
)

var allowedVaultSectionAuthorizationTypeEnumValues = []VaultSectionAuthorizationType{
	"ViewVaultSection",
	"ChangeVaultSection",
}

func (v *VaultSectionAuthorizationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VaultSectionAuthorizationType(value)
	for _, existing := range allowedVaultSectionAuthorizationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VaultSectionAuthorizationType", value)
}

// NewVaultSectionAuthorizationTypeFromValue returns a pointer to a valid VaultSectionAuthorizationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVaultSectionAuthorizationTypeFromValue(v string) (*VaultSectionAuthorizationType, error) {
	ev := VaultSectionAuthorizationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VaultSectionAuthorizationType: valid values are %v", v, allowedVaultSectionAuthorizationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VaultSectionAuthorizationType) IsValid() bool {
	for _, existing := range allowedVaultSectionAuthorizationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VaultSectionAuthorizationType value
func (v VaultSectionAuthorizationType) Ptr() *VaultSectionAuthorizationType {
	return &v
}

type NullableVaultSectionAuthorizationType struct {
	value *VaultSectionAuthorizationType
	isSet bool
}

func (v NullableVaultSectionAuthorizationType) Get() *VaultSectionAuthorizationType {
	return v.value
}

func (v *NullableVaultSectionAuthorizationType) Set(val *VaultSectionAuthorizationType) {
	v.value = val
	v.isSet = true
}

func (v NullableVaultSectionAuthorizationType) IsSet() bool {
	return v.isSet
}

func (v *NullableVaultSectionAuthorizationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVaultSectionAuthorizationType(val *VaultSectionAuthorizationType) *NullableVaultSectionAuthorizationType {
	return &NullableVaultSectionAuthorizationType{value: val, isSet: true}
}

func (v NullableVaultSectionAuthorizationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVaultSectionAuthorizationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

