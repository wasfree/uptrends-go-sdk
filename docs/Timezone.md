# Timezone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimezoneId** | **int32** | The time zone unique identifier | 
**Description** | Pointer to **string** | The description of the time zone | [optional] 
**OffsetFromUtc** | **int32** | The offset from UTC in minutes (if this time zone runs behind UTC, the number is negative) | 
**HasDaylightSaving** | **bool** | Indicates whether or not this time zone uses Daylight Saving Time | 
**DaylightSavingOffset** | Pointer to **int32** | The time offset for Daylight Saving Time in minutes | [optional] 

## Methods

### NewTimezone

`func NewTimezone(timezoneId int32, offsetFromUtc int32, hasDaylightSaving bool, ) *Timezone`

NewTimezone instantiates a new Timezone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimezoneWithDefaults

`func NewTimezoneWithDefaults() *Timezone`

NewTimezoneWithDefaults instantiates a new Timezone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezoneId

`func (o *Timezone) GetTimezoneId() int32`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *Timezone) GetTimezoneIdOk() (*int32, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *Timezone) SetTimezoneId(v int32)`

SetTimezoneId sets TimezoneId field to given value.


### GetDescription

`func (o *Timezone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Timezone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Timezone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Timezone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOffsetFromUtc

`func (o *Timezone) GetOffsetFromUtc() int32`

GetOffsetFromUtc returns the OffsetFromUtc field if non-nil, zero value otherwise.

### GetOffsetFromUtcOk

`func (o *Timezone) GetOffsetFromUtcOk() (*int32, bool)`

GetOffsetFromUtcOk returns a tuple with the OffsetFromUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetFromUtc

`func (o *Timezone) SetOffsetFromUtc(v int32)`

SetOffsetFromUtc sets OffsetFromUtc field to given value.


### GetHasDaylightSaving

`func (o *Timezone) GetHasDaylightSaving() bool`

GetHasDaylightSaving returns the HasDaylightSaving field if non-nil, zero value otherwise.

### GetHasDaylightSavingOk

`func (o *Timezone) GetHasDaylightSavingOk() (*bool, bool)`

GetHasDaylightSavingOk returns a tuple with the HasDaylightSaving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDaylightSaving

`func (o *Timezone) SetHasDaylightSaving(v bool)`

SetHasDaylightSaving sets HasDaylightSaving field to given value.


### GetDaylightSavingOffset

`func (o *Timezone) GetDaylightSavingOffset() int32`

GetDaylightSavingOffset returns the DaylightSavingOffset field if non-nil, zero value otherwise.

### GetDaylightSavingOffsetOk

`func (o *Timezone) GetDaylightSavingOffsetOk() (*int32, bool)`

GetDaylightSavingOffsetOk returns a tuple with the DaylightSavingOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaylightSavingOffset

`func (o *Timezone) SetDaylightSavingOffset(v int32)`

SetDaylightSavingOffset sets DaylightSavingOffset field to given value.

### HasDaylightSavingOffset

`func (o *Timezone) HasDaylightSavingOffset() bool`

HasDaylightSavingOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


