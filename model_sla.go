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

// Sla struct for Sla
type Sla struct {
	// The unique key of this sla.
	SlaGuid *string `json:"SlaGuid,omitempty"`
	// Hash corresponding with this sla.
	Hash *string `json:"Hash,omitempty"`
	// The description name of this sla.
	Description *string `json:"Description,omitempty"`
	UptimeErrorThreshold *float64 `json:"UptimeErrorThreshold,omitempty"`
	UptimeWarningThreshold *float64 `json:"UptimeWarningThreshold,omitempty"`
	LoadTimeErrorThreshold *float64 `json:"LoadTimeErrorThreshold,omitempty"`
	OperatorReponseTimeThreshold *int32 `json:"OperatorReponseTimeThreshold,omitempty"`
}

// NewSla instantiates a new Sla object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSla() *Sla {
	this := Sla{}
	return &this
}

// NewSlaWithDefaults instantiates a new Sla object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlaWithDefaults() *Sla {
	this := Sla{}
	return &this
}

// GetSlaGuid returns the SlaGuid field value if set, zero value otherwise.
func (o *Sla) GetSlaGuid() string {
	if o == nil || o.SlaGuid == nil {
		var ret string
		return ret
	}
	return *o.SlaGuid
}

// GetSlaGuidOk returns a tuple with the SlaGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sla) GetSlaGuidOk() (*string, bool) {
	if o == nil || o.SlaGuid == nil {
		return nil, false
	}
	return o.SlaGuid, true
}

// HasSlaGuid returns a boolean if a field has been set.
func (o *Sla) HasSlaGuid() bool {
	if o != nil && o.SlaGuid != nil {
		return true
	}

	return false
}

// SetSlaGuid gets a reference to the given string and assigns it to the SlaGuid field.
func (o *Sla) SetSlaGuid(v string) {
	o.SlaGuid = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *Sla) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sla) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *Sla) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *Sla) SetHash(v string) {
	o.Hash = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Sla) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sla) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Sla) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Sla) SetDescription(v string) {
	o.Description = &v
}

// GetUptimeErrorThreshold returns the UptimeErrorThreshold field value if set, zero value otherwise.
func (o *Sla) GetUptimeErrorThreshold() float64 {
	if o == nil || o.UptimeErrorThreshold == nil {
		var ret float64
		return ret
	}
	return *o.UptimeErrorThreshold
}

// GetUptimeErrorThresholdOk returns a tuple with the UptimeErrorThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sla) GetUptimeErrorThresholdOk() (*float64, bool) {
	if o == nil || o.UptimeErrorThreshold == nil {
		return nil, false
	}
	return o.UptimeErrorThreshold, true
}

// HasUptimeErrorThreshold returns a boolean if a field has been set.
func (o *Sla) HasUptimeErrorThreshold() bool {
	if o != nil && o.UptimeErrorThreshold != nil {
		return true
	}

	return false
}

// SetUptimeErrorThreshold gets a reference to the given float64 and assigns it to the UptimeErrorThreshold field.
func (o *Sla) SetUptimeErrorThreshold(v float64) {
	o.UptimeErrorThreshold = &v
}

// GetUptimeWarningThreshold returns the UptimeWarningThreshold field value if set, zero value otherwise.
func (o *Sla) GetUptimeWarningThreshold() float64 {
	if o == nil || o.UptimeWarningThreshold == nil {
		var ret float64
		return ret
	}
	return *o.UptimeWarningThreshold
}

// GetUptimeWarningThresholdOk returns a tuple with the UptimeWarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sla) GetUptimeWarningThresholdOk() (*float64, bool) {
	if o == nil || o.UptimeWarningThreshold == nil {
		return nil, false
	}
	return o.UptimeWarningThreshold, true
}

// HasUptimeWarningThreshold returns a boolean if a field has been set.
func (o *Sla) HasUptimeWarningThreshold() bool {
	if o != nil && o.UptimeWarningThreshold != nil {
		return true
	}

	return false
}

// SetUptimeWarningThreshold gets a reference to the given float64 and assigns it to the UptimeWarningThreshold field.
func (o *Sla) SetUptimeWarningThreshold(v float64) {
	o.UptimeWarningThreshold = &v
}

// GetLoadTimeErrorThreshold returns the LoadTimeErrorThreshold field value if set, zero value otherwise.
func (o *Sla) GetLoadTimeErrorThreshold() float64 {
	if o == nil || o.LoadTimeErrorThreshold == nil {
		var ret float64
		return ret
	}
	return *o.LoadTimeErrorThreshold
}

// GetLoadTimeErrorThresholdOk returns a tuple with the LoadTimeErrorThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sla) GetLoadTimeErrorThresholdOk() (*float64, bool) {
	if o == nil || o.LoadTimeErrorThreshold == nil {
		return nil, false
	}
	return o.LoadTimeErrorThreshold, true
}

// HasLoadTimeErrorThreshold returns a boolean if a field has been set.
func (o *Sla) HasLoadTimeErrorThreshold() bool {
	if o != nil && o.LoadTimeErrorThreshold != nil {
		return true
	}

	return false
}

// SetLoadTimeErrorThreshold gets a reference to the given float64 and assigns it to the LoadTimeErrorThreshold field.
func (o *Sla) SetLoadTimeErrorThreshold(v float64) {
	o.LoadTimeErrorThreshold = &v
}

// GetOperatorReponseTimeThreshold returns the OperatorReponseTimeThreshold field value if set, zero value otherwise.
func (o *Sla) GetOperatorReponseTimeThreshold() int32 {
	if o == nil || o.OperatorReponseTimeThreshold == nil {
		var ret int32
		return ret
	}
	return *o.OperatorReponseTimeThreshold
}

// GetOperatorReponseTimeThresholdOk returns a tuple with the OperatorReponseTimeThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sla) GetOperatorReponseTimeThresholdOk() (*int32, bool) {
	if o == nil || o.OperatorReponseTimeThreshold == nil {
		return nil, false
	}
	return o.OperatorReponseTimeThreshold, true
}

// HasOperatorReponseTimeThreshold returns a boolean if a field has been set.
func (o *Sla) HasOperatorReponseTimeThreshold() bool {
	if o != nil && o.OperatorReponseTimeThreshold != nil {
		return true
	}

	return false
}

// SetOperatorReponseTimeThreshold gets a reference to the given int32 and assigns it to the OperatorReponseTimeThreshold field.
func (o *Sla) SetOperatorReponseTimeThreshold(v int32) {
	o.OperatorReponseTimeThreshold = &v
}

func (o Sla) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SlaGuid != nil {
		toSerialize["SlaGuid"] = o.SlaGuid
	}
	if o.Hash != nil {
		toSerialize["Hash"] = o.Hash
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.UptimeErrorThreshold != nil {
		toSerialize["UptimeErrorThreshold"] = o.UptimeErrorThreshold
	}
	if o.UptimeWarningThreshold != nil {
		toSerialize["UptimeWarningThreshold"] = o.UptimeWarningThreshold
	}
	if o.LoadTimeErrorThreshold != nil {
		toSerialize["LoadTimeErrorThreshold"] = o.LoadTimeErrorThreshold
	}
	if o.OperatorReponseTimeThreshold != nil {
		toSerialize["OperatorReponseTimeThreshold"] = o.OperatorReponseTimeThreshold
	}
	return json.Marshal(toSerialize)
}

type NullableSla struct {
	value *Sla
	isSet bool
}

func (v NullableSla) Get() *Sla {
	return v.value
}

func (v *NullableSla) Set(val *Sla) {
	v.value = val
	v.isSet = true
}

func (v NullableSla) IsSet() bool {
	return v.isSet
}

func (v *NullableSla) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSla(val *Sla) *NullableSla {
	return &NullableSla{value: val, isSet: true}
}

func (v NullableSla) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSla) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


