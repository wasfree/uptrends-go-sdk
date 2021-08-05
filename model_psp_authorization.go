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

// PSPAuthorization struct for PSPAuthorization
type PSPAuthorization struct {
	AuthorizationId *string `json:"AuthorizationId,omitempty"`
	AuthorizationType PSPAuthorizationType `json:"AuthorizationType"`
	OperatorGuid *string `json:"OperatorGuid,omitempty"`
	OperatorGroupGuid *string `json:"OperatorGroupGuid,omitempty"`
}

// NewPSPAuthorization instantiates a new PSPAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPSPAuthorization(authorizationType PSPAuthorizationType) *PSPAuthorization {
	this := PSPAuthorization{}
	this.AuthorizationType = authorizationType
	return &this
}

// NewPSPAuthorizationWithDefaults instantiates a new PSPAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPSPAuthorizationWithDefaults() *PSPAuthorization {
	this := PSPAuthorization{}
	return &this
}

// GetAuthorizationId returns the AuthorizationId field value if set, zero value otherwise.
func (o *PSPAuthorization) GetAuthorizationId() string {
	if o == nil || o.AuthorizationId == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationId
}

// GetAuthorizationIdOk returns a tuple with the AuthorizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PSPAuthorization) GetAuthorizationIdOk() (*string, bool) {
	if o == nil || o.AuthorizationId == nil {
		return nil, false
	}
	return o.AuthorizationId, true
}

// HasAuthorizationId returns a boolean if a field has been set.
func (o *PSPAuthorization) HasAuthorizationId() bool {
	if o != nil && o.AuthorizationId != nil {
		return true
	}

	return false
}

// SetAuthorizationId gets a reference to the given string and assigns it to the AuthorizationId field.
func (o *PSPAuthorization) SetAuthorizationId(v string) {
	o.AuthorizationId = &v
}

// GetAuthorizationType returns the AuthorizationType field value
func (o *PSPAuthorization) GetAuthorizationType() PSPAuthorizationType {
	if o == nil {
		var ret PSPAuthorizationType
		return ret
	}

	return o.AuthorizationType
}

// GetAuthorizationTypeOk returns a tuple with the AuthorizationType field value
// and a boolean to check if the value has been set.
func (o *PSPAuthorization) GetAuthorizationTypeOk() (*PSPAuthorizationType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuthorizationType, true
}

// SetAuthorizationType sets field value
func (o *PSPAuthorization) SetAuthorizationType(v PSPAuthorizationType) {
	o.AuthorizationType = v
}

// GetOperatorGuid returns the OperatorGuid field value if set, zero value otherwise.
func (o *PSPAuthorization) GetOperatorGuid() string {
	if o == nil || o.OperatorGuid == nil {
		var ret string
		return ret
	}
	return *o.OperatorGuid
}

// GetOperatorGuidOk returns a tuple with the OperatorGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PSPAuthorization) GetOperatorGuidOk() (*string, bool) {
	if o == nil || o.OperatorGuid == nil {
		return nil, false
	}
	return o.OperatorGuid, true
}

// HasOperatorGuid returns a boolean if a field has been set.
func (o *PSPAuthorization) HasOperatorGuid() bool {
	if o != nil && o.OperatorGuid != nil {
		return true
	}

	return false
}

// SetOperatorGuid gets a reference to the given string and assigns it to the OperatorGuid field.
func (o *PSPAuthorization) SetOperatorGuid(v string) {
	o.OperatorGuid = &v
}

// GetOperatorGroupGuid returns the OperatorGroupGuid field value if set, zero value otherwise.
func (o *PSPAuthorization) GetOperatorGroupGuid() string {
	if o == nil || o.OperatorGroupGuid == nil {
		var ret string
		return ret
	}
	return *o.OperatorGroupGuid
}

// GetOperatorGroupGuidOk returns a tuple with the OperatorGroupGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PSPAuthorization) GetOperatorGroupGuidOk() (*string, bool) {
	if o == nil || o.OperatorGroupGuid == nil {
		return nil, false
	}
	return o.OperatorGroupGuid, true
}

// HasOperatorGroupGuid returns a boolean if a field has been set.
func (o *PSPAuthorization) HasOperatorGroupGuid() bool {
	if o != nil && o.OperatorGroupGuid != nil {
		return true
	}

	return false
}

// SetOperatorGroupGuid gets a reference to the given string and assigns it to the OperatorGroupGuid field.
func (o *PSPAuthorization) SetOperatorGroupGuid(v string) {
	o.OperatorGroupGuid = &v
}

func (o PSPAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizationId != nil {
		toSerialize["AuthorizationId"] = o.AuthorizationId
	}
	if true {
		toSerialize["AuthorizationType"] = o.AuthorizationType
	}
	if o.OperatorGuid != nil {
		toSerialize["OperatorGuid"] = o.OperatorGuid
	}
	if o.OperatorGroupGuid != nil {
		toSerialize["OperatorGroupGuid"] = o.OperatorGroupGuid
	}
	return json.Marshal(toSerialize)
}

type NullablePSPAuthorization struct {
	value *PSPAuthorization
	isSet bool
}

func (v NullablePSPAuthorization) Get() *PSPAuthorization {
	return v.value
}

func (v *NullablePSPAuthorization) Set(val *PSPAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullablePSPAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullablePSPAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePSPAuthorization(val *PSPAuthorization) *NullablePSPAuthorization {
	return &NullablePSPAuthorization{value: val, isSet: true}
}

func (v NullablePSPAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePSPAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

