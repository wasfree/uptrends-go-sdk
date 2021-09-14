/*
Uptrends API v4

This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package uptrends

import (
	"encoding/json"
	"time"
)

// StepTimingInfo struct for StepTimingInfo
type StepTimingInfo struct {
	Description *string `json:"Description,omitempty"`
	StartUtc time.Time `json:"StartUtc"`
	EndUtc time.Time `json:"EndUtc"`
	ElapsedMilliseconds int64 `json:"ElapsedMilliseconds"`
	DelayMilliseconds int64 `json:"DelayMilliseconds"`
	SubTimingInfos *[]StepTimingInfo `json:"SubTimingInfos,omitempty"`
	// If true, this TimingInfo should be counted as part of the sum of its siblings. If false, the TimingInfo should be discarded (e.g. for PreDelays we don't want to count).
	IsValid bool `json:"IsValid"`
}

// NewStepTimingInfo instantiates a new StepTimingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStepTimingInfo(startUtc time.Time, endUtc time.Time, elapsedMilliseconds int64, delayMilliseconds int64, isValid bool) *StepTimingInfo {
	this := StepTimingInfo{}
	this.StartUtc = startUtc
	this.EndUtc = endUtc
	this.ElapsedMilliseconds = elapsedMilliseconds
	this.DelayMilliseconds = delayMilliseconds
	this.IsValid = isValid
	return &this
}

// NewStepTimingInfoWithDefaults instantiates a new StepTimingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStepTimingInfoWithDefaults() *StepTimingInfo {
	this := StepTimingInfo{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StepTimingInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StepTimingInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StepTimingInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StepTimingInfo) SetDescription(v string) {
	o.Description = &v
}

// GetStartUtc returns the StartUtc field value
func (o *StepTimingInfo) GetStartUtc() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartUtc
}

// GetStartUtcOk returns a tuple with the StartUtc field value
// and a boolean to check if the value has been set.
func (o *StepTimingInfo) GetStartUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartUtc, true
}

// SetStartUtc sets field value
func (o *StepTimingInfo) SetStartUtc(v time.Time) {
	o.StartUtc = v
}

// GetEndUtc returns the EndUtc field value
func (o *StepTimingInfo) GetEndUtc() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndUtc
}

// GetEndUtcOk returns a tuple with the EndUtc field value
// and a boolean to check if the value has been set.
func (o *StepTimingInfo) GetEndUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndUtc, true
}

// SetEndUtc sets field value
func (o *StepTimingInfo) SetEndUtc(v time.Time) {
	o.EndUtc = v
}

// GetElapsedMilliseconds returns the ElapsedMilliseconds field value
func (o *StepTimingInfo) GetElapsedMilliseconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ElapsedMilliseconds
}

// GetElapsedMillisecondsOk returns a tuple with the ElapsedMilliseconds field value
// and a boolean to check if the value has been set.
func (o *StepTimingInfo) GetElapsedMillisecondsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ElapsedMilliseconds, true
}

// SetElapsedMilliseconds sets field value
func (o *StepTimingInfo) SetElapsedMilliseconds(v int64) {
	o.ElapsedMilliseconds = v
}

// GetDelayMilliseconds returns the DelayMilliseconds field value
func (o *StepTimingInfo) GetDelayMilliseconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DelayMilliseconds
}

// GetDelayMillisecondsOk returns a tuple with the DelayMilliseconds field value
// and a boolean to check if the value has been set.
func (o *StepTimingInfo) GetDelayMillisecondsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DelayMilliseconds, true
}

// SetDelayMilliseconds sets field value
func (o *StepTimingInfo) SetDelayMilliseconds(v int64) {
	o.DelayMilliseconds = v
}

// GetSubTimingInfos returns the SubTimingInfos field value if set, zero value otherwise.
func (o *StepTimingInfo) GetSubTimingInfos() []StepTimingInfo {
	if o == nil || o.SubTimingInfos == nil {
		var ret []StepTimingInfo
		return ret
	}
	return *o.SubTimingInfos
}

// GetSubTimingInfosOk returns a tuple with the SubTimingInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StepTimingInfo) GetSubTimingInfosOk() (*[]StepTimingInfo, bool) {
	if o == nil || o.SubTimingInfos == nil {
		return nil, false
	}
	return o.SubTimingInfos, true
}

// HasSubTimingInfos returns a boolean if a field has been set.
func (o *StepTimingInfo) HasSubTimingInfos() bool {
	if o != nil && o.SubTimingInfos != nil {
		return true
	}

	return false
}

// SetSubTimingInfos gets a reference to the given []StepTimingInfo and assigns it to the SubTimingInfos field.
func (o *StepTimingInfo) SetSubTimingInfos(v []StepTimingInfo) {
	o.SubTimingInfos = &v
}

// GetIsValid returns the IsValid field value
func (o *StepTimingInfo) GetIsValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value
// and a boolean to check if the value has been set.
func (o *StepTimingInfo) GetIsValidOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsValid, true
}

// SetIsValid sets field value
func (o *StepTimingInfo) SetIsValid(v bool) {
	o.IsValid = v
}

func (o StepTimingInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if true {
		toSerialize["StartUtc"] = o.StartUtc
	}
	if true {
		toSerialize["EndUtc"] = o.EndUtc
	}
	if true {
		toSerialize["ElapsedMilliseconds"] = o.ElapsedMilliseconds
	}
	if true {
		toSerialize["DelayMilliseconds"] = o.DelayMilliseconds
	}
	if o.SubTimingInfos != nil {
		toSerialize["SubTimingInfos"] = o.SubTimingInfos
	}
	if true {
		toSerialize["IsValid"] = o.IsValid
	}
	return json.Marshal(toSerialize)
}

type NullableStepTimingInfo struct {
	value *StepTimingInfo
	isSet bool
}

func (v NullableStepTimingInfo) Get() *StepTimingInfo {
	return v.value
}

func (v *NullableStepTimingInfo) Set(val *StepTimingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStepTimingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStepTimingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStepTimingInfo(val *StepTimingInfo) *NullableStepTimingInfo {
	return &NullableStepTimingInfo{value: val, isSet: true}
}

func (v NullableStepTimingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStepTimingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


