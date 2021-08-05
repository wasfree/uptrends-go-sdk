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

// ApiComparisonType 
type ApiComparisonType string

// List of ApiComparisonType
const (
	APICOMPARISONTYPE_EQUAL ApiComparisonType = "Equal"
	APICOMPARISONTYPE_DOES_NOT_EQUAL ApiComparisonType = "DoesNotEqual"
	APICOMPARISONTYPE_CONTAINS_TEXT ApiComparisonType = "ContainsText"
	APICOMPARISONTYPE_DOES_NOT_CONTAIN_TEXT ApiComparisonType = "DoesNotContainText"
	APICOMPARISONTYPE_SHOULD_BE_IGNORED ApiComparisonType = "ShouldBeIgnored"
	APICOMPARISONTYPE_LESS_THAN ApiComparisonType = "LessThan"
	APICOMPARISONTYPE_GREATER_THAN ApiComparisonType = "GreaterThan"
	APICOMPARISONTYPE_LESS_THAN_OR_EQUAL ApiComparisonType = "LessThanOrEqual"
	APICOMPARISONTYPE_GREATER_THAN_OR_EQUAL ApiComparisonType = "GreaterThanOrEqual"
	APICOMPARISONTYPE_IS_NOT_EMPTY ApiComparisonType = "IsNotEmpty"
)

var allowedApiComparisonTypeEnumValues = []ApiComparisonType{
	"Equal",
	"DoesNotEqual",
	"ContainsText",
	"DoesNotContainText",
	"ShouldBeIgnored",
	"LessThan",
	"GreaterThan",
	"LessThanOrEqual",
	"GreaterThanOrEqual",
	"IsNotEmpty",
}

func (v *ApiComparisonType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApiComparisonType(value)
	for _, existing := range allowedApiComparisonTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApiComparisonType", value)
}

// NewApiComparisonTypeFromValue returns a pointer to a valid ApiComparisonType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApiComparisonTypeFromValue(v string) (*ApiComparisonType, error) {
	ev := ApiComparisonType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApiComparisonType: valid values are %v", v, allowedApiComparisonTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApiComparisonType) IsValid() bool {
	for _, existing := range allowedApiComparisonTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApiComparisonType value
func (v ApiComparisonType) Ptr() *ApiComparisonType {
	return &v
}

type NullableApiComparisonType struct {
	value *ApiComparisonType
	isSet bool
}

func (v NullableApiComparisonType) Get() *ApiComparisonType {
	return v.value
}

func (v *NullableApiComparisonType) Set(val *ApiComparisonType) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComparisonType) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComparisonType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComparisonType(val *ApiComparisonType) *NullableApiComparisonType {
	return &NullableApiComparisonType{value: val, isSet: true}
}

func (v NullableApiComparisonType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComparisonType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

