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

// AlertDefinition struct for AlertDefinition
type AlertDefinition struct {
	// The unique key of this Alert Definition.
	AlertDefinitionGuid *string `json:"AlertDefinitionGuid,omitempty"`
	// Hash corresponding with this alert definition.
	Hash *string `json:"Hash,omitempty"`
	// Name of this Alert Definition.
	AlertName *string `json:"AlertName,omitempty"`
	// Indicates whether this Alert Definition is active.
	IsActive *bool `json:"IsActive,omitempty"`
}

// NewAlertDefinition instantiates a new AlertDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertDefinition() *AlertDefinition {
	this := AlertDefinition{}
	return &this
}

// NewAlertDefinitionWithDefaults instantiates a new AlertDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertDefinitionWithDefaults() *AlertDefinition {
	this := AlertDefinition{}
	return &this
}

// GetAlertDefinitionGuid returns the AlertDefinitionGuid field value if set, zero value otherwise.
func (o *AlertDefinition) GetAlertDefinitionGuid() string {
	if o == nil || o.AlertDefinitionGuid == nil {
		var ret string
		return ret
	}
	return *o.AlertDefinitionGuid
}

// GetAlertDefinitionGuidOk returns a tuple with the AlertDefinitionGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinition) GetAlertDefinitionGuidOk() (*string, bool) {
	if o == nil || o.AlertDefinitionGuid == nil {
		return nil, false
	}
	return o.AlertDefinitionGuid, true
}

// HasAlertDefinitionGuid returns a boolean if a field has been set.
func (o *AlertDefinition) HasAlertDefinitionGuid() bool {
	if o != nil && o.AlertDefinitionGuid != nil {
		return true
	}

	return false
}

// SetAlertDefinitionGuid gets a reference to the given string and assigns it to the AlertDefinitionGuid field.
func (o *AlertDefinition) SetAlertDefinitionGuid(v string) {
	o.AlertDefinitionGuid = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *AlertDefinition) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinition) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *AlertDefinition) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *AlertDefinition) SetHash(v string) {
	o.Hash = &v
}

// GetAlertName returns the AlertName field value if set, zero value otherwise.
func (o *AlertDefinition) GetAlertName() string {
	if o == nil || o.AlertName == nil {
		var ret string
		return ret
	}
	return *o.AlertName
}

// GetAlertNameOk returns a tuple with the AlertName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinition) GetAlertNameOk() (*string, bool) {
	if o == nil || o.AlertName == nil {
		return nil, false
	}
	return o.AlertName, true
}

// HasAlertName returns a boolean if a field has been set.
func (o *AlertDefinition) HasAlertName() bool {
	if o != nil && o.AlertName != nil {
		return true
	}

	return false
}

// SetAlertName gets a reference to the given string and assigns it to the AlertName field.
func (o *AlertDefinition) SetAlertName(v string) {
	o.AlertName = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *AlertDefinition) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinition) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *AlertDefinition) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *AlertDefinition) SetIsActive(v bool) {
	o.IsActive = &v
}

func (o AlertDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AlertDefinitionGuid != nil {
		toSerialize["AlertDefinitionGuid"] = o.AlertDefinitionGuid
	}
	if o.Hash != nil {
		toSerialize["Hash"] = o.Hash
	}
	if o.AlertName != nil {
		toSerialize["AlertName"] = o.AlertName
	}
	if o.IsActive != nil {
		toSerialize["IsActive"] = o.IsActive
	}
	return json.Marshal(toSerialize)
}

type NullableAlertDefinition struct {
	value *AlertDefinition
	isSet bool
}

func (v NullableAlertDefinition) Get() *AlertDefinition {
	return v.value
}

func (v *NullableAlertDefinition) Set(val *AlertDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertDefinition(val *AlertDefinition) *NullableAlertDefinition {
	return &NullableAlertDefinition{value: val, isSet: true}
}

func (v NullableAlertDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

