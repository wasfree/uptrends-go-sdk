# MetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Period** | Pointer to [**PeriodMetaData**](PeriodMetaData.md) |  | [optional] 

## Methods

### NewMetaData

`func NewMetaData() *MetaData`

NewMetaData instantiates a new MetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaDataWithDefaults

`func NewMetaDataWithDefaults() *MetaData`

NewMetaDataWithDefaults instantiates a new MetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *MetaData) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MetaData) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MetaData) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MetaData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetPeriod

`func (o *MetaData) GetPeriod() PeriodMetaData`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *MetaData) GetPeriodOk() (*PeriodMetaData, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *MetaData) SetPeriod(v PeriodMetaData)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *MetaData) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


