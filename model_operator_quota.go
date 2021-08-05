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

// OperatorQuota struct for OperatorQuota
type OperatorQuota struct {
	Operators *int32 `json:"Operators,omitempty"`
	OperatorsInUse *int32 `json:"OperatorsInUse,omitempty"`
}

// NewOperatorQuota instantiates a new OperatorQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorQuota() *OperatorQuota {
	this := OperatorQuota{}
	return &this
}

// NewOperatorQuotaWithDefaults instantiates a new OperatorQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorQuotaWithDefaults() *OperatorQuota {
	this := OperatorQuota{}
	return &this
}

// GetOperators returns the Operators field value if set, zero value otherwise.
func (o *OperatorQuota) GetOperators() int32 {
	if o == nil || o.Operators == nil {
		var ret int32
		return ret
	}
	return *o.Operators
}

// GetOperatorsOk returns a tuple with the Operators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorQuota) GetOperatorsOk() (*int32, bool) {
	if o == nil || o.Operators == nil {
		return nil, false
	}
	return o.Operators, true
}

// HasOperators returns a boolean if a field has been set.
func (o *OperatorQuota) HasOperators() bool {
	if o != nil && o.Operators != nil {
		return true
	}

	return false
}

// SetOperators gets a reference to the given int32 and assigns it to the Operators field.
func (o *OperatorQuota) SetOperators(v int32) {
	o.Operators = &v
}

// GetOperatorsInUse returns the OperatorsInUse field value if set, zero value otherwise.
func (o *OperatorQuota) GetOperatorsInUse() int32 {
	if o == nil || o.OperatorsInUse == nil {
		var ret int32
		return ret
	}
	return *o.OperatorsInUse
}

// GetOperatorsInUseOk returns a tuple with the OperatorsInUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorQuota) GetOperatorsInUseOk() (*int32, bool) {
	if o == nil || o.OperatorsInUse == nil {
		return nil, false
	}
	return o.OperatorsInUse, true
}

// HasOperatorsInUse returns a boolean if a field has been set.
func (o *OperatorQuota) HasOperatorsInUse() bool {
	if o != nil && o.OperatorsInUse != nil {
		return true
	}

	return false
}

// SetOperatorsInUse gets a reference to the given int32 and assigns it to the OperatorsInUse field.
func (o *OperatorQuota) SetOperatorsInUse(v int32) {
	o.OperatorsInUse = &v
}

func (o OperatorQuota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Operators != nil {
		toSerialize["Operators"] = o.Operators
	}
	if o.OperatorsInUse != nil {
		toSerialize["OperatorsInUse"] = o.OperatorsInUse
	}
	return json.Marshal(toSerialize)
}

type NullableOperatorQuota struct {
	value *OperatorQuota
	isSet bool
}

func (v NullableOperatorQuota) Get() *OperatorQuota {
	return v.value
}

func (v *NullableOperatorQuota) Set(val *OperatorQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorQuota(val *OperatorQuota) *NullableOperatorQuota {
	return &NullableOperatorQuota{value: val, isSet: true}
}

func (v NullableOperatorQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

