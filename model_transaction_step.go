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

// TransactionStep Describes one step in a transaction
type TransactionStep struct {
	// Step (index) number
	StepNumber int32 `json:"StepNumber"`
	// The name of the step
	StepName *string `json:"StepName,omitempty"`
	// Number of milliseconds it took for this step to succeed
	TotalTime float64 `json:"TotalTime"`
	// Text representation of the step element results
	Elements *[]string `json:"Elements,omitempty"`
	// Did this step result in an error?
	HasError bool `json:"HasError"`
}

// NewTransactionStep instantiates a new TransactionStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionStep(stepNumber int32, totalTime float64, hasError bool) *TransactionStep {
	this := TransactionStep{}
	this.StepNumber = stepNumber
	this.TotalTime = totalTime
	this.HasError = hasError
	return &this
}

// NewTransactionStepWithDefaults instantiates a new TransactionStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionStepWithDefaults() *TransactionStep {
	this := TransactionStep{}
	return &this
}

// GetStepNumber returns the StepNumber field value
func (o *TransactionStep) GetStepNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StepNumber
}

// GetStepNumberOk returns a tuple with the StepNumber field value
// and a boolean to check if the value has been set.
func (o *TransactionStep) GetStepNumberOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StepNumber, true
}

// SetStepNumber sets field value
func (o *TransactionStep) SetStepNumber(v int32) {
	o.StepNumber = v
}

// GetStepName returns the StepName field value if set, zero value otherwise.
func (o *TransactionStep) GetStepName() string {
	if o == nil || o.StepName == nil {
		var ret string
		return ret
	}
	return *o.StepName
}

// GetStepNameOk returns a tuple with the StepName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionStep) GetStepNameOk() (*string, bool) {
	if o == nil || o.StepName == nil {
		return nil, false
	}
	return o.StepName, true
}

// HasStepName returns a boolean if a field has been set.
func (o *TransactionStep) HasStepName() bool {
	if o != nil && o.StepName != nil {
		return true
	}

	return false
}

// SetStepName gets a reference to the given string and assigns it to the StepName field.
func (o *TransactionStep) SetStepName(v string) {
	o.StepName = &v
}

// GetTotalTime returns the TotalTime field value
func (o *TransactionStep) GetTotalTime() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.TotalTime
}

// GetTotalTimeOk returns a tuple with the TotalTime field value
// and a boolean to check if the value has been set.
func (o *TransactionStep) GetTotalTimeOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalTime, true
}

// SetTotalTime sets field value
func (o *TransactionStep) SetTotalTime(v float64) {
	o.TotalTime = v
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *TransactionStep) GetElements() []string {
	if o == nil || o.Elements == nil {
		var ret []string
		return ret
	}
	return *o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionStep) GetElementsOk() (*[]string, bool) {
	if o == nil || o.Elements == nil {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *TransactionStep) HasElements() bool {
	if o != nil && o.Elements != nil {
		return true
	}

	return false
}

// SetElements gets a reference to the given []string and assigns it to the Elements field.
func (o *TransactionStep) SetElements(v []string) {
	o.Elements = &v
}

// GetHasError returns the HasError field value
func (o *TransactionStep) GetHasError() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasError
}

// GetHasErrorOk returns a tuple with the HasError field value
// and a boolean to check if the value has been set.
func (o *TransactionStep) GetHasErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasError, true
}

// SetHasError sets field value
func (o *TransactionStep) SetHasError(v bool) {
	o.HasError = v
}

func (o TransactionStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["StepNumber"] = o.StepNumber
	}
	if o.StepName != nil {
		toSerialize["StepName"] = o.StepName
	}
	if true {
		toSerialize["TotalTime"] = o.TotalTime
	}
	if o.Elements != nil {
		toSerialize["Elements"] = o.Elements
	}
	if true {
		toSerialize["HasError"] = o.HasError
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionStep struct {
	value *TransactionStep
	isSet bool
}

func (v NullableTransactionStep) Get() *TransactionStep {
	return v.value
}

func (v *NullableTransactionStep) Set(val *TransactionStep) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionStep) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionStep(val *TransactionStep) *NullableTransactionStep {
	return &NullableTransactionStep{value: val, isSet: true}
}

func (v NullableTransactionStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


