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

// HttpAttributes Http details attributes
type HttpAttributes struct {
	// The HTML code retrieved from the target
	ResponseBody *string `json:"ResponseBody,omitempty"`
	// The HTTP response headers retrieved from the target 
	ResponseHeaders *string `json:"ResponseHeaders,omitempty"`
}

// NewHttpAttributes instantiates a new HttpAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpAttributes() *HttpAttributes {
	this := HttpAttributes{}
	return &this
}

// NewHttpAttributesWithDefaults instantiates a new HttpAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpAttributesWithDefaults() *HttpAttributes {
	this := HttpAttributes{}
	return &this
}

// GetResponseBody returns the ResponseBody field value if set, zero value otherwise.
func (o *HttpAttributes) GetResponseBody() string {
	if o == nil || o.ResponseBody == nil {
		var ret string
		return ret
	}
	return *o.ResponseBody
}

// GetResponseBodyOk returns a tuple with the ResponseBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpAttributes) GetResponseBodyOk() (*string, bool) {
	if o == nil || o.ResponseBody == nil {
		return nil, false
	}
	return o.ResponseBody, true
}

// HasResponseBody returns a boolean if a field has been set.
func (o *HttpAttributes) HasResponseBody() bool {
	if o != nil && o.ResponseBody != nil {
		return true
	}

	return false
}

// SetResponseBody gets a reference to the given string and assigns it to the ResponseBody field.
func (o *HttpAttributes) SetResponseBody(v string) {
	o.ResponseBody = &v
}

// GetResponseHeaders returns the ResponseHeaders field value if set, zero value otherwise.
func (o *HttpAttributes) GetResponseHeaders() string {
	if o == nil || o.ResponseHeaders == nil {
		var ret string
		return ret
	}
	return *o.ResponseHeaders
}

// GetResponseHeadersOk returns a tuple with the ResponseHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpAttributes) GetResponseHeadersOk() (*string, bool) {
	if o == nil || o.ResponseHeaders == nil {
		return nil, false
	}
	return o.ResponseHeaders, true
}

// HasResponseHeaders returns a boolean if a field has been set.
func (o *HttpAttributes) HasResponseHeaders() bool {
	if o != nil && o.ResponseHeaders != nil {
		return true
	}

	return false
}

// SetResponseHeaders gets a reference to the given string and assigns it to the ResponseHeaders field.
func (o *HttpAttributes) SetResponseHeaders(v string) {
	o.ResponseHeaders = &v
}

func (o HttpAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseBody != nil {
		toSerialize["ResponseBody"] = o.ResponseBody
	}
	if o.ResponseHeaders != nil {
		toSerialize["ResponseHeaders"] = o.ResponseHeaders
	}
	return json.Marshal(toSerialize)
}

type NullableHttpAttributes struct {
	value *HttpAttributes
	isSet bool
}

func (v NullableHttpAttributes) Get() *HttpAttributes {
	return v.value
}

func (v *NullableHttpAttributes) Set(val *HttpAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpAttributes(val *HttpAttributes) *NullableHttpAttributes {
	return &NullableHttpAttributes{value: val, isSet: true}
}

func (v NullableHttpAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


