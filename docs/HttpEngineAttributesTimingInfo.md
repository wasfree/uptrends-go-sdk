# HttpEngineAttributesTimingInfo

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

### NewHttpEngineAttributesTimingInfo

`func NewHttpEngineAttributesTimingInfo(startUtc map[string]interface{}, endUtc map[string]interface{}, elapsedMilliseconds int64, delayMilliseconds int64, isValid bool, ) *HttpEngineAttributesTimingInfo`

NewHttpEngineAttributesTimingInfo instantiates a new HttpEngineAttributesTimingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpEngineAttributesTimingInfoWithDefaults

`func NewHttpEngineAttributesTimingInfoWithDefaults() *HttpEngineAttributesTimingInfo`

NewHttpEngineAttributesTimingInfoWithDefaults instantiates a new HttpEngineAttributesTimingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *HttpEngineAttributesTimingInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HttpEngineAttributesTimingInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HttpEngineAttributesTimingInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HttpEngineAttributesTimingInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartUtc

`func (o *HttpEngineAttributesTimingInfo) GetStartUtc() map[string]interface{}`

GetStartUtc returns the StartUtc field if non-nil, zero value otherwise.

### GetStartUtcOk

`func (o *HttpEngineAttributesTimingInfo) GetStartUtcOk() (*map[string]interface{}, bool)`

GetStartUtcOk returns a tuple with the StartUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartUtc

`func (o *HttpEngineAttributesTimingInfo) SetStartUtc(v map[string]interface{})`

SetStartUtc sets StartUtc field to given value.


### GetEndUtc

`func (o *HttpEngineAttributesTimingInfo) GetEndUtc() map[string]interface{}`

GetEndUtc returns the EndUtc field if non-nil, zero value otherwise.

### GetEndUtcOk

`func (o *HttpEngineAttributesTimingInfo) GetEndUtcOk() (*map[string]interface{}, bool)`

GetEndUtcOk returns a tuple with the EndUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUtc

`func (o *HttpEngineAttributesTimingInfo) SetEndUtc(v map[string]interface{})`

SetEndUtc sets EndUtc field to given value.


### GetElapsedMilliseconds

`func (o *HttpEngineAttributesTimingInfo) GetElapsedMilliseconds() int64`

GetElapsedMilliseconds returns the ElapsedMilliseconds field if non-nil, zero value otherwise.

### GetElapsedMillisecondsOk

`func (o *HttpEngineAttributesTimingInfo) GetElapsedMillisecondsOk() (*int64, bool)`

GetElapsedMillisecondsOk returns a tuple with the ElapsedMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedMilliseconds

`func (o *HttpEngineAttributesTimingInfo) SetElapsedMilliseconds(v int64)`

SetElapsedMilliseconds sets ElapsedMilliseconds field to given value.


### GetDelayMilliseconds

`func (o *HttpEngineAttributesTimingInfo) GetDelayMilliseconds() int64`

GetDelayMilliseconds returns the DelayMilliseconds field if non-nil, zero value otherwise.

### GetDelayMillisecondsOk

`func (o *HttpEngineAttributesTimingInfo) GetDelayMillisecondsOk() (*int64, bool)`

GetDelayMillisecondsOk returns a tuple with the DelayMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMilliseconds

`func (o *HttpEngineAttributesTimingInfo) SetDelayMilliseconds(v int64)`

SetDelayMilliseconds sets DelayMilliseconds field to given value.


### GetSubTimingInfos

`func (o *HttpEngineAttributesTimingInfo) GetSubTimingInfos() []StepTimingInfo`

GetSubTimingInfos returns the SubTimingInfos field if non-nil, zero value otherwise.

### GetSubTimingInfosOk

`func (o *HttpEngineAttributesTimingInfo) GetSubTimingInfosOk() (*[]StepTimingInfo, bool)`

GetSubTimingInfosOk returns a tuple with the SubTimingInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTimingInfos

`func (o *HttpEngineAttributesTimingInfo) SetSubTimingInfos(v []StepTimingInfo)`

SetSubTimingInfos sets SubTimingInfos field to given value.

### HasSubTimingInfos

`func (o *HttpEngineAttributesTimingInfo) HasSubTimingInfos() bool`

HasSubTimingInfos returns a boolean if a field has been set.

### GetIsValid

`func (o *HttpEngineAttributesTimingInfo) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *HttpEngineAttributesTimingInfo) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *HttpEngineAttributesTimingInfo) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


