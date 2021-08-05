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

// IntegrationServiceMap struct for IntegrationServiceMap
type IntegrationServiceMap struct {
	MonitorGuid string `json:"MonitorGuid"`
	IntegrationServiceGuid string `json:"IntegrationServiceGuid"`
}

// NewIntegrationServiceMap instantiates a new IntegrationServiceMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationServiceMap(monitorGuid string, integrationServiceGuid string) *IntegrationServiceMap {
	this := IntegrationServiceMap{}
	this.MonitorGuid = monitorGuid
	this.IntegrationServiceGuid = integrationServiceGuid
	return &this
}

// NewIntegrationServiceMapWithDefaults instantiates a new IntegrationServiceMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationServiceMapWithDefaults() *IntegrationServiceMap {
	this := IntegrationServiceMap{}
	return &this
}

// GetMonitorGuid returns the MonitorGuid field value
func (o *IntegrationServiceMap) GetMonitorGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorGuid
}

// GetMonitorGuidOk returns a tuple with the MonitorGuid field value
// and a boolean to check if the value has been set.
func (o *IntegrationServiceMap) GetMonitorGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MonitorGuid, true
}

// SetMonitorGuid sets field value
func (o *IntegrationServiceMap) SetMonitorGuid(v string) {
	o.MonitorGuid = v
}

// GetIntegrationServiceGuid returns the IntegrationServiceGuid field value
func (o *IntegrationServiceMap) GetIntegrationServiceGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationServiceGuid
}

// GetIntegrationServiceGuidOk returns a tuple with the IntegrationServiceGuid field value
// and a boolean to check if the value has been set.
func (o *IntegrationServiceMap) GetIntegrationServiceGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IntegrationServiceGuid, true
}

// SetIntegrationServiceGuid sets field value
func (o *IntegrationServiceMap) SetIntegrationServiceGuid(v string) {
	o.IntegrationServiceGuid = v
}

func (o IntegrationServiceMap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["MonitorGuid"] = o.MonitorGuid
	}
	if true {
		toSerialize["IntegrationServiceGuid"] = o.IntegrationServiceGuid
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationServiceMap struct {
	value *IntegrationServiceMap
	isSet bool
}

func (v NullableIntegrationServiceMap) Get() *IntegrationServiceMap {
	return v.value
}

func (v *NullableIntegrationServiceMap) Set(val *IntegrationServiceMap) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationServiceMap) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationServiceMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationServiceMap(val *IntegrationServiceMap) *NullableIntegrationServiceMap {
	return &NullableIntegrationServiceMap{value: val, isSet: true}
}

func (v NullableIntegrationServiceMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationServiceMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

