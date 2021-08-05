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

// CursorsData Cursors can be used to navigate the dataset in a fixed manner
type CursorsData struct {
	// Cursor for next data set
	Next *string `json:"Next,omitempty"`
	// Cursor for this data set (data might get updated, depending on your parameters)
	Self *string `json:"Self,omitempty"`
}

// NewCursorsData instantiates a new CursorsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCursorsData() *CursorsData {
	this := CursorsData{}
	return &this
}

// NewCursorsDataWithDefaults instantiates a new CursorsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCursorsDataWithDefaults() *CursorsData {
	this := CursorsData{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *CursorsData) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CursorsData) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *CursorsData) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *CursorsData) SetNext(v string) {
	o.Next = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *CursorsData) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CursorsData) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *CursorsData) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *CursorsData) SetSelf(v string) {
	o.Self = &v
}

func (o CursorsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next != nil {
		toSerialize["Next"] = o.Next
	}
	if o.Self != nil {
		toSerialize["Self"] = o.Self
	}
	return json.Marshal(toSerialize)
}

type NullableCursorsData struct {
	value *CursorsData
	isSet bool
}

func (v NullableCursorsData) Get() *CursorsData {
	return v.value
}

func (v *NullableCursorsData) Set(val *CursorsData) {
	v.value = val
	v.isSet = true
}

func (v NullableCursorsData) IsSet() bool {
	return v.isSet
}

func (v *NullableCursorsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCursorsData(val *CursorsData) *NullableCursorsData {
	return &NullableCursorsData{value: val, isSet: true}
}

func (v NullableCursorsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCursorsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

