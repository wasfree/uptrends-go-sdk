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

// PresetPeriodType 
type PresetPeriodType string

// List of PresetPeriodType
const (
	CURRENT_DAY PresetPeriodType = "CurrentDay"
	CURRENT_WEEK PresetPeriodType = "CurrentWeek"
	CURRENT_MONTH PresetPeriodType = "CurrentMonth"
	CURRENT_YEAR PresetPeriodType = "CurrentYear"
	PREVIOUS_DAY PresetPeriodType = "PreviousDay"
	PREVIOUS_WEEK PresetPeriodType = "PreviousWeek"
	PREVIOUS_MONTH PresetPeriodType = "PreviousMonth"
	PREVIOUS_YEAR PresetPeriodType = "PreviousYear"
	LAST24_HOURS PresetPeriodType = "Last24Hours"
	LAST48_HOURS PresetPeriodType = "Last48Hours"
	LAST7_DAYS PresetPeriodType = "Last7Days"
	LAST30_DAYS PresetPeriodType = "Last30Days"
	LAST60_DAYS PresetPeriodType = "Last60Days"
	LAST90_DAYS PresetPeriodType = "Last90Days"
	LAST365_DAYS PresetPeriodType = "Last365Days"
	LAST3_MONTHS PresetPeriodType = "Last3Months"
	LAST6_MONTHS PresetPeriodType = "Last6Months"
	LAST12_MONTHS PresetPeriodType = "Last12Months"
	LAST24_MONTHS PresetPeriodType = "Last24Months"
	ALL PresetPeriodType = "All"
)

var allowedPresetPeriodTypeEnumValues = []PresetPeriodType{
	"CurrentDay",
	"CurrentWeek",
	"CurrentMonth",
	"CurrentYear",
	"PreviousDay",
	"PreviousWeek",
	"PreviousMonth",
	"PreviousYear",
	"Last24Hours",
	"Last48Hours",
	"Last7Days",
	"Last30Days",
	"Last60Days",
	"Last90Days",
	"Last365Days",
	"Last3Months",
	"Last6Months",
	"Last12Months",
	"Last24Months",
	"All",
}

func (v *PresetPeriodType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PresetPeriodType(value)
	for _, existing := range allowedPresetPeriodTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PresetPeriodType", value)
}

// NewPresetPeriodTypeFromValue returns a pointer to a valid PresetPeriodType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPresetPeriodTypeFromValue(v string) (*PresetPeriodType, error) {
	ev := PresetPeriodType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PresetPeriodType: valid values are %v", v, allowedPresetPeriodTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PresetPeriodType) IsValid() bool {
	for _, existing := range allowedPresetPeriodTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PresetPeriodType value
func (v PresetPeriodType) Ptr() *PresetPeriodType {
	return &v
}

type NullablePresetPeriodType struct {
	value *PresetPeriodType
	isSet bool
}

func (v NullablePresetPeriodType) Get() *PresetPeriodType {
	return v.value
}

func (v *NullablePresetPeriodType) Set(val *PresetPeriodType) {
	v.value = val
	v.isSet = true
}

func (v NullablePresetPeriodType) IsSet() bool {
	return v.isSet
}

func (v *NullablePresetPeriodType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresetPeriodType(val *PresetPeriodType) *NullablePresetPeriodType {
	return &NullablePresetPeriodType{value: val, isSet: true}
}

func (v NullablePresetPeriodType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresetPeriodType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

