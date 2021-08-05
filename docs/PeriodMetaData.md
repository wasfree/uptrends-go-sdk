# PeriodMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warning** | Pointer to **string** |  | [optional] 
**MaximumRetentionDays** | **int32** |  | 

## Methods

### NewPeriodMetaData

`func NewPeriodMetaData(maximumRetentionDays int32, ) *PeriodMetaData`

NewPeriodMetaData instantiates a new PeriodMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeriodMetaDataWithDefaults

`func NewPeriodMetaDataWithDefaults() *PeriodMetaData`

NewPeriodMetaDataWithDefaults instantiates a new PeriodMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarning

`func (o *PeriodMetaData) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *PeriodMetaData) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *PeriodMetaData) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *PeriodMetaData) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetMaximumRetentionDays

`func (o *PeriodMetaData) GetMaximumRetentionDays() int32`

GetMaximumRetentionDays returns the MaximumRetentionDays field if non-nil, zero value otherwise.

### GetMaximumRetentionDaysOk

`func (o *PeriodMetaData) GetMaximumRetentionDaysOk() (*int32, bool)`

GetMaximumRetentionDaysOk returns a tuple with the MaximumRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRetentionDays

`func (o *PeriodMetaData) SetMaximumRetentionDays(v int32)`

SetMaximumRetentionDays sets MaximumRetentionDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


