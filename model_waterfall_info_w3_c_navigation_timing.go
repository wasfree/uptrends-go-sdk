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

// WaterfallInfoW3CNavigationTiming W3C Navigation Timing metrics
type WaterfallInfoW3CNavigationTiming struct {
	// Equal to the requestStart as defined by the W3C.  It is a timestamp indicating when the browser starts requesting the resource from the webserver after the DNS lookup and TCP connection.
	RequestStart int32 `json:"RequestStart"`
	// Equal to the difference between navigationStart and responseStart as defined by the W3C.  In short, it's the time between when the first request was sent from browser to server, and when the first bytes of the following response were received by the browser.
	TimeToFirstByte int32 `json:"TimeToFirstByte"`
	// Equal to domInteractive as defined by W3C.  It is a timestamp, indicating the document readiness is set to 'interactive', to indicate that the browser has stopped parsing the page and the user can start interacting with it.  Resources such as scripts, images, stylesheets, or frames may still be loading.
	DomInteractive int32 `json:"DomInteractive"`
	// Equal to the domComplete as defined by W3C.  It is a timestamp, indicating when the main document has been parsed, the DOM has been fully loaded, and the page readiness is set to 'complete'.
	DomComplete int32 `json:"DomComplete"`
	// Equal to loadEventEnd as defined by W3C.  It is a timestamp, indicating when the load event of the current document has completed, including all dependent resources such as stylesheets and images.
	LoadEvent int32 `json:"LoadEvent"`
}

// NewWaterfallInfoW3CNavigationTiming instantiates a new WaterfallInfoW3CNavigationTiming object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWaterfallInfoW3CNavigationTiming(requestStart int32, timeToFirstByte int32, domInteractive int32, domComplete int32, loadEvent int32) *WaterfallInfoW3CNavigationTiming {
	this := WaterfallInfoW3CNavigationTiming{}
	this.RequestStart = requestStart
	this.TimeToFirstByte = timeToFirstByte
	this.DomInteractive = domInteractive
	this.DomComplete = domComplete
	this.LoadEvent = loadEvent
	return &this
}

// NewWaterfallInfoW3CNavigationTimingWithDefaults instantiates a new WaterfallInfoW3CNavigationTiming object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWaterfallInfoW3CNavigationTimingWithDefaults() *WaterfallInfoW3CNavigationTiming {
	this := WaterfallInfoW3CNavigationTiming{}
	return &this
}

// GetRequestStart returns the RequestStart field value
func (o *WaterfallInfoW3CNavigationTiming) GetRequestStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequestStart
}

// GetRequestStartOk returns a tuple with the RequestStart field value
// and a boolean to check if the value has been set.
func (o *WaterfallInfoW3CNavigationTiming) GetRequestStartOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RequestStart, true
}

// SetRequestStart sets field value
func (o *WaterfallInfoW3CNavigationTiming) SetRequestStart(v int32) {
	o.RequestStart = v
}

// GetTimeToFirstByte returns the TimeToFirstByte field value
func (o *WaterfallInfoW3CNavigationTiming) GetTimeToFirstByte() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimeToFirstByte
}

// GetTimeToFirstByteOk returns a tuple with the TimeToFirstByte field value
// and a boolean to check if the value has been set.
func (o *WaterfallInfoW3CNavigationTiming) GetTimeToFirstByteOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TimeToFirstByte, true
}

// SetTimeToFirstByte sets field value
func (o *WaterfallInfoW3CNavigationTiming) SetTimeToFirstByte(v int32) {
	o.TimeToFirstByte = v
}

// GetDomInteractive returns the DomInteractive field value
func (o *WaterfallInfoW3CNavigationTiming) GetDomInteractive() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DomInteractive
}

// GetDomInteractiveOk returns a tuple with the DomInteractive field value
// and a boolean to check if the value has been set.
func (o *WaterfallInfoW3CNavigationTiming) GetDomInteractiveOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DomInteractive, true
}

// SetDomInteractive sets field value
func (o *WaterfallInfoW3CNavigationTiming) SetDomInteractive(v int32) {
	o.DomInteractive = v
}

// GetDomComplete returns the DomComplete field value
func (o *WaterfallInfoW3CNavigationTiming) GetDomComplete() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DomComplete
}

// GetDomCompleteOk returns a tuple with the DomComplete field value
// and a boolean to check if the value has been set.
func (o *WaterfallInfoW3CNavigationTiming) GetDomCompleteOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DomComplete, true
}

// SetDomComplete sets field value
func (o *WaterfallInfoW3CNavigationTiming) SetDomComplete(v int32) {
	o.DomComplete = v
}

// GetLoadEvent returns the LoadEvent field value
func (o *WaterfallInfoW3CNavigationTiming) GetLoadEvent() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LoadEvent
}

// GetLoadEventOk returns a tuple with the LoadEvent field value
// and a boolean to check if the value has been set.
func (o *WaterfallInfoW3CNavigationTiming) GetLoadEventOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LoadEvent, true
}

// SetLoadEvent sets field value
func (o *WaterfallInfoW3CNavigationTiming) SetLoadEvent(v int32) {
	o.LoadEvent = v
}

func (o WaterfallInfoW3CNavigationTiming) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["RequestStart"] = o.RequestStart
	}
	if true {
		toSerialize["TimeToFirstByte"] = o.TimeToFirstByte
	}
	if true {
		toSerialize["DomInteractive"] = o.DomInteractive
	}
	if true {
		toSerialize["DomComplete"] = o.DomComplete
	}
	if true {
		toSerialize["LoadEvent"] = o.LoadEvent
	}
	return json.Marshal(toSerialize)
}

type NullableWaterfallInfoW3CNavigationTiming struct {
	value *WaterfallInfoW3CNavigationTiming
	isSet bool
}

func (v NullableWaterfallInfoW3CNavigationTiming) Get() *WaterfallInfoW3CNavigationTiming {
	return v.value
}

func (v *NullableWaterfallInfoW3CNavigationTiming) Set(val *WaterfallInfoW3CNavigationTiming) {
	v.value = val
	v.isSet = true
}

func (v NullableWaterfallInfoW3CNavigationTiming) IsSet() bool {
	return v.isSet
}

func (v *NullableWaterfallInfoW3CNavigationTiming) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWaterfallInfoW3CNavigationTiming(val *WaterfallInfoW3CNavigationTiming) *NullableWaterfallInfoW3CNavigationTiming {
	return &NullableWaterfallInfoW3CNavigationTiming{value: val, isSet: true}
}

func (v NullableWaterfallInfoW3CNavigationTiming) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWaterfallInfoW3CNavigationTiming) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

