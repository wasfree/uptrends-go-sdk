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

// WaterfallInfoPageLoadMetrics Metrics for describing the load cycle of the entire web page.  Includes CoreWebVitals and other related metrics.
type WaterfallInfoPageLoadMetrics struct {
	// Measures the sum total of all individual layout shift scores for every unexpected layout shift that occurs during the entire lifespan of the page. A score lower than 0.1 is considered a good score. Between 0.1 and 0.25 could use some improvement. 0.25 or higher is a poor score. https://web.dev/cls/?utm_source=devtools  In case the data was inconclusive null is returned.
	CumulativeLayoutShift *float32 `json:"CumulativeLayoutShift,omitempty"`
	// Measures the time from when the page starts loading to when any of the page's content is rendered on screen. A score lower than 1000ms is considered a good score. https://web.dev/fcp/  In case the data was inconclusive null is returned.
	FirstContentfulPaint *int32 `json:"FirstContentfulPaint,omitempty"`
	// The render timestamp of the largest image or text block visible within the viewport. A score lower than 2500ms is good. Between 2500ms and 4000ms needs improvement. 4000ms or higher is poor https://web.dev/lcp/  In case the data was inconclusive null is returned.
	LargestContentfulPaint *int32 `json:"LargestContentfulPaint,omitempty"`
	// The Total Blocking Time (TBT) metric measures the total amount of time between First Contentful Paint (FCP) and Time to Interactive (TTI) where the main thread was blocked for long enough to prevent input responsiveness.  https://web.dev/tbt/  In case the data was inconclusive null is returned.
	TotalBlockingTime *int32 `json:"TotalBlockingTime,omitempty"`
	// TTI measures how long it takes a page to become fully interactive.   https://web.dev/interactive/  In case the data was inconclusive null is returned.
	TimeToInteractive *int32 `json:"TimeToInteractive,omitempty"`
}

// NewWaterfallInfoPageLoadMetrics instantiates a new WaterfallInfoPageLoadMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWaterfallInfoPageLoadMetrics() *WaterfallInfoPageLoadMetrics {
	this := WaterfallInfoPageLoadMetrics{}
	return &this
}

// NewWaterfallInfoPageLoadMetricsWithDefaults instantiates a new WaterfallInfoPageLoadMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWaterfallInfoPageLoadMetricsWithDefaults() *WaterfallInfoPageLoadMetrics {
	this := WaterfallInfoPageLoadMetrics{}
	return &this
}

// GetCumulativeLayoutShift returns the CumulativeLayoutShift field value if set, zero value otherwise.
func (o *WaterfallInfoPageLoadMetrics) GetCumulativeLayoutShift() float32 {
	if o == nil || isNil(o.CumulativeLayoutShift) {
		var ret float32
		return ret
	}
	return *o.CumulativeLayoutShift
}

// GetCumulativeLayoutShiftOk returns a tuple with the CumulativeLayoutShift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaterfallInfoPageLoadMetrics) GetCumulativeLayoutShiftOk() (*float32, bool) {
	if o == nil || isNil(o.CumulativeLayoutShift) {
    return nil, false
	}
	return o.CumulativeLayoutShift, true
}

// HasCumulativeLayoutShift returns a boolean if a field has been set.
func (o *WaterfallInfoPageLoadMetrics) HasCumulativeLayoutShift() bool {
	if o != nil && !isNil(o.CumulativeLayoutShift) {
		return true
	}

	return false
}

// SetCumulativeLayoutShift gets a reference to the given float32 and assigns it to the CumulativeLayoutShift field.
func (o *WaterfallInfoPageLoadMetrics) SetCumulativeLayoutShift(v float32) {
	o.CumulativeLayoutShift = &v
}

// GetFirstContentfulPaint returns the FirstContentfulPaint field value if set, zero value otherwise.
func (o *WaterfallInfoPageLoadMetrics) GetFirstContentfulPaint() int32 {
	if o == nil || isNil(o.FirstContentfulPaint) {
		var ret int32
		return ret
	}
	return *o.FirstContentfulPaint
}

// GetFirstContentfulPaintOk returns a tuple with the FirstContentfulPaint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaterfallInfoPageLoadMetrics) GetFirstContentfulPaintOk() (*int32, bool) {
	if o == nil || isNil(o.FirstContentfulPaint) {
    return nil, false
	}
	return o.FirstContentfulPaint, true
}

// HasFirstContentfulPaint returns a boolean if a field has been set.
func (o *WaterfallInfoPageLoadMetrics) HasFirstContentfulPaint() bool {
	if o != nil && !isNil(o.FirstContentfulPaint) {
		return true
	}

	return false
}

// SetFirstContentfulPaint gets a reference to the given int32 and assigns it to the FirstContentfulPaint field.
func (o *WaterfallInfoPageLoadMetrics) SetFirstContentfulPaint(v int32) {
	o.FirstContentfulPaint = &v
}

// GetLargestContentfulPaint returns the LargestContentfulPaint field value if set, zero value otherwise.
func (o *WaterfallInfoPageLoadMetrics) GetLargestContentfulPaint() int32 {
	if o == nil || isNil(o.LargestContentfulPaint) {
		var ret int32
		return ret
	}
	return *o.LargestContentfulPaint
}

// GetLargestContentfulPaintOk returns a tuple with the LargestContentfulPaint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaterfallInfoPageLoadMetrics) GetLargestContentfulPaintOk() (*int32, bool) {
	if o == nil || isNil(o.LargestContentfulPaint) {
    return nil, false
	}
	return o.LargestContentfulPaint, true
}

// HasLargestContentfulPaint returns a boolean if a field has been set.
func (o *WaterfallInfoPageLoadMetrics) HasLargestContentfulPaint() bool {
	if o != nil && !isNil(o.LargestContentfulPaint) {
		return true
	}

	return false
}

// SetLargestContentfulPaint gets a reference to the given int32 and assigns it to the LargestContentfulPaint field.
func (o *WaterfallInfoPageLoadMetrics) SetLargestContentfulPaint(v int32) {
	o.LargestContentfulPaint = &v
}

// GetTotalBlockingTime returns the TotalBlockingTime field value if set, zero value otherwise.
func (o *WaterfallInfoPageLoadMetrics) GetTotalBlockingTime() int32 {
	if o == nil || isNil(o.TotalBlockingTime) {
		var ret int32
		return ret
	}
	return *o.TotalBlockingTime
}

// GetTotalBlockingTimeOk returns a tuple with the TotalBlockingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaterfallInfoPageLoadMetrics) GetTotalBlockingTimeOk() (*int32, bool) {
	if o == nil || isNil(o.TotalBlockingTime) {
    return nil, false
	}
	return o.TotalBlockingTime, true
}

// HasTotalBlockingTime returns a boolean if a field has been set.
func (o *WaterfallInfoPageLoadMetrics) HasTotalBlockingTime() bool {
	if o != nil && !isNil(o.TotalBlockingTime) {
		return true
	}

	return false
}

// SetTotalBlockingTime gets a reference to the given int32 and assigns it to the TotalBlockingTime field.
func (o *WaterfallInfoPageLoadMetrics) SetTotalBlockingTime(v int32) {
	o.TotalBlockingTime = &v
}

// GetTimeToInteractive returns the TimeToInteractive field value if set, zero value otherwise.
func (o *WaterfallInfoPageLoadMetrics) GetTimeToInteractive() int32 {
	if o == nil || isNil(o.TimeToInteractive) {
		var ret int32
		return ret
	}
	return *o.TimeToInteractive
}

// GetTimeToInteractiveOk returns a tuple with the TimeToInteractive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaterfallInfoPageLoadMetrics) GetTimeToInteractiveOk() (*int32, bool) {
	if o == nil || isNil(o.TimeToInteractive) {
    return nil, false
	}
	return o.TimeToInteractive, true
}

// HasTimeToInteractive returns a boolean if a field has been set.
func (o *WaterfallInfoPageLoadMetrics) HasTimeToInteractive() bool {
	if o != nil && !isNil(o.TimeToInteractive) {
		return true
	}

	return false
}

// SetTimeToInteractive gets a reference to the given int32 and assigns it to the TimeToInteractive field.
func (o *WaterfallInfoPageLoadMetrics) SetTimeToInteractive(v int32) {
	o.TimeToInteractive = &v
}

func (o WaterfallInfoPageLoadMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CumulativeLayoutShift) {
		toSerialize["CumulativeLayoutShift"] = o.CumulativeLayoutShift
	}
	if !isNil(o.FirstContentfulPaint) {
		toSerialize["FirstContentfulPaint"] = o.FirstContentfulPaint
	}
	if !isNil(o.LargestContentfulPaint) {
		toSerialize["LargestContentfulPaint"] = o.LargestContentfulPaint
	}
	if !isNil(o.TotalBlockingTime) {
		toSerialize["TotalBlockingTime"] = o.TotalBlockingTime
	}
	if !isNil(o.TimeToInteractive) {
		toSerialize["TimeToInteractive"] = o.TimeToInteractive
	}
	return json.Marshal(toSerialize)
}

type NullableWaterfallInfoPageLoadMetrics struct {
	value *WaterfallInfoPageLoadMetrics
	isSet bool
}

func (v NullableWaterfallInfoPageLoadMetrics) Get() *WaterfallInfoPageLoadMetrics {
	return v.value
}

func (v *NullableWaterfallInfoPageLoadMetrics) Set(val *WaterfallInfoPageLoadMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableWaterfallInfoPageLoadMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableWaterfallInfoPageLoadMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWaterfallInfoPageLoadMetrics(val *WaterfallInfoPageLoadMetrics) *NullableWaterfallInfoPageLoadMetrics {
	return &NullableWaterfallInfoPageLoadMetrics{value: val, isSet: true}
}

func (v NullableWaterfallInfoPageLoadMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWaterfallInfoPageLoadMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


