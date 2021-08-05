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

// DayOfWeek 
type DayOfWeek string

// List of DayOfWeek
const (
	SUNDAY DayOfWeek = "Sunday"
	MONDAY DayOfWeek = "Monday"
	TUESDAY DayOfWeek = "Tuesday"
	WEDNESDAY DayOfWeek = "Wednesday"
	THURSDAY DayOfWeek = "Thursday"
	FRIDAY DayOfWeek = "Friday"
	SATURDAY DayOfWeek = "Saturday"
)

var allowedDayOfWeekEnumValues = []DayOfWeek{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

func (v *DayOfWeek) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DayOfWeek(value)
	for _, existing := range allowedDayOfWeekEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DayOfWeek", value)
}

// NewDayOfWeekFromValue returns a pointer to a valid DayOfWeek
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDayOfWeekFromValue(v string) (*DayOfWeek, error) {
	ev := DayOfWeek(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DayOfWeek: valid values are %v", v, allowedDayOfWeekEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DayOfWeek) IsValid() bool {
	for _, existing := range allowedDayOfWeekEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DayOfWeek value
func (v DayOfWeek) Ptr() *DayOfWeek {
	return &v
}

type NullableDayOfWeek struct {
	value *DayOfWeek
	isSet bool
}

func (v NullableDayOfWeek) Get() *DayOfWeek {
	return v.value
}

func (v *NullableDayOfWeek) Set(val *DayOfWeek) {
	v.value = val
	v.isSet = true
}

func (v NullableDayOfWeek) IsSet() bool {
	return v.isSet
}

func (v *NullableDayOfWeek) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDayOfWeek(val *DayOfWeek) *NullableDayOfWeek {
	return &NullableDayOfWeek{value: val, isSet: true}
}

func (v NullableDayOfWeek) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDayOfWeek) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

