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

// Dashboard struct for Dashboard
type Dashboard struct {
	DashboardGuid *string `json:"DashboardGuid,omitempty"`
	Name *string `json:"Name,omitempty"`
	DashboardFilter *DashboardFilter `json:"DashboardFilter,omitempty"`
	AutoRefresh *bool `json:"AutoRefresh,omitempty"`
}

// NewDashboard instantiates a new Dashboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboard() *Dashboard {
	this := Dashboard{}
	return &this
}

// NewDashboardWithDefaults instantiates a new Dashboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardWithDefaults() *Dashboard {
	this := Dashboard{}
	return &this
}

// GetDashboardGuid returns the DashboardGuid field value if set, zero value otherwise.
func (o *Dashboard) GetDashboardGuid() string {
	if o == nil || isNil(o.DashboardGuid) {
		var ret string
		return ret
	}
	return *o.DashboardGuid
}

// GetDashboardGuidOk returns a tuple with the DashboardGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboard) GetDashboardGuidOk() (*string, bool) {
	if o == nil || isNil(o.DashboardGuid) {
    return nil, false
	}
	return o.DashboardGuid, true
}

// HasDashboardGuid returns a boolean if a field has been set.
func (o *Dashboard) HasDashboardGuid() bool {
	if o != nil && !isNil(o.DashboardGuid) {
		return true
	}

	return false
}

// SetDashboardGuid gets a reference to the given string and assigns it to the DashboardGuid field.
func (o *Dashboard) SetDashboardGuid(v string) {
	o.DashboardGuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dashboard) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboard) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dashboard) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dashboard) SetName(v string) {
	o.Name = &v
}

// GetDashboardFilter returns the DashboardFilter field value if set, zero value otherwise.
func (o *Dashboard) GetDashboardFilter() DashboardFilter {
	if o == nil || isNil(o.DashboardFilter) {
		var ret DashboardFilter
		return ret
	}
	return *o.DashboardFilter
}

// GetDashboardFilterOk returns a tuple with the DashboardFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboard) GetDashboardFilterOk() (*DashboardFilter, bool) {
	if o == nil || isNil(o.DashboardFilter) {
    return nil, false
	}
	return o.DashboardFilter, true
}

// HasDashboardFilter returns a boolean if a field has been set.
func (o *Dashboard) HasDashboardFilter() bool {
	if o != nil && !isNil(o.DashboardFilter) {
		return true
	}

	return false
}

// SetDashboardFilter gets a reference to the given DashboardFilter and assigns it to the DashboardFilter field.
func (o *Dashboard) SetDashboardFilter(v DashboardFilter) {
	o.DashboardFilter = &v
}

// GetAutoRefresh returns the AutoRefresh field value if set, zero value otherwise.
func (o *Dashboard) GetAutoRefresh() bool {
	if o == nil || isNil(o.AutoRefresh) {
		var ret bool
		return ret
	}
	return *o.AutoRefresh
}

// GetAutoRefreshOk returns a tuple with the AutoRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboard) GetAutoRefreshOk() (*bool, bool) {
	if o == nil || isNil(o.AutoRefresh) {
    return nil, false
	}
	return o.AutoRefresh, true
}

// HasAutoRefresh returns a boolean if a field has been set.
func (o *Dashboard) HasAutoRefresh() bool {
	if o != nil && !isNil(o.AutoRefresh) {
		return true
	}

	return false
}

// SetAutoRefresh gets a reference to the given bool and assigns it to the AutoRefresh field.
func (o *Dashboard) SetAutoRefresh(v bool) {
	o.AutoRefresh = &v
}

func (o Dashboard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DashboardGuid) {
		toSerialize["DashboardGuid"] = o.DashboardGuid
	}
	if !isNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !isNil(o.DashboardFilter) {
		toSerialize["DashboardFilter"] = o.DashboardFilter
	}
	if !isNil(o.AutoRefresh) {
		toSerialize["AutoRefresh"] = o.AutoRefresh
	}
	return json.Marshal(toSerialize)
}

type NullableDashboard struct {
	value *Dashboard
	isSet bool
}

func (v NullableDashboard) Get() *Dashboard {
	return v.value
}

func (v *NullableDashboard) Set(val *Dashboard) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboard) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboard(val *Dashboard) *NullableDashboard {
	return &NullableDashboard{value: val, isSet: true}
}

func (v NullableDashboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


