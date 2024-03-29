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

// AlertDefinitionOperatorGroup struct for AlertDefinitionOperatorGroup
type AlertDefinitionOperatorGroup struct {
	AlertDefinition string `json:"AlertDefinition"`
	Escalationlevel int32 `json:"Escalationlevel"`
	OperatorGroup string `json:"OperatorGroup"`
}

// NewAlertDefinitionOperatorGroup instantiates a new AlertDefinitionOperatorGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertDefinitionOperatorGroup(alertDefinition string, escalationlevel int32, operatorGroup string) *AlertDefinitionOperatorGroup {
	this := AlertDefinitionOperatorGroup{}
	this.AlertDefinition = alertDefinition
	this.Escalationlevel = escalationlevel
	this.OperatorGroup = operatorGroup
	return &this
}

// NewAlertDefinitionOperatorGroupWithDefaults instantiates a new AlertDefinitionOperatorGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertDefinitionOperatorGroupWithDefaults() *AlertDefinitionOperatorGroup {
	this := AlertDefinitionOperatorGroup{}
	return &this
}

// GetAlertDefinition returns the AlertDefinition field value
func (o *AlertDefinitionOperatorGroup) GetAlertDefinition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlertDefinition
}

// GetAlertDefinitionOk returns a tuple with the AlertDefinition field value
// and a boolean to check if the value has been set.
func (o *AlertDefinitionOperatorGroup) GetAlertDefinitionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AlertDefinition, true
}

// SetAlertDefinition sets field value
func (o *AlertDefinitionOperatorGroup) SetAlertDefinition(v string) {
	o.AlertDefinition = v
}

// GetEscalationlevel returns the Escalationlevel field value
func (o *AlertDefinitionOperatorGroup) GetEscalationlevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Escalationlevel
}

// GetEscalationlevelOk returns a tuple with the Escalationlevel field value
// and a boolean to check if the value has been set.
func (o *AlertDefinitionOperatorGroup) GetEscalationlevelOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Escalationlevel, true
}

// SetEscalationlevel sets field value
func (o *AlertDefinitionOperatorGroup) SetEscalationlevel(v int32) {
	o.Escalationlevel = v
}

// GetOperatorGroup returns the OperatorGroup field value
func (o *AlertDefinitionOperatorGroup) GetOperatorGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperatorGroup
}

// GetOperatorGroupOk returns a tuple with the OperatorGroup field value
// and a boolean to check if the value has been set.
func (o *AlertDefinitionOperatorGroup) GetOperatorGroupOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OperatorGroup, true
}

// SetOperatorGroup sets field value
func (o *AlertDefinitionOperatorGroup) SetOperatorGroup(v string) {
	o.OperatorGroup = v
}

func (o AlertDefinitionOperatorGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["AlertDefinition"] = o.AlertDefinition
	}
	if true {
		toSerialize["Escalationlevel"] = o.Escalationlevel
	}
	if true {
		toSerialize["OperatorGroup"] = o.OperatorGroup
	}
	return json.Marshal(toSerialize)
}

type NullableAlertDefinitionOperatorGroup struct {
	value *AlertDefinitionOperatorGroup
	isSet bool
}

func (v NullableAlertDefinitionOperatorGroup) Get() *AlertDefinitionOperatorGroup {
	return v.value
}

func (v *NullableAlertDefinitionOperatorGroup) Set(val *AlertDefinitionOperatorGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertDefinitionOperatorGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertDefinitionOperatorGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertDefinitionOperatorGroup(val *AlertDefinitionOperatorGroup) *NullableAlertDefinitionOperatorGroup {
	return &NullableAlertDefinitionOperatorGroup{value: val, isSet: true}
}

func (v NullableAlertDefinitionOperatorGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertDefinitionOperatorGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


