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
)

// PeriodMetaData struct for PeriodMetaData
type PeriodMetaData struct {
	Warning *string `json:"Warning,omitempty"`
	MaximumRetentionDays int32 `json:"MaximumRetentionDays"`
}

// NewPeriodMetaData instantiates a new PeriodMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeriodMetaData(maximumRetentionDays int32) *PeriodMetaData {
	this := PeriodMetaData{}
	this.MaximumRetentionDays = maximumRetentionDays
	return &this
}

// NewPeriodMetaDataWithDefaults instantiates a new PeriodMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeriodMetaDataWithDefaults() *PeriodMetaData {
	this := PeriodMetaData{}
	return &this
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *PeriodMetaData) GetWarning() string {
	if o == nil || o.Warning == nil {
		var ret string
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodMetaData) GetWarningOk() (*string, bool) {
	if o == nil || o.Warning == nil {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *PeriodMetaData) HasWarning() bool {
	if o != nil && o.Warning != nil {
		return true
	}

	return false
}

// SetWarning gets a reference to the given string and assigns it to the Warning field.
func (o *PeriodMetaData) SetWarning(v string) {
	o.Warning = &v
}

// GetMaximumRetentionDays returns the MaximumRetentionDays field value
func (o *PeriodMetaData) GetMaximumRetentionDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaximumRetentionDays
}

// GetMaximumRetentionDaysOk returns a tuple with the MaximumRetentionDays field value
// and a boolean to check if the value has been set.
func (o *PeriodMetaData) GetMaximumRetentionDaysOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaximumRetentionDays, true
}

// SetMaximumRetentionDays sets field value
func (o *PeriodMetaData) SetMaximumRetentionDays(v int32) {
	o.MaximumRetentionDays = v
}

func (o PeriodMetaData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Warning != nil {
		toSerialize["Warning"] = o.Warning
	}
	if true {
		toSerialize["MaximumRetentionDays"] = o.MaximumRetentionDays
	}
	return json.Marshal(toSerialize)
}

type NullablePeriodMetaData struct {
	value *PeriodMetaData
	isSet bool
}

func (v NullablePeriodMetaData) Get() *PeriodMetaData {
	return v.value
}

func (v *NullablePeriodMetaData) Set(val *PeriodMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullablePeriodMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullablePeriodMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeriodMetaData(val *PeriodMetaData) *NullablePeriodMetaData {
	return &NullablePeriodMetaData{value: val, isSet: true}
}

func (v NullablePeriodMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeriodMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

