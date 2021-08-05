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

// AlertResponse struct for AlertResponse
type AlertResponse struct {
	Data *[]Alert `json:"Data,omitempty"`
	Links *LinksData `json:"Links,omitempty"`
	Relationships *[]RelationObject `json:"Relationships,omitempty"`
	Meta *MetaData `json:"Meta,omitempty"`
	// Cursors can be used to navigate the dataset in a fixed manner
	Cursors *CursorsData `json:"Cursors,omitempty"`
}

// NewAlertResponse instantiates a new AlertResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertResponse() *AlertResponse {
	this := AlertResponse{}
	return &this
}

// NewAlertResponseWithDefaults instantiates a new AlertResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertResponseWithDefaults() *AlertResponse {
	this := AlertResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AlertResponse) GetData() []Alert {
	if o == nil || o.Data == nil {
		var ret []Alert
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponse) GetDataOk() (*[]Alert, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AlertResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Alert and assigns it to the Data field.
func (o *AlertResponse) SetData(v []Alert) {
	o.Data = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlertResponse) GetLinks() LinksData {
	if o == nil || o.Links == nil {
		var ret LinksData
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponse) GetLinksOk() (*LinksData, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlertResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksData and assigns it to the Links field.
func (o *AlertResponse) SetLinks(v LinksData) {
	o.Links = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AlertResponse) GetRelationships() []RelationObject {
	if o == nil || o.Relationships == nil {
		var ret []RelationObject
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponse) GetRelationshipsOk() (*[]RelationObject, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AlertResponse) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given []RelationObject and assigns it to the Relationships field.
func (o *AlertResponse) SetRelationships(v []RelationObject) {
	o.Relationships = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AlertResponse) GetMeta() MetaData {
	if o == nil || o.Meta == nil {
		var ret MetaData
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponse) GetMetaOk() (*MetaData, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AlertResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaData and assigns it to the Meta field.
func (o *AlertResponse) SetMeta(v MetaData) {
	o.Meta = &v
}

// GetCursors returns the Cursors field value if set, zero value otherwise.
func (o *AlertResponse) GetCursors() CursorsData {
	if o == nil || o.Cursors == nil {
		var ret CursorsData
		return ret
	}
	return *o.Cursors
}

// GetCursorsOk returns a tuple with the Cursors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponse) GetCursorsOk() (*CursorsData, bool) {
	if o == nil || o.Cursors == nil {
		return nil, false
	}
	return o.Cursors, true
}

// HasCursors returns a boolean if a field has been set.
func (o *AlertResponse) HasCursors() bool {
	if o != nil && o.Cursors != nil {
		return true
	}

	return false
}

// SetCursors gets a reference to the given CursorsData and assigns it to the Cursors field.
func (o *AlertResponse) SetCursors(v CursorsData) {
	o.Cursors = &v
}

func (o AlertResponse) MarshalJSON() ([]byte, error) {
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
	if o.Cursors != nil {
		toSerialize["Cursors"] = o.Cursors
	}
	return json.Marshal(toSerialize)
}

type NullableAlertResponse struct {
	value *AlertResponse
	isSet bool
}

func (v NullableAlertResponse) Get() *AlertResponse {
	return v.value
}

func (v *NullableAlertResponse) Set(val *AlertResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertResponse(val *AlertResponse) *NullableAlertResponse {
	return &NullableAlertResponse{value: val, isSet: true}
}

func (v NullableAlertResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

