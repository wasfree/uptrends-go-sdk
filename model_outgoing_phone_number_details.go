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

// OutgoingPhoneNumberDetails struct for OutgoingPhoneNumberDetails
type OutgoingPhoneNumberDetails struct {
	Id int32 `json:"Id"`
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	Description *string `json:"Description,omitempty"`
}

// NewOutgoingPhoneNumberDetails instantiates a new OutgoingPhoneNumberDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutgoingPhoneNumberDetails(id int32) *OutgoingPhoneNumberDetails {
	this := OutgoingPhoneNumberDetails{}
	this.Id = id
	return &this
}

// NewOutgoingPhoneNumberDetailsWithDefaults instantiates a new OutgoingPhoneNumberDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutgoingPhoneNumberDetailsWithDefaults() *OutgoingPhoneNumberDetails {
	this := OutgoingPhoneNumberDetails{}
	return &this
}

// GetId returns the Id field value
func (o *OutgoingPhoneNumberDetails) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OutgoingPhoneNumberDetails) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OutgoingPhoneNumberDetails) SetId(v int32) {
	o.Id = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *OutgoingPhoneNumberDetails) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingPhoneNumberDetails) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *OutgoingPhoneNumberDetails) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *OutgoingPhoneNumberDetails) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OutgoingPhoneNumberDetails) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingPhoneNumberDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OutgoingPhoneNumberDetails) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OutgoingPhoneNumberDetails) SetDescription(v string) {
	o.Description = &v
}

func (o OutgoingPhoneNumberDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Id"] = o.Id
	}
	if o.PhoneNumber != nil {
		toSerialize["PhoneNumber"] = o.PhoneNumber
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableOutgoingPhoneNumberDetails struct {
	value *OutgoingPhoneNumberDetails
	isSet bool
}

func (v NullableOutgoingPhoneNumberDetails) Get() *OutgoingPhoneNumberDetails {
	return v.value
}

func (v *NullableOutgoingPhoneNumberDetails) Set(val *OutgoingPhoneNumberDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOutgoingPhoneNumberDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOutgoingPhoneNumberDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutgoingPhoneNumberDetails(val *OutgoingPhoneNumberDetails) *NullableOutgoingPhoneNumberDetails {
	return &NullableOutgoingPhoneNumberDetails{value: val, isSet: true}
}

func (v NullableOutgoingPhoneNumberDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutgoingPhoneNumberDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


