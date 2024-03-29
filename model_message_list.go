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

// MessageList struct for MessageList
type MessageList struct {
	Messages []MessageInfo `json:"Messages,omitempty"`
}

// NewMessageList instantiates a new MessageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageList() *MessageList {
	this := MessageList{}
	return &this
}

// NewMessageListWithDefaults instantiates a new MessageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageListWithDefaults() *MessageList {
	this := MessageList{}
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *MessageList) GetMessages() []MessageInfo {
	if o == nil || isNil(o.Messages) {
		var ret []MessageInfo
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageList) GetMessagesOk() ([]MessageInfo, bool) {
	if o == nil || isNil(o.Messages) {
    return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *MessageList) HasMessages() bool {
	if o != nil && !isNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []MessageInfo and assigns it to the Messages field.
func (o *MessageList) SetMessages(v []MessageInfo) {
	o.Messages = v
}

func (o MessageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Messages) {
		toSerialize["Messages"] = o.Messages
	}
	return json.Marshal(toSerialize)
}

type NullableMessageList struct {
	value *MessageList
	isSet bool
}

func (v NullableMessageList) Get() *MessageList {
	return v.value
}

func (v *NullableMessageList) Set(val *MessageList) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageList) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageList(val *MessageList) *NullableMessageList {
	return &NullableMessageList{value: val, isSet: true}
}

func (v NullableMessageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


