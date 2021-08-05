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

// BrowserWindowDimensions struct for BrowserWindowDimensions
type BrowserWindowDimensions struct {
	IsMobile bool `json:"IsMobile"`
	Width int32 `json:"Width"`
	Height int32 `json:"Height"`
	PixelRatio int32 `json:"PixelRatio"`
	MobileDevice *string `json:"MobileDevice,omitempty"`
}

// NewBrowserWindowDimensions instantiates a new BrowserWindowDimensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrowserWindowDimensions(isMobile bool, width int32, height int32, pixelRatio int32) *BrowserWindowDimensions {
	this := BrowserWindowDimensions{}
	this.IsMobile = isMobile
	this.Width = width
	this.Height = height
	this.PixelRatio = pixelRatio
	return &this
}

// NewBrowserWindowDimensionsWithDefaults instantiates a new BrowserWindowDimensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrowserWindowDimensionsWithDefaults() *BrowserWindowDimensions {
	this := BrowserWindowDimensions{}
	return &this
}

// GetIsMobile returns the IsMobile field value
func (o *BrowserWindowDimensions) GetIsMobile() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMobile
}

// GetIsMobileOk returns a tuple with the IsMobile field value
// and a boolean to check if the value has been set.
func (o *BrowserWindowDimensions) GetIsMobileOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsMobile, true
}

// SetIsMobile sets field value
func (o *BrowserWindowDimensions) SetIsMobile(v bool) {
	o.IsMobile = v
}

// GetWidth returns the Width field value
func (o *BrowserWindowDimensions) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *BrowserWindowDimensions) GetWidthOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *BrowserWindowDimensions) SetWidth(v int32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *BrowserWindowDimensions) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *BrowserWindowDimensions) GetHeightOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *BrowserWindowDimensions) SetHeight(v int32) {
	o.Height = v
}

// GetPixelRatio returns the PixelRatio field value
func (o *BrowserWindowDimensions) GetPixelRatio() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PixelRatio
}

// GetPixelRatioOk returns a tuple with the PixelRatio field value
// and a boolean to check if the value has been set.
func (o *BrowserWindowDimensions) GetPixelRatioOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PixelRatio, true
}

// SetPixelRatio sets field value
func (o *BrowserWindowDimensions) SetPixelRatio(v int32) {
	o.PixelRatio = v
}

// GetMobileDevice returns the MobileDevice field value if set, zero value otherwise.
func (o *BrowserWindowDimensions) GetMobileDevice() string {
	if o == nil || o.MobileDevice == nil {
		var ret string
		return ret
	}
	return *o.MobileDevice
}

// GetMobileDeviceOk returns a tuple with the MobileDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserWindowDimensions) GetMobileDeviceOk() (*string, bool) {
	if o == nil || o.MobileDevice == nil {
		return nil, false
	}
	return o.MobileDevice, true
}

// HasMobileDevice returns a boolean if a field has been set.
func (o *BrowserWindowDimensions) HasMobileDevice() bool {
	if o != nil && o.MobileDevice != nil {
		return true
	}

	return false
}

// SetMobileDevice gets a reference to the given string and assigns it to the MobileDevice field.
func (o *BrowserWindowDimensions) SetMobileDevice(v string) {
	o.MobileDevice = &v
}

func (o BrowserWindowDimensions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["IsMobile"] = o.IsMobile
	}
	if true {
		toSerialize["Width"] = o.Width
	}
	if true {
		toSerialize["Height"] = o.Height
	}
	if true {
		toSerialize["PixelRatio"] = o.PixelRatio
	}
	if o.MobileDevice != nil {
		toSerialize["MobileDevice"] = o.MobileDevice
	}
	return json.Marshal(toSerialize)
}

type NullableBrowserWindowDimensions struct {
	value *BrowserWindowDimensions
	isSet bool
}

func (v NullableBrowserWindowDimensions) Get() *BrowserWindowDimensions {
	return v.value
}

func (v *NullableBrowserWindowDimensions) Set(val *BrowserWindowDimensions) {
	v.value = val
	v.isSet = true
}

func (v NullableBrowserWindowDimensions) IsSet() bool {
	return v.isSet
}

func (v *NullableBrowserWindowDimensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrowserWindowDimensions(val *BrowserWindowDimensions) *NullableBrowserWindowDimensions {
	return &NullableBrowserWindowDimensions{value: val, isSet: true}
}

func (v NullableBrowserWindowDimensions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrowserWindowDimensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

