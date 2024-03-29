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

// OperatorGroup struct for OperatorGroup
type OperatorGroup struct {
	// The unique identifier of this Operator group
	OperatorGroupGuid *string `json:"OperatorGroupGuid,omitempty"`
	// The descriptive name of this operator group
	Description *string `json:"Description,omitempty"`
	// Indicates whether this is the default group for all operators
	IsEveryone *bool `json:"IsEveryone,omitempty"`
	IsAdministratorsGroup *bool `json:"IsAdministratorsGroup,omitempty"`
}

// NewOperatorGroup instantiates a new OperatorGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorGroup() *OperatorGroup {
	this := OperatorGroup{}
	return &this
}

// NewOperatorGroupWithDefaults instantiates a new OperatorGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorGroupWithDefaults() *OperatorGroup {
	this := OperatorGroup{}
	return &this
}

// GetOperatorGroupGuid returns the OperatorGroupGuid field value if set, zero value otherwise.
func (o *OperatorGroup) GetOperatorGroupGuid() string {
	if o == nil || isNil(o.OperatorGroupGuid) {
		var ret string
		return ret
	}
	return *o.OperatorGroupGuid
}

// GetOperatorGroupGuidOk returns a tuple with the OperatorGroupGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorGroup) GetOperatorGroupGuidOk() (*string, bool) {
	if o == nil || isNil(o.OperatorGroupGuid) {
    return nil, false
	}
	return o.OperatorGroupGuid, true
}

// HasOperatorGroupGuid returns a boolean if a field has been set.
func (o *OperatorGroup) HasOperatorGroupGuid() bool {
	if o != nil && !isNil(o.OperatorGroupGuid) {
		return true
	}

	return false
}

// SetOperatorGroupGuid gets a reference to the given string and assigns it to the OperatorGroupGuid field.
func (o *OperatorGroup) SetOperatorGroupGuid(v string) {
	o.OperatorGroupGuid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OperatorGroup) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OperatorGroup) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OperatorGroup) SetDescription(v string) {
	o.Description = &v
}

// GetIsEveryone returns the IsEveryone field value if set, zero value otherwise.
func (o *OperatorGroup) GetIsEveryone() bool {
	if o == nil || isNil(o.IsEveryone) {
		var ret bool
		return ret
	}
	return *o.IsEveryone
}

// GetIsEveryoneOk returns a tuple with the IsEveryone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorGroup) GetIsEveryoneOk() (*bool, bool) {
	if o == nil || isNil(o.IsEveryone) {
    return nil, false
	}
	return o.IsEveryone, true
}

// HasIsEveryone returns a boolean if a field has been set.
func (o *OperatorGroup) HasIsEveryone() bool {
	if o != nil && !isNil(o.IsEveryone) {
		return true
	}

	return false
}

// SetIsEveryone gets a reference to the given bool and assigns it to the IsEveryone field.
func (o *OperatorGroup) SetIsEveryone(v bool) {
	o.IsEveryone = &v
}

// GetIsAdministratorsGroup returns the IsAdministratorsGroup field value if set, zero value otherwise.
func (o *OperatorGroup) GetIsAdministratorsGroup() bool {
	if o == nil || isNil(o.IsAdministratorsGroup) {
		var ret bool
		return ret
	}
	return *o.IsAdministratorsGroup
}

// GetIsAdministratorsGroupOk returns a tuple with the IsAdministratorsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorGroup) GetIsAdministratorsGroupOk() (*bool, bool) {
	if o == nil || isNil(o.IsAdministratorsGroup) {
    return nil, false
	}
	return o.IsAdministratorsGroup, true
}

// HasIsAdministratorsGroup returns a boolean if a field has been set.
func (o *OperatorGroup) HasIsAdministratorsGroup() bool {
	if o != nil && !isNil(o.IsAdministratorsGroup) {
		return true
	}

	return false
}

// SetIsAdministratorsGroup gets a reference to the given bool and assigns it to the IsAdministratorsGroup field.
func (o *OperatorGroup) SetIsAdministratorsGroup(v bool) {
	o.IsAdministratorsGroup = &v
}

func (o OperatorGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OperatorGroupGuid) {
		toSerialize["OperatorGroupGuid"] = o.OperatorGroupGuid
	}
	if !isNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !isNil(o.IsEveryone) {
		toSerialize["IsEveryone"] = o.IsEveryone
	}
	if !isNil(o.IsAdministratorsGroup) {
		toSerialize["IsAdministratorsGroup"] = o.IsAdministratorsGroup
	}
	return json.Marshal(toSerialize)
}

type NullableOperatorGroup struct {
	value *OperatorGroup
	isSet bool
}

func (v NullableOperatorGroup) Get() *OperatorGroup {
	return v.value
}

func (v *NullableOperatorGroup) Set(val *OperatorGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorGroup(val *OperatorGroup) *NullableOperatorGroup {
	return &NullableOperatorGroup{value: val, isSet: true}
}

func (v NullableOperatorGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


