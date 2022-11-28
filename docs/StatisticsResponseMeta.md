# StatisticsResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **map[string]interface{}** |  | [optional] 
**Period** | Pointer to [**PeriodMetaData**](PeriodMetaData.md) |  | [optional] 

## Methods

### NewStatisticsResponseMeta

`func NewStatisticsResponseMeta() *StatisticsResponseMeta`

NewStatisticsResponseMeta instantiates a new StatisticsResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsResponseMetaWithDefaults

`func NewStatisticsResponseMetaWithDefaults() *StatisticsResponseMeta`

NewStatisticsResponseMetaWithDefaults instantiates a new StatisticsResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *StatisticsResponseMeta) GetTimestamp() map[string]interface{}`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StatisticsResponseMeta) GetTimestampOk() (*map[string]interface{}, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StatisticsResponseMeta) SetTimestamp(v map[string]interface{})`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *StatisticsResponseMeta) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetPeriod

`func (o *StatisticsResponseMeta) GetPeriod() PeriodMetaData`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *StatisticsResponseMeta) GetPeriodOk() (*PeriodMetaData, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *StatisticsResponseMeta) SetPeriod(v PeriodMetaData)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *StatisticsResponseMeta) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


