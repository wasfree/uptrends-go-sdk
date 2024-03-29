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

// RumWebsiteWithMetricValues struct for RumWebsiteWithMetricValues
type RumWebsiteWithMetricValues struct {
	RumWebsiteId string `json:"RumWebsiteId"`
	Description *string `json:"Description,omitempty"`
	MetricValues []RumMetricValues `json:"MetricValues,omitempty"`
}

// NewRumWebsiteWithMetricValues instantiates a new RumWebsiteWithMetricValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRumWebsiteWithMetricValues(rumWebsiteId string) *RumWebsiteWithMetricValues {
	this := RumWebsiteWithMetricValues{}
	this.RumWebsiteId = rumWebsiteId
	return &this
}

// NewRumWebsiteWithMetricValuesWithDefaults instantiates a new RumWebsiteWithMetricValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRumWebsiteWithMetricValuesWithDefaults() *RumWebsiteWithMetricValues {
	this := RumWebsiteWithMetricValues{}
	return &this
}

// GetRumWebsiteId returns the RumWebsiteId field value
func (o *RumWebsiteWithMetricValues) GetRumWebsiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RumWebsiteId
}

// GetRumWebsiteIdOk returns a tuple with the RumWebsiteId field value
// and a boolean to check if the value has been set.
func (o *RumWebsiteWithMetricValues) GetRumWebsiteIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RumWebsiteId, true
}

// SetRumWebsiteId sets field value
func (o *RumWebsiteWithMetricValues) SetRumWebsiteId(v string) {
	o.RumWebsiteId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RumWebsiteWithMetricValues) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumWebsiteWithMetricValues) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RumWebsiteWithMetricValues) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RumWebsiteWithMetricValues) SetDescription(v string) {
	o.Description = &v
}

// GetMetricValues returns the MetricValues field value if set, zero value otherwise.
func (o *RumWebsiteWithMetricValues) GetMetricValues() []RumMetricValues {
	if o == nil || isNil(o.MetricValues) {
		var ret []RumMetricValues
		return ret
	}
	return o.MetricValues
}

// GetMetricValuesOk returns a tuple with the MetricValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumWebsiteWithMetricValues) GetMetricValuesOk() ([]RumMetricValues, bool) {
	if o == nil || isNil(o.MetricValues) {
    return nil, false
	}
	return o.MetricValues, true
}

// HasMetricValues returns a boolean if a field has been set.
func (o *RumWebsiteWithMetricValues) HasMetricValues() bool {
	if o != nil && !isNil(o.MetricValues) {
		return true
	}

	return false
}

// SetMetricValues gets a reference to the given []RumMetricValues and assigns it to the MetricValues field.
func (o *RumWebsiteWithMetricValues) SetMetricValues(v []RumMetricValues) {
	o.MetricValues = v
}

func (o RumWebsiteWithMetricValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["RumWebsiteId"] = o.RumWebsiteId
	}
	if !isNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !isNil(o.MetricValues) {
		toSerialize["MetricValues"] = o.MetricValues
	}
	return json.Marshal(toSerialize)
}

type NullableRumWebsiteWithMetricValues struct {
	value *RumWebsiteWithMetricValues
	isSet bool
}

func (v NullableRumWebsiteWithMetricValues) Get() *RumWebsiteWithMetricValues {
	return v.value
}

func (v *NullableRumWebsiteWithMetricValues) Set(val *RumWebsiteWithMetricValues) {
	v.value = val
	v.isSet = true
}

func (v NullableRumWebsiteWithMetricValues) IsSet() bool {
	return v.isSet
}

func (v *NullableRumWebsiteWithMetricValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRumWebsiteWithMetricValues(val *RumWebsiteWithMetricValues) *NullableRumWebsiteWithMetricValues {
	return &NullableRumWebsiteWithMetricValues{value: val, isSet: true}
}

func (v NullableRumWebsiteWithMetricValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRumWebsiteWithMetricValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


