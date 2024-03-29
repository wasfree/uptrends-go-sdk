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

// MonitorGroup Monitor group
type MonitorGroup struct {
	// The unique identifier of this monitor group
	MonitorGroupGuid *string `json:"MonitorGroupGuid,omitempty"`
	// The descriptive name of this probe group
	Description *string `json:"Description,omitempty"`
	// Indicates whether this is the default group for all probes
	IsAll bool `json:"IsAll"`
	// Indicates whether the monitor quota is unlimited Only administrators can change this
	IsQuotaUnlimited *bool `json:"IsQuotaUnlimited,omitempty"`
	// The basic monitor quota for the monitor group Only administrators can change this
	BasicMonitorQuota *int32 `json:"BasicMonitorQuota,omitempty"`
	// The browser monitor quota for the monitor group Only administrators can change this
	BrowserMonitorQuota *int32 `json:"BrowserMonitorQuota,omitempty"`
	// The transaction monitor quota for the monitor group Only administrators can change this
	TransactionMonitorQuota *int32 `json:"TransactionMonitorQuota,omitempty"`
	// The api monitor quota for the monitor group Only administrators can change this
	ApiMonitorQuota *int32 `json:"ApiMonitorQuota,omitempty"`
	// The classic quota for the monitor group Only administrators can change this
	ClassicQuota *int32 `json:"ClassicQuota,omitempty"`
}

// NewMonitorGroup instantiates a new MonitorGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorGroup(isAll bool) *MonitorGroup {
	this := MonitorGroup{}
	this.IsAll = isAll
	return &this
}

// NewMonitorGroupWithDefaults instantiates a new MonitorGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorGroupWithDefaults() *MonitorGroup {
	this := MonitorGroup{}
	return &this
}

// GetMonitorGroupGuid returns the MonitorGroupGuid field value if set, zero value otherwise.
func (o *MonitorGroup) GetMonitorGroupGuid() string {
	if o == nil || isNil(o.MonitorGroupGuid) {
		var ret string
		return ret
	}
	return *o.MonitorGroupGuid
}

// GetMonitorGroupGuidOk returns a tuple with the MonitorGroupGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorGroup) GetMonitorGroupGuidOk() (*string, bool) {
	if o == nil || isNil(o.MonitorGroupGuid) {
    return nil, false
	}
	return o.MonitorGroupGuid, true
}

// HasMonitorGroupGuid returns a boolean if a field has been set.
func (o *MonitorGroup) HasMonitorGroupGuid() bool {
	if o != nil && !isNil(o.MonitorGroupGuid) {
		return true
	}

	return false
}

// SetMonitorGroupGuid gets a reference to the given string and assigns it to the MonitorGroupGuid field.
func (o *MonitorGroup) SetMonitorGroupGuid(v string) {
	o.MonitorGroupGuid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MonitorGroup) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MonitorGroup) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MonitorGroup) SetDescription(v string) {
	o.Description = &v
}

// GetIsAll returns the IsAll field value
func (o *MonitorGroup) GetIsAll() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAll
}

// GetIsAllOk returns a tuple with the IsAll field value
// and a boolean to check if the value has been set.
func (o *MonitorGroup) GetIsAllOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsAll, true
}

// SetIsAll sets field value
func (o *MonitorGroup) SetIsAll(v bool) {
	o.IsAll = v
}

// GetIsQuotaUnlimited returns the IsQuotaUnlimited field value if set, zero value otherwise.
func (o *MonitorGroup) GetIsQuotaUnlimited() bool {
	if o == nil || isNil(o.IsQuotaUnlimited) {
		var ret bool
		return ret
	}
	return *o.IsQuotaUnlimited
}

// GetIsQuotaUnlimitedOk returns a tuple with the IsQuotaUnlimited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorGroup) GetIsQuotaUnlimitedOk() (*bool, bool) {
	if o == nil || isNil(o.IsQuotaUnlimited) {
    return nil, false
	}
	return o.IsQuotaUnlimited, true
}

// HasIsQuotaUnlimited returns a boolean if a field has been set.
func (o *MonitorGroup) HasIsQuotaUnlimited() bool {
	if o != nil && !isNil(o.IsQuotaUnlimited) {
		return true
	}

	return false
}

// SetIsQuotaUnlimited gets a reference to the given bool and assigns it to the IsQuotaUnlimited field.
func (o *MonitorGroup) SetIsQuotaUnlimited(v bool) {
	o.IsQuotaUnlimited = &v
}

// GetBasicMonitorQuota returns the BasicMonitorQuota field value if set, zero value otherwise.
func (o *MonitorGroup) GetBasicMonitorQuota() int32 {
	if o == nil || isNil(o.BasicMonitorQuota) {
		var ret int32
		return ret
	}
	return *o.BasicMonitorQuota
}

// GetBasicMonitorQuotaOk returns a tuple with the BasicMonitorQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorGroup) GetBasicMonitorQuotaOk() (*int32, bool) {
	if o == nil || isNil(o.BasicMonitorQuota) {
    return nil, false
	}
	return o.BasicMonitorQuota, true
}

// HasBasicMonitorQuota returns a boolean if a field has been set.
func (o *MonitorGroup) HasBasicMonitorQuota() bool {
	if o != nil && !isNil(o.BasicMonitorQuota) {
		return true
	}

	return false
}

// SetBasicMonitorQuota gets a reference to the given int32 and assigns it to the BasicMonitorQuota field.
func (o *MonitorGroup) SetBasicMonitorQuota(v int32) {
	o.BasicMonitorQuota = &v
}

// GetBrowserMonitorQuota returns the BrowserMonitorQuota field value if set, zero value otherwise.
func (o *MonitorGroup) GetBrowserMonitorQuota() int32 {
	if o == nil || isNil(o.BrowserMonitorQuota) {
		var ret int32
		return ret
	}
	return *o.BrowserMonitorQuota
}

// GetBrowserMonitorQuotaOk returns a tuple with the BrowserMonitorQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorGroup) GetBrowserMonitorQuotaOk() (*int32, bool) {
	if o == nil || isNil(o.BrowserMonitorQuota) {
    return nil, false
	}
	return o.BrowserMonitorQuota, true
}

// HasBrowserMonitorQuota returns a boolean if a field has been set.
func (o *MonitorGroup) HasBrowserMonitorQuota() bool {
	if o != nil && !isNil(o.BrowserMonitorQuota) {
		return true
	}

	return false
}

// SetBrowserMonitorQuota gets a reference to the given int32 and assigns it to the BrowserMonitorQuota field.
func (o *MonitorGroup) SetBrowserMonitorQuota(v int32) {
	o.BrowserMonitorQuota = &v
}

// GetTransactionMonitorQuota returns the TransactionMonitorQuota field value if set, zero value otherwise.
func (o *MonitorGroup) GetTransactionMonitorQuota() int32 {
	if o == nil || isNil(o.TransactionMonitorQuota) {
		var ret int32
		return ret
	}
	return *o.TransactionMonitorQuota
}

// GetTransactionMonitorQuotaOk returns a tuple with the TransactionMonitorQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorGroup) GetTransactionMonitorQuotaOk() (*int32, bool) {
	if o == nil || isNil(o.TransactionMonitorQuota) {
    return nil, false
	}
	return o.TransactionMonitorQuota, true
}

// HasTransactionMonitorQuota returns a boolean if a field has been set.
func (o *MonitorGroup) HasTransactionMonitorQuota() bool {
	if o != nil && !isNil(o.TransactionMonitorQuota) {
		return true
	}

	return false
}

// SetTransactionMonitorQuota gets a reference to the given int32 and assigns it to the TransactionMonitorQuota field.
func (o *MonitorGroup) SetTransactionMonitorQuota(v int32) {
	o.TransactionMonitorQuota = &v
}

// GetApiMonitorQuota returns the ApiMonitorQuota field value if set, zero value otherwise.
func (o *MonitorGroup) GetApiMonitorQuota() int32 {
	if o == nil || isNil(o.ApiMonitorQuota) {
		var ret int32
		return ret
	}
	return *o.ApiMonitorQuota
}

// GetApiMonitorQuotaOk returns a tuple with the ApiMonitorQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorGroup) GetApiMonitorQuotaOk() (*int32, bool) {
	if o == nil || isNil(o.ApiMonitorQuota) {
    return nil, false
	}
	return o.ApiMonitorQuota, true
}

// HasApiMonitorQuota returns a boolean if a field has been set.
func (o *MonitorGroup) HasApiMonitorQuota() bool {
	if o != nil && !isNil(o.ApiMonitorQuota) {
		return true
	}

	return false
}

// SetApiMonitorQuota gets a reference to the given int32 and assigns it to the ApiMonitorQuota field.
func (o *MonitorGroup) SetApiMonitorQuota(v int32) {
	o.ApiMonitorQuota = &v
}

// GetClassicQuota returns the ClassicQuota field value if set, zero value otherwise.
func (o *MonitorGroup) GetClassicQuota() int32 {
	if o == nil || isNil(o.ClassicQuota) {
		var ret int32
		return ret
	}
	return *o.ClassicQuota
}

// GetClassicQuotaOk returns a tuple with the ClassicQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorGroup) GetClassicQuotaOk() (*int32, bool) {
	if o == nil || isNil(o.ClassicQuota) {
    return nil, false
	}
	return o.ClassicQuota, true
}

// HasClassicQuota returns a boolean if a field has been set.
func (o *MonitorGroup) HasClassicQuota() bool {
	if o != nil && !isNil(o.ClassicQuota) {
		return true
	}

	return false
}

// SetClassicQuota gets a reference to the given int32 and assigns it to the ClassicQuota field.
func (o *MonitorGroup) SetClassicQuota(v int32) {
	o.ClassicQuota = &v
}

func (o MonitorGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MonitorGroupGuid) {
		toSerialize["MonitorGroupGuid"] = o.MonitorGroupGuid
	}
	if !isNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if true {
		toSerialize["IsAll"] = o.IsAll
	}
	if !isNil(o.IsQuotaUnlimited) {
		toSerialize["IsQuotaUnlimited"] = o.IsQuotaUnlimited
	}
	if !isNil(o.BasicMonitorQuota) {
		toSerialize["BasicMonitorQuota"] = o.BasicMonitorQuota
	}
	if !isNil(o.BrowserMonitorQuota) {
		toSerialize["BrowserMonitorQuota"] = o.BrowserMonitorQuota
	}
	if !isNil(o.TransactionMonitorQuota) {
		toSerialize["TransactionMonitorQuota"] = o.TransactionMonitorQuota
	}
	if !isNil(o.ApiMonitorQuota) {
		toSerialize["ApiMonitorQuota"] = o.ApiMonitorQuota
	}
	if !isNil(o.ClassicQuota) {
		toSerialize["ClassicQuota"] = o.ClassicQuota
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorGroup struct {
	value *MonitorGroup
	isSet bool
}

func (v NullableMonitorGroup) Get() *MonitorGroup {
	return v.value
}

func (v *NullableMonitorGroup) Set(val *MonitorGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorGroup(val *MonitorGroup) *NullableMonitorGroup {
	return &NullableMonitorGroup{value: val, isSet: true}
}

func (v NullableMonitorGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


