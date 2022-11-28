# StepTimingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**StartUtc** | **map[string]interface{}** |  | 
**EndUtc** | **map[string]interface{}** |  | 
**ElapsedMilliseconds** | **int64** |  | 
**DelayMilliseconds** | **int64** |  | 
**SubTimingInfos** | Pointer to [**[]StepTimingInfo**](StepTimingInfo.md) |  | [optional] 
**IsValid** | **bool** | If true, this TimingInfo should be counted as part of the sum of its siblings. If false, the TimingInfo should be discarded (e.g. for PreDelays we don&#39;t want to count). | 

## Methods

### NewStepTimingInfo

`func NewStepTimingInfo(startUtc map[string]interface{}, endUtc map[string]interface{}, elapsedMilliseconds int64, delayMilliseconds int64, isValid bool, ) *StepTimingInfo`

NewStepTimingInfo instantiates a new StepTimingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepTimingInfoWithDefaults

`func NewStepTimingInfoWithDefaults() *StepTimingInfo`

NewStepTimingInfoWithDefaults instantiates a new StepTimingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StepTimingInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StepTimingInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StepTimingInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StepTimingInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartUtc

`func (o *StepTimingInfo) GetStartUtc() map[string]interface{}`

GetStartUtc returns the StartUtc field if non-nil, zero value otherwise.

### GetStartUtcOk

`func (o *StepTimingInfo) GetStartUtcOk() (*map[string]interface{}, bool)`

GetStartUtcOk returns a tuple with the StartUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartUtc

`func (o *StepTimingInfo) SetStartUtc(v map[string]interface{})`

SetStartUtc sets StartUtc field to given value.


### GetEndUtc

`func (o *StepTimingInfo) GetEndUtc() map[string]interface{}`

GetEndUtc returns the EndUtc field if non-nil, zero value otherwise.

### GetEndUtcOk

`func (o *StepTimingInfo) GetEndUtcOk() (*map[string]interface{}, bool)`

GetEndUtcOk returns a tuple with the EndUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUtc

`func (o *StepTimingInfo) SetEndUtc(v map[string]interface{})`

SetEndUtc sets EndUtc field to given value.


### GetElapsedMilliseconds

`func (o *StepTimingInfo) GetElapsedMilliseconds() int64`

GetElapsedMilliseconds returns the ElapsedMilliseconds field if non-nil, zero value otherwise.

### GetElapsedMillisecondsOk

`func (o *StepTimingInfo) GetElapsedMillisecondsOk() (*int64, bool)`

GetElapsedMillisecondsOk returns a tuple with the ElapsedMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedMilliseconds

`func (o *StepTimingInfo) SetElapsedMilliseconds(v int64)`

SetElapsedMilliseconds sets ElapsedMilliseconds field to given value.


### GetDelayMilliseconds

`func (o *StepTimingInfo) GetDelayMilliseconds() int64`

GetDelayMilliseconds returns the DelayMilliseconds field if non-nil, zero value otherwise.

### GetDelayMillisecondsOk

`func (o *StepTimingInfo) GetDelayMillisecondsOk() (*int64, bool)`

GetDelayMillisecondsOk returns a tuple with the DelayMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMilliseconds

`func (o *StepTimingInfo) SetDelayMilliseconds(v int64)`

SetDelayMilliseconds sets DelayMilliseconds field to given value.


### GetSubTimingInfos

`func (o *StepTimingInfo) GetSubTimingInfos() []StepTimingInfo`

GetSubTimingInfos returns the SubTimingInfos field if non-nil, zero value otherwise.

### GetSubTimingInfosOk

`func (o *StepTimingInfo) GetSubTimingInfosOk() (*[]StepTimingInfo, bool)`

GetSubTimingInfosOk returns a tuple with the SubTimingInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTimingInfos

`func (o *StepTimingInfo) SetSubTimingInfos(v []StepTimingInfo)`

SetSubTimingInfos sets SubTimingInfos field to given value.

### HasSubTimingInfos

`func (o *StepTimingInfo) HasSubTimingInfos() bool`

HasSubTimingInfos returns a boolean if a field has been set.

### GetIsValid

`func (o *StepTimingInfo) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *StepTimingInfo) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *StepTimingInfo) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


