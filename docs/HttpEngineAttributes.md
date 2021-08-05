# HttpEngineAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepResults** | Pointer to [**[]HttpEngineStep**](HttpEngineStep.md) | The results of the steps  | [optional] 
**TimingInfo** | Pointer to [**StepTimingInfo**](StepTimingInfo.md) | Timing info | [optional] 
**TotalSteps** | **int32** | Number of total steps | 
**PassedSteps** | **int32** | Number of passed/succeed tests | 

## Methods

### NewHttpEngineAttributes

`func NewHttpEngineAttributes(totalSteps int32, passedSteps int32, ) *HttpEngineAttributes`

NewHttpEngineAttributes instantiates a new HttpEngineAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpEngineAttributesWithDefaults

`func NewHttpEngineAttributesWithDefaults() *HttpEngineAttributes`

NewHttpEngineAttributesWithDefaults instantiates a new HttpEngineAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepResults

`func (o *HttpEngineAttributes) GetStepResults() []HttpEngineStep`

GetStepResults returns the StepResults field if non-nil, zero value otherwise.

### GetStepResultsOk

`func (o *HttpEngineAttributes) GetStepResultsOk() (*[]HttpEngineStep, bool)`

GetStepResultsOk returns a tuple with the StepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepResults

`func (o *HttpEngineAttributes) SetStepResults(v []HttpEngineStep)`

SetStepResults sets StepResults field to given value.

### HasStepResults

`func (o *HttpEngineAttributes) HasStepResults() bool`

HasStepResults returns a boolean if a field has been set.

### GetTimingInfo

`func (o *HttpEngineAttributes) GetTimingInfo() StepTimingInfo`

GetTimingInfo returns the TimingInfo field if non-nil, zero value otherwise.

### GetTimingInfoOk

`func (o *HttpEngineAttributes) GetTimingInfoOk() (*StepTimingInfo, bool)`

GetTimingInfoOk returns a tuple with the TimingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimingInfo

`func (o *HttpEngineAttributes) SetTimingInfo(v StepTimingInfo)`

SetTimingInfo sets TimingInfo field to given value.

### HasTimingInfo

`func (o *HttpEngineAttributes) HasTimingInfo() bool`

HasTimingInfo returns a boolean if a field has been set.

### GetTotalSteps

`func (o *HttpEngineAttributes) GetTotalSteps() int32`

GetTotalSteps returns the TotalSteps field if non-nil, zero value otherwise.

### GetTotalStepsOk

`func (o *HttpEngineAttributes) GetTotalStepsOk() (*int32, bool)`

GetTotalStepsOk returns a tuple with the TotalSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSteps

`func (o *HttpEngineAttributes) SetTotalSteps(v int32)`

SetTotalSteps sets TotalSteps field to given value.


### GetPassedSteps

`func (o *HttpEngineAttributes) GetPassedSteps() int32`

GetPassedSteps returns the PassedSteps field if non-nil, zero value otherwise.

### GetPassedStepsOk

`func (o *HttpEngineAttributes) GetPassedStepsOk() (*int32, bool)`

GetPassedStepsOk returns a tuple with the PassedSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassedSteps

`func (o *HttpEngineAttributes) SetPassedSteps(v int32)`

SetPassedSteps sets PassedSteps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


