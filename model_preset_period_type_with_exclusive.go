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

// PresetPeriodTypeWithExclusive 
type PresetPeriodTypeWithExclusive string

// List of PresetPeriodTypeWithExclusive
const (
	PRESETPERIODTYPEWITHEXCLUSIVE_CURRENT_DAY PresetPeriodTypeWithExclusive = "CurrentDay"
	PRESETPERIODTYPEWITHEXCLUSIVE_CURRENT_WEEK PresetPeriodTypeWithExclusive = "CurrentWeek"
	PRESETPERIODTYPEWITHEXCLUSIVE_CURRENT_MONTH PresetPeriodTypeWithExclusive = "CurrentMonth"
	PRESETPERIODTYPEWITHEXCLUSIVE_CURRENT_QUARTER PresetPeriodTypeWithExclusive = "CurrentQuarter"
	PRESETPERIODTYPEWITHEXCLUSIVE_CURRENT_YEAR PresetPeriodTypeWithExclusive = "CurrentYear"
	PRESETPERIODTYPEWITHEXCLUSIVE_PREVIOUS_DAY PresetPeriodTypeWithExclusive = "PreviousDay"
	PRESETPERIODTYPEWITHEXCLUSIVE_PREVIOUS_WEEK PresetPeriodTypeWithExclusive = "PreviousWeek"
	PRESETPERIODTYPEWITHEXCLUSIVE_PREVIOUS_MONTH PresetPeriodTypeWithExclusive = "PreviousMonth"
	PRESETPERIODTYPEWITHEXCLUSIVE_PREVIOUS_QUARTER PresetPeriodTypeWithExclusive = "PreviousQuarter"
	PRESETPERIODTYPEWITHEXCLUSIVE_PREVIOUS_YEAR PresetPeriodTypeWithExclusive = "PreviousYear"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST2_HOURS PresetPeriodTypeWithExclusive = "Last2Hours"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST6_HOURS PresetPeriodTypeWithExclusive = "Last6Hours"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST12_HOURS PresetPeriodTypeWithExclusive = "Last12Hours"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST24_HOURS PresetPeriodTypeWithExclusive = "Last24Hours"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST48_HOURS PresetPeriodTypeWithExclusive = "Last48Hours"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST7_DAYS PresetPeriodTypeWithExclusive = "Last7Days"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST30_DAYS PresetPeriodTypeWithExclusive = "Last30Days"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST60_DAYS PresetPeriodTypeWithExclusive = "Last60Days"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST90_DAYS PresetPeriodTypeWithExclusive = "Last90Days"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST365_DAYS PresetPeriodTypeWithExclusive = "Last365Days"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST3_MONTHS PresetPeriodTypeWithExclusive = "Last3Months"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST6_MONTHS PresetPeriodTypeWithExclusive = "Last6Months"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST12_MONTHS PresetPeriodTypeWithExclusive = "Last12Months"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST24_MONTHS PresetPeriodTypeWithExclusive = "Last24Months"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST2_HOURS_EXCL PresetPeriodTypeWithExclusive = "Last2HoursExcl"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST6_HOURS_EXCL PresetPeriodTypeWithExclusive = "Last6HoursExcl"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST12_HOURS_EXCL PresetPeriodTypeWithExclusive = "Last12HoursExcl"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST24_HOURS_EXCL PresetPeriodTypeWithExclusive = "Last24HoursExcl"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST48_HOURS_EXCL PresetPeriodTypeWithExclusive = "Last48HoursExcl"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST7_DAYS_EXCL PresetPeriodTypeWithExclusive = "Last7DaysExcl"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST30_DAYS_EXCL PresetPeriodTypeWithExclusive = "Last30DaysExcl"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST60_DAYS_EXCL PresetPeriodTypeWithExclusive = "Last60DaysExcl"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST90_DAYS_EXCL PresetPeriodTypeWithExclusive = "Last90DaysExcl"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST365_DAYS_EXCL PresetPeriodTypeWithExclusive = "Last365DaysExcl"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST3_MONTHS_EXCL PresetPeriodTypeWithExclusive = "Last3MonthsExcl"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST6_MONTHS_EXCL PresetPeriodTypeWithExclusive = "Last6MonthsExcl"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST12_MONTHS_EXCL PresetPeriodTypeWithExclusive = "Last12MonthsExcl"
	PRESETPERIODTYPEWITHEXCLUSIVE_LAST24_MONTHS_EXCL PresetPeriodTypeWithExclusive = "Last24MonthsExcl"
	PRESETPERIODTYPEWITHEXCLUSIVE_ALL PresetPeriodTypeWithExclusive = "All"
)

// All allowed values of PresetPeriodTypeWithExclusive enum
var AllowedPresetPeriodTypeWithExclusiveEnumValues = []PresetPeriodTypeWithExclusive{
	"CurrentDay",
	"CurrentWeek",
	"CurrentMonth",
	"CurrentQuarter",
	"CurrentYear",
	"PreviousDay",
	"PreviousWeek",
	"PreviousMonth",
	"PreviousQuarter",
	"PreviousYear",
	"Last2Hours",
	"Last6Hours",
	"Last12Hours",
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
	"Last2HoursExcl",
	"Last6HoursExcl",
	"Last12HoursExcl",
	"Last24HoursExcl",
	"Last48HoursExcl",
	"Last7DaysExcl",
	"Last30DaysExcl",
	"Last60DaysExcl",
	"Last90DaysExcl",
	"Last365DaysExcl",
	"Last3MonthsExcl",
	"Last6MonthsExcl",
	"Last12MonthsExcl",
	"Last24MonthsExcl",
	"All",
}

func (v *PresetPeriodTypeWithExclusive) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PresetPeriodTypeWithExclusive(value)
	for _, existing := range AllowedPresetPeriodTypeWithExclusiveEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PresetPeriodTypeWithExclusive", value)
}

// NewPresetPeriodTypeWithExclusiveFromValue returns a pointer to a valid PresetPeriodTypeWithExclusive
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPresetPeriodTypeWithExclusiveFromValue(v string) (*PresetPeriodTypeWithExclusive, error) {
	ev := PresetPeriodTypeWithExclusive(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PresetPeriodTypeWithExclusive: valid values are %v", v, AllowedPresetPeriodTypeWithExclusiveEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PresetPeriodTypeWithExclusive) IsValid() bool {
	for _, existing := range AllowedPresetPeriodTypeWithExclusiveEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PresetPeriodTypeWithExclusive value
func (v PresetPeriodTypeWithExclusive) Ptr() *PresetPeriodTypeWithExclusive {
	return &v
}

type NullablePresetPeriodTypeWithExclusive struct {
	value *PresetPeriodTypeWithExclusive
	isSet bool
}

func (v NullablePresetPeriodTypeWithExclusive) Get() *PresetPeriodTypeWithExclusive {
	return v.value
}

func (v *NullablePresetPeriodTypeWithExclusive) Set(val *PresetPeriodTypeWithExclusive) {
	v.value = val
	v.isSet = true
}

func (v NullablePresetPeriodTypeWithExclusive) IsSet() bool {
	return v.isSet
}

func (v *NullablePresetPeriodTypeWithExclusive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresetPeriodTypeWithExclusive(val *PresetPeriodTypeWithExclusive) *NullablePresetPeriodTypeWithExclusive {
	return &NullablePresetPeriodTypeWithExclusive{value: val, isSet: true}
}

func (v NullablePresetPeriodTypeWithExclusive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresetPeriodTypeWithExclusive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

