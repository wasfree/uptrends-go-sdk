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

// SelectedPeriod struct for SelectedPeriod
type SelectedPeriod struct {
	// The type of period
	SelectedPeriodType *SelectedPeriodType `json:"SelectedPeriodType,omitempty"`
	// The start of a custom period (can't be used together with the SelectedPeriodType parameter)
	Start map[string]interface{} `json:"Start,omitempty"`
	// The end of a custom period
	End map[string]interface{} `json:"End,omitempty"`
	// The requested time period.
	PresetPeriod *PresetPeriodType `json:"PresetPeriod,omitempty"`
}

// NewSelectedPeriod instantiates a new SelectedPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectedPeriod() *SelectedPeriod {
	this := SelectedPeriod{}
	return &this
}

// NewSelectedPeriodWithDefaults instantiates a new SelectedPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectedPeriodWithDefaults() *SelectedPeriod {
	this := SelectedPeriod{}
	return &this
}

// GetSelectedPeriodType returns the SelectedPeriodType field value if set, zero value otherwise.
func (o *SelectedPeriod) GetSelectedPeriodType() SelectedPeriodType {
	if o == nil || isNil(o.SelectedPeriodType) {
		var ret SelectedPeriodType
		return ret
	}
	return *o.SelectedPeriodType
}

// GetSelectedPeriodTypeOk returns a tuple with the SelectedPeriodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedPeriod) GetSelectedPeriodTypeOk() (*SelectedPeriodType, bool) {
	if o == nil || isNil(o.SelectedPeriodType) {
    return nil, false
	}
	return o.SelectedPeriodType, true
}

// HasSelectedPeriodType returns a boolean if a field has been set.
func (o *SelectedPeriod) HasSelectedPeriodType() bool {
	if o != nil && !isNil(o.SelectedPeriodType) {
		return true
	}

	return false
}

// SetSelectedPeriodType gets a reference to the given SelectedPeriodType and assigns it to the SelectedPeriodType field.
func (o *SelectedPeriod) SetSelectedPeriodType(v SelectedPeriodType) {
	o.SelectedPeriodType = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SelectedPeriod) GetStart() map[string]interface{} {
	if o == nil || isNil(o.Start) {
		var ret map[string]interface{}
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedPeriod) GetStartOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Start) {
    return map[string]interface{}{}, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SelectedPeriod) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given map[string]interface{} and assigns it to the Start field.
func (o *SelectedPeriod) SetStart(v map[string]interface{}) {
	o.Start = v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *SelectedPeriod) GetEnd() map[string]interface{} {
	if o == nil || isNil(o.End) {
		var ret map[string]interface{}
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedPeriod) GetEndOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.End) {
    return map[string]interface{}{}, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *SelectedPeriod) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given map[string]interface{} and assigns it to the End field.
func (o *SelectedPeriod) SetEnd(v map[string]interface{}) {
	o.End = v
}

// GetPresetPeriod returns the PresetPeriod field value if set, zero value otherwise.
func (o *SelectedPeriod) GetPresetPeriod() PresetPeriodType {
	if o == nil || isNil(o.PresetPeriod) {
		var ret PresetPeriodType
		return ret
	}
	return *o.PresetPeriod
}

// GetPresetPeriodOk returns a tuple with the PresetPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedPeriod) GetPresetPeriodOk() (*PresetPeriodType, bool) {
	if o == nil || isNil(o.PresetPeriod) {
    return nil, false
	}
	return o.PresetPeriod, true
}

// HasPresetPeriod returns a boolean if a field has been set.
func (o *SelectedPeriod) HasPresetPeriod() bool {
	if o != nil && !isNil(o.PresetPeriod) {
		return true
	}

	return false
}

// SetPresetPeriod gets a reference to the given PresetPeriodType and assigns it to the PresetPeriod field.
func (o *SelectedPeriod) SetPresetPeriod(v PresetPeriodType) {
	o.PresetPeriod = &v
}

func (o SelectedPeriod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SelectedPeriodType) {
		toSerialize["SelectedPeriodType"] = o.SelectedPeriodType
	}
	if !isNil(o.Start) {
		toSerialize["Start"] = o.Start
	}
	if !isNil(o.End) {
		toSerialize["End"] = o.End
	}
	if !isNil(o.PresetPeriod) {
		toSerialize["PresetPeriod"] = o.PresetPeriod
	}
	return json.Marshal(toSerialize)
}

type NullableSelectedPeriod struct {
	value *SelectedPeriod
	isSet bool
}

func (v NullableSelectedPeriod) Get() *SelectedPeriod {
	return v.value
}

func (v *NullableSelectedPeriod) Set(val *SelectedPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectedPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectedPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectedPeriod(val *SelectedPeriod) *NullableSelectedPeriod {
	return &NullableSelectedPeriod{value: val, isSet: true}
}

func (v NullableSelectedPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectedPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


