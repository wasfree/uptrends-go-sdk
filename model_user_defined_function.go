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

// UserDefinedFunction struct for UserDefinedFunction
type UserDefinedFunction struct {
	Name *string `json:"Name,omitempty"`
	Type UserDefinedFunctionType `json:"Type"`
	Mappings *[]UserDefinedFunctionMapping `json:"Mappings,omitempty"`
	Regex *string `json:"Regex,omitempty"`
	// This property is not supported yet
	JwtSigningKey *string `json:"JwtSigningKey,omitempty"`
	// This property is not supported yet
	JwtAlgorithm *JwtAlgorithm `json:"JwtAlgorithm,omitempty"`
}

// NewUserDefinedFunction instantiates a new UserDefinedFunction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDefinedFunction(type_ UserDefinedFunctionType) *UserDefinedFunction {
	this := UserDefinedFunction{}
	this.Type = type_
	return &this
}

// NewUserDefinedFunctionWithDefaults instantiates a new UserDefinedFunction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDefinedFunctionWithDefaults() *UserDefinedFunction {
	this := UserDefinedFunction{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserDefinedFunction) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFunction) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserDefinedFunction) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserDefinedFunction) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value
func (o *UserDefinedFunction) GetType() UserDefinedFunctionType {
	if o == nil {
		var ret UserDefinedFunctionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserDefinedFunction) GetTypeOk() (*UserDefinedFunctionType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserDefinedFunction) SetType(v UserDefinedFunctionType) {
	o.Type = v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *UserDefinedFunction) GetMappings() []UserDefinedFunctionMapping {
	if o == nil || o.Mappings == nil {
		var ret []UserDefinedFunctionMapping
		return ret
	}
	return *o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFunction) GetMappingsOk() (*[]UserDefinedFunctionMapping, bool) {
	if o == nil || o.Mappings == nil {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *UserDefinedFunction) HasMappings() bool {
	if o != nil && o.Mappings != nil {
		return true
	}

	return false
}

// SetMappings gets a reference to the given []UserDefinedFunctionMapping and assigns it to the Mappings field.
func (o *UserDefinedFunction) SetMappings(v []UserDefinedFunctionMapping) {
	o.Mappings = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *UserDefinedFunction) GetRegex() string {
	if o == nil || o.Regex == nil {
		var ret string
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFunction) GetRegexOk() (*string, bool) {
	if o == nil || o.Regex == nil {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *UserDefinedFunction) HasRegex() bool {
	if o != nil && o.Regex != nil {
		return true
	}

	return false
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *UserDefinedFunction) SetRegex(v string) {
	o.Regex = &v
}

// GetJwtSigningKey returns the JwtSigningKey field value if set, zero value otherwise.
func (o *UserDefinedFunction) GetJwtSigningKey() string {
	if o == nil || o.JwtSigningKey == nil {
		var ret string
		return ret
	}
	return *o.JwtSigningKey
}

// GetJwtSigningKeyOk returns a tuple with the JwtSigningKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFunction) GetJwtSigningKeyOk() (*string, bool) {
	if o == nil || o.JwtSigningKey == nil {
		return nil, false
	}
	return o.JwtSigningKey, true
}

// HasJwtSigningKey returns a boolean if a field has been set.
func (o *UserDefinedFunction) HasJwtSigningKey() bool {
	if o != nil && o.JwtSigningKey != nil {
		return true
	}

	return false
}

// SetJwtSigningKey gets a reference to the given string and assigns it to the JwtSigningKey field.
func (o *UserDefinedFunction) SetJwtSigningKey(v string) {
	o.JwtSigningKey = &v
}

// GetJwtAlgorithm returns the JwtAlgorithm field value if set, zero value otherwise.
func (o *UserDefinedFunction) GetJwtAlgorithm() JwtAlgorithm {
	if o == nil || o.JwtAlgorithm == nil {
		var ret JwtAlgorithm
		return ret
	}
	return *o.JwtAlgorithm
}

// GetJwtAlgorithmOk returns a tuple with the JwtAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFunction) GetJwtAlgorithmOk() (*JwtAlgorithm, bool) {
	if o == nil || o.JwtAlgorithm == nil {
		return nil, false
	}
	return o.JwtAlgorithm, true
}

// HasJwtAlgorithm returns a boolean if a field has been set.
func (o *UserDefinedFunction) HasJwtAlgorithm() bool {
	if o != nil && o.JwtAlgorithm != nil {
		return true
	}

	return false
}

// SetJwtAlgorithm gets a reference to the given JwtAlgorithm and assigns it to the JwtAlgorithm field.
func (o *UserDefinedFunction) SetJwtAlgorithm(v JwtAlgorithm) {
	o.JwtAlgorithm = &v
}

func (o UserDefinedFunction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if true {
		toSerialize["Type"] = o.Type
	}
	if o.Mappings != nil {
		toSerialize["Mappings"] = o.Mappings
	}
	if o.Regex != nil {
		toSerialize["Regex"] = o.Regex
	}
	if o.JwtSigningKey != nil {
		toSerialize["JwtSigningKey"] = o.JwtSigningKey
	}
	if o.JwtAlgorithm != nil {
		toSerialize["JwtAlgorithm"] = o.JwtAlgorithm
	}
	return json.Marshal(toSerialize)
}

type NullableUserDefinedFunction struct {
	value *UserDefinedFunction
	isSet bool
}

func (v NullableUserDefinedFunction) Get() *UserDefinedFunction {
	return v.value
}

func (v *NullableUserDefinedFunction) Set(val *UserDefinedFunction) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDefinedFunction) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDefinedFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDefinedFunction(val *UserDefinedFunction) *NullableUserDefinedFunction {
	return &NullableUserDefinedFunction{value: val, isSet: true}
}

func (v NullableUserDefinedFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDefinedFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


