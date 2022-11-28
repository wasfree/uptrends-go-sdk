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

// MonitorStatusAttributes Object attributes 
type MonitorStatusAttributes struct {
	// The error level for the monitor status
	ErrorLevel LastErrorLevel `json:"ErrorLevel"`
	// Last checked timeStamp for this monitor
	LastCheck map[string]interface{} `json:"LastCheck,omitempty"`
	// Checkpoint id for the monitor status
	CheckpointId *int32 `json:"CheckpointId,omitempty"`
	// Checkpoint name for the monitor status
	CheckpointName *string `json:"CheckpointName,omitempty"`
	// Error description for the monitor status
	ErrorDescription *string `json:"ErrorDescription,omitempty"`
	// Uptime percentage for the monitor status
	UptimePercentage float64 `json:"UptimePercentage"`
	// Error code for the monitor status
	ErrorCode int32 `json:"ErrorCode"`
	// Last monitor check id
	LastMonitorCheckId *int64 `json:"LastMonitorCheckId,omitempty"`
	// Total time for the monitor status
	TotalTime *int32 `json:"TotalTime,omitempty"`
}

// NewMonitorStatusAttributes instantiates a new MonitorStatusAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorStatusAttributes(errorLevel LastErrorLevel, uptimePercentage float64, errorCode int32) *MonitorStatusAttributes {
	this := MonitorStatusAttributes{}
	this.ErrorLevel = errorLevel
	this.UptimePercentage = uptimePercentage
	this.ErrorCode = errorCode
	return &this
}

// NewMonitorStatusAttributesWithDefaults instantiates a new MonitorStatusAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorStatusAttributesWithDefaults() *MonitorStatusAttributes {
	this := MonitorStatusAttributes{}
	return &this
}

// GetErrorLevel returns the ErrorLevel field value
func (o *MonitorStatusAttributes) GetErrorLevel() LastErrorLevel {
	if o == nil {
		var ret LastErrorLevel
		return ret
	}

	return o.ErrorLevel
}

// GetErrorLevelOk returns a tuple with the ErrorLevel field value
// and a boolean to check if the value has been set.
func (o *MonitorStatusAttributes) GetErrorLevelOk() (*LastErrorLevel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ErrorLevel, true
}

// SetErrorLevel sets field value
func (o *MonitorStatusAttributes) SetErrorLevel(v LastErrorLevel) {
	o.ErrorLevel = v
}

// GetLastCheck returns the LastCheck field value if set, zero value otherwise.
func (o *MonitorStatusAttributes) GetLastCheck() map[string]interface{} {
	if o == nil || isNil(o.LastCheck) {
		var ret map[string]interface{}
		return ret
	}
	return o.LastCheck
}

// GetLastCheckOk returns a tuple with the LastCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStatusAttributes) GetLastCheckOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.LastCheck) {
    return map[string]interface{}{}, false
	}
	return o.LastCheck, true
}

// HasLastCheck returns a boolean if a field has been set.
func (o *MonitorStatusAttributes) HasLastCheck() bool {
	if o != nil && !isNil(o.LastCheck) {
		return true
	}

	return false
}

// SetLastCheck gets a reference to the given map[string]interface{} and assigns it to the LastCheck field.
func (o *MonitorStatusAttributes) SetLastCheck(v map[string]interface{}) {
	o.LastCheck = v
}

// GetCheckpointId returns the CheckpointId field value if set, zero value otherwise.
func (o *MonitorStatusAttributes) GetCheckpointId() int32 {
	if o == nil || isNil(o.CheckpointId) {
		var ret int32
		return ret
	}
	return *o.CheckpointId
}

// GetCheckpointIdOk returns a tuple with the CheckpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStatusAttributes) GetCheckpointIdOk() (*int32, bool) {
	if o == nil || isNil(o.CheckpointId) {
    return nil, false
	}
	return o.CheckpointId, true
}

// HasCheckpointId returns a boolean if a field has been set.
func (o *MonitorStatusAttributes) HasCheckpointId() bool {
	if o != nil && !isNil(o.CheckpointId) {
		return true
	}

	return false
}

// SetCheckpointId gets a reference to the given int32 and assigns it to the CheckpointId field.
func (o *MonitorStatusAttributes) SetCheckpointId(v int32) {
	o.CheckpointId = &v
}

// GetCheckpointName returns the CheckpointName field value if set, zero value otherwise.
func (o *MonitorStatusAttributes) GetCheckpointName() string {
	if o == nil || isNil(o.CheckpointName) {
		var ret string
		return ret
	}
	return *o.CheckpointName
}

// GetCheckpointNameOk returns a tuple with the CheckpointName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStatusAttributes) GetCheckpointNameOk() (*string, bool) {
	if o == nil || isNil(o.CheckpointName) {
    return nil, false
	}
	return o.CheckpointName, true
}

// HasCheckpointName returns a boolean if a field has been set.
func (o *MonitorStatusAttributes) HasCheckpointName() bool {
	if o != nil && !isNil(o.CheckpointName) {
		return true
	}

	return false
}

// SetCheckpointName gets a reference to the given string and assigns it to the CheckpointName field.
func (o *MonitorStatusAttributes) SetCheckpointName(v string) {
	o.CheckpointName = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *MonitorStatusAttributes) GetErrorDescription() string {
	if o == nil || isNil(o.ErrorDescription) {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStatusAttributes) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.ErrorDescription) {
    return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *MonitorStatusAttributes) HasErrorDescription() bool {
	if o != nil && !isNil(o.ErrorDescription) {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *MonitorStatusAttributes) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetUptimePercentage returns the UptimePercentage field value
func (o *MonitorStatusAttributes) GetUptimePercentage() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.UptimePercentage
}

// GetUptimePercentageOk returns a tuple with the UptimePercentage field value
// and a boolean to check if the value has been set.
func (o *MonitorStatusAttributes) GetUptimePercentageOk() (*float64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UptimePercentage, true
}

// SetUptimePercentage sets field value
func (o *MonitorStatusAttributes) SetUptimePercentage(v float64) {
	o.UptimePercentage = v
}

// GetErrorCode returns the ErrorCode field value
func (o *MonitorStatusAttributes) GetErrorCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *MonitorStatusAttributes) GetErrorCodeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *MonitorStatusAttributes) SetErrorCode(v int32) {
	o.ErrorCode = v
}

// GetLastMonitorCheckId returns the LastMonitorCheckId field value if set, zero value otherwise.
func (o *MonitorStatusAttributes) GetLastMonitorCheckId() int64 {
	if o == nil || isNil(o.LastMonitorCheckId) {
		var ret int64
		return ret
	}
	return *o.LastMonitorCheckId
}

// GetLastMonitorCheckIdOk returns a tuple with the LastMonitorCheckId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStatusAttributes) GetLastMonitorCheckIdOk() (*int64, bool) {
	if o == nil || isNil(o.LastMonitorCheckId) {
    return nil, false
	}
	return o.LastMonitorCheckId, true
}

// HasLastMonitorCheckId returns a boolean if a field has been set.
func (o *MonitorStatusAttributes) HasLastMonitorCheckId() bool {
	if o != nil && !isNil(o.LastMonitorCheckId) {
		return true
	}

	return false
}

// SetLastMonitorCheckId gets a reference to the given int64 and assigns it to the LastMonitorCheckId field.
func (o *MonitorStatusAttributes) SetLastMonitorCheckId(v int64) {
	o.LastMonitorCheckId = &v
}

// GetTotalTime returns the TotalTime field value if set, zero value otherwise.
func (o *MonitorStatusAttributes) GetTotalTime() int32 {
	if o == nil || isNil(o.TotalTime) {
		var ret int32
		return ret
	}
	return *o.TotalTime
}

// GetTotalTimeOk returns a tuple with the TotalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStatusAttributes) GetTotalTimeOk() (*int32, bool) {
	if o == nil || isNil(o.TotalTime) {
    return nil, false
	}
	return o.TotalTime, true
}

// HasTotalTime returns a boolean if a field has been set.
func (o *MonitorStatusAttributes) HasTotalTime() bool {
	if o != nil && !isNil(o.TotalTime) {
		return true
	}

	return false
}

// SetTotalTime gets a reference to the given int32 and assigns it to the TotalTime field.
func (o *MonitorStatusAttributes) SetTotalTime(v int32) {
	o.TotalTime = &v
}

func (o MonitorStatusAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ErrorLevel"] = o.ErrorLevel
	}
	if !isNil(o.LastCheck) {
		toSerialize["LastCheck"] = o.LastCheck
	}
	if !isNil(o.CheckpointId) {
		toSerialize["CheckpointId"] = o.CheckpointId
	}
	if !isNil(o.CheckpointName) {
		toSerialize["CheckpointName"] = o.CheckpointName
	}
	if !isNil(o.ErrorDescription) {
		toSerialize["ErrorDescription"] = o.ErrorDescription
	}
	if true {
		toSerialize["UptimePercentage"] = o.UptimePercentage
	}
	if true {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	if !isNil(o.LastMonitorCheckId) {
		toSerialize["LastMonitorCheckId"] = o.LastMonitorCheckId
	}
	if !isNil(o.TotalTime) {
		toSerialize["TotalTime"] = o.TotalTime
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorStatusAttributes struct {
	value *MonitorStatusAttributes
	isSet bool
}

func (v NullableMonitorStatusAttributes) Get() *MonitorStatusAttributes {
	return v.value
}

func (v *NullableMonitorStatusAttributes) Set(val *MonitorStatusAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorStatusAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorStatusAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorStatusAttributes(val *MonitorStatusAttributes) *NullableMonitorStatusAttributes {
	return &NullableMonitorStatusAttributes{value: val, isSet: true}
}

func (v NullableMonitorStatusAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorStatusAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


