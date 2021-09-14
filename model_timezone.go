/*
Uptrends API v4

This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package uptrends

import (
	"encoding/json"
)

// Timezone Time zone available in Uptrends
type Timezone struct {
	// The time zone unique identifier
	TimezoneId int32 `json:"TimezoneId"`
	// The description of the time zone
	Description *string `json:"Description,omitempty"`
	// The offset from UTC in minutes (if this time zone runs behind UTC, the number is negative)
	OffsetFromUtc int32 `json:"OffsetFromUtc"`
	// Indicates whether or not this time zone uses Daylight Saving Time
	HasDaylightSaving bool `json:"HasDaylightSaving"`
	// The time offset for Daylight Saving Time in minutes
	DaylightSavingOffset *int32 `json:"DaylightSavingOffset,omitempty"`
}

// NewTimezone instantiates a new Timezone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimezone(timezoneId int32, offsetFromUtc int32, hasDaylightSaving bool) *Timezone {
	this := Timezone{}
	this.TimezoneId = timezoneId
	this.OffsetFromUtc = offsetFromUtc
	this.HasDaylightSaving = hasDaylightSaving
	return &this
}

// NewTimezoneWithDefaults instantiates a new Timezone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimezoneWithDefaults() *Timezone {
	this := Timezone{}
	return &this
}

// GetTimezoneId returns the TimezoneId field value
func (o *Timezone) GetTimezoneId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimezoneId
}

// GetTimezoneIdOk returns a tuple with the TimezoneId field value
// and a boolean to check if the value has been set.
func (o *Timezone) GetTimezoneIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TimezoneId, true
}

// SetTimezoneId sets field value
func (o *Timezone) SetTimezoneId(v int32) {
	o.TimezoneId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Timezone) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timezone) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Timezone) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Timezone) SetDescription(v string) {
	o.Description = &v
}

// GetOffsetFromUtc returns the OffsetFromUtc field value
func (o *Timezone) GetOffsetFromUtc() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OffsetFromUtc
}

// GetOffsetFromUtcOk returns a tuple with the OffsetFromUtc field value
// and a boolean to check if the value has been set.
func (o *Timezone) GetOffsetFromUtcOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OffsetFromUtc, true
}

// SetOffsetFromUtc sets field value
func (o *Timezone) SetOffsetFromUtc(v int32) {
	o.OffsetFromUtc = v
}

// GetHasDaylightSaving returns the HasDaylightSaving field value
func (o *Timezone) GetHasDaylightSaving() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasDaylightSaving
}

// GetHasDaylightSavingOk returns a tuple with the HasDaylightSaving field value
// and a boolean to check if the value has been set.
func (o *Timezone) GetHasDaylightSavingOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasDaylightSaving, true
}

// SetHasDaylightSaving sets field value
func (o *Timezone) SetHasDaylightSaving(v bool) {
	o.HasDaylightSaving = v
}

// GetDaylightSavingOffset returns the DaylightSavingOffset field value if set, zero value otherwise.
func (o *Timezone) GetDaylightSavingOffset() int32 {
	if o == nil || o.DaylightSavingOffset == nil {
		var ret int32
		return ret
	}
	return *o.DaylightSavingOffset
}

// GetDaylightSavingOffsetOk returns a tuple with the DaylightSavingOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timezone) GetDaylightSavingOffsetOk() (*int32, bool) {
	if o == nil || o.DaylightSavingOffset == nil {
		return nil, false
	}
	return o.DaylightSavingOffset, true
}

// HasDaylightSavingOffset returns a boolean if a field has been set.
func (o *Timezone) HasDaylightSavingOffset() bool {
	if o != nil && o.DaylightSavingOffset != nil {
		return true
	}

	return false
}

// SetDaylightSavingOffset gets a reference to the given int32 and assigns it to the DaylightSavingOffset field.
func (o *Timezone) SetDaylightSavingOffset(v int32) {
	o.DaylightSavingOffset = &v
}

func (o Timezone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["TimezoneId"] = o.TimezoneId
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if true {
		toSerialize["OffsetFromUtc"] = o.OffsetFromUtc
	}
	if true {
		toSerialize["HasDaylightSaving"] = o.HasDaylightSaving
	}
	if o.DaylightSavingOffset != nil {
		toSerialize["DaylightSavingOffset"] = o.DaylightSavingOffset
	}
	return json.Marshal(toSerialize)
}

type NullableTimezone struct {
	value *Timezone
	isSet bool
}

func (v NullableTimezone) Get() *Timezone {
	return v.value
}

func (v *NullableTimezone) Set(val *Timezone) {
	v.value = val
	v.isSet = true
}

func (v NullableTimezone) IsSet() bool {
	return v.isSet
}

func (v *NullableTimezone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimezone(val *Timezone) *NullableTimezone {
	return &NullableTimezone{value: val, isSet: true}
}

func (v NullableTimezone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimezone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


