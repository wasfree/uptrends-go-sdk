# RumMetricValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | [**RumMetric**](RumMetric.md) |  | 
**Median** | Pointer to **int32** |  | [optional] 
**Minimum** | Pointer to **int32** |  | [optional] 
**Maximum** | Pointer to **int32** |  | [optional] 
**Average** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewRumMetricValues

`func NewRumMetricValues(metric RumMetric, ) *RumMetricValues`

NewRumMetricValues instantiates a new RumMetricValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRumMetricValuesWithDefaults

`func NewRumMetricValuesWithDefaults() *RumMetricValues`

NewRumMetricValuesWithDefaults instantiates a new RumMetricValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *RumMetricValues) GetMetric() RumMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *RumMetricValues) GetMetricOk() (*RumMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *RumMetricValues) SetMetric(v RumMetric)`

SetMetric sets Metric field to given value.


### GetMedian

`func (o *RumMetricValues) GetMedian() int32`

GetMedian returns the Median field if non-nil, zero value otherwise.

### GetMedianOk

`func (o *RumMetricValues) GetMedianOk() (*int32, bool)`

GetMedianOk returns a tuple with the Median field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedian

`func (o *RumMetricValues) SetMedian(v int32)`

SetMedian sets Median field to given value.

### HasMedian

`func (o *RumMetricValues) HasMedian() bool`

HasMedian returns a boolean if a field has been set.

### GetMinimum

`func (o *RumMetricValues) GetMinimum() int32`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *RumMetricValues) GetMinimumOk() (*int32, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *RumMetricValues) SetMinimum(v int32)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *RumMetricValues) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetMaximum

`func (o *RumMetricValues) GetMaximum() int32`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *RumMetricValues) GetMaximumOk() (*int32, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *RumMetricValues) SetMaximum(v int32)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *RumMetricValues) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetAverage

`func (o *RumMetricValues) GetAverage() int32`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *RumMetricValues) GetAverageOk() (*int32, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *RumMetricValues) SetAverage(v int32)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *RumMetricValues) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### GetCount

`func (o *RumMetricValues) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RumMetricValues) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RumMetricValues) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *RumMetricValues) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


