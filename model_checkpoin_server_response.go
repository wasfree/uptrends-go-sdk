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

// CheckpoinServerResponse struct for CheckpoinServerResponse
type CheckpoinServerResponse struct {
	Data *CheckpointServer `json:"Data,omitempty"`
	Links *LinksData `json:"Links,omitempty"`
	Relationships *[]RelationObject `json:"Relationships,omitempty"`
	Meta *MetaData `json:"Meta,omitempty"`
}

// NewCheckpoinServerResponse instantiates a new CheckpoinServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckpoinServerResponse() *CheckpoinServerResponse {
	this := CheckpoinServerResponse{}
	return &this
}

// NewCheckpoinServerResponseWithDefaults instantiates a new CheckpoinServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckpoinServerResponseWithDefaults() *CheckpoinServerResponse {
	this := CheckpoinServerResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CheckpoinServerResponse) GetData() CheckpointServer {
	if o == nil || o.Data == nil {
		var ret CheckpointServer
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckpoinServerResponse) GetDataOk() (*CheckpointServer, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CheckpoinServerResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CheckpointServer and assigns it to the Data field.
func (o *CheckpoinServerResponse) SetData(v CheckpointServer) {
	o.Data = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CheckpoinServerResponse) GetLinks() LinksData {
	if o == nil || o.Links == nil {
		var ret LinksData
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckpoinServerResponse) GetLinksOk() (*LinksData, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CheckpoinServerResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksData and assigns it to the Links field.
func (o *CheckpoinServerResponse) SetLinks(v LinksData) {
	o.Links = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CheckpoinServerResponse) GetRelationships() []RelationObject {
	if o == nil || o.Relationships == nil {
		var ret []RelationObject
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckpoinServerResponse) GetRelationshipsOk() (*[]RelationObject, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CheckpoinServerResponse) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given []RelationObject and assigns it to the Relationships field.
func (o *CheckpoinServerResponse) SetRelationships(v []RelationObject) {
	o.Relationships = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CheckpoinServerResponse) GetMeta() MetaData {
	if o == nil || o.Meta == nil {
		var ret MetaData
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckpoinServerResponse) GetMetaOk() (*MetaData, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CheckpoinServerResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaData and assigns it to the Meta field.
func (o *CheckpoinServerResponse) SetMeta(v MetaData) {
	o.Meta = &v
}

func (o CheckpoinServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["Data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["Links"] = o.Links
	}
	if o.Relationships != nil {
		toSerialize["Relationships"] = o.Relationships
	}
	if o.Meta != nil {
		toSerialize["Meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableCheckpoinServerResponse struct {
	value *CheckpoinServerResponse
	isSet bool
}

func (v NullableCheckpoinServerResponse) Get() *CheckpoinServerResponse {
	return v.value
}

func (v *NullableCheckpoinServerResponse) Set(val *CheckpoinServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckpoinServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckpoinServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckpoinServerResponse(val *CheckpoinServerResponse) *NullableCheckpoinServerResponse {
	return &NullableCheckpoinServerResponse{value: val, isSet: true}
}

func (v NullableCheckpoinServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckpoinServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

