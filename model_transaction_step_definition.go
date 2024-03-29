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

// TransactionStepDefinition struct for TransactionStepDefinition
type TransactionStepDefinition struct {
	Steps []TransactionStep2 `json:"Steps,omitempty"`
}

// NewTransactionStepDefinition instantiates a new TransactionStepDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionStepDefinition() *TransactionStepDefinition {
	this := TransactionStepDefinition{}
	return &this
}

// NewTransactionStepDefinitionWithDefaults instantiates a new TransactionStepDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionStepDefinitionWithDefaults() *TransactionStepDefinition {
	this := TransactionStepDefinition{}
	return &this
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *TransactionStepDefinition) GetSteps() []TransactionStep2 {
	if o == nil || isNil(o.Steps) {
		var ret []TransactionStep2
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionStepDefinition) GetStepsOk() ([]TransactionStep2, bool) {
	if o == nil || isNil(o.Steps) {
    return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *TransactionStepDefinition) HasSteps() bool {
	if o != nil && !isNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []TransactionStep2 and assigns it to the Steps field.
func (o *TransactionStepDefinition) SetSteps(v []TransactionStep2) {
	o.Steps = v
}

func (o TransactionStepDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Steps) {
		toSerialize["Steps"] = o.Steps
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionStepDefinition struct {
	value *TransactionStepDefinition
	isSet bool
}

func (v NullableTransactionStepDefinition) Get() *TransactionStepDefinition {
	return v.value
}

func (v *NullableTransactionStepDefinition) Set(val *TransactionStepDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionStepDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionStepDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionStepDefinition(val *TransactionStepDefinition) *NullableTransactionStepDefinition {
	return &NullableTransactionStepDefinition{value: val, isSet: true}
}

func (v NullableTransactionStepDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionStepDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


